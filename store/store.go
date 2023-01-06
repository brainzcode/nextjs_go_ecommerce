package store

import (
	"context"

	"github.com/brainzcode/nextjs_go_ecommerce/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetByID(string) (context.Context, *types.Product, error)
}
