package fflogs

import (
	"context"
	"testing"
)

func TestReportsUser(t *testing.T) {
	c := newTestClient()

	opt := ReportsUserOptions{
		UserName: testUserName,
	}

	_, err := c.ReportsUser(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
