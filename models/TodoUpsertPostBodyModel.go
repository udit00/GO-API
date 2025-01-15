package models

import "time"

type TodoUpsertPostBodyModel struct {
	TodoID               int
	Title                string
	Description          string
	UserID               int
	SessionID            string
	TodoTypeID           int
	TargetDateTimeString string
	TargetDateTime       time.Time
	CompletionStatusID   int
}
