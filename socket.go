package hyprland

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// SocketPath is the path of a hyprland socket
type SocketPath string

// GetRequestSocket returns hyprland request socket (socket) path
func GetRequestSocket() (SocketPath, error) { return getSocket(".socket.sock") }

// GetEventSocket returns hyprland event socket (socket2) path
func GetEventSocket() (SocketPath, error) { return getSocket(".socket2.sock") }

func should[T any](v T, _ error) T {
	return v
}

func getSocket(name string) (SocketPath, error) {
	his, ok := os.LookupEnv("HYPRLAND_INSTANCE_SIGNATURE")
	if !ok {
		return "", errors.New("HYPRLAND_INSTANCE_SIGNATURE env does not exist")
	}

	xdgRuntimeDir, ok := os.LookupEnv("XDG_RUNTIME_DIR")
	if !ok {
		u, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("failed to get XDG_RUNTIME_DIR: %w", err)
		}
		xdgRuntimeDir = filepath.Join("/run/user", u.Uid)
	}

	return SocketPath(filepath.Join(xdgRuntimeDir, "hypr", his, name)), nil
}
