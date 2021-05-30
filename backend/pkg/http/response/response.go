package response

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func NewFailedResponse(message string) *ApiResponse {
	return &ApiResponse{
		Data:    nil,
		Message: message,
		Success: false,
	}
}

func NewSuccessResponse(data interface{}, message string) *ApiResponse {
	return &ApiResponse{
		Data:    data,
		Message: message,
		Success: true,
	}
}
