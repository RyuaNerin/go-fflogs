package fflogs

import (
	"context"
	"testing"
)

var (
	optReportEvents = ReportEventsOptions{
		Code:  testCode,
		Start: &testStart,
		End:   &testEnd,
	}
)

func TestReportEventsSummary(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsSummary(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsDamageDone(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsDamageDone(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsDamageTaken(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsDamageTaken(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsHealing(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsHealing(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsCasts(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsCasts(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsSummons(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsSummons(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsBuffs(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsBuffs(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsDebuffs(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsDebuffs(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsDeaths(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsDeaths(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

func TestReportEventsThreat(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsThreat(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}

/* TODO
func TestReportEventsResources(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsResources(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}
*/

func TestReportEventsInterrupts(t *testing.T) {
	c := newTestClient()
	_, err := c.ReportEventsInterrupts(context.Background(), &optReportEvents)
	if err != nil {
		panic(err)
	}
}
