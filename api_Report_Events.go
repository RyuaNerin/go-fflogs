package fflogs

import (
	"context"

	"github.com/RyuaNerin/go-fflogs/structure/reportevents"
)

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsSummary(context context.Context, opt *ReportEventsOptions) (*reportevents.Summary, error) {
	var resp reportevents.Summary
	err := c.Raw.ReportEventsSummary(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsDamageDone(context context.Context, opt *ReportEventsOptions) (*reportevents.DamageDone, error) {
	var resp reportevents.DamageDone
	err := c.Raw.ReportEventsDamageDone(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsDamageTaken(context context.Context, opt *ReportEventsOptions) (*reportevents.DamageTaken, error) {
	var resp reportevents.DamageTaken
	err := c.Raw.ReportEventsDamageTaken(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsHealing(context context.Context, opt *ReportEventsOptions) (*reportevents.Healing, error) {
	var resp reportevents.Healing
	err := c.Raw.ReportEventsHealing(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsCasts(context context.Context, opt *ReportEventsOptions) (*reportevents.Casts, error) {
	var resp reportevents.Casts
	err := c.Raw.ReportEventsCasts(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsSummons(context context.Context, opt *ReportEventsOptions) (*reportevents.Summons, error) {
	var resp reportevents.Summons
	err := c.Raw.ReportEventsSummons(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsBuffs(context context.Context, opt *ReportEventsOptions) (*reportevents.Buffs, error) {
	var resp reportevents.Buffs
	err := c.Raw.ReportEventsBuffs(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsDebuffs(context context.Context, opt *ReportEventsOptions) (*reportevents.Debuffs, error) {
	var resp reportevents.Debuffs
	err := c.Raw.ReportEventsDebuffs(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsDeaths(context context.Context, opt *ReportEventsOptions) (*reportevents.Deaths, error) {
	var resp reportevents.Deaths
	err := c.Raw.ReportEventsDeaths(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsThreat(context context.Context, opt *ReportEventsOptions) (*reportevents.Threat, error) {
	var resp reportevents.Threat
	err := c.Raw.ReportEventsThreat(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsResources(context context.Context, opt *ReportEventsOptions) (*reportevents.Resources, error) {
	var resp reportevents.Resources
	err := c.Raw.ReportEventsResources(context, opt, &resp)
	return &resp, err
}

// Gets a set of events based off the view you're asking for. This exactly corresponds to the Events view on the site.
func (c *Client) ReportEventsInterrupts(context context.Context, opt *ReportEventsOptions) (*reportevents.Interrupts, error) {
	var resp reportevents.Interrupts
	err := c.Raw.ReportEventsInterrupts(context, opt, &resp)
	return &resp, err
}
