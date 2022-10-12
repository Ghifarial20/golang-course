package controllers

import (
	"net/http"
	"rest-api-golang/structs"

	"github.com/gin-gonic/gin"
)

// get data person by id
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id =?", id).First(&person).Error // get Id by gin

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	}
	result = gin.H{
		"result": person,
		"count":  1,
	}
	c.JSON(http.StatusOK, result)
}

// get all data person
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons) // get person by gin
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	}
	result = gin.H{
		"result": persons,
		"count":  len(persons),
	}

	c.JSON(http.StatusOK, result)
}

// Create person data
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	person.First_Name = firstName
	person.Last_Name = lastName
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// Update person
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	newPerson.First_Name = firstName
	newPerson.Last_Name = lastName
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
	}
	result = gin.H{
		"result": "Update successfully",
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var (
		person structs.Person
		result gin.H
	)
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data Not Found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "Delete Unsuccessful",
		}
	}
	result = gin.H{
		"result": "Delete successful",
	}

	c.JSON(http.StatusOK, result)
}
