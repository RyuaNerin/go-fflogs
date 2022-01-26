package fflogs

import (
	"context"
	"testing"
)

func TestClasses(t *testing.T) {
	c := newTestClient()
	_, err := c.Classes(context.Background())
	if err != nil {
		panic(err)
	}
}
