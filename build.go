/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/r3labs/broadcast"
)

// Build : holds builds values
type Build struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Subject string      `json:"_subject"`
	Changes []Component `json:"changes"`
}

func processBuild(msg *nats.Msg) {
	var s Build

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
	case "build.create", "build.delete", "build.import", "environment.sync":
		if !bc.StreamExists(id) {
			log.Println("Creating stream: ", id)
			bc.CreateStream(id)
		}
		bc.Publish(id, data)
	case "build.create.done", "build.create.error", "build.delete.done", "build.delete.error", "build.import.done", "build.import.error", "environment.sync.done", "environment.sync.error":
		bc.Publish(id, data)
		go func(bc *broadcast.Server) {
			// Wait for any late connecting clients before closing stream
			time.Sleep(broadcast.DefaultMaxInactivity)
			log.Println("Closing stream: ", id)
		}(bc)
	}
}

func (b *Build) getID() string {
	return b.ID
}
