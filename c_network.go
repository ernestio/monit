/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Network : ...
type Network struct {
}

// Handle : ...
func (n *Network) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "network.create":
		lines = n.getSingleDetail(c, "Created Network")
	case "network.delete":
		lines = n.getSingleDetail(c, "Deleted Network")
	case "networks.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found Network")...)
		}
	}
	return lines
}

func (n *Network) getSingleDetail(c component, prefix string) (lines []Message) {
	ip, _ := c["range"].(string)
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
	lines = append(lines, Message{Body: "   IP     : " + ip, Level: ""})
	id, _ := c["network_aws_id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   AWS ID : " + id, Level: ""})
	}
	lines = append(lines, Message{Body: "   Status : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
