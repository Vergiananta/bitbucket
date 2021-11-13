package handler

import (
	"bitbucket/controllers"
	"bitbucket/middlewares"
	"bitbucket/models"
	"bitbucket/models/dto"
	"bitbucket/usecase"
	"bitbucket/utils/httpParse"
	"bitbucket/utils/httpResponse"
	"bitbucket/utils/status"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service usecase.IUserUseCase
}

func (s *UserController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := s.router.PathPrefix("/users").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	u.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	u.HandleFunc("", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.EditUser))).Methods("PUT")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.FindById))).Methods("GET")
}


func NewUserController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IUserUseCase) controllers.IDelivery  {
	return &UserController{
		router, parse, responder, service,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	newUser, err := u.service.Register(&user)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), newUser)
}

func (u *UserController) EditUser(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userEdit, err := u.service.UpdateUser(&user)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}

	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userEdit)
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var login dto.LoginRequest
	err = json.Unmarshal(body, &login)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	loginUser, err := u.service.LoginUser(&login)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, http.StatusOK, status.StatusText(status.Success), loginUser)

}

func (u *UserController) FindById(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	userId, err := u.service.FindUserById(param["id"])
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, status.Success, status.StatusText(status.Success), userId)
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	err := u.service.DeleteUser(param["id"])
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, status.Success, status.StatusText(status.Success), "Data has been deleted")
}
