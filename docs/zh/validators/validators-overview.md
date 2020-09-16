# 验证者概述
## 简介

BHP 基于 Cosmos SDK 和 Tendermint 实现，它依靠一组验证人节点来维护网络的安全。验证人节点的作用是运行一个全节点并通过广播投票（包含由其私钥签名的密文）来参与共识。验证人节点在区块链中提交新区块并获得收入以激励他们的工作。他们还必须对提案进行投票来参与链上治理。验证人节点的投票权重取决于他们的BHP总质押数量。

## 验证人的责任

验证者有两个主要职责：

- 能够持续运行正确版本的软件： 验证人需要确保其服务器始终在线并且其私钥不受到损害。
- 积极参与治理： 要求验证人对每个提案进行投票。

此外，验证人将是社区的积极成员。他们应该始终与生态系统的当前状态保持同步，以便轻松地适应任何变化。


## 验证人的部分操作

### 创建验证人

[bhpcli tx staking create-validator](https://docs.bhpnet.io/zh/cli-client/staking.html#bhpcli-tx-staking-create-validator)

### 编辑验证人

[bhpcli tx staking edit-validator](https://docs.bhpnet.io/zh/cli-client/staking.html#bhpcli-tx-staking-edit-validator)

### 委托BHP到验证人节点

[bhpcli tx staking delegate](https://docs.bhpnet.io/zh/cli-client/staking.html#bhpcli-tx-staking-delegate)

### 获取佣金和手续费收益

从某个验证者地址提取奖励，如果你的地址是验证人身份还可以提取佣金收益

[bhpcli tx distribution withdraw-rewards](https://docs.bhpnet.io/zh/cli-client/slashing.html)

### 解除监禁

[bhpcli tx slashing unjail](https://docs.bhpnet.io/zh/cli-client/slashing.html)

### 其他命令 

更多验证人命令可[查看](/zh/cli-client/staking.md)