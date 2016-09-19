/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

type genericMessage struct {
	ID         string        `json:"service"`
	Components []interface{} `json:"components"`
}

func genericHandler(msg *nats.Msg) {
	var msgLines []Message
	var input genericMessage

	if err := json.Unmarshal(msg.Data, &input); err != nil {
		return
	}
	notification, err := processNotification(msg)
	if err != nil {
		return
	}
	input.ID = notification.getServiceID()

	switch msg.Subject {
	case "routers.create.done":
		msgLines = routersCreateHandler(input.Components)
	case "routers.delete.done":
		msgLines = routersDeleteHandler(input.Components)
	case "networks.create.done":
		msgLines = networksCreateHandler(input.Components)
	case "networks.delete.done":
		msgLines = networksDeleteHandler(input.Components)
	case "instances.create.done":
		msgLines = instancesCreateHandler(input.Components)
	case "instances.update.done":
		msgLines = instancesUpdateHandler(input.Components)
	case "instances.delete.done":
		msgLines = instancesDeleteHandler(input.Components)
	case "firewalls.create.done":
		msgLines = firewallsCreateHandler(input.Components)
	case "firewalls.update.done":
		msgLines = firewallsUpdateHandler(input.Components)
	case "firewalls.delete.done":
		msgLines = firewallsDeleteHandler(input.Components)
	case "nats.create.done":
		msgLines = natsCreateHandler(input.Components)
	case "nats.update.done":
		msgLines = natsUpdateHandler(input.Components)
	case "nats.delete.done":
		msgLines = natsDeleteHandler(input.Components)
	case "executions.create.done":
		msgLines = executionsCreateHandler(input.Components)
	case "bootstraps.create.done":
		msgLines = bootstrapsCreateHandler(input.Components)
	}
	for _, v := range msgLines {
		publishMessage(input.ID, &v)
	}
}

func routersCreateHandler(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		ip := r["ip"].(string)
		lines = append(lines, Message{Body: "\t" + ip, Level: ""})
	}

	lines = append(lines, Message{Body: "Routers successfully created", Level: "INFO"})

	return lines
}

func routersDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Routers deleted", Level: "INFO"})
}

func networksCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Networks successfully created", Level: "INFO"})
}

func networksDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Networks deleted", Level: "INFO"})
}

func instancesCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Instances successfully created", Level: "INFO"})
}

func instancesUpdateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Instances successfully updated", Level: "INFO"})
}

func instancesDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Instances deleted", Level: "INFO"})
}

func firewallsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Firewalls Created", Level: "INFO"})
}

func firewallsUpdateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Firewalls Updated", Level: "INFO"})
}

func firewallsDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Firewalls Deleted", Level: "INFO"})
}

func natsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Nats Created", Level: "INFO"})
}

func natsUpdateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Nats Updated", Level: "INFO"})
}

func natsDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Nats Deleted", Level: "INFO"})
}

func executionsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Executions ran", Level: "INFO"})
}

func bootstrapsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Instances bootstrapped", Level: "INFO"})
}
