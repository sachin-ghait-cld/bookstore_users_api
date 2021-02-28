package app

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/controller/ping"
	"github.com/sachin-ghait-cld/bookstore_users_api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", user.GetUser)
	router.POST("/user", user.CreateUser)
	router.PUT("/user/:user_id", user.UpdateUser)
	router.PATCH("/user/:user_id", user.UpdateUser)
	router.DELETE("/user/:user_id", user.DeleteUser)

	router.GET("/users/search", user.FindUser)
}
