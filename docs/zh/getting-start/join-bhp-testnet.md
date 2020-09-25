---
order: 3
---

# 加入星空测试网

有关最新测试网络的信息，请参阅[测试网络存储库](https://github.com/bhpnet/bhp)，包含了所使用的BHP的正确版本和genesis文件。

:::tip
需要先 [安装 bhp](install-bhp.md)
:::

## 创建一个新节点

这些指令适用于从头开始设置一个全节点。

首先，初始化节点并创建必要的配置文件：

```shell script
bhpd init <your_custom_moniker> --chain-id=testing
```

:::tip
注意 moniker只能包含ASCII字符。使用Unicode字符会使得你的节点不可访问
:::

你可以稍后在`~/.bhpd/config/config.toml`文件中编辑`moniker`:

```shell script
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

你的全节点已经初始化成功！

## Genesis & Seeds

### 复制genesis文件

将测试网的`genesis.json`文件放置在`bhpd`的配置文件夹中

```shell script
cd $HOME/.bhpd/config
curl -o ~/.bhpd/config/genesis.json https://raw.githubusercontent.com/bhpnet/testnet/master/genesis.json
```

### 添加持久节点

你的节点需要知道如何寻找伙伴节点。你需要添加有用的持久节点到`$HOME/.bhpd/config/config.toml`文件中。[testnet repo](https://github.com/bhpnet/testnet)包含了一些种子节点的链接。

## 关于 Gas 和 Fees

:::warning
在BHP星空测试网中，可接受的通证是`abhp`,`1bhp = 100000000abhp`
:::

BHP网络中的交易需要支付一笔交易手续费以得到处理。手续费支付执行交易所消耗的gas。计算公式如下：

```shell script
fees = gas * gasPrices
```

`gas`由交易本身决定。不同的交易需要不同数量的`gas`。一笔交易的gas数量在它被执行时计算，但有一种方式可以提前估算，那就是把标识`gas` 的值设置为`auto`。当然，这只是给出一个预估值。如果你想要确保为交易提供足够的gas，你可以使用`--gas-adjustment`标识来调整预估值(默认是`1.0`)。

`gasPrice`是每个单位gas的单价。每个验证人节点可以设置`min-gas-price`，只会把那些`gasPrice`高于自己设置的`min-gas-price`的交易打包。

交易的`fees`是`gas`与`gasPrice`的结果。作为一个用户，你必须输入三者中的两者。更高的`gasPrice/fees`，将提高你的交易被打包的机会。

:::tip
BHP网络中推荐的`gas-prices`是`2.5abhp`
:::

### 设置`minimum-gas-prices`

你可以编辑`~/.bhpd/config/app.toml`文件来开启垃圾交易过滤机制以拒绝收到的手续费过低的交易：

```shell script
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10abhp).

minimum-gas-prices = ""
```

你的全节点可以在交易池中放入未确认的交易。为了保护其免受Spam攻击，最好设置一个`minimum-gas-prices`来过滤交易以决定是否要放入交易池。这个参数可以在`~/.bhpd/config/bhpd.toml`文件中配置。

推荐的初始`minimum-gas-prices`是`2.5abhp`，如果你愿意可以稍后再修改它。

## 运行一个全节点

通过这条命令开始运行全节点：

```
# 启动节点（也可使用 nohup 或 systemd 等方式后台运行）
# bhpd start

# 后台启动
nohup bhpd start 1>bhpd.log 2>&1 &
```

检查一切是否平稳运行中:

```shell script
bhpcli status
```

使用[BHP Explorer](https://scan.bhpnet.io/)查看网络状态。

:::tip
从零开始追赶区块需要很长时间，您也可以下载[测试网数据快照](#TODO)以减少同步时间
:::

## 升级为验证人节点

你现在有了一个运行状态的全节点。接下来，你可以升级你的全节点，成为一个BHP验证人。排名前21的验证人节点可以向BHP网络提议新的区块。

### 创建钱包

您可以[创建新的钱包](../cli-client/keys.md#创建密钥)或[导入现有的钱包](../cli-client/keys.md#通过助记词恢复密钥)，然后从交易所或其他任何地方转入一些BHP到您刚刚创建的钱包中：

```bash
# 创建一个新钱包
bhpcli keys add <key-name>
```

:::warning
在安全的地方备份好助记词！如果您忘记密码，这是恢复帐户的唯一方法。
:::

### 获得测试token

浏览器访问该地址 [https://faucet.bhpnet.io/](https://faucet.bhpnet.io/) 领取测试token

### 确认节点同步状态

```bash
# 可以使用此命令安装 jq
# apt-get update && apt-get install -y jq

# 如果输出为 false, 则表明您的节点已经完成同步
bhpcli status | jq .sync_info.catching_up
```

### 创建验证人

只有节点已完成同步时，才可以运行以下命令将您的节点升级为验证人：

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
**重要**

一定要备份好 home（默认为〜/.bhpd/）目录中的 `config` 目录！如果您的服务器磁盘损坏或您准备迁移服务器，这是恢复验证人的唯一方法。
:::

如果以上命令没有出现错误，则您的节点已经是验证人或候选人了（取决于您的Voting Power是否在前21名中）

### 验证人命令指南

更多验证人命令可[查看](/zh/cli-client/staking.md)