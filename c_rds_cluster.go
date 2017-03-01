/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// RDSCluster : ...
type RDSCluster struct {
}

// Handle : ...
func (n *RDSCluster) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "rds_cluster.create.done", "rds_cluster.create.error":
		lines = n.getSingleDetail(components, "RDS cluster created")
	case "rds_cluster.update.done", "rds_cluster.update.error":
		lines = n.getSingleDetail(components, "RDS cluster updated")
	case "rds_cluster.delete.done", "rds_cluster.delete.error":
		lines = n.getSingleDetail(components, "RDS cluster deleted")
	case "rds_cluster.find.done", "rds_cluster.find.error":
		lines = n.getSingleDetail(components, "RDS cluster found")
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
