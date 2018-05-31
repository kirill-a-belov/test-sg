package models

type ByURLRequest struct {
	URLs []string `json:"urls"`
}

type ByIdRequest struct {
	RequestID string `json:"req_id"`
}
