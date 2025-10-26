package hyprland

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// EventContext is context.Context with event metadata
type EventContext struct {
	context.Context // default context implementation
	// RawEvent is the line sent by hyprland. Format: [Event]>>[RawData].
	RawEvent string
	// Event is the name of the event
	Event Event
	// RawData is is the [RawData] part of '[Event]>>[RawData]'.
	RawData string
	// Time is the time when when event is emitted
	Time time.Time
}

// EventSeparator is the hyprland event separator
const EventSeparator string = ">>"

// ParseEvent parses raw event line into strcutered data
func ParseEvent(raw string) (*EventContext, error) {
	ctx := new(EventContext)
	ctx.Time = time.Now()
	ctx.RawEvent = raw
	event, data, found := strings.Cut(raw, EventSeparator)
	if !found {
		return ctx, fmt.Errorf("invalid event format: %q", raw)
	}
	ctx.Event = Event(event)
	ctx.RawData = data
	return ctx, nil
}

// Just for fun
var none = struct{}{}

// EventListener is the high-level interface for working with hyprland events
type EventListener struct {
	// socket is the hyprland socket path
	Socket SocketPath
	// handler is the handler for all events
	handler EventHandler
	// conn is socket connection
	conn net.Conn
	// subscribed is a map of
	subscribed map[Event]struct{}
	// mu for sync safety
	mu sync.Mutex

	// on handlers
	onAllEvents          OnAllEventsFunc
	onWorkspace          OnWorkspaceFunc
	onWorkspaceV2        OnWorkspaceV2Func
	onFocusedMon         OnFocusedMonFunc
	onFocusedMonV2       OnFocusedMonV2Func
	onActiveWindow       OnActiveWindowFunc
	onActiveWindowV2     OnActiveWindowV2Func
	onFullscreen         OnFullscreenFunc
	onMonitorRemoved     OnMonitorRemovedFunc
	onMonitorRemovedV2   OnMonitorRemovedV2Func
	onMonitorAdded       OnMonitorAddedFunc
	onMonitorAddedV2     OnMonitorAddedV2Func
	onCreateWorkspace    OnCreateWorkspaceFunc
	onCreateWorkspaceV2  OnCreateWorkspaceV2Func
	onDestroyWorkspace   OnDestroyWorkspaceFunc
	onDestroyWorkspaceV2 OnDestroyWorkspaceV2Func
	onMoveWorkspace      OnMoveWorkspaceFunc
	onMoveWorkspaceV2    OnMoveWorkspaceV2Func
	onRenameWorkspace    OnRenameWorkspaceFunc
	onActiveSpecial      OnActiveSpecialFunc
	onActiveSpecialV2    OnActiveSpecialV2Func
	onActiveLayout       OnActiveLayoutFunc
	onOpenWindow         OnOpenWindowFunc
	onCloseWindow        OnCloseWindowFunc
	onMoveWindow         OnMoveWindowFunc
	onMoveWindowV2       OnMoveWindowV2Func
	onOpenLayer          OnOpenLayerFunc
	onCloseLayer         OnCloseLayerFunc
	onSubmap             OnSubmapFunc
	onChangeFloatingMode OnChangeFloatingModeFunc
	onUrgent             OnUrgentFunc
	onScreencast         OnScreencastFunc
	onWindowTitle        OnWindowTitleFunc
	onWindowTitleV2      OnWindowTitleV2Func
	onToggleGroup        OnToggleGroupFunc
	onMoveIntoGroup      OnMoveIntoGroupFunc
	onMoveOutOfGroup     OnMoveOutOfGroupFunc
	onIgnoreGroupLock    OnIgnoreGroupLockFunc
	onLockGroups         OnLockGroupsFunc
	onConfigReloaded     OnConfigReloadedFunc
	onPin                OnPinFunc
	onMinimized          OnMinimizedFunc
	onBell               OnBellFunc
	onUnknown            OnUnknownFunc
}

// NewEventListener creates a new EventListener.
func NewEventListener() *EventListener {
	l := new(EventListener)
	l.Socket = should(GetEventSocket())
	l.subscribed = map[Event]struct{}{}
	return l
}

// HasHandler returns if given event as a handler.
func (l *EventListener) HasHandler(event Event) bool {
	if l.onAllEvents != nil && l.handler != nil {
		return true
	}

	if event.IsKnown() {
		_, ok := l.subscribed[event]
		return ok
	}

	return l.onUnknown != nil
}

// IsConnected returns if the listener is connected to hyprland socket.
func (l *EventListener) IsConnected() bool {
	return l.conn != nil
}

// Close closes the underlying socket connection
func (l *EventListener) Close() error {
	if l.conn != nil {
		return l.conn.Close()
	}
	return nil
}

// SetHandler sets the event handler
func (l *EventListener) SetHandler(handler EventHandler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handler = handler
}

// OnAllEvents sets the handler for all events
func (l *EventListener) OnAllEvents(fn OnAllEventsFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.onAllEvents = fn
}

// OnWorkspace sets the handler for Workspace events
func (l *EventListener) OnWorkspace(fn OnWorkspaceFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventWorkspace] = none
	l.onWorkspace = fn
}

// OnWorkspaceV2 sets the handler for WorkspaceV2 events
func (l *EventListener) OnWorkspaceV2(fn OnWorkspaceV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventWorkspace] = none
	l.onWorkspaceV2 = fn
}

// OnFocusedMon sets the handler for FocusedMon events
func (l *EventListener) OnFocusedMon(fn OnFocusedMonFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventFocusedMonitor] = none
	l.onFocusedMon = fn
}

// OnFocusedMonV2 sets the handler for FocusedMonV2 events
func (l *EventListener) OnFocusedMonV2(fn OnFocusedMonV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventFocusedMonitorV2] = none
	l.onFocusedMonV2 = fn
}

// OnActiveWindow sets the handler for ActiveWindow events
func (l *EventListener) OnActiveWindow(fn OnActiveWindowFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventActiveWindow] = none
	l.onActiveWindow = fn
}

// OnActiveWindowV2 sets the handler for ActiveWindowV2 events
func (l *EventListener) OnActiveWindowV2(fn OnActiveWindowV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventActiveWindowV2] = none
	l.onActiveWindowV2 = fn
}

// OnFullscreen sets the handler for Fullscreen events
func (l *EventListener) OnFullscreen(fn OnFullscreenFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventFullscreen] = none
	l.onFullscreen = fn
}

// OnMonitorRemoved sets the handler for MonitorRemoved events
func (l *EventListener) OnMonitorRemoved(fn OnMonitorRemovedFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMonitorRemoved] = none
	l.onMonitorRemoved = fn
}

// OnMonitorRemovedV2 sets the handler for MonitorRemovedV2 events
func (l *EventListener) OnMonitorRemovedV2(fn OnMonitorRemovedV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMonitorRemovedV2] = none
	l.onMonitorRemovedV2 = fn
}

// OnMonitorAdded sets the handler for MonitorAdded events
func (l *EventListener) OnMonitorAdded(fn OnMonitorAddedFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMonitorAdded] = none
	l.onMonitorAdded = fn
}

// OnMonitorAddedV2 sets the handler for MonitorAddedV2 events
func (l *EventListener) OnMonitorAddedV2(fn OnMonitorAddedV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMonitorAddedV2] = none
	l.onMonitorAddedV2 = fn
}

// OnCreateWorkspace sets the handler for CreateWorkspace events
func (l *EventListener) OnCreateWorkspace(fn OnCreateWorkspaceFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventCreateWorkspace] = none
	l.onCreateWorkspace = fn
}

// OnCreateWorkspaceV2 sets the handler for CreateWorkspaceV2 events
func (l *EventListener) OnCreateWorkspaceV2(fn OnCreateWorkspaceV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventCreateWorkspaceV2] = none
	l.onCreateWorkspaceV2 = fn
}

// OnDestroyWorkspace sets the handler for DestroyWorkspace events
func (l *EventListener) OnDestroyWorkspace(fn OnDestroyWorkspaceFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventDestroyWorkspace] = none
	l.onDestroyWorkspace = fn
}

// OnDestroyWorkspaceV2 sets the handler for DestroyWorkspaceV2 events
func (l *EventListener) OnDestroyWorkspaceV2(fn OnDestroyWorkspaceV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventDestroyWorkspaceV2] = none
	l.onDestroyWorkspaceV2 = fn
}

// OnMoveWorkspace sets the handler for MoveWorkspace events
func (l *EventListener) OnMoveWorkspace(fn OnMoveWorkspaceFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveWorkspace] = none
	l.onMoveWorkspace = fn
}

// OnMoveWorkspaceV2 sets the handler for MoveWorkspaceV2 events
func (l *EventListener) OnMoveWorkspaceV2(fn OnMoveWorkspaceV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveWorkspaceV2] = none
	l.onMoveWorkspaceV2 = fn
}

// OnRenameWorkspace sets the handler for RenameWorkspace events
func (l *EventListener) OnRenameWorkspace(fn OnRenameWorkspaceFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventRenameWorkspace] = none
	l.onRenameWorkspace = fn
}

// OnActiveSpecial sets the handler for ActiveSpecial events
func (l *EventListener) OnActiveSpecial(fn OnActiveSpecialFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventActiveSpecial] = none
	l.onActiveSpecial = fn
}

// OnActiveSpecialV2 sets the handler for ActiveSpecialV2 events
func (l *EventListener) OnActiveSpecialV2(fn OnActiveSpecialV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventActiveSpecialV2] = none
	l.onActiveSpecialV2 = fn
}

// OnActiveLayout sets the handler for ActiveLayout events
func (l *EventListener) OnActiveLayout(fn OnActiveLayoutFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventActiveLayout] = none
	l.onActiveLayout = fn
}

// OnOpenWindow sets the handler for OpenWindow events
func (l *EventListener) OnOpenWindow(fn OnOpenWindowFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventOpenWindow] = none
	l.onOpenWindow = fn
}

// OnCloseWindow sets the handler for CloseWindow events
func (l *EventListener) OnCloseWindow(fn OnCloseWindowFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventCloseWindow] = none
	l.onCloseWindow = fn
}

// OnMoveWindow sets the handler for MoveWindow events
func (l *EventListener) OnMoveWindow(fn OnMoveWindowFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveWindow] = none
	l.onMoveWindow = fn
}

// OnMoveWindowV2 sets the handler for MoveWindowV2 events
func (l *EventListener) OnMoveWindowV2(fn OnMoveWindowV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveWindowV2] = none
	l.onMoveWindowV2 = fn
}

// OnOpenLayer sets the handler for OpenLayer events
func (l *EventListener) OnOpenLayer(fn OnOpenLayerFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventOpenLayer] = none
	l.onOpenLayer = fn
}

// OnCloseLayer sets the handler for CloseLayer events
func (l *EventListener) OnCloseLayer(fn OnCloseLayerFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventCloseLayer] = none
	l.onCloseLayer = fn
}

// OnSubmap sets the handler for Submap events
func (l *EventListener) OnSubmap(fn OnSubmapFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventSubmap] = none
	l.onSubmap = fn
}

// OnChangeFloatingMode sets the handler for ChangeFloatingMode events
func (l *EventListener) OnChangeFloatingMode(fn OnChangeFloatingModeFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventChangeFloatingMode] = none
	l.onChangeFloatingMode = fn
}

// OnUrgent sets the handler for Urgent events
func (l *EventListener) OnUrgent(fn OnUrgentFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventUrgent] = none
	l.onUrgent = fn
}

// OnScreencast sets the handler for Screencast events
func (l *EventListener) OnScreencast(fn OnScreencastFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventScreencast] = none
	l.onScreencast = fn
}

// OnWindowTitle sets the handler for WindowTitle events
func (l *EventListener) OnWindowTitle(fn OnWindowTitleFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventWindowTitle] = none
	l.onWindowTitle = fn
}

// OnWindowTitleV2 sets the handler for WindowTitleV2 events
func (l *EventListener) OnWindowTitleV2(fn OnWindowTitleV2Func) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventWindowTitleV2] = none
	l.onWindowTitleV2 = fn
}

// OnToggleGroup sets the handler for ToggleGroup events
func (l *EventListener) OnToggleGroup(fn OnToggleGroupFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventToggleGroup] = none
	l.onToggleGroup = fn
}

// OnMoveIntoGroup sets the handler for MoveIntoGroup events
func (l *EventListener) OnMoveIntoGroup(fn OnMoveIntoGroupFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveIntoGroup] = none
	l.onMoveIntoGroup = fn
}

// OnMoveOutOfGroup sets the handler for MoveOutOfGroup events
func (l *EventListener) OnMoveOutOfGroup(fn OnMoveOutOfGroupFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMoveOutOfGroup] = none
	l.onMoveOutOfGroup = fn
}

// OnIgnoreGroupLock sets the handler for IgnoreGroupLock events
func (l *EventListener) OnIgnoreGroupLock(fn OnIgnoreGroupLockFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventIgnoreGroupLock] = none
	l.onIgnoreGroupLock = fn
}

// OnLockGroups sets the handler for LockGroups events
func (l *EventListener) OnLockGroups(fn OnLockGroupsFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventLockGroups] = none
	l.onLockGroups = fn
}

// OnConfigReloaded sets the handler for ConfigReloaded events
func (l *EventListener) OnConfigReloaded(fn OnConfigReloadedFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventConfigReloaded] = none
	l.onConfigReloaded = fn
}

// OnPin sets the handler for Pin events
func (l *EventListener) OnPin(fn OnPinFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventPin] = none
	l.onPin = fn
}

// OnMinimized sets the handler for Minimized events
func (l *EventListener) OnMinimized(fn OnMinimizedFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventMinimized] = none
	l.onMinimized = fn
}

// OnBell sets the handler for Bell events
func (l *EventListener) OnBell(fn OnBellFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventBell] = none
	l.onBell = fn
}

// OnUnknown sets the handler for Unknown events
func (l *EventListener) OnUnknown(fn OnUnknownFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.subscribed[EventBell] = none
	l.onUnknown = fn
}

// Listen is a dials the socket2 connection and start listening for events
// synchronously
func (l *EventListener) Listen(ctx context.Context) error {
	l.mu.Lock()
	socket, err := GetEventSocket()
	if err != nil {
		return err
	}
	l.Socket = socket

	conn, err := net.Dial("unix", string(l.Socket))
	if err != nil {
		return err
	}
	l.conn = conn
	defer func() {
		l.mu.Lock()
		l.conn = nil
		l.mu.Unlock()
	}()
	defer conn.Close()

	l.mu.Unlock()

	events := make(chan string, 10)

	scanner := bufio.NewScanner(conn)

	var wg sync.WaitGroup
	wg.Go(func() {
		defer close(events)
		for scanner.Scan() {
			events <- scanner.Text()
		}
	})

	for {
		select {
		case <-ctx.Done():
			conn.Close()
			wg.Wait()
			return ctx.Err()
		case rawData, ok := <-events:
			if !ok {
				return errors.New("event channel closed expectedly")
			}
			eventCtx, err := ParseEvent(rawData)
			if err != nil {
				return err
			}
			eventCtx.Context = ctx
			if err := l.processEvent(eventCtx); err != nil {
				return err
			}
		}
	}
}

func (l *EventListener) processEvent(ctx *EventContext) error {
	if l.onAllEvents != nil {
		l.onAllEvents(ctx)
	}
	if l.handler != nil {
		l.handler.All(ctx)
	}
	switch ctx.Event {
	case EventWorkspace:
		if l.onWorkspace != nil {
			l.onWorkspace(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.Workspace(ctx, ctx.RawData)
		}
	case EventWorkspaceV2:
		id, name, err := cast2[int, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onWorkspaceV2 != nil {
			l.onWorkspaceV2(ctx, id, name)
		}
		if l.handler != nil {
			l.handler.WorkspaceV2(ctx, id, name)
		}
	case EventFocusedMonitor:
		mon, name, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onFocusedMon != nil {
			l.onFocusedMon(ctx, mon, name)
		}
		if l.handler != nil {
			l.handler.FocusedMon(ctx, mon, name)
		}
	case EventFocusedMonitorV2:
		mon, id, err := cast2[string, int](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onFocusedMonV2 != nil {
			l.onFocusedMonV2(ctx, mon, id)
		}
		if l.handler != nil {
			l.handler.FocusedMonV2(ctx, mon, id)
		}
	case EventActiveWindow:
		class, title, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onActiveWindow != nil {
			l.onActiveWindow(ctx, class, title)
		}
		if l.handler != nil {
			l.handler.ActiveWindow(ctx, class, title)
		}
	case EventActiveWindowV2:
		if l.onActiveWindowV2 != nil {
			l.onActiveWindowV2(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.ActiveWindowV2(ctx, ctx.RawData)
		}
	case EventFullscreen:
		fullscreen, err := cast[bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onFullscreen != nil {
			l.onFullscreen(ctx, fullscreen)
		}
		if l.handler != nil {
			l.handler.Fullscreen(ctx, fullscreen)
		}
	case EventMonitorRemoved:
		if l.onMonitorRemoved != nil {
			l.onMonitorRemoved(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.MonitorRemoved(ctx, ctx.RawData)
		}
	case EventMonitorRemovedV2:
		id, name, desc, err := cast3[int, string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMonitorRemovedV2 != nil {
			l.onMonitorRemovedV2(ctx, id, name, desc)
		}
		if l.handler != nil {
			l.handler.MonitorRemovedV2(ctx, id, name, desc)
		}
	case EventMonitorAdded:
		if l.onMonitorAdded != nil {
			l.onMonitorAdded(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.MonitorAdded(ctx, ctx.RawData)
		}
	case EventMonitorAddedV2:
		id, name, desc, err := cast3[int, string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMonitorAddedV2 != nil {
			l.onMonitorAddedV2(ctx, id, name, desc)
		}
		if l.handler != nil {
			l.handler.MonitorAddedV2(ctx, id, name, desc)
		}
	case EventCreateWorkspace:
		if l.onCreateWorkspace != nil {
			l.onCreateWorkspace(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.CreateWorkspace(ctx, ctx.RawData)
		}
	case EventCreateWorkspaceV2:
		id, name, err := cast2[int, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onCreateWorkspaceV2 != nil {
			l.onCreateWorkspaceV2(ctx, id, name)
		}
		if l.handler != nil {
			l.handler.CreateWorkspaceV2(ctx, id, name)
		}
	case EventDestroyWorkspace:
		if l.onDestroyWorkspace != nil {
			l.onDestroyWorkspace(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.DestroyWorkspace(ctx, ctx.RawData)
		}
	case EventDestroyWorkspaceV2:
		id, name, err := cast2[int, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onDestroyWorkspaceV2 != nil {
			l.onDestroyWorkspaceV2(ctx, id, name)
		}
		if l.handler != nil {
			l.handler.DestroyWorkspaceV2(ctx, id, name)
		}
	case EventMoveWorkspace:
		name, mon, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMoveWorkspace != nil {
			l.onMoveWorkspace(ctx, name, mon)
		}
		if l.handler != nil {
			l.handler.MoveWorkspace(ctx, name, mon)
		}
	case EventMoveWorkspaceV2:
		id, name, mon, err := cast3[int, string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMoveWorkspaceV2 != nil {
			l.onMoveWorkspaceV2(ctx, id, name, mon)
		}
		if l.handler != nil {
			l.handler.MoveWorkspaceV2(ctx, id, name, mon)
		}
	case EventRenameWorkspace:
		id, name, err := cast2[int, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onRenameWorkspace != nil {
			l.onRenameWorkspace(ctx, id, name)
		}
		if l.handler != nil {
			l.handler.RenameWorkspace(ctx, id, name)
		}
	case EventActiveSpecial:
		name, mon, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onActiveSpecial != nil {
			l.onActiveSpecial(ctx, name, mon)
		}
		if l.handler != nil {
			l.handler.ActiveSpecial(ctx, name, mon)
		}
	case EventActiveSpecialV2:
		id, name, mon, err := cast3[int, string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onActiveSpecialV2 != nil {
			l.onActiveSpecialV2(ctx, id, name, mon)
		}
		if l.handler != nil {
			l.handler.ActiveSpecialV2(ctx, id, name, mon)
		}
	case EventActiveLayout:
		kb, layout, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onActiveLayout != nil {
			l.onActiveLayout(ctx, kb, layout)
		}
		if l.handler != nil {
			l.handler.ActiveLayout(ctx, kb, layout)
		}
	case EventOpenWindow:
		addr, workspace, class, title, err := cast4[string, string, string, string](
			ctx.RawData,
		)
		if err != nil {
			return err
		}
		if l.onOpenWindow != nil {
			l.onOpenWindow(ctx, addr, workspace, class, title)
		}
		if l.handler != nil {
			l.handler.OpenWindow(ctx, addr, workspace, class, title)
		}
	case EventCloseWindow:
		if l.onCloseWindow != nil {
			l.onCloseWindow(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.CloseWindow(ctx, ctx.RawData)
		}
	case EventMoveWindow:
		addr, workspace, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMoveWindow != nil {
			l.onMoveWindow(ctx, addr, workspace)
		}
		if l.handler != nil {
			l.handler.MoveWindow(ctx, addr, workspace)
		}
	case EventMoveWindowV2:
		addr, id, workspace, err := cast3[string, int, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMoveWindowV2 != nil {
			l.onMoveWindowV2(ctx, addr, id, workspace)
		}
		if l.handler != nil {
			l.handler.MoveWindowV2(ctx, addr, id, workspace)
		}
	case EventOpenLayer:
		if l.onOpenLayer != nil {
			l.onOpenLayer(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.OpenLayer(ctx, ctx.RawData)
		}
	case EventCloseLayer:
		if l.onCloseLayer != nil {
			l.onCloseLayer(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.CloseLayer(ctx, ctx.RawData)
		}
	case EventSubmap:
		if l.onSubmap != nil {
			l.onSubmap(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.Submap(ctx, ctx.RawData)
		}
	case EventChangeFloatingMode:
		addr, floating, err := cast2[string, bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onChangeFloatingMode != nil {
			l.onChangeFloatingMode(ctx, addr, floating)
		}
		if l.handler != nil {
			l.handler.ChangeFloatingMode(ctx, addr, floating)
		}
	case EventUrgent:
		if l.onUrgent != nil {
			l.onUrgent(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.Urgent(ctx, ctx.RawData)
		}
	case EventScreencast:
		state, owner, err := cast2[bool, bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onScreencast != nil {
			l.onScreencast(ctx, state, owner)
		}
		if l.handler != nil {
			l.handler.Screencast(ctx, state, owner)
		}
	case EventWindowTitle:
		if l.onWindowTitle != nil {
			l.onWindowTitle(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.WindowTitle(ctx, ctx.RawData)
		}
	case EventWindowTitleV2:
		addr, title, err := cast2[string, string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onWindowTitleV2 != nil {
			l.onWindowTitleV2(ctx, addr, title)
		}
		if l.handler != nil {
			l.handler.WindowTitleV2(ctx, addr, title)
		}
	case EventToggleGroup:
		parts := strings.Split(ctx.RawData, ",")
		state, err := strconv.ParseBool(parts[0])
		if err != nil {
			return err
		}
		if l.onToggleGroup != nil {
			l.onToggleGroup(ctx, state, parts[1:])
		}
		if l.handler != nil {
			l.handler.ToggleGroup(ctx, state, parts[1:])
		}
	case EventMoveIntoGroup:
		if l.onMoveIntoGroup != nil {
			l.onMoveIntoGroup(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.MoveIntoGroup(ctx, ctx.RawData)
		}
	case EventMoveOutOfGroup:
		if l.onMoveOutOfGroup != nil {
			l.onMoveOutOfGroup(ctx, ctx.RawData)
		}
		if l.handler != nil {
			l.handler.MoveOutOfGroup(ctx, ctx.RawData)
		}
	case EventIgnoreGroupLock:
		state, err := cast[bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onIgnoreGroupLock != nil {
			l.onIgnoreGroupLock(ctx, state)
		}
		if l.handler != nil {
			l.handler.IgnoreGroupLock(ctx, state)
		}
	case EventLockGroups:
		state, err := cast[bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onLockGroups != nil {
			l.onLockGroups(ctx, state)
		}
		if l.handler != nil {
			l.handler.LockGroups(ctx, state)
		}
	case EventConfigReloaded:
		if l.onIgnoreGroupLock != nil {
			l.onConfigReloaded(ctx)
		}
		if l.handler != nil {
			l.handler.ConfigReloaded(ctx)
		}
	case EventPin:
		addr, state, err := cast2[string, bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onPin != nil {
			l.onPin(ctx, addr, state)
		}
		if l.handler != nil {
			l.handler.Pin(ctx, addr, state)
		}
	case EventMinimized:
		addr, state, err := cast2[string, bool](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onMinimized != nil {
			l.onMinimized(ctx, addr, state)
		}
		if l.handler != nil {
			l.handler.Minimized(ctx, addr, state)
		}
	case EventBell:
		addr, err := cast[string](ctx.RawData)
		if err != nil {
			return err
		}
		if l.onBell != nil {
			l.onBell(ctx, addr)
		}
		if l.handler != nil {
			l.handler.Bell(ctx, addr)
		}
	default:
		if l.onUnknown != nil {
			l.onUnknown(ctx)
		}
		if l.handler != nil {
			l.handler.Unknown(ctx)
		}
	}
	return nil
}
