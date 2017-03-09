/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// ELB : ...
type ELB struct {
}

// Handle : ...
func (n *ELB) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "elb.create":
		lines = n.getSingleDetail(c, "Created ELB")
	case "elb.update":
		lines = n.getSingleDetail(c, "Updated ELB")
	case "elb.delete":
		lines = n.getSingleDetail(c, "Deleted ELB")
	case "elbs.find":
		lines = n.getSingleDetail(c, "Found ELB")

	}
	return lines
}

func (n *ELB) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if c["dns_name"] != nil {
		dnsName, _ := c["dns_name"].(string)
		if dnsName != "" {
			lines = append(lines, Message{Body: "   DNS    : " + dnsName, Level: ""})
		}
	}
	lines = append(lines)
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}

	return lines
}
