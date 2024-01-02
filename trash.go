package yadisk

import (
	"context"
	"net/http"
)

const trashResourcesAPIPath = "/v1/disk/trash/resources"

type CleanTrashParams struct {
	Fields []string `param:"fields"`
	Async  bool     `param:"force_async"`
	Path   string   `param:"path"`
}

func (c *Client) CleanTrash(ctx context.Context, params CleanTrashParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodDelete, trashResourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type TrashContentParams struct {
	Path        string   `param:"path,required"`
	Fields      []string `param:"fields"`
	Limit       int      `param:"limit"`
	Offset      int      `param:"offset"`
	PreviewCrop bool     `param:"preview_crop"`
	PreviewSize string   `param:"preview_size"`
	Sort        string   `param:"sort"`
}

func (c *Client) TrashContent(ctx context.Context, params TrashContentParams) (*Resource, error) { // todo use correct response
	resp := &Resource{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, trashResourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type RestoreResourceParams struct {
	Path      string   `param:"path,required"`
	Fields    []string `param:"fields"`
	Async     bool     `param:"force_async"`
	Name      string   `param:"name"`
	Overwrite bool     `param:"overwrite"`
}

func (c *Client) RestoreResource(ctx context.Context, params RestoreResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPut, trashResourcesAPIPath+"/restore", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
