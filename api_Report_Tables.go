package fflogs

import (
	"context"

	"github.com/RyuaNerin/go-fflogs/structure/reporttables"
)

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesSummary(context context.Context, opt *ReportTablesOptions) (*reporttables.Summary, error) {
	var resp reporttables.Summary
	err := c.Raw.ReportTablesSummary(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesDamageDone(context context.Context, opt *ReportTablesOptions) (*reporttables.DamageDone, error) {
	var resp reporttables.DamageDone
	err := c.Raw.ReportTablesDamageDone(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesDamageTaken(context context.Context, opt *ReportTablesOptions) (*reporttables.DamageTaken, error) {
	var resp reporttables.DamageTaken
	err := c.Raw.ReportTablesDamageTaken(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesHealing(context context.Context, opt *ReportTablesOptions) (*reporttables.Healing, error) {
	var resp reporttables.Healing
	err := c.Raw.ReportTablesHealing(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesCasts(context context.Context, opt *ReportTablesOptions) (*reporttables.Casts, error) {
	var resp reporttables.Casts
	err := c.Raw.ReportTablesCasts(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesSummons(context context.Context, opt *ReportTablesOptions) (*reporttables.Summons, error) {
	var resp reporttables.Summons
	err := c.Raw.ReportTablesSummons(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesBuffs(context context.Context, opt *ReportTablesOptions) (*reporttables.Buffs, error) {
	var resp reporttables.Buffs
	err := c.Raw.ReportTablesBuffs(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesDebuffs(context context.Context, opt *ReportTablesOptions) (*reporttables.Debuffs, error) {
	var resp reporttables.Debuffs
	err := c.Raw.ReportTablesDebuffs(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesDeaths(context context.Context, opt *ReportTablesOptions) (*reporttables.Deaths, error) {
	var resp reporttables.Deaths
	err := c.Raw.ReportTablesDeaths(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesSurvivability(context context.Context, opt *ReportTablesOptions) (*reporttables.Survivability, error) {
	var resp reporttables.Survivability
	err := c.Raw.ReportTablesSurvivability(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesResources(context context.Context, opt *ReportTablesOptions) (*reporttables.Resources, error) {
	var resp reporttables.Resources
	err := c.Raw.ReportTablesResources(context, opt, &resp)
	return &resp, err
}

// Gets a table of entries, either by actor or ability, of damage, healing and cast totals for each entry. This API exactly follows what is returned for the Tables panes on the site. It can and will change as the needs of those panes do, and as such should never be considered a frozen API. Use at your own risk.
func (c *Client) ReportTablesResourcesGains(context context.Context, opt *ReportTablesOptions) (*reporttables.ResourcesGains, error) {
	var resp reporttables.ResourcesGains
	err := c.Raw.ReportTablesResourcesGains(context, opt, &resp)
	return &resp, err
}
