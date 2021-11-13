package main

import (
	"bitbucket/connect"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type majooApp struct {
	connect connect.Connect
	router  *mux.Router
}

func (app *majooApp) run() {
	h := app.connect.ApiServer([]string{})
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRoutes()
	err := http.ListenAndServe(h, app.router)
	if err != nil {
		log.Fatalln(err)
	}

}

func NewBriApp() *majooApp {
	r := mux.NewRouter()
	var appConnect = connect.NewConnect()

	return &majooApp{
		connect: appConnect,
		router:  r,
	}
}

func main() {
	NewBriApp().run()

}
