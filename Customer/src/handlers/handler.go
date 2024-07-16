package handlers

import (
	"CustomerApi/Customer/src/entities"
	upRequestModel "CustomerApi/Customer/src/entities/requestModels"
	"CustomerApi/Customer/src/services"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *services.Service
	echo    *echo.Echo
	v       *validator.Validate
}

func NewHandler(service *services.Service, echo *echo.Echo) *Handler {
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())
	return &Handler{service: service, echo: echo, v: validate} //
}

func (h *Handler) InitEndpoints() {
	e := h.echo
	g := e.Group("/customers")
	g.GET("/", h.All)
	g.GET("/:id", h.GetByID) // buradaki id ile paramdan al dediğimde yazım hatası yapmışım.
	g.POST("/", h.Create)
	g.DELETE("/:id", h.Delete)
	g.POST("/update/:id", h.Update)
}

func (h *Handler) All(c echo.Context) error {

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	offsetInt, _ := strconv.ParseInt(offset, 10, 64)
	customers, err := h.service.GetAll(limitInt, offsetInt)
	if err != nil {
		fmt.Println(err, "getall handler")
		return err
	}

	return c.JSON(http.StatusOK, &customers)
}

func (h *Handler) Create(c echo.Context) error {

	customerData := entities.Customer{}
	err := c.Bind(&customerData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "Body")
	}

	Id, err := h.service.Create(customerData)
	if err != nil {
		fmt.Println(err, "- Create Handler")
		return c.JSON(http.StatusInternalServerError, "Error")
	}

	return c.JSON(http.StatusCreated, Id)
}

func (h *Handler) GetByID(c echo.Context) error {

	id := c.Param("id")

	customer, err := h.service.GetByID(id)
	if err != nil {
		fmt.Println(err, "- Handler: GetByID")
		return c.JSON(http.StatusBadRequest, "Error getting customer")
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *Handler) Delete(c echo.Context) error {

	id := c.Param("id")

	success, _ := h.service.Delete(id)
	if !success {
		fmt.Println(success, "- Handler: GetByID")
		return c.JSON(http.StatusInternalServerError, "Error deleting customer")
	}

	return c.JSON(http.StatusOK, "Customer deleted successfully")
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")

	updateData := upRequestModel.UpdateRequestModel{}

	if err := c.Bind(&updateData); err != nil {
		fmt.Println(err, "handler update bind")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	//todo : validasyon yapılacak validator kullanarak.

	if err := h.v.Struct(updateData); err != nil {
		fmt.Println(err, "validate struct error")
		return c.JSON(http.StatusBadRequest, "validation error")
	}

	//	if len(updateData.Phone) != 10 {
	//		return fmt.Errorf("phone must 10 character")
	//	}

	err := h.service.Update(id, &updateData)
	if err != nil {
		fmt.Println(err, "handler update")
		return c.JSON(http.StatusInternalServerError, "Error updating customer")
	}

	return c.JSON(http.StatusOK, "Customer updated successfully")
}
