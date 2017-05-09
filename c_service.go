/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Service : ...
type Service struct {
}

// Handle : ...
func (n *Service) Handle(subject string, lines []Message) []Message {
	switch subject {
	case "service.create":
		lines = append(lines, Message{Body: "Applying your definition", Level: "INFO"})
	case "service.delete":
		lines = append(lines, Message{Body: "Starting environment deletion", Level: "INFO"})
	case "service.create.done":
		lines = append(lines, Message{Body: "SUCCESS: rules successfully applied", Level: "SUCCESS"})
		lines = append(lines, Message{Body: "error", Level: "ERROR"})
	case "service.create.error":
		lines = append(lines, Message{Body: "\nOops! Something went wrong. Please manually fix any errors shown above and re-apply your definition.", Level: "INFO"})
		lines = append(lines, Message{Body: "error", Level: "ERROR"})
	case "service.delete.done":
		lines = append(lines, Message{Body: "SUCCESS: your environment has been successfully deleted", Level: "SUCCESS"})
		lines = append(lines, Message{Body: "success", Level: "SUCCESS"})
	case "service.delete.error":
		lines = append(lines, Message{Body: "\nOops! Something went wrong. Please manually fix any errors shown above and re-apply your service deletion.", Level: "INFO"})
		lines = append(lines, Message{Body: "error", Level: "ERROR"})
	case "service.import.done":
		lines = append(lines, Message{Body: "SUCCESS: service successfully imported", Level: "SUCCESS"})
		lines = append(lines, Message{Body: "error", Level: "ERROR"})
	case "service.import.error":
		lines = append(lines, Message{Body: "\nOops! Something went wrong. Please manually fix any errors shown above and re-apply your definition.", Level: "INFO"})
		lines = append(lines, Message{Body: "error", Level: "ERROR"})
	}
	return lines
}
