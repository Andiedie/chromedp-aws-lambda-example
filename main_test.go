package main

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
	err := Handler(context.Background(), nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
