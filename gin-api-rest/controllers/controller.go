package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-braga-souza/challenge_api-go/database"
	"github.com/lucas-braga-souza/challenge_api-go/models"
)

func ExibeTodosUsuario(c *gin.Context) {
	var usuario []models.Usuario
	database.DB.Find(&usuario)
	c.JSON(200, usuario)

}

func CriaNovoUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)

}

func BuscaUsuarioPorID(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func DeletaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario deletado com sucesso"})
}

func EditaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(*&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)
}

func BuscaUsuarioPorUsername(c *gin.Context) {
	var usuario models.Usuario
	username := c.Param("username")
	database.DB.Where(&models.Usuario{Username: username}).First(&usuario)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)

}
