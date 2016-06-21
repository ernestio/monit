/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func unauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func authMiddleware(w http.ResponseWriter, r *http.Request) {
	// Check Auth, Until Proper Auth Service is implemented
	authToken := strings.Trim(r.Header.Get("Authorization"), "Bearer ")
	if authToken == "" {
		unauthorized(w)
		return
	}

	token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		unauthorized(w)
		return
	}

	if token.Valid != true {
		unauthorized(w)
		return
	}

	// Pass to sse server
	s.HTTPHandler(w, r)
}
