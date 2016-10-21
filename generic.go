/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"strings"

	"github.com/nats-io/nats"
)

type genericMessage struct {
	ID         string        `json:"service"`
	Components []interface{} `json:"components"`
}

func genericHandler(msg *nats.Msg) {
	var msgLines []Message
	var input genericMessage
	var notification Notification

	if err := json.Unmarshal(msg.Data, &input); err != nil {
		return
	}
	if err := processNotification(&notification, msg); err != nil {
		return
	}
	input.ID = notification.getServiceID()

	parts := strings.Split(msg.Subject, ".")
	component := parts[0]

	switch component {
	case "instances":
		var n Instance
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "networks":
		var n Network
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "firewalls":
		var n Firewall
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "nats":
		var n Nat
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "routers":
		var n Router
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "vpcs":
		var n Vpc
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "executions":
		var n Execution
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "bootstraps":
		var n Bootstrap
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "elbs":
		var n ELB
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	case "s3s":
		var n S3Bucket
		msgLines = n.Handle(msg.Subject, input.Components, msgLines)
	default:
		switch msg.Subject {
		case "executions.create.done":
			msgLines = executionsCreateHandler(input.Components)
		case "bootstraps.create.done":
			msgLines = bootstrapsCreateHandler(input.Components)
		case "bootstraps.create.error":
			msgLines = genericErrorMessageHandler(input.Components, "Bootstraping", "")
		case "executions.create.error":
			msgLines = genericErrorMessageHandler(input.Components, "Execution", "")
		}
	}
	for _, v := range msgLines {
		publishMessage(input.ID, &v)
	}
}

func executionsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Executions ran", Level: "INFO"})
}

func bootstrapsCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Bootstraps ran", Level: "INFO"})
}

func genericErrorMessageHandler(components []interface{}, cType, cAction string) (lines []Message) {
	for _, c := range components {
		component := c.(map[string]interface{})
		if component["status"].(string) == "errored" {
			name := component["name"].(string)
			msg := component["error_message"].(string)
			msg = strings.Replace(msg, ":", " -", -1)
			line := cType + " " + name + " " + cAction + " failed with: \n" + msg
			lines = append(lines, Message{Body: line, Level: "ERROR"})
		}
	}

	return lines
}
