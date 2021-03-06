package worldserver

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/gophercraft/packet/update"
	"github.com/superp00t/gophercraft/worldserver/wdb"
)

func x_Tele(c *C) {
	// port string
	if len(c.Args) < 5 && len(c.Args) != 1 {
		c.Session.Sysf(".go <x> <y> <z> <o> <map>")
		return
	}

	pos := update.Position{}

	var mapID uint32

	if len(c.Args) == 1 {
		portID := c.String(0)

		var port *wdb.PortLocation
		wdb.GetData(portID, &port)

		if port == nil {
			var ports []*wdb.PortLocation
			if err := wdb.SearchTemplates(portID, 1, &ports); err != nil {
				c.Session.Warnf("%s", err)
				return
			}
			if len(ports) == 0 {
				c.Session.Warnf("could not find port location: '%s'", portID)
				return
			}
			port = ports[0]
			c.Session.Warnf("Could not find teleport location %s, sending you to %s.", portID, port.ID)
		}

		mapID = port.Map
		pos.X = port.X
		pos.Y = port.Y
		pos.Z = port.Z
		pos.O = port.O
		c.Session.Warnf("%s", spew.Sdump(port))
	} else {
		pos.X = c.Float32(0)
		pos.Y = c.Float32(1)
		pos.Z = c.Float32(2)
		pos.O = c.Float32(3)
		mapID = c.Uint32(4)
	}

	c.Session.TeleportTo(mapID, pos)
}
