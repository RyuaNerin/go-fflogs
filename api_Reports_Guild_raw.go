package fflogs

import (
	"context"
)

type ReportsGuildOptions struct {
	GuildName    string `path:"guildName"`    // The name of the guild to collect reports for.
	ServerName   string `path:"serverName"`   // The server that the guild is found on. For World of Warcraft this is the 'slug' field returned from their realm status API.
	ServerRegion Region `path:"serverRegion"` // The short region name for the server on which the guild is located: US, EU, KR, TW, CN.
	Start        *int   `query:"start"`       // An optional start time. This is a UNIX timestamp but with millisecond precision. If omitted, 0 is assumed.
	End          *int   `query:"end"`         // An optional end time. This is a UNIX timestamp but with millisecond precision. If omitted, the current time is assumed.
}

// Gets an object that contains a total count and an array of EncounterRanking objects and a total number of rankings for that encounter. Each EncounterRanking corresponds to a single character or guild/team.
func (c *Raw) ReportsGuild(context context.Context, opt *ReportsGuildOptions, resp interface{}) error {
	return c.call(
		context,
		"/reports/guild/{guildName}/{serverName}/{serverRegion}",
		opt,
		resp,
	)
}
