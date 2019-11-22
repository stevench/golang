package apis

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "newland/models"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}

func AddPersonApi(c *gin.Context) {
	var p Person
	c.Bind(&p)
	//firstName := c.Request.FormValue("first_name")
	//lastName := c.Request.FormValue("last_name")
	//city := c.Request.FormValue("city")

	//p := &Person{FirstName: firstName, LastName: lastName, City: city}
	beego.Info(p)
	err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", p.ID)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func ModPersonApi(c *gin.Context) {
	//firstName := c.Request.FormValue("first_name")
	//lastName := c.Request.FormValue("last_name")
	//city := c.Request.FormValue("city")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{ID: id}

	p.GetPerson()
	if p.FirstName != "" {
		//p.FirstName = firstName
		//p.LastName = lastName
		//p.City = city
		c.Bind(&p)
		p.ModPerson()
		msg := fmt.Sprintf("update successful %d", p.ID)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{ID: id}
	p.DelPerson()
	msg := fmt.Sprintf("delete successful %d", p.ID)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	p := &Person{ID: id}
	p.GetPerson()
	if p.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": p,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func GetPersonsApi(c *gin.Context) {
	p := &Person{}
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
