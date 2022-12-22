package controllers

import (
	"github.com/Api_Go_Gin/models"
	"github.com/gin-gonic/gin"
)

// Aqui definimos todos as funções que utilizaremos no projeto mas, poderiamos colocar cada função em um file específico dentro da pasta controllers, modularizando ainda mais nosso código. Neste projeto faremos todas as funções dentro deste único file.
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos) //Aqui fazemos uso da variável "Alunos" que nos trás uma lista da struct "Aluno" criada em "models", para ser exibida pelo JSON.
}

// Saliento que todo o contexto de nossas requisições é controlada por nossa propriedade "c".
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome") // Está variável recupera o parametro "nome" do endpoint criado no file "routes.go" através do "c.Params.ByName("nome")".
	c.JSON(200, gin.H{
		"API diz:": "Iae! " + nome + ", tudo bem com você?",
	})
}
