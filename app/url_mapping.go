package app

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/controller/ping"
	"github.com/sachin-ghait-cld/bookstore_users_api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/search", user.FindUser)
	router.GET("/user/:user_id", user.GetUser)
	router.POST("/user", user.CreateUser)
}
