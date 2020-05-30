package main

import (
	"ccg/ccg"
	"net/http"
)

func main() {
	r := ccg.New()
	r.GET("/", func(c *ccg.Context) {
		c.JSON(http.StatusOK, ccg.H{
			"test": c.PostForm("test"),
			"aaa":  c.PostForm("aaa"),
		})
	})
	r.Run(":9999")
}
