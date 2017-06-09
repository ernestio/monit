/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"strings"

	"github.com/nats-io/nats"
)

type Component struct {
	ID       string `json:"_component_id"`
	Type     string `json:"_component"`
	State    string `json:"_state"`
	Action   string `json:"_action"`
	Provider string `json:"_provider"`
	Name     string `json:"name"`
	Error    string `json:"error,omitempty"`
	Service  string `json:"service,omitempty"`
}

func componentHandler(msg *nats.Msg) {
	var c Component
	if err := json.Unmarshal(msg.Data, &c); err != nil {
		panic(err)
	}

	id := c.getID()

	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	publishEvent(id, data)
}

func (c *Component) getID() string {
	var pieces []string
	pieces = strings.Split(c.Service, "-")

	return pieces[len(pieces)-1]
}
