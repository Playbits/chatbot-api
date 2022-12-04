package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {}
func (p *PongController) Reply(c *gin.Context) {
	 c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
}
