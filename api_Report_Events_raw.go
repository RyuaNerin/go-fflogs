package fflogs

import (
	"context"
)

type ReportEventsOptions struct {
	//View               string `path:"view"`                // The type of data requested. Supported values are 'summary', 'damage-done', 'damage-taken', 'healing', 'casts', 'summons', 'buffs', 'debuffs', 'deaths', 'threat', 'resources', 'interrupts' and 'dispels'.
	Code               string  `path:"code"`                // The specific report to collect table entries for.
	Start              *int    `query:"start"`              // A start time. This is a time from the start of the report in milliseconds. If omitted, 0 is assumed.
	End                *int    `query:"end"`                // An end time. This is a time from the start of the report in milliseconds. If omitted, 0 is assumed.
	Hostility          *int    `query:"hostility"`          // An optional hostility value of 0 or 1. The default is 0. A value of 0 means to collect data for Friendlies. A value of 1 means to collect data for Enemies.
	Sourceid           *int    `query:"sourceid"`           // An optional actor ID to filter to. If set, only events where the ID matches the source (or target for damage-taken) of the event will be returned. The actor's pets will also be included (unless the options field overrides).
	Sourceinstance     *int    `query:"sourceinstance"`     // An optional actor instance ID to filter to. If set, only events where the instance ID matches the source (or target for damage-taken) of the event will be returned. This is useful to look for all events involving NPC N, where N is the actor instance ID.
	Sourceclass        *string `query:"sourceclass"`        // An optional actor class to filter to. If set, only events where the source (or target for damage-taken) involves that class (e.g., Mage) will be returned.
	Targetid           *int    `query:"targetid"`           // An optional actor ID to filter to. If set, only events where the ID matches the target (or source for damage-taken) of the event will be returned. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Targetinstance     *int    `query:"targetinstance"`     // An optional actor instance ID to filter to. If set, only events where the instance ID matches the target (or source for damage-taken) of the event will be returned. This is useful to look for all events involving NPC N, where N is the actor instance ID. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Targetclass        *string `query:"targetclass"`        // An optional actor class to filter to. If set, only events where the target (or source for damage-taken) involves that class (e.g., Mage) will be returned. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	SourceAurasPresent *string `query:"sourceAurasPresent"` // A comma-separated string of aura game IDs. Only matches if the aura is present on the source.
	SourceAurasAbsent  *string `query:"sourceAurasAbsent"`  // A comma-separated string of aura game IDs. Only matches if the aura is absent on the source.
	TargetAurasPresent *string `query:"targetAurasPresent"` // A comma-separated string of aura game IDs. Only matches if the aura is present on the target.
	TargetAurasAbsent  *string `query:"targetAurasAbsent"`  // A comma-separated string of aura game IDs. Only matches if the aura is absent on the target.
	Abilityid          *int    `query:"abilityid"`          // An optional ability ID to filter to. If set, only events where the ability matches will be returned. Consolidated abilities (WCL only) are represented using a negative number that matches the ability ID that everything is consolidated under. For the 'deaths' view, this represents a specific killing blow. For the resources views, the abilityid is not an ability but a resource type. Valid resource types can be viewed at https://www.fflogs.com/reports/resource_types/
	Death              *int    `query:"death"`              // An optional death to filter to. Only used for the deaths command. Select the Nth death in the time range that matches all the other filters.
	Options            *int    `query:"options"`            // A set of options for what to include/exclude. These correspond to options like Include Overkill in the Damage Done pane. Complete list will be forthcoming. If omitted, appropriate defaults that match WCL's default behavior will be chosen. This value is not used in the 'deaths', 'survivability', 'resources' and 'resources-gains' views.
	Cutoff             *int    `query:"cutoff"`             // An optional death cutoff. If set, events after that number of deaths have occurred will not be examined.
	Encounter          *int    `query:"encounter"`          // An optional encounter filter. If set to a specific encounter ID, only fights involving a specific encounter will be considered. The encounter IDs match those used in rankings/statistics.
	Wipes              *int    `query:"wipes"`              // An optional wipes filter. If set to 1, only wipes will be considered.
	Difficulty         *int    `query:"difficulty"`         // An optional difficulty filter.
	Filter             *string `query:"filter"`             // An optional filter written in WCL's expression language. Events must match the filter to be included.
	Translate          *bool   `query:"translate"`          // An optional flag indicating that the results should be translated into the language of the host (e.g., cn.warcraftlogs.com would get Chinese results).
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsSummary(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/summary/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsDamageDone(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/damage-done/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsDamageTaken(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/damage-taken/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsHealing(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/healing/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsCasts(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/casts/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsSummons(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/summons/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsBuffs(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/buffs/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsDebuffs(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/debuffs/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsDeaths(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/deaths/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsThreat(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/threat/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsResources(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/resources/{code}",
		opt,
		resp,
	)
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Raw) ReportEventsInterrupts(context context.Context, opt *ReportEventsOptions, resp interface{}) error {
	return c.call(
		context,
		"/report/events/interrupts/{code}",
		opt,
		resp,
	)
}
