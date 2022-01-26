package fflogs

import (
	"context"
	"testing"
)

func TestReportFights(t *testing.T) {
	c := newTestClient()

	opt := ReportFightsOptions{
		Code: testCode,
	}

	_, err := c.ReportFights(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
