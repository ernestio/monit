/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// EBSVolume : ...
type EBSVolume struct {
}

// Handle : ...
func (n *EBSVolume) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "ebs_volume.create":
		lines = n.getSingleDetail(c, "Created EBS volume ")
	case "ebs_volume.delete":
		lines = n.getSingleDetail(c, "Deleted EBS volume ")
	case "ebs_volumes.find":
		lines = n.getSingleDetail(c, "Found EBS volume ")
	}
	return lines
}

func (n *EBSVolume) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	id, _ := c["volume_aws_id"].(string)
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
