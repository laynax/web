package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

const mainDir = `/home/daniel/workspace/go/src/web/.temp/`

func main() {
	if err := os.Mkdir(mainDir, 0777); !os.IsExist(err) {
		assert(err)
	}

	e := gin.New()
	gp := e.Group("/api")
	gp.Use(recovery)

	fs := gp.Group("/filesystem")
	fs.POST("", create)
	fs.GET("/:id", get)
	fs.PUT("/:id", update)
	fs.DELETE("/:id", deleteHandler)

	e.Run(":3412")
}

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			println(err.(error).Error())
			c.JSON(http.StatusInternalServerError, nil)
		}
	}()
	c.Next()
}

func assert(err error) {
	if err != nil {
		println(err.Error())
	}
}