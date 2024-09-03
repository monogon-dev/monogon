package tconsole

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gdamore/tcell/v2"

	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/roleserve"
	cpb "source.monogon.dev/metropolis/proto/common"
	"source.monogon.dev/osbase/event"
	"source.monogon.dev/osbase/supervisor"
)

// Console is a Terminal Console (TConsole), a user-interactive informational
// display visible on the TTY of a running Metropolis instance.
type Console struct {
	// Quit will be closed when the user press CTRL-C. A new channel will be created
	// on each New call.

	Quit    chan struct{}
	ttyPath string
	tty     tcell.Tty
	screen  tcell.Screen
	width   int
	height  int
	// palette chosen for the given terminal.
	palette palette

	network     event.Value[*network.Status]
	roles       event.Value[*cpb.NodeRoles]
	curatorConn event.Value[*roleserve.CuratorConnection]
}

// New creates a new Console, taking over the TTY at the given path. The given
// Terminal type selects between a Linux terminal (VTY) and a generic terminal
// for testing.
//
// network, roles, curatorConn point to various Metropolis subsystems that are
// used to populate the console data.
func New(terminal Terminal, ttyPath string, network event.Value[*network.Status], roles event.Value[*cpb.NodeRoles], curatorConn event.Value[*roleserve.CuratorConnection]) (*Console, error) {
	tty, err := tcell.NewDevTtyFromDev(ttyPath)
	if err != nil {
		return nil, err
	}
	screen, err := tcell.NewTerminfoScreenFromTty(tty)
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}
	screen.SetStyle(tcell.StyleDefault)

	var pal palette
	switch terminal {
	case TerminalLinux:
		pal = paletteLinux
		tty.Write([]byte(pal.setupLinuxConsole()))
	case TerminalGeneric:
		pal = paletteGeneric
	}

	width, height := screen.Size()
	return &Console{
		ttyPath: ttyPath,
		tty:     tty,
		screen:  screen,
		width:   width,
		height:  height,
		network: network,
		palette: pal,
		Quit:    make(chan struct{}),

		roles:       roles,
		curatorConn: curatorConn,
	}, nil
}

// Cleanup should be called when the console exits. This is only used in testing,
// the Metropolis console always runs.
func (c *Console) Cleanup() {
	c.screen.Fini()
}

func (c *Console) processEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyCtrlC {
			close(c.Quit)
		}
	case *tcell.EventResize:
		c.width, c.height = ev.Size()
	}
}

// Run blocks while displaying the Console to the user.
func (c *Console) Run(ctx context.Context) error {
	// Build channel for console event processing.
	evC := make(chan tcell.Event)
	evQuitC := make(chan struct{})
	defer close(evQuitC)
	go c.screen.ChannelEvents(evC, evQuitC)

	// Pipe event values into channels.
	netAddrC := make(chan *network.Status)
	rolesC := make(chan *cpb.NodeRoles)
	curatorConnC := make(chan *roleserve.CuratorConnection)
	if err := supervisor.Run(ctx, "netpipe", event.Pipe(c.network, netAddrC)); err != nil {
		return err
	}
	if err := supervisor.Run(ctx, "rolespipe", event.Pipe(c.roles, rolesC)); err != nil {
		return err
	}
	if err := supervisor.Run(ctx, "curatorpipe", event.Pipe(c.curatorConn, curatorConnC)); err != nil {
		return err
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	pageStatus := pageStatusData{
		netAddr:     "Waiting...",
		roles:       "Waiting...",
		id:          "Waiting...",
		fingerprint: "Waiting...",
	}

	for {
		c.pageStatus(&pageStatus)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case ev := <-evC:
			c.processEvent(ev)
		case t := <-netAddrC:
			pageStatus.netAddr = t.ExternalAddress.String()
		case t := <-rolesC:
			var rlist []string
			if t.ConsensusMember != nil {
				rlist = append(rlist, "ConsensusMember")
			}
			if t.KubernetesController != nil {
				rlist = append(rlist, "KubernetesController")
			}
			if t.KubernetesWorker != nil {
				rlist = append(rlist, "KubernetesWorker")
			}
			pageStatus.roles = strings.Join(rlist, ", ")
			if pageStatus.roles == "" {
				pageStatus.roles = "none"
			}
		case t := <-curatorConnC:
			pageStatus.id = t.Credentials.ID()
			cert := t.Credentials.ClusterCA()
			sum := sha256.New()
			sum.Write(cert.Raw)
			pageStatus.fingerprint = hex.EncodeToString(sum.Sum(nil))
		}
	}
}
