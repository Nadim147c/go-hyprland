package hyprland

import (
	"encoding/json"
	"fmt"
)

type SimpleWorkspace struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Clients []Client

type Client struct {
	Address          string          `json:"address"`
	Mapped           bool            `json:"mapped"`
	Hidden           bool            `json:"hidden"`
	At               []int64         `json:"at"`
	Size             []int64         `json:"size"`
	Workspace        SimpleWorkspace `json:"workspace"`
	Floating         bool            `json:"floating"`
	Pseudo           bool            `json:"pseudo"`
	Monitor          int64           `json:"monitor"`
	Class            string          `json:"class"`
	Title            string          `json:"title"`
	InitialClass     string          `json:"initialClass"`
	InitialTitle     string          `json:"initialTitle"`
	PID              int64           `json:"pid"`
	Xwayland         bool            `json:"xwayland"`
	Pinned           bool            `json:"pinned"`
	Fullscreen       int64           `json:"fullscreen"`
	FullscreenClient int64           `json:"fullscreenClient"`
	Grouped          []string        `json:"grouped"`
	Tags             []string        `json:"tags"`
	Swallowing       string          `json:"swallowing"`
	FocusHistoryID   int64           `json:"focusHistoryID"`
	InhibitingIdle   bool            `json:"inhibitingIdle"`
	XdgTag           string          `json:"xdgTag"`
	XdgDescription   string          `json:"xdgDescription"`
}

type Monitors []Monitor

type Monitor struct {
	ID               int64           `json:"id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Make             string          `json:"make"`
	Model            string          `json:"model"`
	Serial           string          `json:"serial"`
	Width            int64           `json:"width"`
	Height           int64           `json:"height"`
	RefreshRate      float64         `json:"refreshRate"`
	X                int64           `json:"x"`
	Y                int64           `json:"y"`
	ActiveWorkspace  SimpleWorkspace `json:"activeWorkspace"`
	SpecialWorkspace SimpleWorkspace `json:"specialWorkspace"`
	Reserved         []int64         `json:"reserved"`
	Scale            float64         `json:"scale"`
	Transform        int64           `json:"transform"`
	Focused          bool            `json:"focused"`
	DPMSStatus       bool            `json:"dpmsStatus"`
	Vrr              bool            `json:"vrr"`
	Solitary         string          `json:"solitary"`
	ActivelyTearing  bool            `json:"activelyTearing"`
	DirectScanoutTo  string          `json:"directScanoutTo"`
	Disabled         bool            `json:"disabled"`
	CurrentFormat    string          `json:"currentFormat"`
	MirrorOf         string          `json:"mirrorOf"`
	AvailableModes   []string        `json:"availableModes"`
}

type Animation struct {
	Name       string  `json:"name"`
	Overridden bool    `json:"overridden"`
	Bezier     string  `json:"bezier"`
	Enabled    bool    `json:"enabled"`
	Speed      float64 `json:"speed"`
	Style      string  `json:"style"`
}

type Bezier struct {
	Name string `json:"name"`
}

type Animations struct {
	List    []Animation
	Beziers []Bezier
}

// UnmarshalJSON implements json.Unmarshaler
func (a *Animations) UnmarshalJSON(data []byte) error {
	var raw [2]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("expected a 2-element JSON array: %w", err)
	}

	// Decode both arrays
	if err := json.Unmarshal(raw[0], &a.List); err != nil {
		return fmt.Errorf("failed to unmarshal animations: %w", err)
	}
	if err := json.Unmarshal(raw[1], &a.Beziers); err != nil {
		return fmt.Errorf("failed to unmarshal beziers: %w", err)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (a Animations) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{a.List, a.Beziers})
}

type Workspaces []Workspace

type Workspace struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Monitor         string `json:"monitor"`
	MonitorID       int64  `json:"monitorID"`
	Windows         int64  `json:"windows"`
	Hasfullscreen   bool   `json:"hasfullscreen"`
	Lastwindow      string `json:"lastwindow"`
	Lastwindowtitle string `json:"lastwindowtitle"`
	Ispersistent    bool   `json:"ispersistent"`
}

type Binds []Bind

type Bind struct {
	Locked         bool   `json:"locked"`
	Mouse          bool   `json:"mouse"`
	Release        bool   `json:"release"`
	Repeat         bool   `json:"repeat"`
	LongPress      bool   `json:"longPress"`
	NonConsuming   bool   `json:"non_consuming"`
	HasDescription bool   `json:"has_description"`
	Modmask        int64  `json:"modmask"`
	Submap         string `json:"submap"`
	Key            string `json:"key"`
	Keycode        int64  `json:"keycode"`
	CatchAll       bool   `json:"catch_all"`
	Description    string `json:"description"`
	Dispatcher     string `json:"dispatcher"`
	Arg            string `json:"arg"`
}

type CursorPosition struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
