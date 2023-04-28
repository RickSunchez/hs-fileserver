package models

type RequestBody struct {
	Command  string `json:"command"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type ResponseBody struct {
	Ok bool `json:"ok"`
}

type ResponseGetBody struct {
	Ok      bool   `json:"ok"`
	Content string `json:"content"`
}

type ResponseListBody struct {
	Ok   bool     `json:"ok"`
	List []string `json:"list"`
}
