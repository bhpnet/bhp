---
order: 3
---

# Join The Testnet

See the [testnet repo](https://github.com/bhpnet/bhp) for information on the latest testnet, including the correct version of bhp to use and details about the genesis file.

:::tip
Requirements: [install bhp](install-bhp.md)
:::

## Setting Up a New Node

These instructions are for setting up a brand new full node from scratch.

First, initialize the node and create the necessary config files:

```shell script
bhpd init <your-custom-name> --chain-id=testing
```

:::tip
Monikers can contain only ASCII characters. Using Unicode characters will render your node unreachable.
:::

You can edit this `moniker` later, in the `~/.bhpd/config/config.toml` file:

```shell script
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the `~/.bhpd/config/app.toml` file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:

```shell script
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10uatom).

minimum-gas-prices = ``""
```

Your full node has been initialized!

## Genesis & Seeds

### Copy the Genesis File

Fetch the testnet's genesis.json file into bhpd's config directory.

```shell script
cd $HOME/.bhpd/config
curl -o ~/.bhpd/config/genesis.json https://raw.githubusercontent.com/bhpnet/testnet/master/genesis.json
```

### Add Seed Nodes

Your node needs to know how to find peers. You'll need to add healthy seed nodes to `$HOME/.bhpd/config/config.toml`. The [testnet repo](https://github.com/bhpnet/testnet) contains links to some seed nodes.

## A Note on Gas and Fees

:::warning
On BHP Network testnet, the accepted denom is `abhp`,`1bhp = 100000000abhp`
:::

Transactions on the BHP network need to include a transaction fee in order to be processed. This fee pays for the gas required to run the transaction. The formula is the following:

```shell script
fees = ceil(gas * gasPrices)
```

The gas is dependent on the transaction. Different transaction require different amount of gas. The gas amount for a transaction is calculated as it is being processed, but there is a way to estimate it beforehand by using the auto value for the gas flag. Of course, this only gives an estimate. You can adjust this estimate with the flag --gas-adjustment (default 1.0) if you want to be sure you provide enough gas for the transaction.

The gasPrice is the price of each unit of gas. Each validator sets a min-gas-price value, and will only include transactions that have a gasPrice greater than their min-gas-price.

The transaction fees are the product of gas and gasPrice. As a user, you have to input 2 out of 3. The higher the gasPrice/fees, the higher the chance that your transaction will get included in a block.``

:::tip
For BHP network, the recommended `gas-prices` is `2.5abhp`.
:::

### Set `minimum-gas-prices`

Your full-node keeps unconfirmed transactions in its mempool. In order to protect it from spam, it is better to set a minimum-gas-prices that the transaction must meet in order to be accepted in your node's mempool. This parameter can be set in the following file `~/.bhpd/config/app.toml`.

The initial recommended min-gas-prices is `2.5abhp`, but you might want to change it later.

:::tip
It will take a long time to sync from scratch, you can also download the [testnet data snapshot](#TODO) to reduce the time spent on synchronization
:::

## Run a Full Node

Start the full node with this command:

```shell script
# start the node (you can also use "nohup" or "systemd" to run in the background)
# bhpd start

# start the node in the background
nohup bhpd start 1>bhpd.log 2>&1 &
```

Check that everything is running smoothly:

```shell script
bhpcli status
```

View the status of the network with the [BHP Explorer](https://scan.bhpnet.io/)

## Upgrade to Validator Node

### Create a Wallet

You can [create a new wallet](../cli-client/keys.md#create-a-new-key) or [import an existing one](../cli-client/keys.md#recover-an-existing-key-from-seed-phrase), then get some BHP from the exchanges or anywhere else into the wallet you just created, .e.g.

```bash
# create a new wallet
bhpcli keys add <key-name>
```

:::warning
write the seed phrase in a safe place! It is the only way to recover your account if you ever forget your password.
:::

### Obtain test Token

To get the test token, please refer to [here](https://faucet.bhpnet.io/)

### Confirm your node has caught-up

```bash
# if you have not installed jq
# apt-get update && apt-get install -y jq

# if the output is false, means your node has caught-up
bhpcli status | jq .sync_info.catching_up
```

### Create Validator

Only if your node has caught-up, you can run the following command to upgrade your node to be a validator.

```bash
bhpcli tx staking create-validator \
    --amount=1000000000abhp \
    --pubkey=$(bhpd tendermint show-validator) \
    --moniker="DemoValidator" \
    --details="I'm a good validator" \
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1" \
    --from=demokey \
    --chain-id=testing \
    --gas=auto \
    --gas-prices="2.5abhp" \
    --gas-adjustment=2
```

:::warning
**Important**

Backup the `config` directory located in your bhp home (default ~/.bhpd/) carefully! It is the only way to recover your validator.
:::

If there are no errors, then your node is now a validator or candidate (depending on whether your delegation amount is in the top 21)


### Validator Command Guide

More validator commands are available [view](/cli-client/staking.md)