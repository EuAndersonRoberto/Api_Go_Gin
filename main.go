package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Anderson Roberto",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos) // Aqui definimos que toda requisão GET feita para o endpoint: /alunos, irá responder com a função "ExibeTodosAlunos", nos devolvendo o código 200 e a mensagem JSON: "id":"1" e "nome":"Anderson Roberto", na URL: http://localhost:8080/alunos
	r.Run()
}
