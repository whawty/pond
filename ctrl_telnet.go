//
// Copyright (c) 2016 Christian Pointner <equinox@spreadspace.org>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of whawty.pond nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"path"
	"strings"

	"github.com/spreadspace/telgo"
)

type TelnetInterface struct {
	server *telgo.Server
}

func telnetAddService(c *telgo.Client, name string, args []string, ctx *Context) bool {
	_, err := ctx.NewService(name, path.Join(volumeBasePath, name, "shared"))
	if err != nil {
		c.Sayln("failed to add service '%s': %v", name, err)
		return false
	}
	c.Sayln("service '%s' successfully added", name)
	return false
}

func telnetRemoveService(c *telgo.Client, name string, args []string, ctx *Context) bool {
	c.Sayln("remove service %s: not yet implemented (args: %+v)", name, args)
	return false
}

func telnetStartService(c *telgo.Client, name string, args []string, ctx *Context) bool {
	c.Sayln("start service %s: not yet implemented (args: %+v)", name, args)
	return false
}

func telnetStopService(c *telgo.Client, name string, args []string, ctx *Context) bool {
	c.Sayln("stop service %s: not yet implemented (args: %+v)", name, args)
	return false
}

func telnetService(c *telgo.Client, args []string, ctx *Context) bool {
	if len(args) >= 3 {
		if !serviceNameRe.MatchString(args[2]) {
			c.Sayln("service name '%s' is invalid", args[2])
			return false
		}

		switch strings.ToLower(args[1]) {
		case "add":
			return telnetAddService(c, args[2], args[3:], ctx)
		case "remove":
			return telnetRemoveService(c, args[2], args[3:], ctx)
		case "start":
			return telnetRemoveService(c, args[2], args[3:], ctx)
		case "stop":
			return telnetRemoveService(c, args[2], args[3:], ctx)
		default:
			c.Sayln("unknown sub command")
		}
	}
	c.Sayln("too few arguments")
	return false
}

func telnetShow(c *telgo.Client, args []string, ctx *Context) bool {
	switch len(args) {
	case 2:
		switch strings.ToLower(args[1]) {
		case "services":
			c.Sayln("Services:")
			for name, svc := range ctx.Services {
				c.Sayln(" - %s: %+v", name, *svc)
			}
			return false
		default:
			c.Sayln("unknown type")
		}
		fallthrough
	default:
		c.Sayln("too few arguments")
	}
	return false

}

func telnetHelp(c *telgo.Client, args []string) bool {
	switch len(args) {
	case 2:
		switch strings.ToLower(args[1]) {
		case "quit":
			c.Sayln("usage: quit")
			c.Sayln("   terminates the client connection. You may also use Ctrl-D to do this.")
			return false
		case "help":
			c.Sayln("usage: help [ <cmd> ]")
			c.Sayln("   prints command overview or detailed info to <cmd>.")
			return false
		case "service":
			c.Sayln("usage: service (add|remove|start|stop) <name>")
			c.Sayln("   ...tba...")
			return false
		case "show":
			c.Sayln("usage: show (services)")
			c.Sayln("   ...tba...")
			return false
		}
		fallthrough
	default:
		c.Sayln("usage: <cmd> [ [ <arg1> ] ... ]")
		c.Sayln("  available commands:")
		c.Sayln("    quit                              close connection (or use Ctrl-D)")
		c.Sayln("    help [ <cmd> ]                    print this, or help for specific command")
		c.Sayln("    service <cmd> <name> [ <args..> ] manage services")
		c.Sayln("    show services                     list all services")
	}
	return false
}

func telnetQuit(c *telgo.Client, args []string) bool {
	return true
}

func (telnet *TelnetInterface) Run() {
	wdl.Printf("telnet: handler running...")
	if err := telnet.server.Run(); err != nil {
		wel.Printf("telnet: server returned: %s", err)
	}
}

func TelnetInit(addr string, ctx *Context) (telnet *TelnetInterface) {
	telnet = &TelnetInterface{}

	cmdlist := make(telgo.CmdList)
	cmdlist["help"] = telnetHelp
	cmdlist["quit"] = telnetQuit
	cmdlist["service"] = func(c *telgo.Client, args []string) bool { return telnetService(c, args, ctx) }
	cmdlist["show"] = func(c *telgo.Client, args []string) bool { return telnetShow(c, args, ctx) }

	telnet.server = telgo.NewServer(addr, "pond> ", cmdlist, nil)

	return
}
