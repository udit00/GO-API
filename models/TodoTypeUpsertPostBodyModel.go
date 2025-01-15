package models

type TodoTypeUpsertPostBodyModel struct {
	UserID       int
	SessionID    string
	TodoTypeID   int
	TodoTypeName string
	ColorID      int
}
