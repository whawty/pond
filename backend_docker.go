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
	"errors"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

type DockerBackend struct {
	client *client.Client
}

func NewDockerBackend() (b *DockerBackend, err error) {
	return &DockerBackend{}, nil
}

func (b *DockerBackend) Init() (err error) {
	if b.client, err = client.NewEnvClient(); err != nil {
		return
	}

	var info types.Info
	if info, err = b.client.Info(context.Background()); err != nil {
		return err
	}

	wdl.Printf("docker connected to: %s (Version: %s), Root-Dir: %s, Driver: %s, %d CPUs, Memory: %d", info.Name, info.ServerVersion, info.DockerRootDir, info.Driver, info.NCPU, info.MemTotal)
	return
}

func (b *DockerBackend) Cleanup() error {
	// TODO: search for dangeling images and remove them
	// this might be unsafe during image build... needs investigation!!!
	return errors.New("not yet implemented")
}

func (b *DockerBackend) GetClient() (Client *client.Client, err error) {
	// TODO: check if connection is alive and reconnect if not
	return b.client, nil
}
