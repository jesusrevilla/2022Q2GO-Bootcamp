package main

import (
	"github.com/jesusrevilla/capstone/internal/app"
	"github.com/jesusrevilla/capstone/internal/inputport"
	"github.com/jesusrevilla/capstone/internal/interfaceadapter"
)

func main() {
	interfaceAdapterServices := interfaceadapter.NewServices()
	appServices := app.NewServices(interfaceAdapterServices.DataRepository)
	inputPortServices := inputport.NewServices(appServices)
	inputPortServices.Server.ListenAndServe(":8080")
}
