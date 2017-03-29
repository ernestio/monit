/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Vpc : ...
type Vpc struct {
}

// Handle : ...
func (n *Vpc) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "vpc.create":
		lines = n.getSingleDetail(c, "Created VPC")
	case "vpc.update":
		lines = n.getSingleDetail(c, "Updated VPC")
	case "vpc.delete":
		lines = n.getSingleDetail(c, "Deleted VPC")
	case "vpcs.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found VPC")...)
		}
	}
	return lines
}

func (n *Vpc) getSingleDetail(c component, prefix string) (lines []Message) {
	id, _ := c["vpc_id"].(string)
	if prefix != "" {
		id = prefix + " " + id
	}
	subnet, _ := c["subnet"].(string)
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	if status != "errored" && status != "completed" && status != "" {
		return lines
	}
	lines = append(lines, Message{Body: " " + id, Level: level})
	lines = append(lines, Message{Body: "   Subnet    : " + subnet, Level: ""})
	if status != "" {
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	}
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
