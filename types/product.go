package types

type Product struct {
	SKU  string `json: "sku"`
	Name string `json: "name"`
	Slug string `json: "slug"`
}
