package api

import (
	"encoding/json"
	"net/http"

	"github.com/brainzcode/nextjs_go_ecommerce/types"

	"github.com/anthdm/weavebox"
)

type CreateProductRequest struct {
	SKU  string `json: "sku"`
	Name string `json: "name"`
}

type ProductStorer interface {
	Insert(*types.Product) error
	GetByID(string) (*types.Product, error)
}

type ProductHandler struct {
	store ProductStorer
}

func NewProductHandler(pStore ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct(c *weavebox.Context) error {
	productReq := &CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, productReq)
}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error {

	return c.JSON(http.StatusOK, &types.Product{SKU: "SHOE-1011"})
}
