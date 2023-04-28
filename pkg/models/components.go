package models

type Motor struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgUrl      string `json:"img_url"`
}

type MotorEditRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MotorAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Encoder struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgUrl      string `json:"img_url"`
}

type EncoderEditRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EncoderAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
