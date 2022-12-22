package controllers

import "github.com/gin-gonic/gin"

// Aqui definimos todos as funções que utilizaremos no projeto mas, poderiamos colocar cada função em um file específico dentro da pasta controllers, modularizando ainda mais nosso código. Neste projeto faremos todas as funções dentro deste único file.
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Anderson Roberto",
	})
}
