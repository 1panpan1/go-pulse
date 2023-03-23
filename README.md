# Go-Pulse: A PulseChain Execution Client Written in Go

This is the core repository for Go-Pulse, a [Golang](https://golang.org/) implementation of the [PulseChain](https://pulsechain.com/) execution client, forked from [Go-Ethereum](https://github.com/ethereum/go-ethereum) to provide native PulseChain support.

> This project is forked from [Go-Ethereum](https://github.com/ethereum/go-ethereum), the official Golang implementation of the Ethereum protocol. Credit goes to the Go-Ethereum developers for the original development. Go-Pulse extends this foundation with PulseChain-specific features and network support.

## PulseChain Features

Go-Pulse extends Go-Ethereum with PulseChain-specific features:

- **Native PulseChain Support**: Built-in support for PulseChain mainnet and testnet networks
- **PrimordialPulse Fork**: Implements sacrifice credits and deposit contract replacement
- **Ethereum Compatibility**: Maintains compatibility with Ethereum networks for cross-chain operations
- **Standard RPC APIs**: Uses standard Ethereum JSON-RPC APIs for all network interactions
- **Fork Support**: Supports PulseChain's unique fork characteristics and consensus rules

## Building the source

For prerequisites and detailed build instructions, please refer to the [Go-Ethereum Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth) as Go-Pulse follows the same build process.

Building `geth` requires both a Go (version 1.23 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-pulse project comes with several wrappers/executables found in the `cmd`
directory.

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| :--------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`geth`** | Our main PulseChain CLI client. It is the entry point into the PulseChain networks (mainnet, testnet, or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the PulseChain networks via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI page](https://geth.ethereum.org/docs/fundamentals/command-line-options) for command line options. |
|   `clef`   | Stand-alone signing tool, which can be used as a backend signer for `geth`.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
|  `devp2p`  | Utilities to interact with nodes on the networking layer, without running a full blockchain.                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|  `abigen`  | Source code generator to convert contract definitions into easy-to-use, compile-time type-safe Go packages. It operates on plain [contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see the [Go-Ethereum documentation](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) for details.                                  |
|   `evm`    | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                               |
| `rlpdump`  | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                |

## Running `geth`

Going through all the possible command line flags is out of scope here (please consult the
[Go-Ethereum CLI documentation](https://geth.ethereum.org/docs/fundamentals/command-line-options) for reference),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance on PulseChain networks.

### Hardware Requirements

**PulseChain Mainnet:**

Minimum:
* CPU with 4+ cores
* 8GB RAM
* 1TB free storage space to sync PulseChain mainnet
* 8 MBit/sec download Internet service

Recommended:
* Fast CPU with 8+ cores
* 16GB+ RAM
* High-performance SSD with at least 1TB of free space
* 25+ MBit/sec download Internet service

### Full node on PulseChain Mainnet

The most common scenario is connecting to the PulseChain mainnet to interact with the network: create accounts; transfer PLS; deploy and interact with contracts. To connect to PulseChain mainnet:

```shell
$ geth --pulsechain console
```

This command will:
 * Start `geth` in snap sync mode (default, can be changed with the `--syncmode` flag),
   causing it to download more data in exchange for avoiding processing the entire history
   of the PulseChain network, which is very CPU intensive.
 * Start the built-in interactive [JavaScript console](https://geth.ethereum.org/docs/interacting-with-geth/javascript-console),
   (via the trailing `console` subcommand) through which you can interact using [`web3` methods](https://github.com/ChainSafe/web3.js/blob/0.20.7/DOCUMENTATION.md) 
   (note: the `web3` version bundled within `geth` may not be up to date with the latest documentation),
   as well as `geth`'s own [management APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc).
   This tool is optional and if you leave it out you can always attach it to an already running
   `geth` instance with `geth attach`.

### Full node on PulseChain Testnet V3

For development and testing, you can connect to the PulseChain Testnet V3:

```shell
$ geth --pulsechain-testnet-v3 console
```

The `console` subcommand works the same way as on mainnet and is equally useful for testing.

Specifying the `--pulsechain-testnet-v3` flag will configure your `geth` instance for the testnet:

 * Connect to the PulseChain Testnet V3, which uses different P2P bootnodes, different network ID and a different fork block
 * Use a separate data directory for testnet data to keep it isolated from mainnet.

*Note: Always use separate accounts for testnet and mainnet. Go-Pulse will correctly separate the two networks and will not make accounts available between them.*

### Full node on Ethereum networks (for compatibility)

Go-Pulse also supports connecting to Ethereum networks for compatibility:

**Ethereum Mainnet:**
```shell
$ geth --ethereum console
```

**Holesky Testnet:**
```shell
$ geth --holesky console
```

### Configuration

As an alternative to passing the numerous flags to the `geth` binary, you can also pass a
configuration file via:

```shell
$ geth --config /path/to/your_config.toml
```

To get an idea of how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ geth --your-favorite-flags dumpconfig
```

#### Docker quick start

One of the quickest ways to get Go-Pulse up and running on your machine is by using
Docker:

**PulseChain Mainnet:**
```shell
docker run -d --name pulsechain-mainnet-node -v /path/to/pulsechain/data:/root \
           -p 8545:8545 -p 30303:30303 \
           registry.gitlab.com/pulsechaincom/go-pulse --pulsechain
```

**PulseChain Testnet V3:**
```shell
docker run -d --name pulsechain-testnet-node -v /path/to/pulsechain-testnet/data:/root \
           -p 8545:8545 -p 30303:30303 \
           registry.gitlab.com/pulsechaincom/go-pulse --pulsechain-testnet-v3
```

This will start `geth` in snap-sync mode and create a persistent volume for
saving your blockchain data as well as map the default ports.

Do not forget `--http.addr 0.0.0.0`, if you want to access RPC from other containers
and/or hosts. By default, `geth` binds to the local interface and RPC endpoints are not
accessible from the outside.

### Programmatically interfacing `geth` nodes

As a developer, sooner rather than later you'll want to start interacting with `geth` and the
PulseChain network via your own programs and not manually through the console. To aid
this, `geth` has built-in support for a JSON-RPC based APIs ([standard APIs](https://ethereum.org/en/developers/docs/apis/json-rpc/)
and [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `geth`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` APIs offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross-origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` APIs offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept WebSocket requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `geth` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
PulseChain nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

### Operating a private network

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

Unfortunately since [the Merge](https://ethereum.org/en/roadmap/merge/) it is no longer possible
to easily set up a network of geth nodes without also setting up a corresponding beacon chain.

There are three different solutions depending on your use case:

  * If you are looking for a simple way to test smart contracts from go in your CI, you can use the [Simulated Backend](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings#blockchain-simulator).
  * If you want a convenient single node environment for testing, you can use our [Dev Mode](https://geth.ethereum.org/docs/developers/dapp-developer/dev-mode).
  * If you are looking for a multiple node test network, you can set one up quite easily with [Kurtosis](https://geth.ethereum.org/docs/fundamentals/kurtosis).

## License

The go-pulse library (i.e. all code outside of the `cmd` directory), like the upstream go-ethereum, is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-pulse binaries (i.e. all code inside of the `cmd` directory), like the upstream go-ethereum, are licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
