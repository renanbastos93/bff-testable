package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Users struct {
	Data []User `json:"data"`
}

var url = "https://reqres.in/api/users"

func getReqresIn() []byte {
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Timeout: 1 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func filterByID(usr Users, id int) User {
	for _, v := range usr.Data {
		if v.ID == id {
			return v
		}
	}
	return User{}
}

func getAll(ctx *fiber.Ctx) {
	b := getReqresIn()
	var res Users
	json.Unmarshal(b, &res)
	ctx.JSON(res)
}

func getID(ctx *fiber.Ctx) {
	if id, _ := strconv.Atoi(ctx.Params("id")); id > 0 {
		b := getReqresIn()
		var res Users
		json.Unmarshal(b, &res)
		ctx.JSON(filterByID(res, id))
		return
	}
	ctx.SendStatus(fiber.StatusInternalServerError)
}

func main() {
	app := fiber.New()

	app.Get("/", getAll)
	app.Get("/:id", getID)

	app.Listen(3000)
}
