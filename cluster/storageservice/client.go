package storageservice

import (
	"api-getaway/cluster"
	"context"
	"net/http"
	"strconv"
)

const (
	_getProductUri     = "/product"
	_getProductsUri    = "/products"
	_saveProductsUri   = "/products"
	_updateProductsUri = "/products"
	_deleteProductsUri = "/products"
)

type Client struct {
	httpClient cluster.BaseClient
}

func NewClient(client cluster.BaseClient) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) GetProduct(ctx context.Context, productId string) (product Product, err error) {
	params := map[string]string{
		"product_id": productId,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodGet, _getProductUri, nil, params)
	if err != nil {
		return Product{}, err
	}

	return product, c.httpClient.Do(httpRequest, body, &product)
}

func (c *Client) GetProducts(ctx context.Context, limit int, cursor string) (products []Product, err error) {
	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"cursor": cursor,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodGet, _getProductsUri, nil, params)
	if err != nil {
		return products, err
	}

	return products, c.httpClient.Do(httpRequest, body, &products)
}

func (c *Client) SaveProducts(ctx context.Context, products []Product) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _saveProductsUri, products, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}

func (c *Client) UpdateProducts(ctx context.Context, products []Product) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _updateProductsUri, products, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}

func (c *Client) DeleteProducts(ctx context.Context, productIds []string) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _deleteProductsUri, productIds, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}
