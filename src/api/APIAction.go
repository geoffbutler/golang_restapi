package api

import (
	"net/http"
)

// Action defines a routed API action
type Action struct {
	Route  string
	Verb   string
	Action func(http.ResponseWriter, *http.Request)
}
