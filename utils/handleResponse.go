package utils

import "udit/api-padhai/models"

func GetSuccessResponse(response any) models.ApiResponse {
	return models.ApiResponse{Status: 1, Message: "Success", Response: response}
}

func GetErrorResponse(message string, response any) models.ApiResponse {
	return models.ApiResponse{Status: -1, Message: message, Response: response}
}
