package gohetz

import (
	"./models"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

const serversUrl = "/servers/"

func (c *Client) GetAllServers() (*models.Servers, error) {

	req, err := c.NewRequest(context.Background(), http.MethodGet, fmt.Sprintf(serversUrl), nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	servers, err := models.UnmarshalServers(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &servers, nil
}

func (c *Client) GetServerBy(ID float64) (*models.Server, error) {
	req, err := c.NewRequest(context.Background(), http.MethodGet, fmt.Sprintf("%s%.f", serversUrl, ID), nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	servers, err := models.UnmarshalServer(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &servers, nil
}

func (c *Client) CreateServerWith(request *models.ServerCreateRequest) (*models.ServerCreateResponse, error) {
	s, err := request.Marshal()
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(context.Background(), http.MethodPost, fmt.Sprintf(serversUrl), bytes.NewReader(s))
	if err != nil {
		return nil, err
	}

	raw, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	server, err := models.UnmarshalCreateServerResponse(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (c *Client) UpdateServerBy(ID float64, request *models.ServerUpdateRequest) (*models.ServerUpdateResponse, error) {
	s, err := request.Marshal()
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(context.Background(), http.MethodPut, fmt.Sprintf("%s%.f", serversUrl, ID), bytes.NewReader(s))
	if err != nil {
		return nil, err
	}

	raw, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	server, err := models.UnmarshalServerUpdateResponse(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (c *Client) DeleteServerBy(ID float64) (*models.ServerDeleteResponse, error) {
	req, err := c.NewRequest(context.Background(), http.MethodDelete, fmt.Sprintf("%s%.f", serversUrl, ID), nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	response, err := models.UnmarshalServerDeleteResponse(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
