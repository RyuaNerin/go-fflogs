package fflogs

import (
	"context"
)

type ParsesCharacterOptions struct {
	CharacterName        string     `path:"characterName"`         // The name of the character to collect rankings for.
	ServerName           string     `path:"serverName"`            // The server that the character is found on. For World of Warcraft this is the 'slug' field returned from their realm status API.
	ServerRegion         Region     `path:"serverRegion"`          // The short region name for the server on which the character is located: US, EU, KR, TW, CN.
	Zone                 *int       `query:"zone"`                 // The zone to fetch rankings for. If omitted, the latest open raid zone is used.
	Encounter            *int       `query:"encounter"`            // An encounter within the zone to fetch parses for. If omitted, all encounters in the zone will be checked. If a valid encounter is specified, then historical graph data for the encounter will also be included in the results.
	Metric               *Metric    `query:"metric"`               // The metric to query for. Valid character metrics are 'dps', 'hps', 'bossdps, 'tankhps', or 'playerspeed'. For WoW only, 'krsi' can be used for tank survivability ranks.
	Bracket              *int       `query:"bracket"`              // The bracket to query for. If omitted or if a value of 0 is specified, then all brackets are examined. Brackets can be obtained from a /zones API request. The special value -1 can be used to obtain scores by bracket %, i.e., to automatically look only at the bracket each parse falls in.
	Compare              *int       `query:"compare"`              // Optional. Whether or not to compare against rankings (0) when computing performance percentiles or to compare against statistics (1). A value of 0 is assumed if omitted.
	Partition            *int       `query:"partition"`            // The partition group to query for. Most zones have only one partition, and this can be omitted. Hellfire Citadel has two partitions (1 for original, 2 for pre-patch). Highmaul and BRF have two partitions (1 for US/EU, 2 for Asia).
	Timeframe            *Timeframe `query:"timeframe"`            // Whether to compare against today's rankings or to return historical information (where the rank was back when it was earned. The accepted values are 'today' and 'historical', with the default being 'today'.
	IncludeCombatantInfo *bool      `query:"includeCombatantInfo"` // Whether or not to include combatant info like gear and talents. Optional. Defaults to false.
}

// Obtains all parses for a character in the zone across all specs. Every parse is included and not just rankings.
func (c *Raw) ParsesCharacter(context context.Context, opt *ParsesCharacterOptions, resp interface{}) error {
	return c.call(
		context,
		"/parses/character/{characterName}/{serverName}/{serverRegion}",
		opt,
		resp,
	)
}
