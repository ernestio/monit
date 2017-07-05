/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
)

type Service struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Subject string      `json:"_subject"`
	Changes []Component `json:"changes"`
}

func processService(msg *nats.Msg) {
	var s Service

	s.Subject = msg.Subject

	if err := json.Unmarshal(msg.Data, &s); err != nil {
		log.Println(err)
		return
	}

	id := s.getID()

	data, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}

	switch msg.Subject {
	case "service.create", "service.delete", "service.import":
		log.Println("Creating stream: ", id)
		ss.CreateStream(id)
		ss.Publish(id, data)
	case "service.create.done", "service.create.error", "service.delete.done", "service.delete.error", "service.import.done", "service.import.error":
		ss.Publish(id, data)
		time.Sleep(10 * time.Millisecond)
		log.Println("Closing stream: ", id)
		go func(ss *sse.Server) {
			ss.RemoveStream(id)
		}(ss)
	}
}

func (s *Service) getID() string {
	var pieces []string
	pieces = strings.Split(s.ID, "-")

	return pieces[len(pieces)-1]
}
