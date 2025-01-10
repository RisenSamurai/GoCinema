package services

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func SendLoginData(c *gin.Context) (interface{}, error) {

	var obj interface{}

	return obj, nil
}

func SendRegisterData(c *gin.Context) (interface{}, error) {

	var formData LoginForm

	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	requstData, err := json.Marshal(formData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	resp, err := http.Post("http://localhost:8082/auth/login",
		"application/json", bytes.NewBuffer(requstData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return result, nil

}
