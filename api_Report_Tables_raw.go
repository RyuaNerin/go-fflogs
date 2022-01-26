package fflogs

import (
	"context"
)

type ReportTablesOptions struct {
	//view               string `path:"view"`                // The type of data requested. Supported values are 'summary', 'damage-done', 'damage-taken', 'healing', 'casts', 'summons', 'buffs', 'debuffs', 'deaths', 'survivability', 'resources' and 'resources-gains'.
	Code               string  `path:"code"`                // The specific report to collect table entries for.
	Start              *int    `query:"start"`              // A start time. This is a time from the start of the report in milliseconds. If omitted, 0 is assumed.
	End                *int    `query:"end"`                // An end time. This is a time from the start of the report in milliseconds. If omitted, 0 is assumed.
	Hostility          *int    `query:"hostility"`          // An optional hostility value of 0 or 1. The default is 0. A value of 0 means to collect data for Friendlies. A value of 1 means to collect data for Enemies.
	By                 *string `query:"by"`                 // An optional parameter indicating how to group entries. They can be grouped by 'source', by 'target', or by 'ability'. This value matches WCL's default behavior if omitted. For buffs and debuffs, a value of 'source' means auras gained by the source, and a value of 'target' means auras cast by the source. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Sourceid           *int    `query:"sourceid"`           // An optional actor ID to filter to. If set, only events where the ID matches the source (or target for damage-taken) of the event will be returned. The actor's pets will also be included (unless the options field overrides).
	Sourceinstance     *int    `query:"sourceinstance"`     // An optional actor instance ID to filter to. If set, only events where the instance ID matches the source (or target for damage-taken) of the event will be returned. This is useful to look for all events involving NPC N, where N is the actor instance ID.
	Sourceclass        *string `query:"sourceclass"`        // An optional actor class to filter to. If set, only events where the source (or target for damage-taken) involves that class (e.g., Mage) will be returned.
	Targetid           *int    `query:"targetid"`           // An optional actor ID to filter to. If set, only events where the ID matches the target (or source for damage-taken) of the event will be returned. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Targetinstance     *int    `query:"targetinstance"`     // An optional actor instance ID to filter to. If set, only events where the instance ID matches the target (or source for damage-taken) of the event will be returned. This is useful to look for all events involving NPC N, where N is the actor instance ID. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Targetclass        *string `query:"targetclass"`        // An optional actor class to filter to. If set, only events where the target (or source for damage-taken) involves that class (e.g., Mage) will be returned. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	SourceAurasAbsent  *string `query:"sourceAurasAbsent"`  // A comma-separated string of aura game IDs. Only matches if the aura is absent on the source.
	TargetAurasPresent *string `query:"targetAurasPresent"` // A comma-separated string of aura game IDs. Only matches if the aura is present on the target.
	TargetAurasAbsent  *string `query:"targetAurasAbsent"`  // A comma-separated string of aura game IDs. Only matches if the aura is absent on the target.
	Abilityid          *int    `query:"abilityid"`          // An optional ability ID to filter to. If set, only events where the ability matches will be returned. Consolidated abilities (WCL only) are represented using a negative number that matches the ability ID that everything is consolidated under. For the 'deaths' view, this represents a specific killing blow. For the resources views, the abilityid is not an ability but a resource type. Valid resource types can be viewed at https://www.fflogs.com/reports/resource_types/
	Options            *int    `query:"options"`            // A set of options for what to include/exclude. These correspond to options like Include Overkill in the Damage Done pane. Complete list will be forthcoming. If omitted, appropriate defaults that match WCL's default behavior will be chosen. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Cutoff             *int    `query:"cutoff"`             // An optional death cutoff. If set, events after that number of deaths have occurred will not be examined.
	Encounter          *int    `query:"encounter"`          // An optional encounter filter. If set to a specific encounter ID, only fights involving a specific encounter will be considered. The encounter IDs match those used in rankings/statistics.
	Wipes              *int    `query:"wipes"`              // An optional wipes filter. If set to 1, only wipes will be considered.
	Difficulty         *int    `query:"difficulty"`         // An optional difficulty filter.
	Filter             *string `query:"filter"`             // An optional filter written in WCL's expression language. Events must match the filter to be included.
	Translate          *bool   `query:"translate"`          // An optional flag indicating that the results should be translated into the language of the host (e.g., cn.warcraftlogs.com would get Chinese results).
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesSummary(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/summary/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesDamageDone(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/damage-done/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesDamageTaken(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/damage-taken/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesHealing(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/healing/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesCasts(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/casts/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesSummons(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/summons/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesBuffs(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/buffs/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesDebuffs(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/debuffs/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesDeaths(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/deaths/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesSurvivability(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/survivability/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesResources(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/resources/{code}",
		opt,
		resp,
	)
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Raw) ReportTablesResourcesGains(context context.Context, opt *ReportTablesOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/tables/resources-gains/{code}",
		opt,
		resp,
	)
}
