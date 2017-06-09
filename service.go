/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"strings"

	"github.com/nats-io/nats"
)

type Service struct {
	ID      string      `json:"id"`
	Changes []Component `json:"changes"`
}

func serviceHandler(msg *nats.Msg) {
	var s Service
	if err := json.Unmarshal(msg.Data, &s); err != nil {
		panic(err)
	}

	id := s.getID()

	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	publishEvent(id, data)
}

func (s *Service) getID() string {
	var pieces []string
	pieces = strings.Split(s.ID, "-")

	return pieces[len(pieces)-1]
}
