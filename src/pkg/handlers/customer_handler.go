package handlers

import (
	"net/http"
	"strconv"

	"github.com/bburaksseyhan/ctmapp/src/cmd/utils"
	"github.com/bburaksseyhan/ctmapp/src/pkg/entities"
	"github.com/bburaksseyhan/ctmapp/src/pkg/models"
	"github.com/bburaksseyhan/ctmapp/src/pkg/repository"
	"github.com/labstack/echo"
)

// CustomerHandler interface
type CustomerHandler interface {
	List(echo.Context) error
	Add(echo.Context) error
	Delete(echo.Context) error
	Get(echo.Context) error

	Health(echo.Context) error
}

// customerHandler hold the repo and settings
type customerHandler struct {
	repo     repository.CustomerRepository
	settings utils.DbSettings
}

// NewCustomerHandler create a new customerHandler
func NewCustomerHandler(customerRepository repository.CustomerRepository, dbSettings *utils.DbSettings) CustomerHandler {
	return &customerHandler{repo: customerRepository, settings: *dbSettings}
}

func (h *customerHandler) List(c echo.Context) error {

	context := c.Request().Context()
	customerModels := []models.Customer{}

	customerEntities, err := h.repo.List(context, h.settings.Timeout)
	if err != nil {
		return err
	}

	for _, c := range customerEntities {
		customerModels = append(customerModels, models.Customer(c))
	}

	return c.JSON(http.StatusOK, customerModels)
}

func (h *customerHandler) Add(c echo.Context) error {

	context := c.Request().Context()

	customerModel := new(models.Customer)

	if err := c.Bind(customerModel); err != nil {
		return err
	}

	data, err := h.repo.Add(entities.CustomerEntity(*customerModel), context, h.settings.Timeout)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *customerHandler) Delete(c echo.Context) error {

	context := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.repo.Delete(id, context, h.settings.Timeout)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, data)
}

func (h *customerHandler) Get(c echo.Context) error {

	context := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.repo.Get(id, context, h.settings.Timeout)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, data)
}

func (h *customerHandler) Health(c echo.Context) error {

	return c.JSON(http.StatusOK, "healty")
}
