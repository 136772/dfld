package controller

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func Index(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"data": "Ahahaha"})
}
