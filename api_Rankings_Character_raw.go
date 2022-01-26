package fflogs

import (
	"context"
)

type RankingCharacterOptions struct {
	CharacterName        string     `path:"characterName"`         // The name of the character to collect rankings for.
	ServerName           string     `path:"serverName"`            // The server that the character is found on. For World of Warcraft this is the 'slug' field returned from their realm status API.
	ServerRegion         Region     `path:"serverRegion"`          // The short region name for the server on which the character is located: US, EU, KR, TW, CN.
	Zone                 *int       `query:"zone"`                 // The zone to fetch rankings for. If omitted, the latest open raid zone is used.
	Encounter            *int       `query:"encounter"`            // An encounter within the zone to fetch rankings for. If omitted, all encounters in the zone will be checked.
	Metric               *Metric    `query:"metric"`               // The metric to query for. Valid character metrics are 'dps', 'hps', 'bossdps, 'tankhps', or 'playerspeed'. For WoW only, 'krsi' can be used for tank survivability ranks.
	Bracket              *int       `query:"bracket"`              // The bracket to query for. If omitted or if a value of 0 is specified, then all brackets are examined. Brackets can be obtained from a /zones API request.
	Partition            *int       `query:"partition"`            // The partition group to query for. Most zones have only one partition, and this can be omitted. Hellfire Citadel has two partitions (1 for original, 2 for pre-patch). Highmaul and BRF have two partitions (1 for US/EU, 2 for Asia).	query	integer
	Timeframe            *Timeframe `query:"timeframe"`            // Whether to compare against today's rankings or to return historical information (where the rank was back when it was earned. The accepted values are 'today' and 'historical', with the default being 'today'.	query	string
	IncludeCombatantInfo *bool      `query:"includeCombatantInfo"` // Whether or not to include combatant info like gear and talents. Optional. Defaults to false.
}

// Gets an array of CharacterRanking objects. Each CharacterRanking corresponds to a single rank on a fight for the specified character.
func (c *Raw) RankingCharacter(context context.Context, opt *RankingCharacterOptions, resp interface{}) error {
	return c.call(
		context,
		"/rankings/character/{characterName}/{serverName}/{serverRegion}",
		opt,
		resp,
	)
}
