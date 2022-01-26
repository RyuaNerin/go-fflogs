package fflogs

import (
	"context"
)

type RankingEncounterOptions struct {
	EncounterID          int     `path:"encounterID"`           // The encounter to collect rankings for. Encounter IDs can be obtained using a /zones reuest.
	Metric               *Metric `query:"metric"`               // The metric to query for. Valid fight metrics are 'speed', 'execution' and 'feats'. Valid character metrics are 'dps', 'hps', 'bossdps, 'tankhps', or 'playerspeed'. For WoW only, 'krsi' can be used for tank survivability ranks and 'progress' can be used for guild progressinfo.
	Size                 *string `query:"size"`                 // The raid size to query for. This is only valid for fixed size raids. Raids with flexible sizing must omit this paraeter.
	Difficulty           *string `query:"difficulty"`           // The difficulty setting to query for. Valid difficulty settings are 1 = LFR, 2 = Flex, 3 = Normal, 4 = Heroic, 5 = Mythic, 10 = Challenge Mode, 100 = WildStar/FF. Can be omitted for encounters with only one difficulty seting.
	Partition            *int    `query:"partition"`            // The partition group to query for. Most zones have only one partition, and this can be omitted. Hellfire Citadel has two partitions (1 for original, 2 for pre-patch). Highmaul and BRF have two partitions (1 for US/EU, 2 for sia).
	Class                *int    `query:"class"`                // The class to query for if a character metric is specified. Valid class IDs can be obtained from a /classes API request. Optonal.
	Spec                 *int    `query:"spec"`                 // The spec to query for if a character metric is specified. Valid spec IDs can be obtained from a /classes API request. Optonal.
	Bracket              *int    `query:"bracket"`              // The bracket to query for. If omitted or if a value of 0 is specified, then all brackets are examined. Brackets can be obtained from a /zones API reuest.
	Server               *string `query:"server"`               // A server to filter on. If set, the region must also be specified. This is the slug field in Blizzard terminlogy.
	Region               *Region `query:"region"`               // The short name of a region to filter on (e.g., US, NAEU)
	Page                 *int    `query:"page"`                 // The page to examine, starting from 1. If the value is omitted, then 1 is assumed. For example, with a page of 2 and a limit of 300, you will be fetching rankings 30-600.
	Filter               *string `query:"filter"`               // A search filter string, limiting the search to specific classes, specs, fight durations, raid sizes, etc. The format should match the string used on the public rankings ages.
	IncludeCombatantInfo *bool   `query:"includeCombatantInfo"` // Whether or not to include combatant info like gear and talents. Optional. Defaults to alse.
}

// Gets an object that contains a total count and an array of EncounterRanking objects and a total number of rankings for that encounter. Each EncounterRanking corresponds to a single character or guild/team.
func (c *Raw) RankingEncounter(context context.Context, opt *RankingEncounterOptions, resp interface{}) error {
	return c.call(
		context,
		"/rankings/encounter/{encounterID}",
		opt,
		resp,
	)
}
