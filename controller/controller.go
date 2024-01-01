package controller

import (
	"errors"
	"go-crawl/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	page := c.Query("page")
	var data interface{}
	var err error

	switch page {
	case "quantrimang":
		data, err = service.FetchQuanTriMangData()
	case "cellphone":
		data, err = service.FetchCellPhoneData()
	// Add more cases as needed...
	default:
		err = errors.New("Invalid page parameter")
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
