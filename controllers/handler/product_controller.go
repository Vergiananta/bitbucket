package handler

import (
	"bitbucket/controllers"
	"bitbucket/middlewares"
	"bitbucket/models"
	"bitbucket/usecase"
	"bitbucket/utils/httpParse"
	"bitbucket/utils/httpResponse"
	"bitbucket/utils/status"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type ProductController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service usecase.IProductUsecase
}

func (p *ProductController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := p.router.PathPrefix("/product").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(p.CreateProduct))).Methods("POST")
	u.HandleFunc("", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(p.EditProduct))).Methods("PUT")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(p.DeleteProduct))).Methods("DELETE")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(p.FindProductById))).Methods("GET")
	u.HandleFunc("", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(p.PaginateProduct))).Methods("GET")
}

func NewProductController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, services usecase.IProductUsecase) controllers.IDelivery  {
	return &ProductController{
		router, parse, responder, services,
	}
}

func (u *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var product models.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	newProduct, err := u.service.SaveProduct(&product)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), newProduct)
}

func (u *ProductController) EditProduct(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var product models.Product

	err = json.Unmarshal(body, &product)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	productEdit, err := u.service.UpdateProduct(&product)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), productEdit)
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	err := p.service.DeleteProduct(param["id"])
	if err != nil {
		p.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	p.responder.Data(w, status.Success, status.StatusText(status.Success), "Data has been deleted")
}

func (p *ProductController) FindProductById(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	productId, err := p.service.FindByIdProduct(param["id"])
	if err != nil {
		p.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	p.responder.Data(w, status.Success, status.StatusText(status.Success), productId)
}

func (p *ProductController) PaginateProduct(w http.ResponseWriter, r *http.Request)  {
	param := r.URL.Query()
	products, err := p.service.FindProductWithPaginate(param.Get("page"),param.Get("size"))
	if err != nil {
		p.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	p.responder.Data(w, status.Success, status.StatusText(http.StatusOK), products)
}