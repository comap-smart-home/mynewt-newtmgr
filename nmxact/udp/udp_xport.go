/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package udp

import (
	"fmt"

	"github.com/comap-smart-home/mynewt-newtmgr/nmxact/nmxutil"
	"github.com/comap-smart-home/mynewt-newtmgr/nmxact/sesn"
)

type UdpXport struct {
	started bool
}

func NewUdpXport() *UdpXport {
	return &UdpXport{}
}

func (ux *UdpXport) BuildSesn(cfg sesn.SesnCfg) (sesn.Sesn, error) {
	return NewUdpSesn(cfg)
}

func (ux *UdpXport) Start() error {
	if ux.started {
		return nmxutil.NewXportError("UDP xport started twice")
	}
	ux.started = true
	return nil
}

func (ux *UdpXport) Stop() error {
	if !ux.started {
		return nmxutil.NewXportError("UDP xport stopped twice")
	}
	ux.started = false
	return nil
}

func (ux *UdpXport) Tx(bytes []byte) error {
	return fmt.Errorf("unsupported")
}
