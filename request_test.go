package hyprland

import (
	"testing"
)

func dropVal[T any](_ T, err error) error { return err }

func TestRequestClient(t *testing.T) {
	c := NewRequestClient()

	tests := []struct {
		name string
		err  error
	}{
		{"GetActiveWindow", dropVal(c.GetActiveWindow())},
		{"GetClients", dropVal(c.GetClients())},
		{"GetMonitors", dropVal(c.GetMonitors())},
		{"GetAnimations", dropVal(c.GetAnimations())},
		{"GetWorkspaces", dropVal(c.GetWorkspaces())},
		{"GetActiveWorkspace", dropVal(c.GetActiveWorkspace())},
		{"GetBinds", dropVal(c.GetBinds())},
		{"GetCursorPosition", dropVal(c.GetCursorPosition())},
	}

	for _, tt := range tests {
		t.Run("RequestClient."+tt.name, func(t *testing.T) {
			if tt.err != nil {
				t.Errorf("%s() failed: %v", tt.name, tt.err)
			}
		})
	}
}
