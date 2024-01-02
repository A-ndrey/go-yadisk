package yadisk

import (
	"context"
	"net/http"
)

const publicResourcesAPIPath = "/v1/disk/public/resources"

type PublicResourceMetaInfoParams struct {
	PublicKey   string   `param:"public_key,required"`
	Fields      []string `param:"fields"`
	Limit       int      `param:"limit"`
	Offset      int      `param:"offset"`
	Path        string   `param:"path"`
	PreviewCrop bool     `param:"preview_crop"`
	PreviewSize string   `param:"preview_size"`
	Sort        string   `param:"sort"`
}

func (c *Client) PublicResourceMetaInfo(ctx context.Context, params PublicResourceMetaInfoParams) (*Resource, error) { // todo use correct response
	resp := &Resource{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, publicResourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type DownloadPublicResourceLinkParams struct {
	PublicKey string   `param:"public_key,required"`
	Fields    []string `param:"fields"`
	Path      string   `param:"path"`
}

func (c *Client) DownloadPublicResourceLink(ctx context.Context, params DownloadPublicResourceLinkParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, publicResourcesAPIPath+"/download", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type SaveToDiskParams struct {
	PublicKey string   `param:"public_key,required"`
	Fields    []string `param:"fields"`
	Async     bool     `param:"force_async"`
	Name      string   `param:"name"`
	Path      string   `param:"path"`
	SavePath  string   `param:"save_path"` // default 'Downloads'
}

func (c *Client) SaveToDisk(ctx context.Context, params SaveToDiskParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPost, publicResourcesAPIPath+"/save-to-disk", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
