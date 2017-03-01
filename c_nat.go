/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Nat : ...
type Nat struct {
}

// Handle : ...
func (n *Nat) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "nats.create":
		lines = append(lines, Message{Body: "Creating nats:", Level: "INFO"})
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
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats deleted", Level: "INFO"})
	case "nats.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats deletion failed", Level: "INFO"})
	case "nats.find":
		lines = append(lines, Message{Body: "Importing nats:", Level: "INFO"})
	case "nats.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats imported", Level: "INFO"})
	case "nats.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Nats import failed", Level: "INFO"})

	case "nat.create.done", "nat.create.error":
		lines = n.getSingleDetail(components, "Nat created")
	case "nat.update.done", "nat.update.error":
		lines = n.getSingleDetail(components, "Nat updated")
	case "nat.delete.done", "nat.delete.error":
		lines = n.getSingleDetail(components, "Nat deleted")
	case "nat.find.done", "nat.find.error":
		lines = n.getSingleDetail(components, "Nat created")
	}
	return lines
}

func (n *Nat) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *Nat) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
