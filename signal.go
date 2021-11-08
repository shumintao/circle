/*
Copyright 2021 The Gridsum Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package circle

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSignals returns a context that is canceled with any signal in sigs.
func WithSignals(ctx context.Context, sigs ...os.Signal) context.Context {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sigs...)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			return
		}
	}()
	return ctx
}

// WithStandardSignals cancels the context on os.Interrupt, syscall.SIGTERM.
func WithStandardSignals(ctx context.Context) context.Context {
	return WithSignals(ctx, os.Interrupt, syscall.SIGTERM)
}

