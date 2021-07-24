package model

type Message struct {
	Message    []byte
	Attributes map[string]string
	Id         uint
}
