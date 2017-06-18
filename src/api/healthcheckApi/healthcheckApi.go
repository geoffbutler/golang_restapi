// Package healthcheckApi defines actions of the Healthcheck API
package healthcheckApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../../api"
	"../../services/dbService"
)

// CheckDbConnectionStatusGetter is a func for overriding logic for dbService.CheckDbConnectionStatus
type CheckDbConnectionStatusGetter func() bool

// DbService is an interface for dbService
type DbService struct {
	checkDbConnectionStatus CheckDbConnectionStatusGetter
}

var dbServiceImpl *DbService

// SetupDbService allows overriding the dbService implementation
func SetupDbService(db *DbService) {
	if db != nil {
		dbServiceImpl = db
		return
	}
	dbServiceImpl = &DbService{
		checkDbConnectionStatus: dbService.CheckDbConnectionStatus,
	}
}

// GetActions returns the actions of the Healthcheck api
func GetActions() []api.Action {
	var healthcheckAPIActions []api.Action
	healthcheckAPIActions = append(healthcheckAPIActions, api.Action{
		Route: "/healthcheck",
		Verb:  "GET",
		Action: func(w http.ResponseWriter, r *http.Request) {
			Get(w, r)
		},
	})
	return healthcheckAPIActions
}

// Get provides implementation of healthcheckApi.Get
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("healthcheckApi: Get: BEGIN")

	dto := HealthcheckDto{
		Status:     "OK",
		DbStatus:   getDbStatus(),
		ServerTime: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)

	fmt.Println("healthcheckApi: Get: END")
}

func getDbStatus() string {
	fmt.Println("healthcheckApi: getDbStatus: BEGIN")
	if dbServiceImpl == nil {
		SetupDbService(nil)
	}

	dbStatus := dbServiceImpl.checkDbConnectionStatus()
	if dbStatus == true {
		fmt.Println("healthcheckApi: getDbStatus: END - OK")
		return "OK"
	}

	fmt.Println("healthcheckApi: getDbStatus: END - BAD")
	return "Bad"
}
