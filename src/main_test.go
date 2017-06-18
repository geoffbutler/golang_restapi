package main

import (
	"testing"

	"./api/"
)

func TestAPIActionsToMapHaveExpectedLength(t *testing.T) {
	apiActionsToMap := GetAPIActionsToMap()
	if len(apiActionsToMap) != 2 {
		t.Fatal("apiActionsToMap length is unexpected")
	}
}

func TestAPIActionsToMapIncludeRootActions(t *testing.T) {
	apiActionsToMap := GetAPIActionsToMap()
	if !apiActionsContainsExpectedRouteAndVerb(apiActionsToMap, "/", "GET") {
		t.Fatal("Root GET is not found")
	}
}

func TestAPIActionsToMapIncludeHealthcheckActions(t *testing.T) {
	apiActionsToMap := GetAPIActionsToMap()
	if !apiActionsContainsExpectedRouteAndVerb(apiActionsToMap, "/healthcheck", "GET") {
		t.Fatal("Healthcheck GET is not found")
	}
}

func apiActionsContainsExpectedRouteAndVerb(
	apiActions []api.Action,
	expectedRoute string,
	expectedVerb string,
) bool {
	for _, apiAction := range apiActions {
		if apiAction.Route == expectedRoute && apiAction.Verb == expectedVerb {
			return true
		}
	}
	return false
}

// func TestAPIActionsToMapIncludeRootActions2(t *testing.T) {
// 	apiActionsToMap := GetAPIActionsToMap()
// 	rootAPIActions := rootApi.GetActions()
// 	if !apiActionsContainsExpectedAction(apiActionsToMap, rootAPIActions[0]) {
// 		t.Fatal("Root GET is not found")
// 	}
// }

// func TestAPIActionsToMapIncludeHealthcheckActions2(t *testing.T) {
// 	apiActionsToMap := GetAPIActionsToMap()
// 	healthcheckAPIActions := healthcheckApi.GetActions()
// 	if !apiActionsContainsExpectedAction(apiActionsToMap, healthcheckAPIActions[0]) {
// 		t.Fatal("Healthcheck GET is not found")
// 	}
// }

// func apiActionsContainsExpectedAction(
// 	apiActions []api.Action,
// 	expectedAPIAction api.Action,
// ) bool {
// 	for _, apiAction := range apiActions {
// 		if apiAction.Route == expectedAPIAction.Route &&
// 			apiAction.Verb == expectedAPIAction.Verb &&
// 			actionFuncsAreEqual(apiAction, expectedAPIAction) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func actionFuncsAreEqual(apiAction1 api.Action, apiAction2 api.Action) bool {
// 	sf1 := reflect.ValueOf(apiAction1.Action)
// 	sf2 := reflect.ValueOf(apiAction2.Action)
// 	return sf1.Pointer() == sf2.Pointer()
// }
