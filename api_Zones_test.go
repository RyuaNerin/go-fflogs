package fflogs

import (
	"context"
	"testing"
)

func TestZones(t *testing.T) {
	c := newTestClient()

	_, err := c.Zones(context.Background())
	if err != nil {
		panic(err)
	}
}
