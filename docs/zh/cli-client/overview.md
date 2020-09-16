---
order: 1
---

# 简介

`bhpcli` 是BHP网络的命令行客户端。BHP用户可以使用`bhpcli`发送交易和查询区块链数据。

## 全局标识

### GET 请求

所有GET请求都有以下全局标识:

| 名称，速记   | 类型   | 必需 | 默认值                | 描述                                                               |
| ------------ | ------ | ---- | --------------------- | ------------------------------------------------------------------ |
| --chain-id   | string |      | ""                    | tendermint节点的Chain ID                                           |
| --height     | int    |      | 0                     | 查询某个高度的区块链数据，如果是0，则返回最新的区块链数据          |
| --help, -h   | string |      |                       | 打印帮助信息                                                       |
| --output     | string |      | text                  | 指定返回的格式 text或者json                                        |
| --indent     | bool   |      | false                 | 格式化json字符串                                                   |
| --ledger     | bool   |      | false                 | 使用ledger设备                                                     |
| --node       | string |      | tcp://localhost:26657 | tendermint节点的rpc地址                                            |
| --trust-node | bool   |      | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |

### POST 请求

所有POST请求都有以下全局标识:

| 名称，速记       | 类型   | 必需  | 默认值                | 描述                                                                                                                |
| ---------------- | ------ | ----- | --------------------- | ------------------------------------------------------------------------------------------------------------------- |
| --account-number | int    | false | 0                     | 发起交易的账户的编号                                                                                                |
| --async          | bool   | false | false                 | 是否异步广播交易（不对交易进行任何验证，立即返回交易的hash，仅当commit为false时有效）                               |
| --commit         | bool   | false | false                 | 广播交易并等到交易被打包再返回                                                                                      |
| --chain-id       | string | true  | ""                    | tendermint节点的`Chain ID`                                                                                          |
| --dry-run        | bool   | false | false                 | 模拟执行交易，并返回消耗的`gas`。`--gas`指定的值会被忽略                                                            |
| --fee            | string | true  | ""                    | 交易费（指定交易费的上限）                                                                                          |
| --from           | string | false | ""                    | 发送交易的账户名称                                                                                                  |
| --from-addr      | string | false | ""                    | 签名地址，在`generate-only`为`true`的情况下有效                                                                     |
| --gas            | int    | false | 50000                 | 交易的gas上限; 设置为"simulate"将自动计算相应的阈值                                                                 |
| --gas-adjustment | int    | false | 1.5                   | gas调整因子，这个值降乘以模拟执行消耗的`gas`，计算的结果返回给用户; 如果`--gas`的值不是`simulate`，这个标志将被忽略 |
| --generate-only  | bool   | false | false                 | 是否仅仅构建一个未签名的交易便返回                                                                                  |
| --help, -h       | string | false |                       | 打印帮助信息                                                                                                        |
| --indent         | bool   | false | false                 | 格式化json字符串                                                                                                    |
| --json           | string | false | false                 | 指定返回结果的格式，`json`或者`text`                                                                                |
| --ledger         | bool   | false | false                 | 使用ledger设备                                                                                                      |
| --memo           | string | false | ""                    | 指定交易的memo字段                                                                                                  |
| --node           | string | false | tcp://localhost:26657 | tendermint节点的rpc地址                                                                                             |
| --print-response | bool   | false | false                 | 是否打印交易返回结果，仅在`async`为true的情况下有效                                                                 |
| --sequence       | int    | false | 0                     | 发起交易的账户的sequence                                                                                            |
| --trust-node     | bool   | false | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性                                                  |


## 模块命令列表

| **子命令**                        | **描述**                        |
| --------------------------------- | ------------------------------|
| [status](./status.md)             | 查询远程节点的状态               |
| [query](./query.md)               | 查询子命令                      |
| [tx](./tx.md)                     | Tx子命令                       |
| [distribution](./distribution.md) | 收益管理子命令                  |
| [slashing](./slashing.md)         | 解禁子命令                     |
| [gov](./gov.md)                   | 治理和投票子命令                 |
| [stake](./staking.md)             | Staking 子命令                 |
| [rest-server](./rest-server.md)   | 客户端API命令                    |
| [keys](./keys.md)                 | 密钥管理子命令                   |
