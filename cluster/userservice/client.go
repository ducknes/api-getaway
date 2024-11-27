package userservice

import (
	"api-getaway/cluster"
	"context"
	"net/http"
	"strconv"
)

const (
	_getUserUri    = "/user"
	_getUsersUri   = "/users"
	_saveUserUri   = "/user"
	_updateUserUri = "/user"
	_deleteUserUri = "/user"
)

type Client struct {
	httpClient cluster.BaseClient
}

func NewClient(client cluster.BaseClient) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) GetUser(ctx context.Context, userId string) (user User, err error) {
	params := map[string]string{
		"id": userId,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodGet, _getUserUri, nil, params)
	if err != nil {
		return
	}

	return user, c.httpClient.Do(httpRequest, body, &user)
}

func (c *Client) GetUsers(ctx context.Context, limit int, cursor string) (users []User, err error) {
	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"cursor": cursor,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodGet, _getUsersUri, nil, params)
	if err != nil {
		return
	}

	return users, c.httpClient.Do(httpRequest, body, &users)
}

func (c *Client) SaveUser(ctx context.Context, user User) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _saveUserUri, user, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}

func (c *Client) UpdateUser(ctx context.Context, user User) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _updateUserUri, user, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}

func (c *Client) DeleteUser(ctx context.Context, userId string) error {
	params := map[string]string{
		"id": userId,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _deleteUserUri, nil, params)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}
