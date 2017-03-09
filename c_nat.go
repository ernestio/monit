/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Nat : ...
type Nat struct {
}

// Handle : ...
func (n *Nat) Handle(subject string, component interface{}, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "nat.create":
		lines = n.getSingleDetail(component, "Nat created")
	case "nat.update":
		lines = n.getSingleDetail(component, "Nat updated")
	case "nat.delete":
		lines = n.getSingleDetail(component, "Nat deleted")
	case "nat.find":
		lines = n.getSingleDetail(component, "Nat created")
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
