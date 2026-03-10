// Package ipam_test exercises the IPAM subcommands that contain custom
// filtering logic not covered by the generic cmdutil helpers.
package ipam

import (
	"context"
	"testing"
)

// ── ipAddressesListCmd ────────────────────────────────────────────────────────

func TestIPAddressesListCmd_Flags(t *testing.T) {
	cmd := ipAddressesListCmd()
	for _, flag := range []string{"virtual-machine", "device", "address"} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("expected --%s flag to be registered", flag)
		}
	}
}

func TestIPAddressesListCmd_MissingClient(t *testing.T) {
	cmd := ipAddressesListCmd()
	cmd.SetContext(context.Background())
	if err := cmd.RunE(cmd, nil); err == nil {
		t.Fatal("expected error when no client in context")
	}
}
