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
	n, err = nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Println("Could not connect to nats")
		return
	}

	// Set the JWT Secret
	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("No JWT secret was set!")
	}

	cfg := monitorConfig{}
	msg, err := n.Request("config.get.monitor", []byte(""), 1*time.Second)
	if err != nil {
		panic("Can't get monitor config")
	}
	json.Unmarshal(msg.Data, &cfg)

	host = cfg.Host
	port = cfg.Port
}
