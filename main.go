package main

import (
	"final_project_1_gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos = map[int]*models.Todos{}
var seq = 1

func createTodos(c *gin.Context) {
	t := &models.Todos{
		ID: seq,
	}

	var result gin.H

	err := c.Bind(t)
	if err != nil {
		result = gin.H{
			"result": "Gagal Membuat data",
		}
	} else {
		result = gin.H{
			"result": t,
		}
	}
	todos[t.ID] = t
	seq++

	c.JSON(http.StatusCreated, result)
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, todos[id])
}

func updateTodo(c *gin.Context) {
	t := new(models.Todos)
	var result gin.H

	err := c.Bind(t)
	if err != nil {
		result = gin.H{
			"result": "Gagal Mengedit data",
		}
	} else {
		result = gin.H{
			"result": t,
		}
	}
	id, _ := strconv.Atoi(c.Param("id"))

	todos[id].Name = t.Name
	todos[id].Tanggal = t.Tanggal
	todos[id].Kegiatan = t.Kegiatan
	c.JSON(http.StatusOK, result)
}

func deleteTodos(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todos, id)
	c.JSON(http.StatusOK, "Delete data Successfully")
}

func main() {
	r := gin.New()

	r.GET("/todos", getAllTodos)
	r.GET("/todos/:id", getTodoById)
	r.POST("/create-todos", createTodos)
	r.PUT("/update-todos/:id", updateTodo)
	r.DELETE("/delete-todos/:id", deleteTodos)
	r.Run(":9000")
}
