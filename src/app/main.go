package main

import (
	"app/router"
)

func main() {
	r := router.RouterRegister()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
