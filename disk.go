package yadisk

import (
	"context"
	"net/http"
)

const diskAPIPath = "/v1/disk"

type DiskMetaInfoParams struct {
	Fields []string `param:"fields"`
}

func (c *Client) DiskMetaInfo(ctx context.Context, params DiskMetaInfoParams) (*Disk, error) {
	resp := &Disk{}

	query, err := queryFromParams(&params)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(ctx, http.MethodGet, diskAPIPath, query, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
