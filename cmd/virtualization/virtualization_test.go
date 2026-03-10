// Package virtualization_test exercises the virtualization subcommands that
// contain custom filtering logic not covered by the generic cmdutil helpers.
package virtualization

import (
	"context"
	"strings"
	"testing"

	netbox "github.com/netbox-community/go-netbox/v4"

	"github.com/kirtis/netbox-cli/internal/clientctx"
)

func newTestClient() *netbox.APIClient {
	return netbox.NewAPIClientFor("http://localhost", "test-token")
}

func ctxWithClient() context.Context {
	return clientctx.WithClient(context.Background(), newTestClient())
}

// ── virtualMachinesListCmd ────────────────────────────────────────────────────

func TestVirtualMachinesListCmd_Flags(t *testing.T) {
	cmd := virtualMachinesListCmd()
	for _, flag := range []string{"name", "site", "role", "cluster", "status", "tag"} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("expected --%s flag to be registered", flag)
		}
	}
}

func TestVirtualMachinesListCmd_MissingClient(t *testing.T) {
	cmd := virtualMachinesListCmd()
	cmd.SetContext(context.Background())
	if err := cmd.RunE(cmd, nil); err == nil {
		t.Fatal("expected error when no client in context")
	}
}

// ── virtualMachinesGetCmd ─────────────────────────────────────────────────────

func TestVirtualMachinesGetCmd_Flags(t *testing.T) {
	cmd := virtualMachinesGetCmd()
	for _, flag := range []string{"id", "name"} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("expected --%s flag to be registered", flag)
		}
	}
}

func TestVirtualMachinesGetCmd_BothIDAndName(t *testing.T) {
	cmd := virtualMachinesGetCmd()
	cmd.SetContext(ctxWithClient())
	if err := cmd.Flags().Set("id", "1"); err != nil {
		t.Fatalf("setting --id: %v", err)
	}
	if err := cmd.Flags().Set("name", "web-01"); err != nil {
		t.Fatalf("setting --name: %v", err)
	}
	err := cmd.RunE(cmd, nil)
	if err == nil {
		t.Fatal("expected error when both --id and --name are set")
	}
	if !strings.Contains(err.Error(), "only one") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVirtualMachinesGetCmd_NeitherIDNorName(t *testing.T) {
	cmd := virtualMachinesGetCmd()
	cmd.SetContext(ctxWithClient())
	err := cmd.RunE(cmd, nil)
	if err == nil {
		t.Fatal("expected error when neither --id nor --name is set")
	}
	if !strings.Contains(err.Error(), "either --id or --name is required") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVirtualMachinesGetCmd_MissingClient(t *testing.T) {
	cmd := virtualMachinesGetCmd()
	cmd.SetContext(context.Background())
	if err := cmd.RunE(cmd, nil); err == nil {
		t.Fatal("expected error when no client in context")
	}
}
