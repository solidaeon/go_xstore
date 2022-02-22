package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go_xstore/api/models/storeitem"
)

type storeItemReq struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func ItemPost(i *storeitem.Inventory) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		reqBody := storeItemReq{}

		c.Bind(&reqBody)
		
		item := storeitem.Item{
			Name: reqBody.Name,
			Price: reqBody.Price,
		}

		i.Add(item)

		c.JSON(http.StatusNoContent, nil)
	}
}