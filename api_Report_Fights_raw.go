package fflogs

import (
	"context"
)

type ReportFightsOptions struct {
	Code      string `path:"code"`       // The specific report to collect fights and participants for.
	Translate *bool  `query:"translate"` // An optional flag indicating that the results should be translated into the language of the host (e.g., cn.warcraftlogs.com would get Chinese results).
}

// Gets arrays of fights and the participants in those fights. Each Fight corresponds to a single pull of a boss.
func (c *Raw) ReportFights(context context.Context, opt *ReportFightsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/fights/{code}",
		opt,
		resp,
	)
}
