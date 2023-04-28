package models

type RequestBody struct {
	Command  string `json:"command"`
	Type     string `json:"type"`
	Filename string `json:"filename"`
	FileId   string `json:"content"`
	Content  []byte `json:"string"`
}

type ResponseBody struct {
	Ok bool `json:"ok"`
}

type ResponseAddBody struct {
	Ok     bool   `json:"ok"`
	FileId uint32 `json:"file_id"`
}

type ResponseGetBody struct {
	Ok       bool   `json:"ok"`
	Filename string `json:"filename"`
	Content  []byte `json:"content"`
}

type ResponseListBody struct {
	Ok   bool              `json:"ok"`
	List map[uint32]string `json:"list"`
}
