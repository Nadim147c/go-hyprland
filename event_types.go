package hyprland

// Event represents a Hyprland event name.
type Event string

const (
	// Workspace is emitted on workspace change. Only emitted when a user
	// requests a workspace change, not on mouse movements (see focusedmon).
	// Args: WORKSPACENAME
	Workspace Event = "workspace"

	// WorkspaceV2 is emitted on workspace change (v2). Only emitted when a user
	// requests a workspace change, not on mouse movements (see focusedmon).
	// Args: WORKSPACEID, WORKSPACENAME
	WorkspaceV2 Event = "workspacev2"

	// FocusedMon is emitted when the active monitor changes.
	// Args: MONNAME, WORKSPACENAME
	FocusedMon Event = "focusedmon"

	// FocusedMonV2 is emitted when the active monitor changes (v2).
	// Args: MONNAME, WORKSPACEID
	FocusedMonV2 Event = "focusedmonv2"

	// ActiveWindow is emitted when the active window changes.
	// Args: WINDOWCLASS, WINDOWTITLE
	ActiveWindow Event = "activewindow"

	// ActiveWindowV2 is emitted when the active window changes (v2).
	// Args: WINDOWADDRESS
	ActiveWindowV2 Event = "activewindowv2"

	// Fullscreen is emitted when a window's fullscreen status changes.
	// Args: 0 (exit fullscreen) / 1 (enter fullscreen)
	Fullscreen Event = "fullscreen"

	// MonitorRemoved is emitted when a monitor is disconnected.
	// Args: MONITORNAME
	MonitorRemoved Event = "monitorremoved"

	// MonitorRemovedV2 is emitted when a monitor is disconnected (v2).
	// Args: MONITORID, MONITORNAME, MONITORDESCRIPTION
	MonitorRemovedV2 Event = "monitorremovedv2"

	// MonitorAdded is emitted when a monitor is connected.
	// Args: MONITORNAME
	MonitorAdded Event = "monitoradded"

	// MonitorAddedV2 is emitted when a monitor is connected (v2).
	// Args: MONITORID, MONITORNAME, MONITORDESCRIPTION
	MonitorAddedV2 Event = "monitoraddedv2"

	// CreateWorkspace is emitted when a workspace is created.
	// Args: WORKSPACENAME
	CreateWorkspace Event = "createworkspace"

	// CreateWorkspaceV2 is emitted when a workspace is created (v2).
	// Args: WORKSPACEID, WORKSPACENAME
	CreateWorkspaceV2 Event = "createworkspacev2"

	// DestroyWorkspace is emitted when a workspace is destroyed.
	// Args: WORKSPACENAME
	DestroyWorkspace Event = "destroyworkspace"

	// DestroyWorkspaceV2 is emitted when a workspace is destroyed (v2).
	// Args: WORKSPACEID, WORKSPACENAME
	DestroyWorkspaceV2 Event = "destroyworkspacev2"

	// MoveWorkspace is emitted when a workspace moves to a different monitor.
	// Args: WORKSPACENAME, MONNAME
	MoveWorkspace Event = "moveworkspace"

	// MoveWorkspaceV2 is emitted when a workspace moves to a different monitor
	// (v2).
	// Args: WORKSPACEID, WORKSPACENAME, MONNAME
	MoveWorkspaceV2 Event = "moveworkspacev2"

	// RenameWorkspace is emitted when a workspace is renamed.
	// Args: WORKSPACEID, NEWNAME
	RenameWorkspace Event = "renameworkspace"

	// ActiveSpecial is emitted when the special workspace on a monitor changes.
	// Closing results in an empty WORKSPACENAME.
	// Args: WORKSPACENAME, MONNAME
	ActiveSpecial Event = "activespecial"

	// ActiveSpecialV2 is emitted when the special workspace on a monitor
	// changes (v2). Closing results in empty WORKSPACEID and WORKSPACENAME.
	// Args: WORKSPACEID, WORKSPACENAME, MONNAME
	ActiveSpecialV2 Event = "activespecialv2"

	// ActiveLayout is emitted when the layout of the active keyboard changes.
	// Args: KEYBOARDNAME, LAYOUTNAME
	ActiveLayout Event = "activelayout"

	// OpenWindow is emitted when a window is opened.
	// Args: WINDOWADDRESS, WORKSPACENAME, WINDOWCLASS, WINDOWTITLE
	OpenWindow Event = "openwindow"

	// CloseWindow is emitted when a window is closed.
	// Args: WINDOWADDRESS
	CloseWindow Event = "closewindow"

	// MoveWindow is emitted when a window moves to a different workspace.
	// Args: WINDOWADDRESS, WORKSPACENAME
	MoveWindow Event = "movewindow"

	// MoveWindowV2 is emitted when a window moves to a different workspace
	// (v2).
	// Args: WINDOWADDRESS, WORKSPACEID, WORKSPACENAME
	MoveWindowV2 Event = "movewindowv2"

	// OpenLayer is emitted when a layer surface is mapped.
	// Args: NAMESPACE
	OpenLayer Event = "openlayer"

	// CloseLayer is emitted when a layer surface is unmapped.
	// Args: NAMESPACE
	CloseLayer Event = "closelayer"

	// Submap is emitted when a keybind submap changes.
	// Empty value means the default submap.
	// Args: SUBMAPNAME
	Submap Event = "submap"

	// ChangeFloatingMode is emitted when a window toggles its floating mode.
	// Args: WINDOWADDRESS, FLOATING (0 or 1)
	ChangeFloatingMode Event = "changefloatingmode"

	// Urgent is emitted when a window requests an urgent state.
	// Args: WINDOWADDRESS
	Urgent Event = "urgent"

	// Screencast is emitted when a client's screencopy state changes. There may
	// be multiple clients.
	// Args: STATE (0/1), OWNER (0 = monitor share, 1 = window share)
	Screencast Event = "screencast"

	// WindowTitle is emitted when a window title changes.
	// Args: WINDOWADDRESS
	WindowTitle Event = "windowtitle"

	// WindowTitleV2 is emitted when a window title changes (v2).
	// Args: WINDOWADDRESS, WINDOWTITLE
	WindowTitleV2 Event = "windowtitlev2"

	// ToggleGroup is emitted when the togglegroup command is used. Returns
	// state and window handles, e.g. "0,64cea2525760,64cea2522380".
	// Args: STATE (0/1), WINDOWADDRESS(ES)
	ToggleGroup Event = "togglegroup"

	// MoveIntoGroup is emitted when a window is merged into a group.
	// Args: WINDOWADDRESS
	MoveIntoGroup Event = "moveintogroup"

	// MoveOutOfGroup is emitted when a window is removed from a group.
	// Args: WINDOWADDRESS
	MoveOutOfGroup Event = "moveoutofgroup"

	// IgnoreGroupLock is emitted when the ignoregrouplock setting is toggled.
	// Args: 0/1
	IgnoreGroupLock Event = "ignoregrouplock"

	// LockGroups is emitted when lockgroups is toggled.
	// Args: 0/1
	LockGroups Event = "lockgroups"

	// ConfigReloaded is emitted when the config finishes reloading.
	// Args: empty
	ConfigReloaded Event = "configreloaded"

	// Pin is emitted when a window is pinned or unpinned.
	// Args: WINDOWADDRESS, PINSTATE
	Pin Event = "pin"

	// Minimized is emitted when an external taskbar-like app requests
	// minimizing a window.
	// Args: WINDOWADDRESS, 0/1
	Minimized Event = "minimized"

	// Bell is emitted when an app rings the system bell via xdg-system-bell-v1.
	// Window address parameter may be empty.
	// Args: WINDOWADDRESS
	Bell Event = "bell"
)

// IsKnown returns if event is a known Hyprland event.
func (e Event) IsKnown() bool {
	_, ok := allEvents[e]
	return ok
}

var allEvents = map[Event]struct{}{
	Workspace: none, WorkspaceV2: none, FocusedMon: none, FocusedMonV2: none,
	ActiveWindow: none, ActiveWindowV2: none, Fullscreen: none,
	MonitorRemoved: none, MonitorRemovedV2: none, MonitorAdded: none,
	MonitorAddedV2: none, CreateWorkspace: none, CreateWorkspaceV2: none,
	DestroyWorkspace: none, DestroyWorkspaceV2: none, MoveWorkspace: none,
	MoveWorkspaceV2: none, RenameWorkspace: none, ActiveSpecial: none,
	ActiveSpecialV2: none, ActiveLayout: none, OpenWindow: none,
	CloseWindow: none, MoveWindow: none, MoveWindowV2: none, OpenLayer: none,
	CloseLayer: none, Submap: none, ChangeFloatingMode: none, Urgent: none,
	Screencast: none, WindowTitle: none, WindowTitleV2: none, ToggleGroup: none,
	MoveIntoGroup: none, MoveOutOfGroup: none, IgnoreGroupLock: none,
	LockGroups: none, ConfigReloaded: none, Pin: none, Minimized: none,
	Bell: none,
}

type (
	// OnAllEventsFunc is called on every event emitted by Hyprland.
	OnAllEventsFunc func(ctx *EventContext)
	// OnWorkspaceFunc is called when a user requests a workspace change. name
	// is the name of the workspace being switched to.
	OnWorkspaceFunc func(ctx *EventContext, name string)
	// OnWorkspaceV2Func is called when a user requests a workspace change. id
	// is the workspace ID and name is the workspace name.
	OnWorkspaceV2Func func(ctx *EventContext, id int, name string)
	// OnFocusedMonFunc is called when the active monitor changes. monitor is
	// the monitor name and workspace is the name of the workspace on that
	// monitor.
	OnFocusedMonFunc func(ctx *EventContext, monitor, workspace string)
	// OnFocusedMonV2Func is called when the active monitor changes. monitor is
	// the monitor name and workspace is the ID of the workspace on that
	// monitor.
	OnFocusedMonV2Func func(ctx *EventContext, monitor string, workspace int)
	// OnActiveWindowFunc is called when the active window changes. class is the
	// window class and title is the window title.
	OnActiveWindowFunc func(ctx *EventContext, class, title string)
	// OnActiveWindowV2Func is called when the active window changes. address is
	// the window address.
	OnActiveWindowV2Func func(ctx *EventContext, address string)
	// OnFullscreenFunc is called when a window enters or exits fullscreen mode.
	// fullscreen is true when entering fullscreen, false when exiting.
	OnFullscreenFunc func(ctx *EventContext, fullscreen bool)
	// OnMonitorRemovedFunc is called when a monitor is disconnected. name is
	// the name of the removed monitor.
	OnMonitorRemovedFunc func(ctx *EventContext, name string)
	// OnMonitorRemovedV2Func is called when a monitor is disconnected. id is
	// the monitor ID, name is the monitor name, and description is the monitor
	// description.
	OnMonitorRemovedV2Func func(ctx *EventContext, id int, name, description string)
	// OnMonitorAddedFunc is called when a monitor is connected. name is the
	// name of the added monitor.
	OnMonitorAddedFunc func(ctx *EventContext, name string)
	// OnMonitorAddedV2Func is called when a monitor is connected. id is the
	// monitor ID, name is the monitor name, and description is the monitor
	// description.
	OnMonitorAddedV2Func func(ctx *EventContext, id int, name, description string)
	// OnCreateWorkspaceFunc is called when a workspace is created. name is the
	// name of the created workspace.
	OnCreateWorkspaceFunc func(ctx *EventContext, name string)
	// OnCreateWorkspaceV2Func is called when a workspace is created. id is the
	// workspace ID and name is the workspace name.
	OnCreateWorkspaceV2Func func(ctx *EventContext, id int, name string)
	// OnDestroyWorkspaceFunc is called when a workspace is destroyed. name is
	// the name of the destroyed workspace.
	OnDestroyWorkspaceFunc func(ctx *EventContext, name string)
	// OnDestroyWorkspaceV2Func is called when a workspace is destroyed. id is
	// the workspace ID and name is the workspace name.
	OnDestroyWorkspaceV2Func func(ctx *EventContext, id int, name string)
	// OnMoveWorkspaceFunc is called when a workspace is moved to a different
	// monitor. name is the workspace name and monitor is the monitor name.
	OnMoveWorkspaceFunc func(ctx *EventContext, name, monitor string)
	// OnMoveWorkspaceV2Func is called when a workspace is moved to a different
	// monitor. id is the workspace ID, name is the workspace name, and monitor
	// is the monitor name.
	OnMoveWorkspaceV2Func func(ctx *EventContext, id int, name, monitor string)
	// OnRenameWorkspaceFunc is called when a workspace is renamed. id is the
	// workspace ID and newName is the new name of the workspace.
	OnRenameWorkspaceFunc func(ctx *EventContext, id int, newName string)
	// OnActiveSpecialFunc is called when the special workspace opened on a
	// monitor changes. name is the workspace name (empty when closing) and
	// monitor is the monitor name.
	OnActiveSpecialFunc func(ctx *EventContext, name, monitor string)
	// OnActiveSpecialV2Func is called when the special workspace opened on a
	// monitor changes. id is the workspace ID (empty when closing), name is the
	// workspace name (empty when closing), and monitor is the monitor name.
	OnActiveSpecialV2Func func(ctx *EventContext, id int, name, monitor string)
	// OnActiveLayoutFunc is called when the active keyboard layout changes.
	// keyboard is the keyboard name and layout is the layout name.
	OnActiveLayoutFunc func(ctx *EventContext, keyboard, layout string)
	// OnOpenWindowFunc is called when a window is opened. address is the window
	// address, workspace is the workspace name, class is the window class, and
	// title is the window title.
	OnOpenWindowFunc func(ctx *EventContext, address string, workspace, class, title string)
	// OnCloseWindowFunc is called when a window is closed. address is the
	// window address.
	OnCloseWindowFunc func(ctx *EventContext, address string)
	// OnMoveWindowFunc is called when a window is moved to a different
	// workspace. address is the window address and workspace is the workspace
	// name.
	OnMoveWindowFunc func(ctx *EventContext, address string, workspace string)
	// OnMoveWindowV2Func is called when a window is moved to a different
	// workspace. address is the window address, workspaceID is the workspace
	// ID, and workspace is the workspace name.
	OnMoveWindowV2Func func(ctx *EventContext, address string, workspaceID int, workspace string)
	// OnOpenLayerFunc is called when a layer surface is mapped. namespace is
	// the namespace of the layer surface.
	OnOpenLayerFunc func(ctx *EventContext, namespace string)
	// OnCloseLayerFunc is called when a layer surface is unmapped. namespace is
	// the namespace of the layer surface.
	OnCloseLayerFunc func(ctx *EventContext, namespace string)
	// OnSubmapFunc is called when a keybind submap changes. name is the submap
	// name (empty string indicates the default submap).
	OnSubmapFunc func(ctx *EventContext, name string)
	// OnChangeFloatingModeFunc is called when a window changes its floating
	// mode. address is the window address and floating is true if the window is
	// now floating, false otherwise.
	OnChangeFloatingModeFunc func(ctx *EventContext, address string, floating bool)
	// OnUrgentFunc is called when a window requests an urgent state. address is
	// the window address.
	OnUrgentFunc func(ctx *EventContext, address string)
	// OnScreencastFunc is called when a screencopy state of a client changes.
	// state is true for screencopy starting, false for stopping. owner is true
	// for window share, false for monitor share. Note: Multiple separate
	// clients may trigger this event independently.
	OnScreencastFunc func(ctx *EventContext, state bool, owner bool)
	// OnWindowTitleFunc is called when a window title changes. address is the
	// window address.
	OnWindowTitleFunc func(ctx *EventContext, address string)
	// OnWindowTitleV2Func is called when a window title changes. address is the
	// window address and title is the new window title.
	OnWindowTitleV2Func func(ctx *EventContext, address string, title string)
	// OnToggleGroupFunc is called when the togglegroup command is used. state
	// is true if a group was created, false if a group was destroyed. addresses
	// is a slice of window addresses in the group.
	OnToggleGroupFunc func(ctx *EventContext, state bool, addresses []string)
	// OnMoveIntoGroupFunc is called when a window is merged into a group.
	// address is the address of the window that was moved into the group.
	OnMoveIntoGroupFunc func(ctx *EventContext, address string)
	// OnMoveOutOfGroupFunc is called when a window is removed from a group.
	// address is the address of the window that was removed from the group.
	OnMoveOutOfGroupFunc func(ctx *EventContext, address string)
	// OnIgnoreGroupLockFunc is called when ignoregrouplock is toggled. state is
	// true if ignore group lock is enabled, false if disabled.
	OnIgnoreGroupLockFunc func(ctx *EventContext, state bool)
	// OnLockGroupsFunc is called when lockgroups is toggled. state is true if
	// group locking is enabled, false if disabled.
	OnLockGroupsFunc func(ctx *EventContext, state bool)
	// OnConfigReloadedFunc is called when the Hyprland config has finished
	// reloading.
	OnConfigReloadedFunc func(ctx *EventContext)
	// OnPinFunc is called when a window is pinned or unpinned. address is the
	// window address and pinned is true if the window is now pinned, false if
	// unpinned.
	OnPinFunc func(ctx *EventContext, address string, pinned bool)
	// OnMinimizedFunc is called when an external taskbar-like app requests a
	// window to be minimized. address is the window address and minimized is
	// true if the window should be minimized, false if restored.
	OnMinimizedFunc func(ctx *EventContext, address string, minimized bool)
	// OnBellFunc is called when an app requests to ring the system bell via
	// xdg-system-bell-v1. address is the window address (may be empty).
	OnBellFunc func(ctx *EventContext, address string)
	// OnUnknownFunc is called for any event that does not have a proper
	// binding. This can occur either from events emitted by a plugin or from
	// new Hyprland events that have not yet been implemented in this handler.
	OnUnknownFunc func(ctx *EventContext)
)

// EventHandler is the interface for handling all Hyprland events. Each method
// corresponds to a specific event type emitted by the Hyprland compositor.
type EventHandler interface {
	// All is called for every event emitted by Hyprland.
	All(ctx *EventContext)
	// Workspace is called when a user requests a workspace change. name is the
	// name of the workspace being switched to.
	Workspace(ctx *EventContext, name string)
	// WorkspaceV2 is called when a user requests a workspace change. id is the
	// workspace ID and name is the workspace name.
	WorkspaceV2(ctx *EventContext, id int, name string)
	// FocusedMon is called when the active monitor changes. monitor is the
	// monitor name and workspaceName is the name of the workspace on that
	// monitor.
	FocusedMon(ctx *EventContext, monitor, workspaceName string)
	// FocusedMonV2 is called when the active monitor changes. monitor is the
	// monitor name and workspaceID is the ID of the workspace on that monitor.
	FocusedMonV2(ctx *EventContext, monitor string, workspaceID int)
	// ActiveWindow is called when the active window changes. class is the
	// window class and title is the window title.
	ActiveWindow(ctx *EventContext, class, title string)
	// ActiveWindowV2 is called when the active window changes. address is the
	// window address.
	ActiveWindowV2(ctx *EventContext, address string)
	// Fullscreen is called when a window enters or exits fullscreen mode.
	// fullscreen is true when entering fullscreen, false when exiting. Note: A
	// fullscreen event is not guaranteed to fire in a strict on/off succession,
	// as some windows may fire multiple requests to be fullscreened.
	Fullscreen(ctx *EventContext, fullscreen bool)
	// MonitorRemoved is called when a monitor is disconnected. name is the name
	// of the removed monitor.
	MonitorRemoved(ctx *EventContext, name string)
	// MonitorRemovedV2 is called when a monitor is disconnected. id is the
	// monitor ID, name is the monitor name, and description is the monitor
	// description.
	MonitorRemovedV2(ctx *EventContext, id int, name, description string)
	// MonitorAdded is called when a monitor is connected. name is the name of
	// the added monitor.
	MonitorAdded(ctx *EventContext, name string)
	// MonitorAddedV2 is called when a monitor is connected. id is the monitor
	// ID, name is the monitor name, and description is the monitor description.
	MonitorAddedV2(ctx *EventContext, id int, name, description string)
	// CreateWorkspace is called when a workspace is created. name is the name
	// of the created workspace.
	CreateWorkspace(ctx *EventContext, name string)
	// CreateWorkspaceV2 is called when a workspace is created. id is the
	// workspace ID and name is the workspace name.
	CreateWorkspaceV2(ctx *EventContext, id int, name string)
	// DestroyWorkspace is called when a workspace is destroyed. name is the
	// name of the destroyed workspace.
	DestroyWorkspace(ctx *EventContext, name string)
	// DestroyWorkspaceV2 is called when a workspace is destroyed. id is the
	// workspace ID and name is the workspace name.
	DestroyWorkspaceV2(ctx *EventContext, id int, name string)
	// MoveWorkspace is called when a workspace is moved to a different monitor.
	// name is the workspace name and monitor is the monitor name.
	MoveWorkspace(ctx *EventContext, name, monitor string)
	// MoveWorkspaceV2 is called when a workspace is moved to a different
	// monitor. id is the workspace ID, name is the workspace name, and monitor
	// is the monitor name.
	MoveWorkspaceV2(ctx *EventContext, id int, name, monitor string)
	// RenameWorkspace is called when a workspace is renamed. id is the
	// workspace ID and newName is the new name of the workspace.
	RenameWorkspace(ctx *EventContext, id int, newName string)
	// ActiveSpecial is called when the special workspace opened on a monitor
	// changes. name is the workspace name (empty when closing) and monitor is
	// the monitor name.
	ActiveSpecial(ctx *EventContext, name, monitor string)
	// ActiveSpecialV2 is called when the special workspace opened on a monitor
	// changes. id is the workspace ID (empty values when closing), name is the
	// workspace name (empty when closing), and monitor is the monitor name.
	ActiveSpecialV2(ctx *EventContext, id int, name, monitor string)
	// ActiveLayout is called when the active keyboard layout changes. keyboard
	// is the keyboard name and layout is the layout name.
	ActiveLayout(ctx *EventContext, keyboard, layout string)
	// OpenWindow is called when a window is opened. address is the window
	// address, workspace is the workspace name, class is the window class, and
	// title is the window title.
	OpenWindow(ctx *EventContext, address, workspace, class, title string)
	// CloseWindow is called when a window is closed. address is the window
	// address.
	CloseWindow(ctx *EventContext, address string)
	// MoveWindow is called when a window is moved to a different workspace.
	// address is the window address and workspace is the workspace name.
	MoveWindow(ctx *EventContext, address string, workspace string)
	// MoveWindowV2 is called when a window is moved to a different workspace.
	// address is the window address, workspaceID is the workspace ID, and
	// workspace is the workspace name.
	MoveWindowV2(
		ctx *EventContext,
		address string,
		workspaceID int,
		workspace string,
	)
	// OpenLayer is called when a layer surface is mapped. namespace is the
	// namespace of the layer shell.
	OpenLayer(ctx *EventContext, namespace string)
	// CloseLayer is called when a layer surface is unmapped. namespace is the
	// namespace of the layer shell.
	CloseLayer(ctx *EventContext, namespace string)
	// Submap is called when a keybind submap changes. name is the submap name
	// (empty string indicates the default submap).
	Submap(ctx *EventContext, name string)
	// ChangeFloatingMode is called when a window changes its floating mode.
	// address is the window address and floating is true if the window is now
	// floating, false otherwise.
	ChangeFloatingMode(ctx *EventContext, address string, floating bool)
	// Urgent is called when a window requests an urgent state. address is the
	// window address.
	Urgent(ctx *EventContext, address string)
	// Screencast is called when a screencopy state of a client changes. state
	// is true for screencopy starting, false for stopping. owner is true for
	// window share, false for monitor share. Note: Multiple separate clients
	// may trigger this event independently.
	Screencast(ctx *EventContext, state bool, owner bool)
	// WindowTitle is called when a window title changes. address is the window
	// address.
	WindowTitle(ctx *EventContext, address string)
	// WindowTitleV2 is called when a window title changes. address is the
	// window address and title is the new window title.
	WindowTitleV2(ctx *EventContext, address string, title string)
	// ToggleGroup is called when the togglegroup command is used.
	// state is true if a group was created, false if a group was destroyed.
	// addresses is a slice of window addresses in the group.
	ToggleGroup(ctx *EventContext, state bool, addresses []string)
	// MoveIntoGroup is called when a window is merged into a group. address is
	// the address of the window that was moved into the group.
	MoveIntoGroup(ctx *EventContext, address string)
	// MoveOutOfGroup is called when a window is removed from a group. address
	// is the address of the window that was removed from the group.
	MoveOutOfGroup(ctx *EventContext, address string)
	// IgnoreGroupLock is called when ignoregrouplock is toggled. state is true
	// if ignore group lock is enabled, false if disabled.
	IgnoreGroupLock(ctx *EventContext, state bool)
	// LockGroups is called when lockgroups is toggled. state is true if group
	// locking is enabled, false if disabled.
	LockGroups(ctx *EventContext, state bool)
	// ConfigReloaded is called when the Hyprland config has finished reloading.
	ConfigReloaded(ctx *EventContext)
	// Pin is called when a window is pinned or unpinned. address is the window
	// address and pinned is true if the window is now
	// pinned, false if unpinned.
	Pin(ctx *EventContext, address string, pinned bool)
	// Minimized is called when an external taskbar-like app requests a window
	// to be minimized. address is the window address and minimized is true if
	// the window should be minimized, false if restored.
	Minimized(ctx *EventContext, address string, minimized bool)
	// Bell is called when an app requests to ring the system bell via
	// xdg-system-bell-v1. address is the window address (may be empty).
	Bell(ctx *EventContext, address string)
	// Unknown is called for any event that does not have a proper binding. This
	// can occur either from events emitted by a plugin or from new Hyprland
	// events that have not yet been implemented in this handler.
	Unknown(ctx *EventContext)
}
