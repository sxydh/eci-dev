package main

import (
	"eci-res/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":20020")
	if err != nil {
		log.Fatalf("gin.Engine.Run error: %v", err)
	}
}
