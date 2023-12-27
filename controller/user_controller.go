package controller

import (
	// "html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/model/request"
	"github.com/jkim7113/centinal/model/response"
	"github.com/jkim7113/centinal/service"
	"github.com/jkim7113/centinal/util"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	UserCreateRequest := request.UserCreateRequest{}
	util.DecodeRequestBody(r, &UserCreateRequest)

	controller.UserService.Create(r.Context(), UserCreateRequest)
	HTTPResponse := response.HTTPResponse{
		Code:   201,
		Status: "Created",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")
	UserUpdateRequest := request.UserUpdateRequest{}
	UserUpdateRequest.UUID = UUID
	util.DecodeRequestBody(r, &UserUpdateRequest)

	controller.UserService.Update(r.Context(), UserUpdateRequest)
	HTTPResponse := response.HTTPResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")

	controller.UserService.Delete(r.Context(), UUID)
	HTTPResponse := response.HTTPResponse{
		Code:   204,
		Status: "No Content",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")

	result := controller.UserService.FindById(r.Context(), UUID)
	HTTPResponse := response.HTTPResponse{
		Code:   204,
		Status: "No Content",
		Data:   result,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}
