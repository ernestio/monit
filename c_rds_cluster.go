/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// RDSCluster : ...
type RDSCluster struct {
}

// Handle : ...
func (n *RDSCluster) Handle(subject string, component interface{}, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "rds_cluster.create":
		lines = n.getSingleDetail(component, "RDS cluster created")
	case "rds_cluster.update":
		lines = n.getSingleDetail(component, "RDS cluster updated")
	case "rds_cluster.delete":
		lines = n.getSingleDetail(component, "RDS cluster deleted")
	case "rds_cluster.find":
		lines = n.getSingleDetail(component, "RDS cluster found")
	}
	return lines
}

func (n *RDSCluster) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	engine, _ := r["engine"].(string)
	endpoint, _ := r["endpoint"].(string)
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Engine    : " + engine, Level: ""})
	lines = append(lines, Message{Body: "   Endpoint  : " + endpoint, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
