// Package clientctx stores and retrieves the go-netbox APIClient from a
// cobra command context, allowing the root command to initialise the client
// once and pass it to all subcommands without global state.
package clientctx

import (
  "context"
  "fmt"

  netbox "github.com/netbox-community/go-netbox/v4"
)

type contextKey struct{}

// WithClient returns a new context carrying the provided NetBox API client.
func WithClient(ctx context.Context, client *netbox.APIClient) context.Context {
  return context.WithValue(ctx, contextKey{}, client)
}

// Client retrieves the NetBox API client from ctx. It returns an error if no
// client has been stored, which indicates a programming error in command setup.
func Client(ctx context.Context) (*netbox.APIClient, error) {
  client, ok := ctx.Value(contextKey{}).(*netbox.APIClient)
  if !ok || client == nil {
    return nil, fmt.Errorf("netbox client not found in context")
  }
  return client, nil
}
