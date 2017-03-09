/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Vpc : ...
type Vpc struct {
}

// Handle : ...
func (n *Vpc) Handle(subject string, component interface{}, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "vpc.create":
		lines = n.getSingleDetail(component, "VPC created")
	case "vpc.update":
		lines = n.getSingleDetail(component, "VPC udpated")
	case "vpc.delete":
		lines = n.getSingleDetail(component, "VPC deleted")
	case "vpcs.find":
		lines = n.getSingleDetail(component, "VPC Found")
	}
	return lines
}

func (n *Vpc) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	id, _ := r["vpc_id"].(string)
	if prefix != "" {
		id = prefix + " " + id
	}
	subnet, _ := r["vpc_subnet"].(string)
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + id, Level: ""})
	lines = append(lines, Message{Body: "   Subnet    : " + subnet, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
