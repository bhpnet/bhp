---
order: 8
---

# distribution

distribution模块用于管理自己的 Staking 收益。

## 可用命令

| 名称                                                            | 描述                                                                                           |
| --------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| [query distribution params](#bhpcli-query-distribution-params)      | 查询奖励参数                                                                          |
| [query distribution validator-outstanding-rewards](#bhpcli-query-distribution-validator-outstanding-rewards)       | 查询指定验证者分配未结清(未提取)的奖励情况                                                              |
| [query distribution commission](#bhpcli-query-distribution-commission) | 从委托人处查询该验证人的佣金奖励。                                                                          |
| [query distribution slashes](#bhpcli-query-distribution-slashes)      | 查询指定验证者在给定区块范围内的监禁记录。 |
| [query distribution rewards](#bhpcli-query-distribution-rewards)| 查询委托人获得的所有奖励，可选择查询为从单个验证人的获取奖励。 |
| [query distribution community-pool](#bhpcli-query-distribution-community-pool)| 查询链上治理的社区池中的所有币。 |
| [tx distribution set-withdraw-addr](#bhpcli-tx-distribution-set-withdraw-addr) | 设置提现地址                |
| [tx distribution withdraw-rewards](#bhpcli-tx-distribution-withdraw-rewards)      | 从某个验证者地址提取奖励，如果你的地址是验证人身份还可以提取佣金收益 |
| [tx distribution withdraw-all-rewards](#bhpcli-tx-distribution-withdraw-all-rewards)|  委托者取回自己的所有奖励 |

## bhpcli query distribution params

查询奖励参数

```shell script
bhpcli query distribution params [flags]
```

响应
```shell script
{
    "community_tax":"0.020000000000000000",
    "base_proposer_reward":"0.010000000000000000",
    "bonus_proposer_reward":"0.040000000000000000",
    "withdraw_addr_enabled":true
}
```

## bhpcli query distribution validator-outstanding-rewards

查询指定验证者分配未结清(未提取)的奖励情况

```shell script
bhpcli query distribution validator-outstanding-rewards [validator] [flags]
```

## bhpcli query distribution commission

查询该验证人的佣金奖励。
```shell script
bhpcli query distribution commission [validator] [flags]
```

示例
```shell script
bhpcli query distribution commission bhpvaloper1rr0u7tgaz07y670tpfqma7p8nmscpk6wcghkfr --home /root/bhp/build/node1/bhpcli --chain-id=testing
```
响应
```shell script
[
    {
        "denom":"abhp",
        "amount":"944987.175516317844464773"
    }
]
```

```shell script
bhpcli query distribution commission [validator] [flags]
```

## bhpcli query distribution slashes

查询指定验证者在给定区块范围内的监禁记录。

```shell script
bhpcli query distribution slashes [validator] [start-height] [end-height] [flags]
```

示例

```shell script
bhpcli query distribution slashes bhpvaloper1rr0u7tgaz07y670tpfqma7p8nmscpk6wcghkfr 0 80000
```

响应

```shell script
[
    {
        "validator_period":"2",
        "fraction":"0.010000000000000000"
    }
]
```

## bhpcli query distribution rewards

查询委托人获得的所有奖励，可选择查询从单个验证人的获取奖励。

```shell script
bhpcli query distribution rewards [delegator-addr] [<validator-addr>] [flags]
```

示例

```shell script
bhpcli query distribution rewards bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y
```

响应

```shell script
{
    "rewards":[
        {
            "validator_address":"bhpvaloper1eesqv2r4v2al6dn5wavndm96fwth3y6s5wfa3y",
            "reward":[
                {
                    "denom":"abhp",
                    "amount":"1907608650290.170200000000000000"
                }
            ]
        }
    ],
    "total":[
        {
            "denom":"abhp",
            "amount":"1907608650290.170200000000000000"
        }
    ]
}
```

```shell script
bhpcli query distribution rewards bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y bhpvaloper1eesqv2r4v2al6dn5wavndm96fwth3y6s5wfa3y
```

响应

```shell script
[
    {
        "denom":"abhp",
        "amount":"1907271195002.995800000000000000"
    }
]
```

## bhpcli query distribution community-pool

查询链上治理的社区池中的所有币。

```shell script
bhpcli query distribution community-pool [flags]
```

响应

```shell script
[
    {
        "denom":"abhp",
        "amount":"30380.356422391918046419"
    }
]
```

## bhpcli tx distribution set-withdraw-addr

设置提现地址

```shell script
bhpcli tx distribution set-withdraw-addr [withdraw-addr] [flags]
```

示例

```shell script
bhpcli tx distribution set-withdraw-addr bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs --from node0 --fees=2abhp --home /root/bhp/build/node0/bhpcli --chain-id=testing
```

响应

```shell script
{
    "height":"0",
    "txhash":"4E875B0039C6856C47E41AE85B3F801B53F34E657CB708BD857CA9462147CFFF",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"set_withdraw_address"}]}]}]",
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
                            "value":"set_withdraw_address"
                        }
                    ]
                }
            ]
        }
    ]
}
```

## bhpcli tx distribution withdraw-rewards

从某个验证者地址提取奖励，如果你的地址是验证人身份还可以提取佣金收益

```shell script
bhpcli tx distribution withdraw-rewards [validator-addr] [flags]
```

示例

```shell script
bhpcli tx distribution withdraw-rewards bhpvaloper1kqezagrup445e8mv4qpy003asfwq0ackkxmrf0 --commission --from link --fees=500000abhp --chain-id testing
```

响应

```shell script
{
    "height":"0",
    "txhash":"E8207DB752839CF8C5340D42E4D58CD4253F19FC265B8AE519D2418C5319AAAC",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"withdraw_delegator_reward"}]}]}]",
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
                            "value":"withdraw_delegator_reward"
                        }
                    ]
                }
            ]
        }
    ]
}
```

## bhpcli tx distribution withdraw-all-rewards

委托者取回自己的所有奖励

```shell script
bhpcli tx distribution withdraw-all-rewards [flags]
```

示例

```shell script
bhpcli tx distribution withdraw-all-rewards --from node0 --home /root/bhp/build/node0/bhpcli --fees=2abhp --chain-id=testing
```

响应

```shell script
{
    "height":"0",
    "txhash":"6D2F4654152BDE49001540A4B917427DAD98CEBF32920A74DC56899A84506791",
    "raw_log":"[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"withdraw_delegator_reward"}]}]}]",
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
                            "value":"withdraw_delegator_reward"
                        }
                    ]
                }
            ]
        }
    ]
}
```
