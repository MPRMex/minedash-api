package minedash

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"

	"encoding/json"
)

const baseAPI = "/api/v1/"

type MineDashServer struct {
	base        string
	credentials string
	client      *http.Client
}

func NewMineDashServer(server, username, password string) *MineDashServer {
	base := server + baseAPI
	credentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return &MineDashServer{base, credentials, &http.Client{}}
}

func (mds *MineDashServer) GetEntities() (*[]Entity, error) {
	res, err := mds.makeRequest("GET", "entity", nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	result := &EntityAllResponse{}
	err = parseJsonBody(res, &result)
	if err != nil {
		return nil, err
	}
	return &result.Entities, nil
}

func (mds *MineDashServer) GetEntity(sysid int) (*Entity, error) {
	res, err := mds.makeRequest("GET", fmt.Sprintf("entity/%d", sysid), nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	result := &Entity{}
	err = parseJsonBody(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (mds *MineDashServer) DeleteEntity(sysid int) error {
	res, _ := mds.makeRequest("DELETE", fmt.Sprintf("entity/%d", sysid), nil)
	if res.StatusCode == http.StatusOK {
		return nil
	}
	result := &DeletedEntityResponse{}
	err := parseJsonBody(res, &result)
	if err != nil {
		return err
	}
	return errors.New(result.ErrorDesc)
}

func (mds *MineDashServer) NewEntity(entity *Entity) (*Entity, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(data)
	res, err := mds.makeRequest("POST", "entity", r)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		return nil, errors.New(res.Status)
	}
	result := &Entity{}
	err = parseJsonBody(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (mds *MineDashServer) createRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, mds.base+url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+mds.credentials)
	return req, nil
}

func (mds *MineDashServer) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := mds.createRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return mds.client.Do(req)
}
