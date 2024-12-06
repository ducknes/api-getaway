package authservice

import (
	"api-getaway/cluster"
	"context"
	"net/http"
)

const (
	_loginUri   = "login"
	_logoutUri  = "logout"
	_refreshUri = "refresh-token"
	_signUpUri  = "sign-up"

	_refreshTokenCookie = "refresh_token"
)

type Client struct {
	httpClient cluster.BaseClient
}

func NewClient(client cluster.BaseClient) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) Login(ctx context.Context, loginData LoginUser) (response LoginResponse, err error) {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _loginUri, loginData, nil)
	if err != nil {
		return LoginResponse{}, err
	}

	return response, c.httpClient.Do(httpRequest, body, &response)
}

func (c *Client) Logout(ctx context.Context, refreshToken string) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _logoutUri, nil, nil)
	if err != nil {
		return err
	}

	cookie := map[string]string{
		_refreshTokenCookie: refreshToken,
	}

	c.httpClient.SetCookie(httpRequest, cookie)

	return c.httpClient.Do(httpRequest, body, nil)
}

func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (response LoginResponse, err error) {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _refreshUri, nil, nil)
	if err != nil {
		return LoginResponse{}, err
	}

	cookie := map[string]string{
		_refreshTokenCookie: refreshToken,
	}

	c.httpClient.SetCookie(httpRequest, cookie)

	return response, c.httpClient.Do(httpRequest, body, &response)
}

func (c *Client) SignUp(ctx context.Context, registrationData LoginUser) (response LoginResponse, err error) {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _signUpUri, registrationData, nil)
	if err != nil {
		return LoginResponse{}, err
	}

	return response, c.httpClient.Do(httpRequest, body, &response)
}
