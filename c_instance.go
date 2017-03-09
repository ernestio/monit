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
		lines = n.getSingleDetail(c, "Instance created")
	case "instance.update":
		lines = n.getSingleDetail(c, "Instance udpated")
	case "instance.delete":
		lines = n.getSingleDetail(c, "Instance delete")
	case "instances.find":
		lines = n.getSingleDetail(c, "Instance find")
	}
	return lines
}

func (n *Instance) getSingleDetail(c component, prefix string) (lines []Message) {
	ip, _ := c["ip"].(string)
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
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
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
