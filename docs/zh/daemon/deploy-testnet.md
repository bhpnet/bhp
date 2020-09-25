---
order: 3
---

# 部署自己的BHP测试网

## 单节点测试网
前提
- [安装 bhp](../getting-start/install-bhp.md)
- [安装 `jq`](https://stedolan.github.io/jq/download/) (可选)

```shell script
# You can run all of these commands from your home directory
cd $HOME

# 初始化genesis.json文件，它将帮助你启动网络
bhpd init --chain-id=testing testing

# 创建一个钱包作为您的验证人帐户
bhpcli keys add validator

# 将该钱包地址添加到genesis文件中的genesis.app_state.accounts数组中
# 注意: 此命令使您可以设置通证数量。确保此帐户有bhp，这是BHPnet上唯一的质押通证
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
bhpd add-genesis-account $(bhpcli keys show validator -a) 100000000000abhp

# 生成创建验证人的交易。gentx存储在`~/.bhpd/config/`中
bhpd gentx --name validator

# 将生成的质押交易添加到创世文件
bhpd collect-gentxs

# 现在可以启动`bhpd`了
bhpd start
```

## bhpd migrate

```shell script
bhpd migrate [target-version] [genesis-file] [flags]
```

## bhpd replay

```shell script
bhpd replay <root-dir> [flags]
```

## bhpd unsafe-reset-all

可以使用此命令来重置节点，包括本地区块链数据库，地址簿文件，并将`priv_validator.json`重置为创世状态。

当本地区块链数据库以某种方式中断和无法同步或参与共识时，这是有用的。

```shell script
bhpd unsafe-reset-all
```

## bhpd tendermint

查询可以在p2p连接中使用的唯一节点ID，例如在[config.toml](intro.md#cnofig-toml)中`seeds`和`persistent_peers`的格式`<node-id>@ip:26656`。

节点ID存储在[node_key.json](intro.md#node_key-json)中。

```shell script
bhpd tendermint show-node-id
```

- 查询[Tendermint Pubkey](../validators/validator-faq.md#tendermint-密钥)，用于[identify your validator](../cli-client/staking.md),并将用于在共识过程中签署Pre-vote/Pre-commit。

[Tendermint Key](../validators/validator-faq.md#tendermint-密钥)存储在[priv_validator.json](intro.md#priv_validator-json)中，创建验证人后，请一定要记得[备份](../validators/validator-faq.md#如何备份验证人节点)。

```bash
bhpd tendermint show-validator
```

- 查询bech32前缀验证人地址

```shell script
bhpvalcons1apd8fcgqa84dg4sagl5cj2yn2z59mvy5p2pn6d
```

- 查询tendermint相关版本

```shell script
bhpd tendermint version
```
响应
```shell script
tendermint: 0.32.12
abci: 0.16.1
blockprotocol: 10
p2pprotocol: 7
```

## 多节点测试网
前提
- [安装 bhp](../getting-start/install-bhp.md)
- [安装 docker](https://docs.docker.com/engine/installation/)
- [安装 docker-compose](https://docs.docker.com/compose/install/)
### 构建和初始化
```shell script
# Clone the bhp repo
git clone -b <latest-release-tag> https://github.com/bhpnet/bhp

# Work from the SDK repo
cd bhp

# Build the linux binary in ./build
make build-linux

# Build okchain/node image
make build-docker-bhpdnode
```
### 运行你的测试网
运行以下命令创建4个节点的测试网
```shell script
make localnet-start
```
该命令将使用`bhpdnode`镜像创建4个节点的测试网。下表列出了每个节点的端口：
| Node      | P2P Port | RPC Port |
| --------- | -------- | -------- |
| bhpdnode0 | 26656    | 26657    |
| bhpdnode1 | 26659    | 26660    |
| bhpdnode2 | 26661    | 26662    |
| bhpdnode3 | 26663    | 26664    |

要更新二进制文件，只需重新构建它并重新启动节点即可：

```bash
make build-linux localnet-start
```
启动客户端
```shell script
nohup bhpcli rest-server --laddr "tcp://0.0.0.0:26690" --home /root/bhp/build/node0/bhpcli/ --chain-id=testing > bhpcli.log &
```
### 停止
停止所有正在运行的节点
```shell script
make localnet-stop
```
### 重置
要停止所有正在运行的节点并将网络重置为创世状态
```shell script
make localnet-unsafe-reset
```
### 清理
要停止所有正在运行的节点并删除build/目录中的所有文件
```shell script
make localnet-clean
```
### 配置
`make localnet-start`将调用`bhpd testnet`命令在`./build`目录下生成4个节点的测试网配置文件。
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
### 日志
日志保存在每个节点`./build/nodeN/bhpd/bhpd.log`。您还可以直接通过Docker查看日志，例如：
```shell script
docker logs -f bhpdnode0
```

```shell script
bhpcli keys list --home /root/bhp/build/node0/bhpcli
```

