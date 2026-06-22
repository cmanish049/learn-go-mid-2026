package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

var todos = []Todo{
	{ID: 1, Title: "Buy groceries", Complete: false},
	{ID: 2, Title: "Walk the dog", Complete: true},
	{ID: 3, Title: "Finish Go project", Complete: false},
	{ID: 4, Title: "Read a book chapter", Complete: false},
	{ID: 5, Title: "Clean the kitchen", Complete: true},
	{ID: 6, Title: "Reply to emails", Complete: false},
	{ID: 7, Title: "Exercise for 30 minutes", Complete: false},
	{ID: 8, Title: "Pay utility bills", Complete: true},
	{ID: 9, Title: "Schedule dentist appointment", Complete: false},
	{ID: 10, Title: "Plan weekend trip", Complete: false},
}

func getTodos(c *gin.Context) {
	completed := c.Query("completed")
	if completed == "" {
		c.JSON(http.StatusOK, todos)
		return
	}
	isCompleted, err := strconv.ParseBool(completed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "complete must be true or false",
		})
		return
	}

	var filtered []Todo

	for _, todo := range todos {
		if todo.Complete == isCompleted {
			filtered = append(filtered, todo)
		}
	}
	c.JSON(http.StatusOK, filtered)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "todo not found",
	})
}
func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	router.GET("todos/:id", getTodo)

	router.Run()
}
