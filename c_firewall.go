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
