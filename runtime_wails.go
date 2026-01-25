//go:build !headless

package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func runtimeEventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	runtime.EventsEmit(ctx, eventName, optionalData...)
}
