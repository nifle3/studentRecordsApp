package service

type StudentDb interface {
}

type Student struct {
	db *StudentDb
}
