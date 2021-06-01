package response

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	// Metadata Metadata    `json:"metadata"`
}

type PaginatedApiResponse struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Success  bool        `json:"success"`
	Metadata Metadata    `json:"metadata"`
}

type Metadata struct {
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	Next       string `json:"next"`
	Previous   string `json:"previous"`
	TotalPages int    `json:"totalPages"`
	TotalSize  int    `json:"totalSize"`
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

func NewPaginatedResponse(data interface{}, page, size, totalPages, totalSize int, message, next, previous string) *PaginatedApiResponse {
	return &PaginatedApiResponse{
		Data:    data,
		Message: message,
		Success: true,
		Metadata: Metadata{
			Page:       page,
			Size:       size,
			TotalPages: totalPages,
			TotalSize:  totalSize,
			Next:       next,
			Previous:   previous,
		},
	}
}
