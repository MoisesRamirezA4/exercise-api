package main

import (
	"github.com/epa-datos/exercise-api/api"
	"github.com/epa-datos/exercise-api/repositories/mysql"
)

func main() {
	mysql.Connect()
	api.RunServer()
}
