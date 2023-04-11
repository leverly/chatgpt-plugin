package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plugin/todo"
)

func AddTodo(c *gin.Context) {
	var task todo.Todo
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.Param("username")
	todo.ToDoTask.Add(username, task)
	c.JSON(http.StatusOK, gin.H{})
}

func GetTodos(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, todo.ToDoTask.Get(username))
}

func DeleteTodo(c *gin.Context) {
	type DeleteTodo struct {
		TodoIdx int `json:"todo_idx"`
	}
	var d DeleteTodo
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.Param("username")
	todo.ToDoTask.Delete(username, d.TodoIdx)
	c.JSON(http.StatusOK, gin.H{})
}
