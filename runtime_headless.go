//go:build headless

package main

import "context"

func runtimeEventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	// No-op in headless mode
}
