/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	//

	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	var session *Session

	reqid := uuid.New().String()

	log.Println("client connected:", reqid)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		upgradefail(w, err)
		return
	}

	defer func() {
		log.Println("client disconnected:", reqid)
		_ = c.Close()

		if session.channel != nil && session.subscriber != nil {
			session.subscriber.Disconnect(
				session.channel,
			)
		}
	}()

	for {
		if !session.authenticated {
			session, err = authenticate(c, reqid)
			if err != nil {
				badrequest(c, reqid, err)
				return
			}
		}

		msg, ok := <-session.channel
		if !ok {
			log.Printf("[%s] event channel closed: %s\n", reqid, *session.Stream)
			return
		}

		log.Println("sending message to:", *session.Stream)

		err := c.WriteMessage(websocket.TextMessage, msg.Data)
		if err != nil {
			badrequest(c, reqid, err)
			return
		}
	}
}

func upgradefail(w http.ResponseWriter, err error) {
	log.Println("Unable to upgrade to websocket connection:", err.Error())
	http.Error(w, "Unable to upgrade to websocket connection", http.StatusBadRequest)
}

func badrequest(c *websocket.Conn, reqid string, err error) {
	log.Printf("[%s] bad request: %s\n", reqid, err.Error())
	_ = c.WriteMessage(websocket.CloseUnsupportedData, []byte(`{"error": "bad request"}`))
	c.Close()
}
