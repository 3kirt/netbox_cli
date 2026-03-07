package clientctx

import (
  "context"
  "fmt"

  netbox "github.com/netbox-community/go-netbox/v4"
)

type contextKey struct{}

func WithClient(ctx context.Context, client *netbox.APIClient) context.Context {
  return context.WithValue(ctx, contextKey{}, client)
}

func Client(ctx context.Context) (*netbox.APIClient, error) {
  client, ok := ctx.Value(contextKey{}).(*netbox.APIClient)
  if !ok || client == nil {
    return nil, fmt.Errorf("netbox client not found in context")
  }
  return client, nil
}
