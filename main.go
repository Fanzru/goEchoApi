package main

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goEchoApi/config"
	"goEchoApi/routes"
)

func main() {
	e := routes.Init()

	config := config.GetDatabaseConfig()
	fmt.Println(config.URL,"=========================")
	_ = mgm.SetDefaultConfig(nil, "myFirstDatabase", options.Client().ApplyURI(config.URL))

	e.Logger.Fatal(e.Start(":1325"))
}
