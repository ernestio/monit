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

	if err := json.Unmarshal(msg.Data, &input); err != nil {
		return
	}
	notification, err := processNotification(msg)
	if err != nil {
		return
	}
	input.ID = notification.getServiceID()

	switch msg.Subject {
	case "routers.create":
		msgLines = routersCreateHandler(input.Components)
	case "routers.delete":
		msgLines = routersDeleteHandler(input.Components)
	case "routers.create.done":
		msgLines = routersCreateDoneHandler(input.Components)
	case "routers.delete.done":
		msgLines = routersDeleteDoneHandler(input.Components)
	case "networks.create":
		msgLines = networksCreateHandler(input.Components)
	case "networks.delete":
		msgLines = networksDeleteHandler(input.Components)
	case "networks.create.done":
		msgLines = networksCreateDoneHandler(input.Components)
	case "networks.delete.done":
		msgLines = networksDeleteDoneHandler(input.Components)
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
	case "vpcs.create.done":
		msgLines = vpcCreateHandler(input.Components)
	case "vpcs.delete.done":
		msgLines = vpcDeleteHandler(input.Components)
	case "routers.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Router", "creation")
	case "routers.delete.error":
		msgLines = genericErrorMessageHandler(input.Components, "Router", "deletion")
	case "networks.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Network", "creation")
	case "networks.delete.error":
		msgLines = genericErrorMessageHandler(input.Components, "Network", "deletion")
	case "instances.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Instance", "creation")
	case "instances.delete.error":
		msgLines = genericErrorMessageHandler(input.Components, "Instance", "deletion")
	case "instances.update.error":
		msgLines = genericErrorMessageHandler(input.Components, "Instance", "modification")
	case "firewalls.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Firewall", "creation")
	case "firewalls.delete.error":
		msgLines = genericErrorMessageHandler(input.Components, "Firewall", "deletion")
	case "firewalls.update.error":
		msgLines = genericErrorMessageHandler(input.Components, "Firewall", "modification")
	case "nats.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Nat", "creation")
	case "nats.delete.error":
		msgLines = genericErrorMessageHandler(input.Components, "Nat", "deletion")
	case "nats.update.error":
		msgLines = genericErrorMessageHandler(input.Components, "Nat", "modification")
	case "bootstraps.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Bootstraping", "")
	case "executions.create.error":
		msgLines = genericErrorMessageHandler(input.Components, "Execution", "")
	case "vpcs.create.error":
		msgLines = vpcErrorMessageHandler(input.Components, "VPC", "creation")
	case "vpcs.delete.error":
		msgLines = vpcErrorMessageHandler(input.Components, "VPC", "deletion")
	}
	for _, v := range msgLines {
		publishMessage(input.ID, &v)
	}
}

func routersCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Creating routers:", Level: "INFO"})
}

func routersDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Deleting router:", Level: "INFO"})
}

func routersCreateDoneHandler(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		ip := r["ip"].(string)
		lines = append(lines, Message{Body: "\t" + ip, Level: ""})
	}

	lines = append(lines, Message{Body: "Routers successfully created", Level: "INFO"})

	return lines
}

func routersDeleteDoneHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Routers deleted", Level: "INFO"})
}

func networksCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Creating networks:", Level: "INFO"})
}

func networksDeleteHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Networks deleted", Level: "INFO"})
}

func networksCreateDoneHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Networks successfully created", Level: "INFO"})
}

func networksDeleteDoneHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "Deleting networks:", Level: "INFO"})
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
	return append(lines, Message{Body: "Bootstraps ran", Level: "INFO"})
}

func vpcCreateHandler(components []interface{}) (lines []Message) {
	return append(lines, Message{Body: "VPC created", Level: "INFO"})
}

func vpcDeleteHandler(components []interface{}) (lines []Message) {
	for _, c := range components {
		component := c.(map[string]interface{})
		msg := component["error_message"].(string)
		if msg != "" {
			return append(lines, Message{Body: msg, Level: "INFO"})
		}
	}
	return append(lines, Message{Body: "VPC deleted", Level: "INFO"})
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

func vpcErrorMessageHandler(components []interface{}, cType, cAction string) (lines []Message) {
	for _, c := range components {
		component := c.(map[string]interface{})
		if component["status"].(string) == "errored" {
			name := component["vpc_id"].(string)
			msg := component["error_message"].(string)
			msg = strings.Replace(msg, ":", " -", -1)
			line := cType + " " + name + " " + cAction + " failed with: \n" + msg
			lines = append(lines, Message{Body: line, Level: "ERROR"})
		}
	}
	return lines
}
