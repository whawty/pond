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
	"io/ioutil"
	"log"
	"os"
	//	"path"
	"regexp"
	"strings"
)

var (
	wil            = log.New(os.Stdout, "[whawty.pond INFO]\t", log.LstdFlags)
	wel            = log.New(os.Stderr, "[whawty.pond ERROR]\t", log.LstdFlags)
	wdl            = log.New(ioutil.Discard, "[whawty.pond DEBUG]\t", log.LstdFlags)
	enableBackends = []string{"docker"}
	volumeBasePath = "/srv/volumes"
	serviceNameRe  = regexp.MustCompile("^[-_.A-Za-z0-9]+$")
	telnetAddr     = "127.0.0.1:9023"
)

func init() {
	if _, exists := os.LookupEnv("WHAWTY_POND_DEBUG"); exists {
		wdl.SetOutput(os.Stderr)
	}

	if value, exists := os.LookupEnv("WHAWTY_POND_BACKENDS"); exists {
		enableBackends = strings.Split(value, ",")
	}

	if value, exists := os.LookupEnv("WHAWTY_POND_VOLUMES_BASE"); exists {
		volumeBasePath = value
	}
}

type Context struct {
	Backends map[string]Backend
	Services map[string]*Service
}

func main() {
	wil.Printf("starting")

	var ctx Context
	ctx.Backends = make(map[string]Backend)
	ctx.Services = make(map[string]*Service)

	for _, name := range enableBackends {
		name = strings.TrimSpace(name)
		backend, err := NewBackend(name)
		if err != nil {
			wel.Printf("Error enabling backend(%s): %v", name, err)
			continue
		}
		if err := backend.Init(); err != nil {
			wel.Printf("backend(%s): can't be enabled: %v", name, err)
		} else {
			ctx.Backends[name] = backend
			wil.Printf("backend(%s): successfully enabled/initialized", name)
		}
	}
	if len(ctx.Backends) == 0 {
		wel.Printf("no backends are enabled, exitting...")
		os.Exit(1)
	}

	stop := make(chan bool)

	if telnetAddr != "" {
		telnet := TelnetInit(telnetAddr, &ctx)
		go func() {
			wil.Printf("starting telnet interface (%s)", telnetAddr)
			telnet.Run()
			wil.Printf("telnet interface just stopped")
			stop <- true
		}()
	}

	<-stop
	wil.Printf("at least one control interface has stopped - bringing down the whole process")

	// TODO: get this from db/config backend
	// var svc_name = "hugo"
	// if !serviceNameRe.MatchString(svc_name) {
	// 	wel.Printf("service name is invalid")
	// 	os.Exit(2)
	// }
	// if _, exists := ctx.Services[svc_name]; exists {
	// 	wel.Printf("Error adding new Service(%s): already exists", svc_name)
	// 	os.Exit(2)
	// }

	// svc, err := NewService(svc_name, path.Join(volumeBasePath, svc_name, "shared"))
	// if err != nil {
	// 	wel.Printf("Error adding new Service(%s): %v", svc_name, err)
	// 	os.Exit(2)
	// }

	// ctx.Services[svc_name] = svc

	// wdl.Printf("Services:")
	// for name, svc := range ctx.Services {
	// 	wdl.Printf(" - %s: %+v", name, *svc)
	// }
}
