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
	case "rds_clusters.create":
		lines = append(lines, Message{Body: "Creating rds clusters:", Level: "INFO"})
	case "rds_clusters.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters created", Level: "INFO"})
	case "rds_clusters.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters creation failed", Level: "INFO"})
	case "rds_clusters.update":
		lines = append(lines, Message{Body: "Updating rds clusters:", Level: "INFO"})
	case "rds_clusters.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters modified", Level: "INFO"})
	case "rds_clusters.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters modification failed", Level: "INFO"})
	case "rds_clusters.delete":
		return append(lines, Message{Body: "Deleting rds clusters:", Level: "INFO"})
	case "rds_clusters.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters deleted", Level: "INFO"})
	case "rds_clusters.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters deletion failed", Level: "INFO"})
	case "rds_clusters.find":
		lines = append(lines, Message{Body: "Importing rds clusters:", Level: "INFO"})
	case "rds_clusters.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters imported", Level: "INFO"})
	case "rds_clusters.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS clusters import failed", Level: "INFO"})

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

func (n *RDSCluster) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
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
