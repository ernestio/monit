/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"net/http"
)

func unauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func authMiddleware(w http.ResponseWriter, r *http.Request) {
	// Check Auth, Until Proper Auth Service is implemented
	authToken := r.Header.Get("X-Auth-Token")
	fmt.Println(authToken)
	if authToken == "" {
		unauthorized(w)
		return
	}

	user, err := db.Get(authToken).Result()
	fmt.Println(user)
	fmt.Println(err)
	if err != nil || user == "" {
		unauthorized(w)
		return
	}

	// Pass to sse server
	s.HTTPHandler(w, r)
}
