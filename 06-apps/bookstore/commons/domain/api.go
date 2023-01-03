package cmsdomain

type ResponseInfo struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}
