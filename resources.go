package yadisk

import (
	"context"
	"encoding/json"
	"net/http"
)

const resourcesAPIPath = "/v1/disk/resources"

type DeleteResourceParams struct {
	Path        string   `param:"path,required"`
	Fields      []string `param:"fields"`
	Async       bool     `param:"force_async"`
	MD5         string   `param:"md5"`
	Permanently bool     `param:"permanently"`
}

func (c *Client) DeleteResource(ctx context.Context, params DeleteResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodDelete, resourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ResourceMetaInfoParams struct {
	Path        string   `param:"path,required"`
	Fields      []string `param:"fields"`
	Limit       int      `param:"limit"`
	Offset      int      `param:"offset"`
	PreviewCrop bool     `param:"preview_crop"`
	PreviewSize string   `param:"preview_size"`
	Sort        string   `param:"sort"`
}

func (c *Client) ResourceMetaInfo(ctx context.Context, params ResourceMetaInfoParams) (*Resource, error) {
	resp := &Resource{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type UpdateResourceUserDataParams struct {
	Path   string   `param:"path,required"`
	Fields []string `param:"fields"`
	Body   struct {
		CustomProperties string `json:"custom_properties"`
	}
}

func (c *Client) UpdateResourceUserData(ctx context.Context, params UpdateResourceUserDataParams) (*Resource, error) {
	resp := &Resource{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(params.Body)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPatch, resourcesAPIPath, query, string(body), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type CreateDirectoryParams struct {
	Path   string   `param:"path,required"`
	Fields []string `param:"fields"`
}

func (c *Client) CreateDirectory(ctx context.Context, params CreateDirectoryParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPut, resourcesAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type CopyResourceParams struct {
	FromPath  string   `param:"from,required"`
	ToPath    string   `param:"path,required"`
	Fields    []string `param:"fields"`
	Async     bool     `param:"force_async"`
	Overwrite bool     `param:"overwrite"`
}

func (c *Client) CopyResource(ctx context.Context, params CopyResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPost, resourcesAPIPath+"/copy", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type DownloadFileLinkParams struct {
	Path   string   `param:"path,required"`
	Fields []string `param:"fields"`
}

func (c *Client) DownloadFileLink(ctx context.Context, params DownloadFileLinkParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath+"/download", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ListFilesParams struct {
	Fields      []string `param:"fields"`
	Limit       int      `param:"limit"`
	Offset      int      `param:"offset"`
	MediaType   string   `param:"media_type"`
	PreviewCrop bool     `param:"preview_crop"`
	PreviewSize string   `param:"preview_size"`
	Sort        string   `param:"sort"`
}

func (c *Client) ListFiles(ctx context.Context, params ListFilesParams) (*FileResourceList, error) {
	resp := &FileResourceList{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath+"/files", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ListFilesLastUploadedParams struct {
	Fields      []string `param:"fields"`
	Limit       int      `param:"limit"`
	MediaType   string   `param:"media_type"`
	PreviewCrop bool     `param:"preview_crop"`
	PreviewSize string   `param:"preview_size"`
	Sort        string   `param:"sort"`
}

func (c *Client) ListFilesLastUploaded(ctx context.Context, params ListFilesLastUploadedParams) (*LastUploadedResourceList, error) {
	resp := &LastUploadedResourceList{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath+"/last-uploaded", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MoveResourceParams struct {
	FromPath  string   `param:"from,required"`
	ToPath    string   `param:"path,required"`
	Fields    []string `param:"fields"`
	Async     bool     `param:"force_async"`
	Overwrite bool     `param:"overwrite"`
}

func (c *Client) MoveResource(ctx context.Context, params MoveResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPost, resourcesAPIPath+"/move", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ListPublicResourcesParams struct {
	Fields       []string `param:"fields"`
	Limit        int      `param:"limit"`
	Offset       int      `param:"offset"`
	PreviewCrop  bool     `param:"preview_crop"`
	PreviewSize  string   `param:"preview_size"`
	ResourceType string   `param:"type"`
}

func (c *Client) ListPublicResources(ctx context.Context, params ListPublicResourcesParams) (*PublicResourceList, error) {
	resp := &PublicResourceList{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath+"/public", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type PublishResourceParams struct {
	Path   string   `param:"path,required"`
	Fields []string `param:"fields"`
}

func (c *Client) PublishResource(ctx context.Context, params PublishResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPut, resourcesAPIPath+"/publish", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type UnpublishResourceParams struct {
	Path   string   `param:"path,required"`
	Fields []string `param:"fields"`
}

func (c *Client) UnpublishResource(ctx context.Context, params UnpublishResourceParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPut, resourcesAPIPath+"/unpublish", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type UploadFileLinkParams struct {
	Path      string   `param:"path,required"`
	Fields    []string `param:"fields"`
	Overwrite bool     `param:"overwrite"`
}

func (c *Client) UploadFileLink(ctx context.Context, params UploadFileLinkParams) (*ResourceUploadLink, error) {
	resp := &ResourceUploadLink{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, resourcesAPIPath+"/upload", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type UploadFileByURLParams struct {
	Path             string   `param:"path,required"`
	URL              string   `param:"url,required"`
	DisableRedirects bool     `param:"disable_redirects"`
	Fields           []string `param:"fields"`
}

func (c *Client) UploadFileByURL(ctx context.Context, params UploadFileByURLParams) (*Link, error) {
	resp := &Link{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodPost, resourcesAPIPath+"/upload", query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
