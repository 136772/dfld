package controller

import (
	"Dfld/utils"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
)

func Login(username, password string) bool {
	if username == "admin" && password == "admin" {
		pwd := utils.PasswordEncode(password)
		log.Println(pwd)
		return true
	} else if username == "guest" && password == "guest" {
		return true
	}
	return false
}

func HandleAuth(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}
