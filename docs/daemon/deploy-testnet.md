---
order: 3
---

# Deploy Your Own BHP Testnet

## Single Node Testnet
Requirements:
- [install bhp](../getting-start/install-bhp.md)
- [install `jq`](https://stedolan.github.io/jq/download/) (optional)

```shell script
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
bhpd init --chain-id=testing testing

# Create a key to hold your validator account
bhpcli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# this command lets you set the number of coins. Make sure this account has some abhp which is the only staking coin on BHP
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
bhpd add-genesis-account $(bhpcli keys show validator -a) 1000000000abhp

# Generate the transaction that creates your validator. The gentxs are stored in `~/.bhpd/config/gentx`
bhpd gentx --name validator

# Add the generated staking transactions to the genesis file
bhpd collect-gentxs

# Now it‘s ready to start `bhpd`
bhpd start
```
## Multiple Nodes Testnet
Requirements:
- [install bhp](../getting-start/install-bhp.md)
- [install docker](https://docs.docker.com/engine/installation/)
- [install docker-compose](https://docs.docker.com/compose/install/)
### Build and Init
```shell script
# Clone the bhp repo
git clone https://github.com/bhpnet/bhp.git

# Work from the SDK repo
cd bhp

# Build the linux binary in ./build
make build-linux

# Build okchain/node image
make build-docker-bhpdnode
```
### Run Your Testnet
To start a 4 node testnet run:
```shell script
make localnet-start
```
This command creates a 4-node network using the bhpdnode image. The ports for each node are found in this table:
| Node      | P2P Port | RPC Port |
| --------- | -------- | -------- |
| bhpdnode0 | 26656    | 26657    |
| bhpdnode1 | 26659    | 26660    |
| bhpdnode2 | 26661    | 26662    |
| bhpdnode3 | 26663    | 26664    |

To update the binary, just rebuild it and restart the nodes:

```bash
make build-linux localnet-start
```
Start bhpcli
```shell script
nohup bhpcli rest-server --laddr "tcp://0.0.0.0:26690" --home /root/bhp/build/node0/bhpcli/ --chain-id=testing > bhpcli.log &
```
### Stop
To stop all the running nodes:
```shell script
make localnet-stop
```
### Reset
To stop all the running nodes and reset the network to the genesis state:
```shell script
make localnet-unsafe-reset
```
### Clean
To stop all the running nodes and delete all the files in the build/ directory:
```shell script
make localnet-clean
```
### Configuration
The `make localnet-start` creates files for a 4-node testnet in `./build` by calling the `bhpd testnet` command. This outputs a handful of files in the `./build` directory:
```shell script
tree -L 3 build/
build/
├── bhpcli
├── bhpd
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── bhpd
│   │   ├── bhpd.log
│   │   ├── config
│   │   ├── data
│   │   └── keys
│   └── bhpcli
│       ├── keys
│       └── key_seed.json
├── node1
│   ├── bhpd
│   │   ├── bhpd.log
│   │   ├── config
│   │   └── data
│   └── bhpcli
│       ├── keys
│       └── key_seed.json
├── node2
│   ├── bhpd
│   │   ├── bhpd.log
│   │   ├── config
│   │   └── data
│   └── bhpcli
│       ├── keys
│       └── key_seed.json
└── node3
    ├── bhpd
    │   ├── bhpd.log
    │   ├── config
    │   └── data
    └── bhpcli
        ├── keys
        └── key_seed.json
```
### Logging
Logs are saved under each `./build/nodeN/bhpd/bhpd.log`. You can also watch logs directly via Docker, for example:
```shell script
docker logs -f bhpdnode0
```
### Keys & Accounts
To interact with `bhpcli` and start querying state or creating txs, you use the `bhpcli` directory of any given node as your `home`, for example:
```shell script
bhpcli keys list --home /root/bhp/build/node0/bhpcli
```

