---
order: 7
---

# gov

该模块提供了[链上治理](../concepts/governance.md)的基本功能。

## 可用命令

| 名称                                            | 描述                                              |
| ----------------------------------------------- | ------------------------------------------------- |
| [query gov proposal](#bhpcli-query-gov-proposal)   | 查询单个提案的详细信息                       |
| [query gov proposals](#bhpcli-query-gov-proposals) | 按条件查询提案                              |
| [query gov vote](#bhpcli-query-gov-vote)           | 查询投票                                   |
| [query gov votes](#bhpcli-query-gov-votes)         | 按条件查询投票                                    |
| [query gov param](#bhpcli-query-gov-param)           | 指定参数类型查询治理过程的参数。                    |
| [query gov params](#bhpcli-query-gov-params)         | 查询治理过程的所有参数。                           |
| [query gov proposer](#bhpcli-query-gov-proposer)           | 查询一个提议信息                            |
| [query gov deposit](#bhpcli-query-gov-votes)         | 查询指定提案的抵押信息。              |
| [query gov deposits](#bhpcli-query-gov-deposits)           | 查询指定提案的所有抵押信息。      |
| [query gov tally](#bhpcli-query-gov-tally)         | 查询提案投票的统计信息        |
| [tx gov submit-proposal](#bhpcli-tx-gov-submit-proposal) | 提交提案以及初始化抵押金额           |
| [tx gov deposit](#bhpcli-tx-gov-deposit)                 | 为有效的提案抵押通证                |
| [tx gov vote](#bhpcli-tx-gov-vote)                       | 为有效的提案投票，选项：Yes/No/NoWithVeto/Abstain |

## bhpcli query gov proposal

查询提案的详细信息。

```shell script
bhpcli query gov proposal [proposal-id] [flags]
```
**标识：**

| 名称, 速记    | 类型 | 必须 | 默认 | 描述     |
| ------------- | ---- | -------- | ---- | -------- |
| proposal-id | uint | 是       |      | 提案的Id |
## bhpcli query gov proposals

按条件查询提案。

```shell script
bhpcli query gov proposals [flags]
```

输出

```shell script
[
    {
        "content":{
            "type":"cosmos-sdk/TextProposal",
            "value":{
                "title":"BHP Test Proposal",
                "description":"My awesome proposal"
            }
        },
        "id":"1",
        "proposal_status":"DepositPeriod",
        "final_tally_result":{
            "yes":"0",
            "abstain":"0",
            "no":"0",
            "no_with_veto":"0"
        },
        "submit_time":"2020-08-18T07:39:32.075083338Z",
        "deposit_end_time":"2020-08-20T07:39:32.075083338Z",
        "total_deposit":[
            {
                "denom":"abhp",
                "amount":"10"
            }
        ],
        "voting_start_time":"0001-01-01T00:00:00Z",
        "voting_end_time":"0001-01-01T00:00:00Z"
    }
]
```

**标识：**

| 名称, 速记  | 类型    | 必须 | 默认 | 描述                                  |
| ----------- | ------- | -------- | ---- | ------------------------------------- |
| --node | host:port |         |      | RPC接口，默认 "tcp://localhost:26657"          |
| --depositor | Address |          |      | 按抵押人地址过滤提案                  |
| --limit     | uint    |          |      | 限制返回提案的个数， 默认返回所有提案 |
| --status    | string  |          |      | 按状态过滤提案 (deposit_period/voting_period/passed/rejected)    |
| --voter     | Address |          |      | 按投票人地址过滤提案                  |
## bhpcli query gov vote

查询投票信息。

```shell script
bhpcli query gov vote [proposal-id] [voter-addr] [flags]
```
## bhpcli query gov votes

查询提案的所有投票信息。

```shell script
bhpcli query gov votes [proposal-id] [flags]
```
## bhpcli query gov param

指定参数类型查询治理过程的参数。 (参数类型为voting：投票时长，tallying：统计参数，deposit：抵押参数)

```shell script
bhpcli query gov param [param-type] [flags]
```
示例
```shell script
bhpcli query gov param voting
{
    "voting_period":"172800000000000"
}
```

```shell script
bhpcli query gov param tallying
{
    "quorum":"0.334000000000000000",
    "threshold":"0.500000000000000000",
    "veto":"0.334000000000000000"
}
```

```shell script
bhpcli query gov param deposit
{
    "min_deposit":[
        {
            "denom":"abhp",
            "amount":"1000000000"
        }
    ],
    "max_deposit_period":"172800000000000"
}
```

## bhpcli query gov params

查询治理过程的所有参数。        

```shell script
bhpcli query gov params [flags]
```

```shell script
{
    "voting_params":{
        "voting_period":"172800000000000"
    },
    "tally_params":{
        "quorum":"0.334000000000000000",
        "threshold":"0.500000000000000000",
        "veto":"0.334000000000000000"
    },
    "deposit_params":{
        "min_deposit":[
            {
                "denom":"abhp",
                "amount":"1000000000"
            }
        ],
        "max_deposit_period":"172800000000000"
    }
}
```

## bhpcli query gov proposer

查询一个提议信息

```shell script
bhpcli query gov proposer [proposal-id] [flags]
```

响应

```shell script
{
    "proposal_id":"1",
    "proposer":"bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y"
}
```
## bhpcli query gov deposit

查询指定提案的抵押信息。

```shell script
bhpcli query gov deposit [proposal-id] [depositer-addr] [flags]
```

示例

```shell script
bhpcli query gov deposit 1 bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs
```

输出

```shell script
{
    "proposal_id":"1",
    "depositor":"bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs",
    "amount":[
        {
            "denom":"abhp",
            "amount":"1000000020"
        }
    ]
}
```

## bhpcli query gov deposits

查询指定提案的所有抵押信息。

```shell script
bhpcli query gov deposits [proposal-id] [flags]
```

示例

```shell script
bhpcli query gov deposits 1
```

响应
```shell script
[
    {
        "proposal_id":"1",
        "depositor":"bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs",
        "amount":[
            {
                "denom":"abhp",
                "amount":"1000000020"
            }
        ]
    },
    {
        "proposal_id":"1",
        "depositor":"bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y",
        "amount":[
            {
                "denom":"abhp",
                "amount":"10"
            }
        ]
    }
]
```

## bhpcli query gov tally

查询提案投票的统计信息

```shell script
bhpcli query gov tally [proposal-id] [flags]
```

示例

```shell script
bhpcli query gov tally 1
```

响应

```shell script
{
    "yes":"200000000000000",
    "abstain":"0",
    "no":"200000000000000",
    "no_with_veto":"0"
}
```

## bhpcli tx gov submit-proposal

提交提案以及初始化抵押金额。

```shell script
 bhpcli tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey
```

```shell script
bhpcli tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
```

**标识：**

| 名称, 速记               | 类型   | 必须 | 默认  | 描述                                                                                           |
| ------------------------ | ------ | -------- | ----- | ---------------------------------------------------------------------------------------------- |
| --deposit                | Coin   | 是       |       | 初始抵押金额(至少最小抵押金额的30% of)                                                         |
| --description            | string | 是       |       | 提案的描述信息                                                                                 |
| --param                  | string |          |       | 提案修改的参数，例如`mint/Inflation=0.050`                                                     |
| --title                  | string | 是       |       | 提案的标题                                                                                     |
| --type                   | string | 是       |       | 提案的类型（PlainText/Parameter/SoftwareUpgrade/SoftwareHalt/CommunityTaxUsage/TokenAddition） |                                                                                 |                                                                      |                                                                                |

### 提交文本交易提案

示例
```shell script
bhpcli tx gov submit-proposal --title="BHP Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10abhp" --fees="2abhp" --from node0 --home /root/bhp/build/node0/bhpcli --chain-id=testing
```
响应
```shell script
{
    "height":"0",
    "txhash":"949AEBF965CB0635D1CD9BC7D088A3505812081D97FE7B7FF061B2F82E23C169",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"submit_proposal"}]}]}]",
    "logs":[
        {
            "msg_index":0,
            "success":true,
            "log":"",
            "events":[
                {
                    "type":"message",
                    "attributes":[
                        {
                            "key":"action",
                            "value":"submit_proposal"
                        }
                    ]
                }
            ]
        }
    ]
}
```

### 提交参数修改提案

## bhpcli tx gov deposit

为有效的提案抵押通证。

```shell script
bhpcli tx gov deposit [proposal-id] [deposit] [flags]
```
示例

```shell script
bhpcli tx gov deposit 1 10abhp --from=node1 --fees=2abhp --chain-id=testing --home /root/bhp/build/node1/bhpcli
```
响应
```shell script
{
    "height":"0",
    "txhash":"2E957C52C000EB0962F3E5CB4F9C9232A344F9276810B8FF2FD7F8FA2C24F877",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"deposit"}]}]}]",
    "logs":[
        {
            "msg_index":0,
            "success":true,
            "log":"",
            "events":[
                {
                    "type":"message",
                    "attributes":[
                        {
                            "key":"action",
                            "value":"deposit"
                        }
                    ]
                }
            ]
        }
    ]
}
```
## bhpcli tx gov vote

对有效的提案投票，选项：Yes/No/NoWithVeto/Abstain。

:::tip
[No VS NoWithVeto](../concepts/governance.md#抵押退款和销毁)

在投票期内，只有验证人和委托人可以对提案进行投票。
:::

```shell script
bhpcli tx gov vote [proposal-id] [option] [flags]
```

**标识：**

| 名称, 速记    | 类型   | 必须 | 默认 | 描述                            |
| ------------- | ------ | -------- | ---- | ------------------------------- |
| option      | string | 是       |      | 选项：Yes/No/NoWithVeto/Abstain |
| proposal-id | uint   | 是       |      | 提案Id                          |

示例
```shell script
bhpcli tx gov vote 1 Yes --from node1 --fees=2abhp --chain-id=testing --home /root/bhp/build/node1/bhpcli
```
响应
```shell script
{
    "height":"0",
    "txhash":"9512E58E298B63CE277CE6D7EB145E0EC41761B9BA1057ADE8CC14C72402C707",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"vote"}]}]}]",
    "logs":[
        {
            "msg_index":0,
            "success":true,
            "log":"",
            "events":[
                {
                    "type":"message",
                    "attributes":[
                        {
                            "key":"action",
                            "value":"vote"
                        }
                    ]
                }
            ]
        }
    ]
}
```
如果查询交易显示以下结果，可能是抵押金额不够，导致投票未开始
```shell script
{
    "height":"68704",
    "txhash":"9512E58E298B63CE277CE6D7EB145E0EC41761B9BA1057ADE8CC14C72402C707",
    "code":2,
    "raw_log":"[{"msg_index":0,"success":false,"log":"{\"codespace\":\"gov\",\"code\":2,\"message\":\"inactive proposal with id 1\"}","events":[{"type":"message","attributes":[{"key":"action","value":"vote"}]}]}]",
    "logs":[
        {
            "msg_index":0,
            "success":false,
            "log":"{"codespace":"gov","code":2,"message":"inactive proposal with id 1"}",
            "events":[
                {
                    "type":"message",
                    "attributes":[
                        {
                            "key":"action",
                            "value":"vote"
                        }
                    ]
                }
            ]
        }
    ],
    "gas_wanted":"200000",
    "gas_used":"26017",
    "tx":{
        "type":"cosmos-sdk/StdTx",
        "value":{
            "msg":[
                {
                    "type":"cosmos-sdk/MsgVote",
                    "value":{
                        "proposal_id":"1",
                        "voter":"bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs",
                        "option":"Yes"
                    }
                }
            ],
            "fee":{
                "amount":[
                    {
                        "denom":"abhp",
                        "amount":"2"
                    }
                ],
                "gas":"200000"
            },
            "signatures":[
                {
                    "pub_key":{
                        "type":"tendermint/PubKeySecp256k1",
                        "value":"ApAtsUCFUxtNuVXGXk3Spe/eSslg8xwiL0b0LIWjZzTY"
                    },
                    "signature":"es/BvznvzbvMPKX9vCR0S37FQFHWHcu907DFpbj0dIFmduDvctL3R0Mt/DQvcwQy4HuHHjmuR+cG7rD9IWUmNA=="
                }
            ],
            "memo":""
        }
    },
    "timestamp":"2020-08-18T08:13:43Z",
    "events":[
        {
            "type":"message",
            "attributes":[
                {
                    "key":"action",
                    "value":"vote"
                }
            ]
        }
    ]
}
```