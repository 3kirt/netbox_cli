// Package cmdutil_test exercises the shared command-builder helpers.
package cmdutil_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	netbox "github.com/netbox-community/go-netbox/v4"

	"github.com/kirtis/netbox-cli/internal/clientctx"
	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// newTestClient returns a minimal APIClient for use in tests.  No real
// network calls are made; the URL and token values are placeholders.
func newTestClient() *netbox.APIClient {
	return netbox.NewAPIClientFor("http://localhost", "test-token")
}

// ctxWithClient returns a context that carries a test APIClient.
func ctxWithClient() context.Context {
	return clientctx.WithClient(context.Background(), newTestClient())
}

// captureStdout replaces os.Stdout with a pipe for the duration of fn, then
// returns everything written to it.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	old := os.Stdout
	os.Stdout = w
	fn()
	if err := w.Close(); err != nil {
		t.Fatalf("closing stdout pipe: %v", err)
	}
	os.Stdout = old
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("reading captured stdout: %v", err)
	}
	return buf.String()
}

// withStdin replaces os.Stdin with a reader containing content for the
// duration of fn.
func withStdin(t *testing.T, content string, fn func()) {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	if _, err := io.WriteString(w, content); err != nil {
		t.Fatalf("writing fake stdin: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("closing stdin pipe: %v", err)
	}
	old := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = old }()
	fn()
}

// ── OutputJSON ────────────────────────────────────────────────────────────────

func TestOutputJSON(t *testing.T) {
	t.Run("writes indented JSON followed by newline", func(t *testing.T) {
		type payload struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		var encErr error
		got := captureStdout(t, func() {
			encErr = cmdutil.OutputJSON(payload{Name: "test", Value: 42})
		})
		if encErr != nil {
			t.Fatalf("OutputJSON returned error: %v", encErr)
		}
		want := "{\n  \"name\": \"test\",\n  \"value\": 42\n}\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("nested struct is indented correctly", func(t *testing.T) {
		type inner struct {
			X int `json:"x"`
		}
		type outer struct {
			Inner inner `json:"inner"`
		}
		var encErr error
		got := captureStdout(t, func() {
			encErr = cmdutil.OutputJSON(outer{Inner: inner{X: 1}})
		})
		if encErr != nil {
			t.Fatalf("OutputJSON returned error: %v", encErr)
		}
		if !strings.Contains(got, "  \"inner\"") {
			t.Errorf("output not indented with two spaces: %q", got)
		}
	})
}

// ── APIError ──────────────────────────────────────────────────────────────────

func TestAPIError(t *testing.T) {
	t.Run("prefixes error message", func(t *testing.T) {
		cause := errors.New("connection refused")
		err := cmdutil.APIError(cause)
		want := "netbox API error: connection refused"
		if err.Error() != want {
			t.Errorf("got %q, want %q", err.Error(), want)
		}
	})

	t.Run("wraps original error for errors.Is", func(t *testing.T) {
		cause := errors.New("some cause")
		err := cmdutil.APIError(cause)
		if !errors.Is(err, cause) {
			t.Error("errors.Is should find the original cause inside APIError")
		}
	})
}

// ── ListCmd ───────────────────────────────────────────────────────────────────

func TestListCmd(t *testing.T) {
	t.Run("Use is list", func(t *testing.T) {
		cmd := cmdutil.ListCmd("widget", nil)
		if cmd.Use != "list" {
			t.Errorf("got Use=%q, want %q", cmd.Use, "list")
		}
	})

	t.Run("Short contains noun", func(t *testing.T) {
		cmd := cmdutil.ListCmd("widget", nil)
		if !strings.Contains(cmd.Short, "widget") {
			t.Errorf("Short %q does not mention noun", cmd.Short)
		}
	})

	t.Run("invokes callback with client from context", func(t *testing.T) {
		client := newTestClient()
		ctx := clientctx.WithClient(context.Background(), client)

		var gotClient *netbox.APIClient
		cmd := cmdutil.ListCmd("widget", func(_ context.Context, cl *netbox.APIClient) error {
			gotClient = cl
			return nil
		})
		cmd.SetContext(ctx)

		if err := cmd.RunE(cmd, nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if gotClient != client {
			t.Error("callback received wrong client pointer")
		}
	})

	t.Run("callback error is propagated", func(t *testing.T) {
		want := errors.New("list failed")
		cmd := cmdutil.ListCmd("widget", func(_ context.Context, _ *netbox.APIClient) error {
			return want
		})
		cmd.SetContext(ctxWithClient())

		if err := cmd.RunE(cmd, nil); !errors.Is(err, want) {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("missing client returns error", func(t *testing.T) {
		cmd := cmdutil.ListCmd("widget", func(_ context.Context, _ *netbox.APIClient) error {
			return nil
		})
		cmd.SetContext(context.Background())

		if err := cmd.RunE(cmd, nil); err == nil {
			t.Fatal("expected error when no client in context")
		}
	})
}

// ── GetCmd ────────────────────────────────────────────────────────────────────

func TestGetCmd(t *testing.T) {
	t.Run("Use is get", func(t *testing.T) {
		cmd := cmdutil.GetCmd("widget", nil)
		if cmd.Use != "get" {
			t.Errorf("got Use=%q, want %q", cmd.Use, "get")
		}
	})

	t.Run("--id flag is registered", func(t *testing.T) {
		cmd := cmdutil.GetCmd("widget", nil)
		if cmd.Flags().Lookup("id") == nil {
			t.Error("--id flag not registered")
		}
	})

	t.Run("passes parsed ID to callback", func(t *testing.T) {
		var gotID int32
		cmd := cmdutil.GetCmd("widget", func(_ context.Context, _ *netbox.APIClient, id int32) error {
			gotID = id
			return nil
		})
		cmd.SetContext(ctxWithClient())
		if err := cmd.Flags().Set("id", "99"); err != nil {
			t.Fatalf("setting --id: %v", err)
		}

		if err := cmd.RunE(cmd, nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if gotID != 99 {
			t.Errorf("got id=%d, want 99", gotID)
		}
	})

	t.Run("callback error is propagated", func(t *testing.T) {
		want := errors.New("get failed")
		cmd := cmdutil.GetCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32) error {
			return want
		})
		cmd.SetContext(ctxWithClient())
		_ = cmd.Flags().Set("id", "1")

		if err := cmd.RunE(cmd, nil); !errors.Is(err, want) {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("missing client returns error", func(t *testing.T) {
		cmd := cmdutil.GetCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32) error {
			return nil
		})
		cmd.SetContext(context.Background())
		_ = cmd.Flags().Set("id", "1")

		if err := cmd.RunE(cmd, nil); err == nil {
			t.Fatal("expected error when no client in context")
		}
	})
}

// ── CreateCmd ─────────────────────────────────────────────────────────────────

func TestCreateCmd(t *testing.T) {
	t.Run("Use is create", func(t *testing.T) {
		cmd := cmdutil.CreateCmd("widget", nil)
		if cmd.Use != "create" {
			t.Errorf("got Use=%q, want %q", cmd.Use, "create")
		}
	})

	t.Run("passes stdin bytes to callback", func(t *testing.T) {
		input := `{"name":"foo"}`
		var gotData []byte
		cmd := cmdutil.CreateCmd("widget", func(_ context.Context, _ *netbox.APIClient, data []byte) error {
			gotData = data
			return nil
		})
		cmd.SetContext(ctxWithClient())

		withStdin(t, input, func() {
			if err := cmd.RunE(cmd, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
		if string(gotData) != input {
			t.Errorf("got %q, want %q", gotData, input)
		}
	})

	t.Run("callback error is propagated", func(t *testing.T) {
		want := errors.New("create failed")
		cmd := cmdutil.CreateCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ []byte) error {
			return want
		})
		cmd.SetContext(ctxWithClient())

		withStdin(t, "{}", func() {
			if err := cmd.RunE(cmd, nil); !errors.Is(err, want) {
				t.Errorf("got %v, want %v", err, want)
			}
		})
	})

	t.Run("missing client returns error", func(t *testing.T) {
		cmd := cmdutil.CreateCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ []byte) error {
			return nil
		})
		cmd.SetContext(context.Background())

		withStdin(t, "{}", func() {
			if err := cmd.RunE(cmd, nil); err == nil {
				t.Fatal("expected error when no client in context")
			}
		})
	})
}

// ── UpdateCmd ─────────────────────────────────────────────────────────────────

func TestUpdateCmd(t *testing.T) {
	t.Run("Use is update", func(t *testing.T) {
		cmd := cmdutil.UpdateCmd("widget", nil)
		if cmd.Use != "update" {
			t.Errorf("got Use=%q, want %q", cmd.Use, "update")
		}
	})

	t.Run("--id flag is registered", func(t *testing.T) {
		cmd := cmdutil.UpdateCmd("widget", nil)
		if cmd.Flags().Lookup("id") == nil {
			t.Error("--id flag not registered")
		}
	})

	t.Run("passes ID and stdin bytes to callback", func(t *testing.T) {
		input := `{"name":"bar"}`
		var gotID int32
		var gotData []byte
		cmd := cmdutil.UpdateCmd("widget", func(_ context.Context, _ *netbox.APIClient, id int32, data []byte) error {
			gotID = id
			gotData = data
			return nil
		})
		cmd.SetContext(ctxWithClient())
		if err := cmd.Flags().Set("id", "7"); err != nil {
			t.Fatalf("setting --id: %v", err)
		}

		withStdin(t, input, func() {
			if err := cmd.RunE(cmd, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
		if gotID != 7 {
			t.Errorf("got id=%d, want 7", gotID)
		}
		if string(gotData) != input {
			t.Errorf("got data=%q, want %q", gotData, input)
		}
	})

	t.Run("callback error is propagated", func(t *testing.T) {
		want := errors.New("update failed")
		cmd := cmdutil.UpdateCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32, _ []byte) error {
			return want
		})
		cmd.SetContext(ctxWithClient())
		_ = cmd.Flags().Set("id", "1")

		withStdin(t, "{}", func() {
			if err := cmd.RunE(cmd, nil); !errors.Is(err, want) {
				t.Errorf("got %v, want %v", err, want)
			}
		})
	})

	t.Run("missing client returns error", func(t *testing.T) {
		cmd := cmdutil.UpdateCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32, _ []byte) error {
			return nil
		})
		cmd.SetContext(context.Background())
		_ = cmd.Flags().Set("id", "1")

		withStdin(t, "{}", func() {
			if err := cmd.RunE(cmd, nil); err == nil {
				t.Fatal("expected error when no client in context")
			}
		})
	})
}

// ── DeleteCmd ─────────────────────────────────────────────────────────────────

func TestDeleteCmd(t *testing.T) {
	t.Run("Use is delete", func(t *testing.T) {
		cmd := cmdutil.DeleteCmd("widget", nil)
		if cmd.Use != "delete" {
			t.Errorf("got Use=%q, want %q", cmd.Use, "delete")
		}
	})

	t.Run("--id flag is registered", func(t *testing.T) {
		cmd := cmdutil.DeleteCmd("widget", nil)
		if cmd.Flags().Lookup("id") == nil {
			t.Error("--id flag not registered")
		}
	})

	t.Run("passes parsed ID to callback", func(t *testing.T) {
		var gotID int32
		cmd := cmdutil.DeleteCmd("widget", func(_ context.Context, _ *netbox.APIClient, id int32) error {
			gotID = id
			return nil
		})
		cmd.SetContext(ctxWithClient())
		if err := cmd.Flags().Set("id", "55"); err != nil {
			t.Fatalf("setting --id: %v", err)
		}

		if err := cmd.RunE(cmd, nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if gotID != 55 {
			t.Errorf("got id=%d, want 55", gotID)
		}
	})

	t.Run("callback error is propagated", func(t *testing.T) {
		want := errors.New("delete failed")
		cmd := cmdutil.DeleteCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32) error {
			return want
		})
		cmd.SetContext(ctxWithClient())
		_ = cmd.Flags().Set("id", "1")

		if err := cmd.RunE(cmd, nil); !errors.Is(err, want) {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("missing client returns error", func(t *testing.T) {
		cmd := cmdutil.DeleteCmd("widget", func(_ context.Context, _ *netbox.APIClient, _ int32) error {
			return nil
		})
		cmd.SetContext(context.Background())
		_ = cmd.Flags().Set("id", "1")

		if err := cmd.RunE(cmd, nil); err == nil {
			t.Fatal("expected error when no client in context")
		}
	})
}
