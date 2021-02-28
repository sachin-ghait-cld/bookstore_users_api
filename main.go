package main

import (
	"os"

	"github.com/sachin-ghait-cld/bookstore_users_api/app"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/env"
)

func main() {
	err := env.LoadEnv()
	if err != nil {
		os.Exit(1)
	}
	app.StartApp()
}
