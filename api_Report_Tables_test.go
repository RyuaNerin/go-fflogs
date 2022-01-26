package fflogs

import (
	"context"
	"testing"
)

var (
	optReportTables = ReportTablesOptions{
		Code:  testCode,
		Start: &testStart,
		End:   &testEnd,
	}
)

func TestReportTablesSummary(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesSummary(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesDamageDone(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesDamageDone(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesDamageTaken(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesDamageTaken(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesHealing(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesHealing(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesCasts(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesCasts(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesSummons(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesSummons(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesBuffs(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesBuffs(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesDebuffs(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesDebuffs(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesDeaths(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesDeaths(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesSurvivability(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesSurvivability(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesResources(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesResources(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}

func TestReportTablesResourcesGains(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportTablesResourcesGains(context.Background(), &optReportTables)
	if err != nil {
		panic(err)
	}
}
