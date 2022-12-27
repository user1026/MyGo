package model

type Page struct {
	Size int `json:"size" form:"size"`
	Page int `json:"page" form:"page"`
}
type House struct {
	id string
}
