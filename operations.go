package yadisk

import (
	"context"
	"net/http"
)

const operationsAPIPath = "/v1/disk/operations"

func (c *Client) OperationStatus(ctx context.Context, operationID string) (*Operation, error) {
	resp := &Operation{}

	err := c.doRequest(ctx, http.MethodGet, operationsAPIPath+"/"+operationID, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
