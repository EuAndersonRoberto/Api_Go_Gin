package routes

import (
	"github.com/Api_Go_Gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos) // Aqui definimos que toda requisão GET feita para o endpoint: /alunos, irá responder com a função "ExibeTodosAlunos", nos devolvendo o código 200 e a mensagem JSON: "id":"1" e "nome":"Anderson Roberto", na URL: http://localhost:8080/alunos
	r.Run()
}
