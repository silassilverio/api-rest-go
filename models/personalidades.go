package models

type Personalidade struct {
	Id       int    `json:"Id"`
	Nome     string `json:"Nome"`
	Historia string `json:"Historia"`
}

var Personalidades []Personalidade
