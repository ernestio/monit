/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type RDSCluster struct {
}

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
	}
	return lines
}

func (n *RDSCluster) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name, _ := r["name"].(string)
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
	}

	return lines
}
