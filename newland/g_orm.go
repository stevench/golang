package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

type Person struct {
	ID        int    `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
}

func (person Person) TableName() string {
	return "person"
}

func init() {
	db, err = gorm.Open("mysql", "root:steven@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&Person{})
}

func add() {
	person := &Person{FirstName: "qiao", LastName: "zhenhua"}
	fmt.Println(person)
	err := db.Create(person)
	fmt.Println(err)
}

func del(id int) {
	person := &Person{ID: id}
	db.Delete(person)
}

func mod(p Person) {
	db.Model(p).Update("LastName", "haha")
}

func main() {
	add()
}
