/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// RDSInstance : ...
type RDSInstance struct {
}

// Handle : ...
func (n *RDSInstance) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "rds_instance.create":
		lines = n.getSingleDetail(c, "Created RDS Instance")
	case "rds_instance.udpate":
		lines = n.getSingleDetail(c, "Updated RDS Instance")
	case "rds_instance.delete":
		lines = n.getSingleDetail(c, "Deleted RDS Instance")
	case "rds_instances.find":
		for _, cx := range c.getFoundComponents() {
			lines = append(lines, n.getSingleDetail(cx, "Found RDS Instances")...)
		}
	}
	return lines
}

func (n *RDSInstance) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	engine, _ := c["engine"].(string)
	cluster, _ := c["cluster"].(string)
	endpoint, _ := c["endpoint"].(string)
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	if status != "errored" && status != "completed" {
		return lines
	}
	lines = append(lines, Message{Body: " " + name, Level: level})
	lines = append(lines, Message{Body: "   Engine    : " + engine, Level: ""})
	lines = append(lines, Message{Body: "   Cluster   : " + cluster, Level: ""})
	lines = append(lines, Message{Body: "   Endpoint  : " + endpoint, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}

	return lines
}
