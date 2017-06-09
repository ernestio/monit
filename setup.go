/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats"
)

type monitorConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func setup() {
	var err error
	// Open Nats connection
	nc, err = nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Println("Could not connect to nats")
		return
	}

	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		token, err := nc.Request("config.get.jwt_token", []byte(""), 1*time.Second)
		if err != nil {
			panic("Can't get jwt_config config")
		}

		secret = string(token.Data)
	}

	cfg := monitorConfig{}
	msg, err := nc.Request("config.get.monitor", []byte(""), 1*time.Second)
	if err != nil {
		panic("Can't get monitor config")
	}
	if err := json.Unmarshal(msg.Data, &cfg); err != nil {
		panic("Can't process monitor config")
	}

	host = cfg.Host
	port = cfg.Port
}
