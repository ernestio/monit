/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type S3Bucket struct {
}

func (n *S3Bucket) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "s3s.create":
		lines = append(lines, Message{Body: "Creating s3 buckets:", Level: "INFO"})
	case "s3s.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets created", Level: "INFO"})
	case "s3s.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets creation failed", Level: "INFO"})
	case "s3s.update":
		lines = append(lines, Message{Body: "Updating s3 buckets:", Level: "INFO"})
	case "s3s.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets modified", Level: "INFO"})
	case "s3s.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets modification failed", Level: "INFO"})
	case "s3s.delete":
		return append(lines, Message{Body: "Deleting s3 buckets:", Level: "INFO"})
	case "s3s.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets deleted", Level: "INFO"})
	case "s3s.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *S3Bucket) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name, _ := r["name"].(string)
		acl, _ := r["acl"].(string)
		if acl == "" {
			acl = "by grantees"
		}
		status, _ := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   ACL       : " + acl, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err, _ := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
