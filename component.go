/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats"
)

// Component : holds component values
type Component struct {
	ID         string      `json:"_component_id"`
	Subject    string      `json:"_subject"`
	Type       string      `json:"_component"`
	State      string      `json:"_state"`
	Action     string      `json:"_action"`
	Provider   string      `json:"_provider"`
	Name       string      `json:"name"`
	Error      string      `json:"error,omitempty"`
	Service    string      `json:"service,omitempty"`
	Components []Component `json:"components,omitempty"`
}

func processComponent(msg *nats.Msg) {
	var c Component

	c.Subject = msg.Subject

	if err := json.Unmarshal(msg.Data, &c); err != nil {
		log.Println(err)
		return
	}

	id := c.getID()
	data, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return
	}

	if bc.StreamExists(id) {
		bc.Publish(id, data)
	}
}

func (c *Component) getID() string {
	return c.Service
}
