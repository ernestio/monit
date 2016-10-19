/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats"
)

// Messages holds a collection of the type Message
type Messages []Message

// Message stores the data of a notification
type Message struct {
	Body  string `json:"body"`
	Level string `json:"level"`
}

// Notification stores any user output sent from the FSM
type Notification struct {
	ID       string   `json:"id"`
	Service  string   `json:"service"`
	Messages Messages `json:"messages"`
}

func (n *Notification) getServiceID() string {
	if n.Service != "" {
		pieces := strings.Split(n.Service, "-")
		return pieces[len(pieces)-1]
	}
	return n.ID
}

func processNotification(notification *Notification, msg *nats.Msg) error {
	return json.Unmarshal(msg.Data, &notification)

}

func publishMessage(service string, msg *Message) {
	data, err := json.Marshal(msg)

	if err != nil {
		log.Println("Could not encode message: ")
		log.Println(err)
	} else {
		s.Publish(service, data)
	}
}
