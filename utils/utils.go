package utils

import (
	"crud-app/instance"
	"github.com/gin-gonic/gin"
)

type ItemManagerHandler func(*gin.Context, instance.ItemManager)

func ItemManagerWrapper(h ItemManagerHandler, param instance.ItemManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c, param) // Call your custom handler with the gin context and your parameter.
	}
}
