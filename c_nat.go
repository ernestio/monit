/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Nat struct {
}

func (n *Nat) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "nats.create":
		lines = append(lines, Message{Body: "Creating firewall:", Level: "INFO"})
	case "nats.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats created", Level: "INFO"})
	case "nats.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats creation failed", Level: "INFO"})
	case "nats.update":
		return append(lines, Message{Body: "Updating nats:", Level: "INFO"})
	case "nats.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats updated", Level: "INFO"})
	case "nats.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats modification failed", Level: "INFO"})
	case "nats.delete":
		return append(lines, Message{Body: "Deleting nats:", Level: "INFO"})
	case "nats.delete.done":
		return append(lines, Message{Body: "Nats deleted", Level: "INFO"})
	case "nats.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *Nat) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name := r["name"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
