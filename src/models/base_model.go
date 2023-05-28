package models

type BaseModel interface {
	FindById(id uint) (interface{}, error)
}
