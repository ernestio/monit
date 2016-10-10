/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Firewall struct {
}

func (n *Firewall) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "firewalls.create":
		lines = append(lines, Message{Body: "Creating firewall:", Level: "INFO"})
	case "firewalls.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls created", Level: "INFO"})
	case "firewalls.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls creation failed", Level: "INFO"})
	case "firewalls.update":
		return append(lines, Message{Body: "Updating firewalls:", Level: "INFO"})
	case "firewalls.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls updated", Level: "INFO"})
	case "firewalls.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls modification failed", Level: "INFO"})
	case "firewalls.delete":
		return append(lines, Message{Body: "Deleting firewalls:", Level: "INFO"})
	case "firewalls.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls deleted", Level: "INFO"})
	case "firewalls.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *Firewall) getDetails(components []interface{}) (lines []Message) {
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
