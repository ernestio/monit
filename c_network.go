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
		lines = n.getSingleDetail(c, "Network created")
	case "network.delete":
		lines = n.getSingleDetail(c, "Network deleted")
	case "networks.find":
		lines = n.getSingleDetail(c, "Network found")
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
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   IP     : " + ip, Level: ""})
	id, _ := c["network_aws_id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   AWS ID : " + id, Level: ""})
	}
	lines = append(lines, Message{Body: "   Status : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
