package main

import (
	"main/pkg/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "alas")
	})
	r.Run("127.0.0.1:8800")
}
