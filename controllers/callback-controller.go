package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallbackController(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Callback received: %s", string(body))
	c.JSON(http.StatusOK, gin.H{"message": "Callback received"})
}
