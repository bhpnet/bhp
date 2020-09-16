---
order: 7
---

# gov

This module provides the basic functionalities for [Governance](../concepts/governance.md).

## Available Commands

| Name                                            | Description                                              |
| ----------------------------------------------- | ------------------------------------------------- |
| [query gov proposal](#bhpcli-query-gov-proposal)   | Query details for a proposal.     |
| [query gov proposals](#bhpcli-query-gov-proposals) | Query proposals by conditions          |
| [query gov vote](#bhpcli-query-gov-vote)           | Query details for a single vote on a proposal given its identifier. |
| [query gov votes](#bhpcli-query-gov-votes)         | Query vote details for a single proposal by its identifier.  |
| [query gov param](#bhpcli-query-gov-param)           | Query the all the parameters for the governance process.  |
| [query gov params](#bhpcli-query-gov-params)         | Query the all the parameters for the governance process.   |
| [query gov proposer](#bhpcli-query-gov-proposer)           | Query which address proposed a proposal with a given ID  |
| [query gov deposit](#bhpcli-query-gov-deposit)         | Query details for a single proposal deposit on a proposal by its identifier. |
| [query gov deposits](#bhpcli-query-gov-deposits)           | Query details for all deposits on a proposal.  |
| [query gov tally](#bhpcli-query-gov-tally)         | Query tally of votes on a proposal.  |
| [tx gov submit-proposal](#bhpcli-tx-gov-submit-proposal) | Submit a proposal along with an initial deposit.  |
| [tx gov deposit](#bhpcli-tx-gov-deposit)                 | Query details for a single proposal deposit on a proposal by its identifier.    |
| [tx gov vote](#bhpcli-tx-gov-vote)                       | Submit a vote for an active proposal, options: Yes/No/NoWithVeto/Abstain |

## bhpcli query gov proposal

Query details for a proposal. 

```shell script
bhpcli query gov proposal [proposal-id] [flags]
```

**Flags:**

| Name, shorthand | Type | Required | Default | Description            |
| --------------- | ---- | -------- | ------- | ---------------------- |
| proposal-id   | uint | Yes      |         | Identity of a proposal |

## bhpcli query gov proposals

Query proposals by conditions    

```shell script
bhpcli query gov proposals [flags]
```

Response

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

**Flags:**

| Name, shorthand | Type    | Required | Default | Description                                                         |
| ----------- | ------- | -------- | ---- | ------------------------------------- |
| --node | host:port |         |      | RPC port，default "tcp://localhost:26657"          |
| --depositor | Address |          |      | Filter proposals by depositor address       |
| --limit     | uint    |          |      | Limit to the latest [number] of proposals. Default to all proposals |
| --status    | string  |          |      | Filter proposals by status (deposit_period/voting_period/passed/rejected)    |
| --voter     | Address |          |      | Filter proposals by voter address              |

## bhpcli query gov vote

Query details for a single vote on a proposal given its identifier.

```shell script
bhpcli query gov vote [proposal-id] [voter-addr] [flags]
```
## bhpcli query gov votes

Query vote details for a single proposal by its identifier.

```shell script
bhpcli query gov votes [proposal-id] [flags]
```
## bhpcli query gov param

Query the all the parameters for the governance process.

```shell script
bhpcli query gov param [param-type] [flags]
```

Example

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

Query the all the parameters for the governance process.      

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

Query which address proposed a proposal with a given ID

```shell script
bhpcli query gov proposer [proposal-id] [flags]
```

Response

```shell script
{
    "proposal_id":"1",
    "proposer":"bhp1eesqv2r4v2al6dn5wavndm96fwth3y6s7dd87y"
}
```
## bhpcli query gov deposit

Query details for a single proposal deposit on a proposal by its identifier.

```shell script
bhpcli query gov deposit [proposal-id] [depositer-addr] [flags]
```

Example

```shell script
bhpcli query gov deposit 1 bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs
```

Response

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

Query details for all deposits on a proposal.

```shell script
bhpcli query gov deposits [proposal-id] [flags]
```

Example

```shell script
bhpcli query gov deposits 1
```

Response

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

Query tally of votes on a proposal.

```shell script
bhpcli query gov tally [proposal-id] [flags]
```

Example

```shell script
bhpcli query gov tally 1
```

Response

```shell script
{
    "yes":"200000000000000",
    "abstain":"0",
    "no":"200000000000000",
    "no_with_veto":"0"
}
```

## bhpcli tx gov submit-proposal

Submit a proposal along with an initial deposit.

```shell script
 bhpcli tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey
```

```shell script
bhpcli tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
```

**Flags:**

| Name, shorthand          | Type   | Required | Default | Description                                                                                           |
| ------------------------ | ------ | -------- | ----- | ---------------------------------------------------------------------------------------------- |
| --deposit                | Coin   | Yes |       | Initial deposit of the proposal        |
| --description            | string | Yes      |         | Description of the proposal                                              |
| --param                  | string |          |         | On-chain Parameter to be changed, eg. mint/Inflation=0.050                                                     |
| --title                  | string | Yes      |         | Title of the proposal                                                   |
| --type                   | string | Yes      |         | ProposalType of the proposal（PlainText/Parameter/SoftwareUpgrade/SoftwareHalt/CommunityTaxUsage/TokenAddition） |       |  |                                                                       

## Submit a Text Proposal

Example

```shell script
bhpcli tx gov submit-proposal --title="BHP Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10abhp" --fees="2abhp" --from node0 --home /root/bhp/build/node0/bhpcli --chain-id=testing
```

Response

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

### Submit a Parameter Change Proposal

## bhpcli tx gov deposit

Submit a deposit for an active proposal.

```shell script
bhpcli tx gov deposit [proposal-id] [deposit] [flags]
```

Example

```shell script
bhpcli tx gov deposit 1 10abhp --from=node1 --fees=2abhp --chain-id=testing --home /root/bhp/build/node1/bhpcli
```

Response

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

Submit a vote for an active proposal, options: Yes/No/NoWithVeto/Abstain

:::tip
[No VS NoWithVeto](../concepts/governance.md)

Only validators and delegators can vote for proposals in the voting period.
:::

```shell script
bhpcli tx gov vote [proposal-id] [option] [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description                            |
| --------------- | ------ | -------- | ------- | -------------------------------------- |
| option        | string | Yes      |         | Vote option: Yes/No/NoWithVeto/Abstain |
| proposal-id   | uint   | Yes      |         | Identity of a proposal                 |

Example

```shell script
bhpcli tx gov vote 1 Yes --from node1 --fees=2abhp --chain-id=testing --home /root/bhp/build/node1/bhpcli
```

Response

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

If the query transaction shows the following results, it may be that the mortgage amount is insufficient, causing the vote not to start

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