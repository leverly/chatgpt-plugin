package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.Param("username")
	if _, ok := todos[username]; !ok {
		todos[username] = make([]string, 0)
	}
	todos[username] = append(todos[username], todo.Text)
	c.JSON(http.StatusOK, gin.H{})
}

func getTodos(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, todos[username])
}

func deleteTodo(c *gin.Context) {
	type DeleteTodo struct {
		TodoIdx int `json:"todo_idx"`
	}
	var d DeleteTodo
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.Param("username")
	if d.TodoIdx >= 0 && d.TodoIdx < len(todos[username]) {
		todos[username] = append(todos[username][:d.TodoIdx], todos[username][d.TodoIdx+1:]...)
	}
	c.JSON(http.StatusOK, gin.H{})
}
