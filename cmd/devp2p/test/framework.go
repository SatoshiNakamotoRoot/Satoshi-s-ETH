// Copyright 2020 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package test

import (
	"crypto/ecdsa"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover/v4wire"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type testenv struct {
	l1, l2     net.PacketConn
	key        *ecdsa.PrivateKey
	remote     *enode.Node
	remoteAddr *net.UDPAddr
	waitTime   int
}

func newTestEnv(remote string, waitTime int, listen1, listen2 string) *testenv {
	l1, err := net.ListenPacket("udp", listen1)
	if err != nil {
		panic(err)
	}
	l2, err := net.ListenPacket("udp", listen2)
	if err != nil {
		panic(err)
	}
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	node, err := enode.Parse(enode.ValidSchemes, remote)
	if err != nil {
		panic(err)
	}
	if node.IP() == nil || node.UDP() == 0 {
		var ip net.IP
		var tcpPort, udpPort int
		if ip = node.IP(); ip == nil {
			ip = net.ParseIP("127.0.0.1")
		}
		if tcpPort = node.TCP(); tcpPort == 0 {
			tcpPort = 30303
		}
		if udpPort = node.TCP(); udpPort == 0 {
			udpPort = 30303
		}
		node = enode.NewV4(node.Pubkey(), ip, tcpPort, udpPort)
	}
	addr := &net.UDPAddr{IP: node.IP(), Port: node.UDP()}
	return &testenv{l1, l2, key, node, addr, waitTime}
}

func (te *testenv) close() {
	te.l1.Close()
	te.l2.Close()
}

func (te *testenv) send(c net.PacketConn, req v4wire.Packet) ([]byte, error) {
	packet, hash, err := v4wire.Encode(te.key, req)
	if err != nil {
		return hash, err
	}
	_, err = c.WriteTo(packet, te.remoteAddr)
	return hash, err
}

func (te *testenv) read(c net.PacketConn) (v4wire.Packet, []byte, error) {
	buf := make([]byte, 2048)
	if err := c.SetReadDeadline(time.Now().Add(time.Duration(te.waitTime) * time.Millisecond)); err != nil {
		return nil, nil, err
	}
	n, _, err := c.ReadFrom(buf)
	if err != nil {
		return nil, nil, err
	}
	p, _, hash, err := v4wire.Decode(buf[:n])
	return p, hash, err
}

func (te *testenv) localEndpoint(c net.PacketConn) v4wire.Endpoint {
	addr := c.LocalAddr().(*net.UDPAddr)
	return v4wire.Endpoint{
		IP:  addr.IP.To4(),
		UDP: uint16(addr.Port),
		TCP: 0,
	}
}

func (te *testenv) remoteEndpoint() v4wire.Endpoint {
	return v4wire.NewEndpoint(te.remoteAddr, 0)
}

func contains(ns []v4wire.Node, key v4wire.Pubkey) bool {
	for _, n := range ns {
		if n.ID == key {
			return true
		}
	}
	return false
}