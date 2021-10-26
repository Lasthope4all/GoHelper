package models

var Boards []Board

type Board struct {
	Id      int64  `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Tasks   []Task `json:"tasks"`
}

type Task struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
