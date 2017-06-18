// Package rootApi defines actions of the Root API
package rootApi

import (
	"fmt"
	"net/http"

	"../../api"
)

// GetActions returns the actions of the Root api
func GetActions() []api.Action {
	var rootAPIActions []api.Action
	rootAPIActions = append(rootAPIActions, api.Action{
		Route: "/",
		Verb:  "GET",
		Action: func(w http.ResponseWriter, r *http.Request) {
			Get(w, r)
		},
	})
	return rootAPIActions
}

// Get provides implementation of rootApi.Get
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rootApi: Get: BEGIN")

	w.WriteHeader(http.StatusOK)

	fmt.Println("rootApi: Get: END")
}
