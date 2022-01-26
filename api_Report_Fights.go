package fflogs

import (
	"context"

	"fflogs/structure"
)

// Gets arrays of fights and the participants in those fights. Each Fight corresponds to a single pull of a boss.
func (c *Client) ReportFights(context context.Context, opt *ReportFightsOptions) (*structure.Fights, error) {
	var resp structure.Fights
	err := c.Raw.ReportFights(context, opt, &resp)
	return &resp, err
}
