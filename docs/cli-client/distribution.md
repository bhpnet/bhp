---
order: 8
---

# distribution

The distribution module allows you to manage your Staking Rewards

## Available Subommands

| 名称                                                            | 描述                                                                                           |
| --------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| [query distribution params](#bhpcli-query-distribution-params)      | Query distribution params                                                                               |
| [query distribution validator-outstanding-rewards](#bhpcli-query-distribution-validator-outstanding-rewards)       | Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.                                                            |
| [query distribution commission](#bhpcli-query-distribution-commission) | Query validator commission rewards from delegators to that validator.       |
| [query distribution slashes](#bhpcli-query-distribution-slashes)      | Query all slashes of a validator for a given block range. |
| [query distribution rewards](#bhpcli-query-distribution-rewards)| Query all rewards earned by a delegator, optionally restrict to rewards from a single validator. |
| [query distribution community-pool](#bhpcli-query-distribution-community-pool)| Query all coins in the community pool which is under Governance control. |
| [tx distribution set-withdraw-addr](#bhpcli-tx-distribution-set-withdraw-addr) | Set the withdraw address for rewards associated with a delegator address.               |
| [tx distribution withdraw-rewards](#bhpcli-tx-distribution-withdraw-rewards)      | Withdraw rewards from a given delegation address,and optionally withdraw validator commission if the delegation address given is a validator operator. |
| [tx distribution withdraw-all-rewards](#bhpcli-tx-distribution-withdraw-all-rewards)|  Withdraw all rewards for a single delegator. |

## bhpcli query distribution params

Query distribution params

```shell script
bhpcli query distribution params [flags]
```

Response
```shell script
{
    "community_tax":"0.020000000000000000",
    "base_proposer_reward":"0.010000000000000000",
    "bonus_proposer_reward":"0.040000000000000000",
    "withdraw_addr_enabled":true
}
```

## bhpcli query distribution validator-outstanding-rewards

Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.

```shell script
bhpcli query distribution validator-outstanding-rewards [validator] [flags]
```

## bhpcli query distribution commission

Query validator commission rewards from delegators to that validator.

```shell script
bhpcli query distribution commission [validator] [flags]
```

Example
```shell script
bhpcli query distribution commission bhpvaloper1rr0u7tgaz07y670tpfqma7p8nmscpk6wcghkfr --home /root/bhp/build/node1/bhpcli --chain-id=testing
```
Response
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

Query all slashes of a validator for a given block range.

```shell script
bhpcli query distribution slashes [validator] [start-height] [end-height] [flags]
```

Example

```shell script
bhpcli query distribution slashes bhpvaloper1rr0u7tgaz07y670tpfqma7p8nmscpk6wcghkfr 0 80000
```

Response

```shell script
[
    {
        "validator_period":"2",
        "fraction":"0.010000000000000000"
    }
]
```

## bhpcli query distribution rewards

Query all rewards earned by a delegator, optionally restrict to rewards from a single validator.

```shell script
bhpcli query distribution rewards [delegator-addr] [<validator-addr>] [flags]
```

Example

```shell script
bhpcli query distribution rewards bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y
```

Response

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

Response

```shell script
[
    {
        "denom":"abhp",
        "amount":"1907271195002.995800000000000000"
    }
]
```

## bhpcli query distribution community-pool

Query all coins in the community pool which is under Governance control.

```shell script
bhpcli query distribution community-pool [flags]
```

Response

```shell script
[
    {
        "denom":"abhp",
        "amount":"30380.356422391918046419"
    }
]
```

## bhpcli tx distribution set-withdraw-addr

Set the withdraw address for rewards associated with a delegator address.

```shell script
bhpcli tx distribution set-withdraw-addr [withdraw-addr] [flags]
```

Example

```shell script
bhpcli tx distribution set-withdraw-addr bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs --from node0 --fees=2abhp --home /root/bhp/build/node0/bhpcli --chain-id=testing
```

Response

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

Withdraw rewards from a given delegation address,and optionally withdraw validator commission if the delegation address given is a validator operator.

```shell script
bhpcli tx distribution withdraw-rewards [validator-addr] [flags]
```

Example

```shell script
bhpcli tx distribution withdraw-rewards bhpvaloper1kqezagrup445e8mv4qpy003asfwq0ackkxmrf0 --commission --from link --fees=500000abhp --chain-id testing
```

Response

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

Withdraw all rewards for a single delegator.

```shell script
bhpcli tx distribution withdraw-all-rewards [flags]
```

Example

```shell script
bhpcli tx distribution withdraw-all-rewards --from node0 --home /root/bhp/build/node0/bhpcli --fees=2abhp --chain-id=testing
```

Response

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
