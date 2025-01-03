package models

type ApiResponse struct {
	Status   int16
	Message  string
	Success  bool
	Response any
}
