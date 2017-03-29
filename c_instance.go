/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Instance : ...
type Instance struct {
}

// Handle : ...
func (n *Instance) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "instance.create":
		lines = n.getSingleDetail(c, "Created Instance")
	case "instance.update":
		lines = n.getSingleDetail(c, "Updated Instance")
	case "instance.delete":
		lines = n.getSingleDetail(c, "Deleted Instance")
	case "instances.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found Instances")...)
		}
	}
	return lines
}

func (n *Instance) getSingleDetail(c component, prefix string) (lines []Message) {
	ip, _ := c["ip"].(string)
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	if status != "errored" && status != "completed" {
		return lines
	}
	lines = append(lines, Message{Body: " " + name, Level: level})
	lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
	publicIP, _ := c["public_ip"].(string)
	if publicIP != "" {
		lines = append(lines, Message{Body: "   PUBLIC IP : " + publicIP, Level: ""})
	}
	id, _ := c["instance_aws_id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   AWS ID    : " + id, Level: ""})
	}
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
