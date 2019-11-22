package models

import (
	db "newland/database"
)

type Person struct {
	ID        uint64 `json: "id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p Person) TableName() string {
	return "person"
}

// 注册Model
func init() {
	db.SqlDB.AutoMigrate(&Person{})
}

// use gorm create new person
func (p *Person) AddPerson() (err error) {
	err = db.SqlDB.Create(p).Error
	return
}

// use gorm update person
func (p *Person) ModPerson() {
	db.SqlDB.Model(&Person{}).Update(p)
}

// use gorm delete person
func (p *Person) DelPerson() {
	db.SqlDB.Delete(p)
}

// use gorm select one person
func (p *Person) GetPerson() {
	db.SqlDB.First(p, p.ID)
}

// use gorm select all person
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	err = db.SqlDB.Find(&persons).Error
	return
}
