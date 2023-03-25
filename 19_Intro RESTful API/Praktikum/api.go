package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Book struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetBookController(c echo.Context) error {
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	defer response.Body.Close()
	responseData, _ := ioutil.ReadAll(response.Body)

	var books []Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get books",
		Status:  "Success",
		Data:    books,
	})
}

func GetBookByIdController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	defer response.Body.Close()
	responseData, _ := ioutil.ReadAll(response.Body)

	var book Book
	json.Unmarshal(responseData, &book)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get book",
		Status:  "Success",
		Data:    book,
	})
}

func CreateBookController(c echo.Context) error {
	var book Book
	id, _ := strconv.Atoi(c.FormValue("id"))
	userId, _ := strconv.Atoi(c.FormValue("userId"))
	book.Id = id
	book.Body = c.FormValue("body")
	book.Title = c.FormValue("title")
	book.UserId = userId

	request, _ := json.Marshal(book)
	byteRequest := bytes.NewBuffer(request)
	response, _ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", byteRequest)
	defer response.Body.Close()

	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to create book",
		Status:  "Success",
		Data:    book,
	})
}

func DeleteBookByIdController(c echo.Context) error {
	id := c.Param("id")
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+id, nil)

	response, _ := client.Do(request)

	defer response.Body.Close()

	responseData, _ := ioutil.ReadAll(response.Body)
	var book Book
	json.Unmarshal(responseData, &book)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to delete book",
		Status:  "Success",
		Data:    book,
	})
}

func main() {

	e := echo.New()
	e.GET("/books", GetBookController)
	e.GET("/books/:id", GetBookByIdController)
	e.POST("/books", CreateBookController)
	e.DELETE("/books/:id", DeleteBookByIdController)
	e.Start(":8000")

}
