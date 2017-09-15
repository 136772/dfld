package controller

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
)

type VerityStruct struct {
	Md5  string
	Code string
	Name string
}

var lock = sync.RWMutex{}

var store = map[string]*VerityStruct{}

func PostVerify(w rest.ResponseWriter, r *rest.Request) {
	verityStruct := VerityStruct{}
	err := r.DecodeJsonPayload(&verityStruct)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if verityStruct.Code == "" {
		rest.Error(w, "country code required", 400)
		return
	}
	if verityStruct.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}

	lock.Lock()
	store[verityStruct.Md5] = &verityStruct
	lock.Unlock()
	w.WriteJson(&verityStruct)

}

func CheckVerify(w rest.ResponseWriter, r *rest.Request) {
	md5 := r.PathParam("md5")
	log.Println(md5)
	lock.RLock()
	var verityStruct *VerityStruct
	if store[md5] != nil {
		verityStruct = &VerityStruct{}
		*verityStruct = *store[md5]
	}
	lock.RUnlock()
	if verityStruct == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(verityStruct)
}
