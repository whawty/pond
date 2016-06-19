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
	"strings"

	"github.com/spreadspace/telgo"
)

type TelnetInterface struct {
	server *telgo.Server
}

func telnetAddService(c *telgo.Client, args []string) bool {
	c.Sayln("add service not yet implemented (args: %+v)", args)
	return false
}

func telnetRemoveService(c *telgo.Client, args []string) bool {
	c.Sayln("remove service not yet implemented (args: %+v)", args)
	return false
}

func telnetService(c *telgo.Client, args []string) bool {
	switch len(args) {
	case 2:
		switch strings.ToLower(args[1]) {
		case "add":
			return telnetAddService(c, args[2:])
		case "remove":
			return telnetRemoveService(c, args[2:])
		}
		fallthrough
	default:
		c.Sayln("unknown sub command")
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
			c.Sayln("usage: service [ <cmd> ]")
			c.Sayln("   ...tba...")
			return false
		}
		fallthrough
	default:
		c.Sayln("usage: <cmd> [ [ <arg1> ] ... ]")
		c.Sayln("  available commands:")
		c.Sayln("    quit                             close connection (or use Ctrl-D)")
		c.Sayln("    help [ <cmd> ]                   print this, or help for specific command")
		c.Sayln("    service (add|remove)             manage services")
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
	cmdlist["service"] = telnetService

	telnet.server = telgo.NewServer(addr, "pond> ", cmdlist, nil)

	return
}
