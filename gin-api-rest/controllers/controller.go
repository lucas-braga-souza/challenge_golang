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
	c.JSON(201, usuario)

}

func CriaNovoUsuario(c *gin.Context) {
	var usuario []models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	result := database.DB.Create(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)

}

func DeletaUsuario(c *gin.Context) {
	var usuario models.Usuario
	username := c.Params.ByName("username")
	database.DB.Where(&models.Usuario{Username: username}).Delete(&usuario)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario deletado com sucesso"})
}

func EditaUsuario(c *gin.Context) {
	var usuario models.Usuario
	username := c.Params.ByName("username")
	result := database.DB.Where(&models.Usuario{Username: username}).First(&usuario)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

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
	result := database.DB.Where(&models.Usuario{Username: username}).First(&usuario)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)

}
