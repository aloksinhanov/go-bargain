//Deals with the app life-cycle management.
// No business logic here whatsoever.
package app

import (
	"context"
	"log"
	"sync"
)

//defining a custom type for context key to avoid collusion
type contextKey string

const (
	keyWaitGroup contextKey = "waitGroup"
)

func SetWaitGroup(ctx context.Context) context.Context {
	//Set a waitgroup in the context which can be uses by gorotines across the application
	return context.WithValue(ctx, keyWaitGroup, &sync.WaitGroup{})
}

func GetWaitGroup(ctx context.Context) *sync.WaitGroup {

	wg, ok := ctx.Value(keyWaitGroup).(*sync.WaitGroup)
	if !ok {
		log.Fatalf("Failed to typecast waitgroup from context.")
	}
	return wg
}
