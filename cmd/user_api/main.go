package main

import (
	"bicingalert/app"
	"net/http"
)

func main() {
	db := app.GetBicingDb()
	googleAuth := app.NewGoogleAuth("http://localhost:8080/auth")

	http.HandleFunc("/login", googleAuth.LoginHandler)
	http.HandleFunc("/logout", app.LogOutHandler)
	http.HandleFunc("/auth", googleAuth.AuthHandler)
	http.Handle("/", app.UserHandler{
		Db: db,
		Users: app.UserRepository{db},
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
