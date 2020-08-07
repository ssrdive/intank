package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/ssrdive/basara/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// user := app.extractUser(r)

	if app.runtimeEnv == "dev" {
		fmt.Fprintf(w, "It works! [dev]")
	} else {
		fmt.Fprintf(w, "It works!")
	}
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	u, err := app.user.Get(username, password)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = u.Username
	claims["name"] = u.Name
	claims["exp"] = time.Now().Add(time.Minute * 180).Unix()

	ts, err := token.SignedString(app.secret)
	if err != nil {
		app.serverError(w, err)
		return
	}

	user := models.UserResponse{u.ID, u.Username, u.Name, u.Type, ts}
	js, err := json.Marshal(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) dropdownHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	items, err := app.dropdown.Get(name)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)

}

func (app *application) dropdownConditionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	where := vars["where"]
	value := vars["value"]
	if name == "" || where == "" || value == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	items, err := app.dropdown.ConditionGet(name, where, value)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)

}

func (app *application) createWarehouse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	requiredParams := []string{"warehouse_type_id", "name", "address", "contact"}
	optionalParams := []string{}
	for _, param := range requiredParams {
		if v := r.PostForm.Get(param); v == "" {
			fmt.Println(param)
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	id, err := app.warehouse.Create(requiredParams, optionalParams, r.PostForm)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%d", id)

}

func (app *application) createModel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	requiredParams := []string{"name", "country", "primary_name", "secondary_name"}
	optionalParams := []string{}
	for _, param := range requiredParams {
		if v := r.PostForm.Get(param); v == "" {
			fmt.Println(param)
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	id, err := app.model.Create(requiredParams, optionalParams, r.PostForm)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%d", id)

}

func (app *application) allItems(w http.ResponseWriter, r *http.Request) {
	results, err := app.model.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) ageWise(w http.ResponseWriter, r *http.Request) {
	model := r.URL.Query().Get("model")
	age := r.URL.Query().Get("age")

	m, err := strconv.Atoi(model)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	a, err := strconv.Atoi(age)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	results, err := app.warehouse.Agewise(m, a)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	results, err := app.warehouse.Search(search)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) warehouseStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	results, err := app.warehouse.Stock(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) allWarehouses(w http.ResponseWriter, r *http.Request) {
	results, err := app.warehouse.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) transaction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	requiredParams := []string{"warehouse_id", "from_warehouse_id", "date", "document_type", "goods"}
	for _, param := range requiredParams {
		if v := r.PostForm.Get(param); v == "" {
			fmt.Println(param)
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	id, err := app.warehouse.Movement(r.PostForm)

	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", id)
}

func (app *application) secNumberModel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	dt, err := app.warehouse.SecNumberModel(r.PostForm.Get("primaryNumber"))

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dt)
}

func (app *application) goodsIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	requiredParams := []string{"warehouse_id", "from_warehouse_id", "date", "goods"}
	for _, param := range requiredParams {
		if v := r.PostForm.Get(param); v == "" {
			fmt.Println(param)
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	id, err := app.warehouse.GoodsIn(r.PostForm)

	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", id)
}
