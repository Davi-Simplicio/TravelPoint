package main

import (
	"TravelPointbackend/internal/controllers"
	"TravelPointbackend/internal/utils"
)


func main() {
	utils.InitGoogleOAuth()
    controllers.Controller()
}
