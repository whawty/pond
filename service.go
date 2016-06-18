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
)

type Service struct {
	Name       string
	VolumePath string
	Images     map[string]*Image
	Container  map[string]*Container
}

// Service Handling
func NewService(name, volumePath string) (s *Service, err error) {
	s = &Service{}
	s.Name = name
	s.VolumePath = volumePath
	s.Images = make(map[string]*Image)
	s.Container = make(map[string]*Container)
	return
}

// Start all Container of the Service
func (s *Service) Start() error {
	return errors.New("not yet implemented")
}

// Stop all Container of the Service
func (s *Service) Stop() error {
	return errors.New("not yet implemented")
}

// Rebuild all Images and Restart Container
func (s *Service) RebuildAndRestart() error {
	return errors.New("not yet implemented")
}

// Remove all Container and Images of this Service
func (s *Service) Wipe() error {
	return errors.New("not yet implemented")
}

// Add an Image to the Service (build if it not exists in backend)
func (s *Service) AddImage(backend, name, version string) (Image, error) {
	return nil, errors.New("not yet implemented")
}

// Return the image with <name>:<version>
func (s *Service) GetImage(name, version string) (Image, error) {
	return nil, errors.New("not yet implemented")
}

// Remove the image from the backend
func (s *Service) RemoveImage(name, version string) error {
	return errors.New("not yet implemented")
}

// Add a container based on <image>
func (s *Service) AddContainer(image, name string) (Container, error) {
	return nil, errors.New("not yet implemented")
}

// Return the container with name <name>
func (s *Service) GetContainer(name string) (Container, error) {
	return nil, errors.New("not yet implemented")
}

// Remove the Container <name> from the service
func (s *Service) RemoveContainer(name string) error {
	return errors.New("not yet implemented")
}

// Start Container with name <name>
func (s *Service) StartContainer(name string) error {
	return errors.New("not yet implemented")
}

// Stop Container with name <name>
func (s *Service) StopContainer(name string) error {
	return errors.New("not yet implemented")
}
