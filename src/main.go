// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"./api/"
	"./api/healthcheckApi/"
	"./api/rootApi/"
)

func main() {
	fmt.Println("main: BEGIN")
	r := mux.NewRouter()

	fmt.Println("main: Map API Actions...")
	mapAPIActions(r, GetAPIActionsToMap())

	fmt.Println("main: Starting server...")
	http.ListenAndServe(":8080", r)
}

// GetAPIActionsToMap returns an array of all API actions that are to be mapped by the Router
func GetAPIActionsToMap() []api.Action {
	fmt.Println("main: GetAPIActionsToMap: BEGIN")
	var apiActions []api.Action

	apiActions = append(apiActions, rootApi.GetActions()...)
	apiActions = append(apiActions, healthcheckApi.GetActions()...)

	fmt.Println("main: GetAPIActionsToMap: END")
	return apiActions
}

func mapAPIActions(r *mux.Router, apiActions []api.Action) {
	fmt.Println("main: mapAPIActions: BEGIN")
	for _, apiAction := range apiActions {
		a := apiAction
		r.HandleFunc(a.Route, func(w http.ResponseWriter, r *http.Request) {
			a.Action(w, r)
		}).Methods(a.Verb)
	}
	fmt.Println("main: mapAPIActions: END")
}
