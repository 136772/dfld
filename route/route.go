package route

import (
	// "github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"Dfld/controller"
	"Dfld/djwt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	// "sync"
	"time"
)

func Route() {
	adminapi := rest.NewApi()
	adminapi.Use(rest.DefaultDevStack...)
	jwtMiddleware := &djwt.JWTMiddleware{
		Key:           []byte("DfldSKey@defengLida"),
		Realm:         "DeFengLiDa",
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: controller.Login,
	}

	adminapi.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/login"
		},
		IfTrue: jwtMiddleware,
	})

	adminRouter, err := rest.MakeRouter(
		rest.Post("/login", jwtMiddleware.LoginHandler),
		rest.Get("/authTest", controller.HandleAuth),
		rest.Get("/refreshToken", jwtMiddleware.RefreshHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	adminapi.SetApp(adminRouter)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", controller.Index),
		rest.Get("/verify/:md5", controller.CheckVerify),
		rest.Post("/verify", controller.PostVerify),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/adminapi/", http.StripPrefix("/adminapi", adminapi.MakeHandler()))
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
