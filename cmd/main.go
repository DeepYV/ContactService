package main

import (
	"log"

	"github.com/gin-gonic/gin"
	RouterHandle "test.com/test/router"
)

func main() {

	router := gin.Default()
	RouterHandle.Router(router)
	log.Fatal(router.Run("localhost:8080"))
	return

}
