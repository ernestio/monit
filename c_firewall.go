/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Firewall : ...
type Firewall struct {
}

// Handle : ...
func (n *Firewall) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "firewalls.create":
		lines = append(lines, Message{Body: "Creating firewalls:", Level: "INFO"})
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
	case "firewalls.find":
		lines = append(lines, Message{Body: "Importing firewalls:", Level: "INFO"})
	case "firewalls.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls imported", Level: "INFO"})
	case "firewalls.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Firewalls import failed", Level: "INFO"})

	case "firewall.create.done", "firewall.create.error":
		lines = n.getSingleDetail(components, "Firewall created")
	case "firewall.update.done", "firewall.update.error":
		lines = n.getSingleDetail(components, "Firewall updated")
	case "firewall.delete.done", "firewall.delete.error":
		lines = n.getSingleDetail(components, "Firewall deleted")
	case "firewall.find.done", "firewall.find.error":
		lines = n.getSingleDetail(components, "Firewall found")
	}
	return lines
}

func (n *Firewall) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *Firewall) getSingleDetail(v interface{}, prefix string) (lines []Message) {
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
