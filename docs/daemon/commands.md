---
order: 4
---

# Commands

## Introduction

BHP Daemon Commands allow you to init, start, reset a node, or generate a genesis file, etc.

## Usage

```bash
bhpd <command>
```

## Available Commands

|Name                       | Description                                        |
| :----------------------------------------------- | :------------------------------------------ |
|                [init](#bhpd-init)                | Initialize private validator, p2p, genesis, and application configuration files |
| [add-genesis-account](#bhpd-add-genesis-account) | Add genesis account to genesis.json                |
|               [gentx](#bhpd-gentx)               | Generate a genesis tx carrying a self delegation                      |
|      [collect-gentxs](#bhpd-collect-gentxs)      | Collect genesis txs and output a genesis.json file              |
| [validate-genesis](#bhpd-validate-genesis) | validates the genesis file at the default location or at the location passed as an arg |



### bhpd init

Initialize private validator, p2p, genesis, and application configuration files

```bash	
bhpd init [moniker] [flags]
```

Example：

```bash
bhpd init testing --chain-id=testing
```

### bhpd add-genesis-account

Add genesis account to genesis.json

```bash
 bhpd add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
 # address_or_key_name => bhp address or account name
 # coin =>  token number，Example:10000000abhp
```

**Flags：**

| Name, shorthand           | Default     | Description                                                         | Required |
| :------------------- | -------- | ------------------------------------------------------------ | ---- |
| --home-client        | /.bhpcli | client's home directory                                                 |      |
| --vesting-amount     |          | amount of coins for vesting accounts |      |
| --vesting-start-time |          | schedule start time (unix epoch) for vesting accounts                             |      |
| --vesting-end-time   |          | schedule end time (unix epoch) for vesting accounts                             |      |

**Example：**

```bash
# token 1000000000abhp
bhpd add-genesis-account mykey2 1000000000abhp

# start time：2020-08-18 15:55:00（1597737830） end time：2020-08-20 16:20:00 （1597911600）
bhpd add-genesis-account mykey 1000000000abhp --vesting-amount 800000000abhp  --vesting-start-time 1597737830  --vesting-end-time 1597911600
```

### bhpd gentx

Generate a genesis tx carrying a self delegation

```bash
bhpd gentx [flags]
```

**Flags：**

| Name, shorthand              | Default         | Description                                             | Required |
| ---------------------------- | --------------- | ------------------------------------------------------- | -------- |
| --name                       |                 | name of private key with which to sign the gentx        |          |
| --amount                     | 10000000000abhp | Amount of coins to bond                                 |          |
| --commission-max-change-rate | 0.01            | The maximum commission change rate percentage (per day) |          |
| --commission-max-rate        | 0.2             | The maximum commission rate percentage                  |          |
| --commission-rate            | 0.1             | The initial commission rate percentage                  |          |
| --min-self-delegation        | 1               | The minimum self delegation required on the validator   |          |

**Example：**

```bash
bhpd gentx --name mykey
```

### bhpd collect-gentxs

Collect genesis txs and output a genesis.json file

```bash
bhpd collect-gentxs [flags]

# Flags:
#     --gentx-dir string => override default "gentx" directory from which collect and execute genesis transactions; default [--home]/config/gentx/
```

**Example：**

```bash
bhpd collect-gentxs
```

### bhpd validate-genesis

validates the genesis file at the default location or at the location passed as an arg

```bash
bhpd validate-genesis [file]
```

**Example：**

```bash
bhpd validate-genesis
# File at \.bhpd\config\genesis.json is a valid genesis file
```

