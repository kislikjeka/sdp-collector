package models

import (
	"database/sql"
)

type CollectorsState struct {
	ID              uint
	CollectorName   string
	LastTimeProceed sql.NullTime
}
