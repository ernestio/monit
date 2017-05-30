/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"strings"
)

// VirtualNetwork : ...
type VirtualNetwork struct {
}

// Handle : ...
func (n *VirtualNetwork) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "virtual_network.create":
		lines = n.getSingleDetail(c, "Created Virtual Network")
	case "virtual_network.delete":
		lines = n.getSingleDetail(c, "Deleted Virtual Network")
	case "virtual_networks.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found Virtual Network")...)
		}
	}
	return lines
}

func (n *VirtualNetwork) getSingleDetail(c component, prefix string) (lines []Message) {
	addressSpace := c["address_space"].([]interface{})
	var netlist []string
	for _, a := range addressSpace {
		netlist = append(netlist, a.(string))
	}
	networks := strings.Join(netlist, ", ")
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
	lines = append(lines, Message{Body: "   Address Space : " + networks, Level: ""})
	id, _ := c["id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   ID : " + id, Level: ""})
	}
	if status != "" {
		lines = append(lines, Message{Body: "   Status : " + status, Level: ""})
	}
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
