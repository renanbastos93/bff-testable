package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func Test_getAll(t *testing.T) {
	defer gock.Off()

	type data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		Avatar    string `json:"avatar"`
	}

	type out struct {
		Page       int    `json:"page"`
		PerPage    int    `json:"per_page"`
		Total      int    `json:"total"`
		TotalPages int    `json:"total_pages"`
		Data       []data `json:"data"`
	}

	dataMock := []data{
		data{
			ID:        1,
			Email:     "garzao@e.o.cara",
			FirstName: "Mr. Garzao",
			Avatar:    "xyz.jpg",
		},
	}

	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"page":     2,
			"per_page": 6,
			"total":    12,
			"data":     dataMock,
		})

	app := fiber.New()
	app.Get("/", getAll)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Status: %v; got: %v", fiber.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res Users
	json.Unmarshal(body, &res)
	for k, v := range res.Data {
		usr := User{ID: dataMock[k].ID, Email: dataMock[k].Email}
		if !reflect.DeepEqual(v, usr) {
			t.Fatalf("User: %v; GOT: %v", usr, v)
		}
	}

	st.Expect(t, gock.IsDone(), true)

}

func Test_getID(t *testing.T) {
	defer gock.Off()

	type data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		Avatar    string `json:"avatar"`
	}

	type out struct {
		Page       int    `json:"page"`
		PerPage    int    `json:"per_page"`
		Total      int    `json:"total"`
		TotalPages int    `json:"total_pages"`
		Data       []data `json:"data"`
	}

	dataMock := []data{
		data{
			ID:        1,
			Email:     "garzao@e.o.cara",
			FirstName: "Mr. Garzao",
			Avatar:    "xyz.jpg",
		},
	}

	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"page":     2,
			"per_page": 6,
			"total":    12,
			"data":     dataMock,
		})

	app := fiber.New()
	app.Get("/:id", getID)

	req := httptest.NewRequest("GET", "/1", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Status: %v; got: %v", fiber.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res User
	json.Unmarshal(body, &res)
	usr := User{ID: dataMock[0].ID, Email: dataMock[0].Email}
	if !reflect.DeepEqual(res, usr) {
		t.Fatalf("User: %v; GOT: %v", usr, res)
	}

	st.Expect(t, gock.IsDone(), true)
}

func Test_getID_not_success(t *testing.T) {
	defer gock.Off()

	type data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		Avatar    string `json:"avatar"`
	}

	type out struct {
		Page       int    `json:"page"`
		PerPage    int    `json:"per_page"`
		Total      int    `json:"total"`
		TotalPages int    `json:"total_pages"`
		Data       []data `json:"data"`
	}

	dataMock := []data{
		data{
			ID:        1,
			Email:     "garzao@e.o.cara",
			FirstName: "Mr. Garzao",
			Avatar:    "xyz.jpg",
		},
	}

	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"page":     2,
			"per_page": 6,
			"total":    12,
			"data":     dataMock,
		})

	app := fiber.New()
	app.Get("/:id", getID)

	req := httptest.NewRequest("GET", "/garzao", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("Status: %v; got: %v", fiber.StatusInternalServerError, resp.StatusCode)
	}

	st.Expect(t, gock.IsDone(), false)
}
