package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go_xstore/api/models/storeitem"
)

func ItemsGet(i *storeitem.Inventory) gin.HandlerFunc {
	return func(c *gin.Context) {
		retval := i.GetItems()
		c.JSON(http.StatusOK, retval)
	}
}