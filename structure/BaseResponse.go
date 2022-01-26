package structure

type BaseResponse struct {
	Status *int   `json:"status"`
	Error  string `json:"error"`
}
