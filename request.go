package hyprland

import (
	"encoding/json"
	"fmt"
	"net"
)

// RequestClient is used to send request to hyprland socket
type RequestClient struct {
	// socket is the hyprland socket path
	Socket SocketPath
	// conn is socket connection
	conn net.Conn
}

// NewRequestClient creates a new RequestClient.
func NewRequestClient() *RequestClient {
	c := new(RequestClient)
	c.Socket = should(GetEventSocket())
	return c
}

func (c *RequestClient) Connect() error {
	if c.conn != nil {
		return nil
	}

	socket, err := GetRequestSocket()
	if err != nil {
		return err
	}
	c.Socket = socket

	conn, err := net.Dial("unix", string(socket))
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *RequestClient) Close() error {
	if c.conn == nil {
		return nil
	}
	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *RequestClient) request(cmd string, v any) error {
	if err := c.Connect(); err != nil {
		return err
	}
	defer c.Close()

	if _, err := fmt.Fprintf(c.conn, "j/%s", cmd); err != nil {
		return err
	}

	return json.NewDecoder(c.conn).Decode(v)
}

func (c *RequestClient) GetActiveWindow() (Client, error) {
	var w Client
	return w, c.request("activewindow", &w)
}

func (c *RequestClient) GetAnimations() (Animations, error) {
	var a Animations
	return a, c.request("animations", &a)
}

func (c *RequestClient) GetBinds() (Binds, error) {
	var a Binds
	return a, c.request("binds", &a)
}

func (c *RequestClient) GetCursorPosition() (CursorPosition, error) {
	var a CursorPosition
	return a, c.request("cursorpos", &a)
}

func (c *RequestClient) GetClients() (Clients, error) {
	var clients Clients
	return clients, c.request("clients", &clients)
}

func (c *RequestClient) GetMonitors() (Monitors, error) {
	var m Monitors
	return m, c.request("monitors", &m)
}

func (c *RequestClient) GetWorkspaces() (Workspaces, error) {
	var w Workspaces
	return w, c.request("workspaces", &w)
}

func (c *RequestClient) GetActiveWorkspace() (Workspace, error) {
	var w Workspace
	return w, c.request("activeworkspace", &w)
}
