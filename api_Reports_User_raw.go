package fflogs

import (
	"context"
)

type ReportsUserOptions struct {
	UserName string `path:"userName"` // The name of the user to collect reports for.
	Start    *int   `query:"start"`   // An optional start time. This is a UNIX timestamp but with millisecond precision. If omitted, 0 is assumed.
	End      *int   `query:"end"`     // An optional end time. This is a UNIX timestamp but with millisecond precision. If omitted, the current time is assumed.
}

// Gets an array of Report objects. Each Report corresponds to a single calendar report for the specified user's personal logs.
func (c *Raw) ReportsUser(context context.Context, opt *ReportsUserOptions, resp interface{}) error {
	return c.call(
		context,
		"/reports/user/{userName}",
		opt,
		resp,
	)
}
