package healthcheckApi

import (
	"time"
)

// HealthcheckDto used by Healthcheck api Get endpoint
type HealthcheckDto struct {
	Status     string
	DbStatus   string
	ServerTime time.Time
}
