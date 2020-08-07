package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(app.home)).Methods("GET")
	r.HandleFunc("/authenticate", http.HandlerFunc(app.authenticate)).Methods("POST")
	r.Handle("/dropdown/{name}", app.validateToken(http.HandlerFunc(app.dropdownHandler))).Methods("GET")
	r.Handle("/dropdown/condition/{name}/{where}/{value}", app.validateToken(http.HandlerFunc(app.dropdownConditionHandler))).Methods("GET")
	r.Handle("/model/create", app.validateToken(http.HandlerFunc(app.createModel))).Methods("POST")
	r.Handle("/model/all", app.validateToken(http.HandlerFunc(app.allItems))).Methods("GET")
	r.Handle("/warehouse/create", app.validateToken(http.HandlerFunc(app.createWarehouse))).Methods("POST")
	r.Handle("/warehouse/all", app.validateToken(http.HandlerFunc(app.allWarehouses))).Methods("GET")
	r.Handle("/warehouse/stock/{id}", app.validateToken(http.HandlerFunc(app.warehouseStock))).Methods("GET")
	r.Handle("/search", app.validateToken(http.HandlerFunc(app.search))).Methods("GET")
	r.Handle("/agewise", app.validateToken(http.HandlerFunc(app.ageWise))).Methods("GET")

	r.Handle("/transactions/goodsin", app.validateToken(http.HandlerFunc(app.goodsIn))).Methods("POST")
	r.Handle("/transactions/movement", app.validateToken(http.HandlerFunc(app.transaction))).Methods("POST")
	r.Handle("/getSecondaryNumberModelName", app.validateToken(http.HandlerFunc(app.secNumberModel))).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
}
