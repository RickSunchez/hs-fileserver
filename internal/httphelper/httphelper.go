package httphelper

type Response struct {
	Message string `json:"message"`
}

type ListResponse struct {
	List []string `json:"list"`
}
