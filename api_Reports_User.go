package fflogs

import (
	"context"

	"fflogs/structure"
)

// Gets an array of Report objects. Each Report corresponds to a single calendar report for the specified user's personal logs.
func (c *Client) ReportsUser(context context.Context, opt *ReportsUserOptions) ([]structure.Report, error) {
	var resp []structure.Report
	err := c.Raw.ReportsUser(context, opt, &resp)
	return resp, err
}
