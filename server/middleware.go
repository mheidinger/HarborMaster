package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NeededHeaderMiddleware(header string) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerVal := c.GetHeader(header)
		if headerVal == "" && c.FullPath() != "/static/*filepath" {
			if strings.Contains(c.Request.UserAgent(), "Mozilla") {
				c.HTML(http.StatusUnauthorized, "unauthorized", gin.H{
					"error":   c.Query(errorURLParam),
					"success": c.Query(successURLParam),
				})
			} else {
				c.String(http.StatusUnauthorized, "Unauthorized", gin.H{})
			}
			c.Abort()
		}
	}
}
