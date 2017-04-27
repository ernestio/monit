/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// StorageContainer : ...
type StorageContainer struct {
}

// Handle : ...
func (n *StorageContainer) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "storage_container.create":
		lines = n.getSingleDetail(c, "Created Storage Container")
	case "storage_container.update":
		lines = n.getSingleDetail(c, "Updated Storage Container")
	case "storage_container.delete":
		lines = n.getSingleDetail(c, "Deleted Storage Container")
	case "storage_containers.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found Storage Container")...)
		}
	}
	return lines
}

func (n *StorageContainer) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	if status != "errored" && status != "completed" && status != "" {
		return lines
	}
	lines = append(lines, Message{Body: " " + name, Level: level})
	id, _ := c["id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   ID    : " + id, Level: ""})
	}
	if status != "" {
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	}
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
