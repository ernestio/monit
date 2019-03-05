/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/r3labs/broadcast"
)

// Session : stores authentication data
type Session struct {
	Token         string  `json:"token"`
	Stream        *string `json:"stream"`
	EventID       *string `json:"event_id"`
	Username      string  `json:"-"`
	authenticated bool
	subscriber    *broadcast.Subscriber
	channel       chan *broadcast.Event
}

func unauthorized(c *websocket.Conn, err error) error {
	if err != nil {
		log.Println("Unauthorized:", err.Error())
	} else {
		log.Println("Unauthorized")
	}
	_ = c.WriteMessage(websocket.CloseMessage, []byte(`{"status": "unauthorized"}`))
	return errors.New("Unauthorized")
}

func getAuthMessage(c *websocket.Conn, s *Session) error {
	// timeout after 2 seconds if no request is sent
	c.SetReadDeadline(time.Now().Add(time.Second * 5))

	_, message, err := c.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(message, &s)
}

func register(stream *string, username, requestID string) (*broadcast.Subscriber, chan *broadcast.Event, error) {
	if stream == nil {
		return nil, nil, errors.New("no stream specified")
	}

	log.Printf("[%s] subscribing to stream: %s\n", requestID, *stream)

	if !bc.StreamExists(*stream) && !bc.AutoStream {
		return nil, nil, errors.New("stream does not exist")
	} else if !bc.StreamExists(*stream) && bc.AutoStream {
		bc.CreateStream(*stream)
	}

	sub := bc.GetStreamSubscriber(*stream, username)
	if sub == nil {
		sub = broadcast.NewSubscriber(username)
		bc.Register(*stream, sub)
	}

	return sub, sub.Connect(), nil
}

func authenticate(c *websocket.Conn, requestID string) (*Session, error) {
	var s Session

	log.Printf("[%s] authenticating user\n", requestID)

	err := getAuthMessage(c, &s)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(s.Token, jwtVerify)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	s.authenticated = true

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		s.Username = claims["username"].(string)
	}

	log.Printf("[%s] user authenticated\n", requestID)
	err = c.WriteMessage(websocket.TextMessage, []byte(`{"status": "ok"}`))
	if err != nil {
		return nil, err
	}

	// register to stream
	s.subscriber, s.channel, err = register(s.Stream, s.Username, requestID)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func jwtVerify(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return []byte(secret), nil
}
