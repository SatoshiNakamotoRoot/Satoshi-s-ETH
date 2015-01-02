// Copyright (c) 2013-2014, Jeffrey Wilcke. All rights reserved.
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this library; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
// MA 02110-1301  USA

package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethutil"
	"github.com/ethereum/go-ethereum/logger"
)

const (
	ClientIdentifier = "Ethereum(G)"
	Version          = "0.8.0"
)

var clilogger = logger.NewLogger("CLI")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer func() {
		logger.Flush()
	}()

	utils.HandleInterrupt()

	// precedence: code-internal flag default < config file < environment variables < command line
	Init() // parsing command line

	utils.InitConfig(VmType, ConfigFile, Datadir, "ETH")

	ethereum, err := eth.New(&eth.Config{
		Name:       ClientIdentifier,
		Version:    Version,
		KeyStore:   KeyStore,
		DataDir:    Datadir,
		LogFile:    LogFile,
		LogLevel:   LogLevel,
		Identifier: Identifier,
		MaxPeers:   MaxPeer,
		Port:       OutboundPort,
		NATType:    PMPGateway,
		PMPGateway: PMPGateway,
		KeyRing:    KeyRing,
	})
	if err != nil {
		clilogger.Fatalln(err)
	}
	utils.KeyTasks(ethereum.KeyManager(), KeyRing, GenAddr, SecretFile, ExportDir, NonInteractive)

	if Dump {
		var block *types.Block

		if len(DumpHash) == 0 && DumpNumber == -1 {
			block = ethereum.ChainManager().CurrentBlock()
		} else if len(DumpHash) > 0 {
			block = ethereum.ChainManager().GetBlock(ethutil.Hex2Bytes(DumpHash))
		} else {
			block = ethereum.ChainManager().GetBlockByNumber(uint64(DumpNumber))
		}

		if block == nil {
			fmt.Fprintln(os.Stderr, "block not found")

			// We want to output valid JSON
			fmt.Println("{}")

			os.Exit(1)
		}

		// Leave the Println. This needs clean output for piping
		fmt.Printf("%s\n", block.State().Dump())

		fmt.Println(block)

		return
	}

	if StartMining {
		utils.StartMining(ethereum)
	}

	if len(ImportChain) > 0 {
		start := time.Now()
		err := utils.ImportChain(ethereum, ImportChain)
		if err != nil {
			clilogger.Infoln(err)
		}
		clilogger.Infoln("import done in", time.Since(start))
		return
	}

	if StartRpc {
		utils.StartRpc(ethereum, RpcPort)
	}

	if StartWebSockets {
		utils.StartWebSockets(ethereum)
	}

	utils.StartEthereum(ethereum, UseSeed)

	if StartJsConsole {
		InitJsConsole(ethereum)
	} else if len(InputFile) > 0 {
		ExecJsFile(ethereum, InputFile)
	}
	// this blocks the thread
	ethereum.WaitForShutdown()
}
