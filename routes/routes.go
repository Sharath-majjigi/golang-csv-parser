package routes

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strconv"

	"sharath/database"
	"sharath/models"

	"github.com/gofiber/fiber/v2"
)

func ImportCSV(c *fiber.Ctx) error {
    file, err := c.FormFile("csv")
    if err != nil {
        log.Println(err)
        return c.SendStatus(http.StatusBadRequest)
    }

	f,err := file.Open()
	if err != nil{
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to open the file"})
	}
	defer f.Close()


    db := database.GetDB()

    // Truncate table because i am using same csv file to test so to avoid duplicates
    db.Exec("TRUNCATE TABLE people")


    r := csv.NewReader(f)
    var people []models.Person
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Println(err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to read the file"})
        }

        age, err := strconv.Atoi(record[2])
        if err != nil {
            log.Println(err)
            return c.SendStatus(http.StatusInternalServerError)
        }

        person := models.Person{
            FirstName: record[0],
            LastName:  record[1],
            Age:       age,
			Gender:    record[3],
        }

        db.Create(&person)
        people = append(people, person)
    }

    return c.JSON(people)
}

func GetData(c *fiber.Ctx) error{
	db := database.GetDB()
	var people []models.Person
	db.Find(&people)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": len(people),
		"data": people,
	})
}

func GetPeopleByFName(c *fiber.Ctx) error{
	name := c.Params("fname")
	db := database.GetDB()
	var people[] models.Person

	db.Where("first_name = ?",name).Find(&people)
	if len(people) == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"No person exist with FirstName: "+name})
	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": len(people),
			"data": people,
		})
	}
}

func GetPeopleWithAge(c *fiber.Ctx) error{
	ageParam := c.Params("age")
	age := 0
	if ageParam != "" {
		var err error
		age , err = strconv.Atoi(ageParam)
		if(err!=nil){
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Conversion of age to string failed" })
		}
	}

	db := database.GetDB()
	var people[] models.Person
	db.Where("age = ?",age).Find(&people)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": len(people),
		"data": people,
	})
}

func DeletePeople(c *fiber.Ctx) error{
	ageParam := c.Params("age")
	age := 0
	if ageParam!= ""{
		var err error
		age,err = strconv.Atoi(ageParam)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Conversion of age to string failed"})
		}
	}
	db := database.GetDB()
	var people[] models.Person

	result:= db.Where("age < ?",age).Delete(&people)
	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error occured while deleting people"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success" :"Removed people below age: "+ageParam,
		}) 
}