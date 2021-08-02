package router

import "github.com/gin-gonic/gin"

// Route api route
func Route(r *gin.Engine) {
	// REST api
	v1 := r.Group("/api/v1")

	v1.Group("/jwt")
}
