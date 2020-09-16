---
order: 5
---
# query
The `bhpcli query` command can query data on the chain.

## Available Subommands

| Name                               | Description                                                         |
| ---------------------------------- | ------------------------------------------------------------ |
| [tx](#bhpcli-query-tx)           | Query transaction by transaction hash |
| [block](#bhpcli-query-block)         | Query block information by transaction height                                                 |
| [txs](#bhpcli-query-txs)         | Paging query transactions by search criteria.                                   |
| [tendermint-validator-set](#bhpcli-query-tendermint-validator-set)     | Query validator details by height                  |
| [account](#bhpcli-query-account)           | Query account details |
| [supply](#bhpcli-query-supply)           | Query the total supply of coins of the chain |

## bhpcli query tx

Query transaction by transaction hash.
```shell script
bhpcli query tx [hash] [flags]
```
Examples
```shell script
bhpcli query tx 2F5EC6DAEB3AA6B6803B3F0A0BE9D6B5B7E849A9CC2BEE6A59A93EE72967995D
```
Response
```shell script
{
	"height": "63139",
	"txhash": "2F5EC6DAEB3AA6B6803B3F0A0BE9D6B5B7E849A9CC2BEE6A59A93EE72967995D",
	"raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq\"},{\"key\":\"module\",\"value\":\"bank\"},{\"key\":\"action\",\"value\":\"send\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"bhp1xmkrs4qh64uvc9rsnhsxulathrel7hrqgkdn7l\"},{\"key\":\"amount\",\"value\":\"1000000000abhp\"}]}]}]",
	"logs": [{
		"msg_index": 0,
		"success": true,
		"log": "",
		"events": [{
			"type": "message",
			"attributes": [{
				"key": "sender",
				"value": "bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq"
			}, {
				"key": "module",
				"value": "bank"
			}, {
				"key": "action",
				"value": "send"
			}]
		}, {
			"type": "transfer",
			"attributes": [{
				"key": "recipient",
				"value": "bhp1xmkrs4qh64uvc9rsnhsxulathrel7hrqgkdn7l"
			}, {
				"key": "amount",
				"value": "1000000000abhp"
			}]
		}]
	}],
	"gas_wanted": "200000",
	"gas_used": "38915",
	"tx": {
		"type": "cosmos-sdk/StdTx",
		"value": {
			"msg": [{
				"type": "cosmos-sdk/MsgSend",
				"value": {
					"from_address": "bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq",
					"to_address": "bhp1xmkrs4qh64uvc9rsnhsxulathrel7hrqgkdn7l",
					"amount": [{
						"denom": "abhp",
						"amount": "1000000000"
					}]
				}
			}],
			"fee": {
				"amount": [{
					"denom": "abhp",
					"amount": "2"
				}],
				"gas": "200000"
			},
			"signatures": [{
				"pub_key": {
					"type": "tendermint/PubKeySecp256k1",
					"value": "AxL2Oz/5Mcb1X6apIUc0kguakJbSIXG5kcYS4fhSraRY"
				},
				"signature": "EoKza4E3Ys3I8Tjle2hk6TiaFqbxgXWx/nu8fWwDFHle9wCJ46WsBQi89mjddBjEFzxrHT+rG3grD50i3pQkyw=="
			}],
			"memo": ""
		}
	},
	"timestamp": "2020-08-17T08:28:20Z",
	"events": [{
		"type": "message",
		"attributes": [{
			"key": "sender",
			"value": "bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq"
		}, {
			"key": "module",
			"value": "bank"
		}, {
			"key": "action",
			"value": "send"
		}]
	}, {
		"type": "transfer",
		"attributes": [{
			"key": "recipient",
			"value": "bhp1xmkrs4qh64uvc9rsnhsxulathrel7hrqgkdn7l"
		}, {
			"key": "amount",
			"value": "1000000000abhp"
		}]
	}]
}
```

## bhpcli query block

Obtain the verification data of the block at a given height. If the height is not specified, the latest height will be used as the default height.

Query block information by transaction height.
```shell script
bhpcli query block <block-height>
```
Examples
```shell script
bhpcli query block 63139
```
Response
```shell script
{
	"block_meta": {
		"block_id": {
			"hash": "2B6FAA90EF7C1859BFA51442755DEA7DE34F3671404B4A72796AD20C96A87291",
			"parts": {
				"total": "1",
				"hash": "5E59EC0BDA45C86A8E8DF6903C8AD0F54AF599FA163DB6EF46CB47B2C31C8B4A"
			}
		},
		"header": {
			"version": {
				"block": "10",
				"app": "0"
			},
			"chain_id": "testing",
			"height": "63139",
			"time": "2020-08-17T08:28:20.116035Z",
			"num_txs": "1",
			"total_txs": "62",
			"last_block_id": {
				"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
				"parts": {
					"total": "1",
					"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
				}
			},
			"last_commit_hash": "8A0317F1C797413EAC2EE7052C421E18A4C6E1BB85FA404E926D75C86FD9210D",
			"data_hash": "C0EEB5FDC736E71773EBB16C7F33B86658FD93F9832BEF28179296F78F857E54",
			"validators_hash": "30B96E1F21AF320E7AE39EC967B914025ECDF4F899D8777DC9CBD5A91BD6B212",
			"next_validators_hash": "30B96E1F21AF320E7AE39EC967B914025ECDF4F899D8777DC9CBD5A91BD6B212",
			"consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
			"app_hash": "3DDF9A24F6E7F6107A85D36724F04137A906D37C50EC1109F55E5D100CE1BDE9",
			"last_results_hash": "",
			"evidence_hash": "",
			"proposer_address": "03A8F953E60C6429D5EE12934D2394F070343893"
		}
	},
	"block": {
		"header": {
			"version": {
				"block": "10",
				"app": "0"
			},
			"chain_id": "testing",
			"height": "63139",
			"time": "2020-08-17T08:28:20.116035Z",
			"num_txs": "1",
			"total_txs": "62",
			"last_block_id": {
				"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
				"parts": {
					"total": "1",
					"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
				}
			},
			"last_commit_hash": "8A0317F1C797413EAC2EE7052C421E18A4C6E1BB85FA404E926D75C86FD9210D",
			"data_hash": "C0EEB5FDC736E71773EBB16C7F33B86658FD93F9832BEF28179296F78F857E54",
			"validators_hash": "30B96E1F21AF320E7AE39EC967B914025ECDF4F899D8777DC9CBD5A91BD6B212",
			"next_validators_hash": "30B96E1F21AF320E7AE39EC967B914025ECDF4F899D8777DC9CBD5A91BD6B212",
			"consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
			"app_hash": "3DDF9A24F6E7F6107A85D36724F04137A906D37C50EC1109F55E5D100CE1BDE9",
			"last_results_hash": "",
			"evidence_hash": "",
			"proposer_address": "03A8F953E60C6429D5EE12934D2394F070343893"
		},
		"data": {
			"txs": ["xwEoKBapCkSoo2GaChQIfyiqKHDgU4ZO5JlvKAxunWsjSRIUNuw4VBfVeMwUcJ3gbn+ruPP/XGAaEgoEYWJocBIKMTAwMDAwMDAwMBIPCgkKBGFiaHASATIQwJoMGmoKJuta6YchAxL2Oz/5Mcb1X6apIUc0kguakJbSIXG5kcYS4fhSraRYEkASgrNrgTdizcjxOOV7aGTpOJoWpvGBdbH+e7x9bAMUeV73AInjpawFCLz2aN10GMQXPGsdP6sbeCsPnSLelCTL"]
		},
		"evidence": {
			"evidence": null
		},
		"last_commit": {
			"block_id": {
				"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
				"parts": {
					"total": "1",
					"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
				}
			},
			"precommits": [{
				"type": 2,
				"height": "63138",
				"round": "0",
				"block_id": {
					"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
					"parts": {
						"total": "1",
						"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
					}
				},
				"timestamp": "2020-08-17T08:28:20.116035Z",
				"validator_address": "03A8F953E60C6429D5EE12934D2394F070343893",
				"validator_index": "0",
				"signature": "k/eYJCoVJKk3NNbNxbNip3q7xu9pvbAdaqwMend8mDQp7fb4WYYUKJ87ckwj/DAjCUjCdHrzFixnGCdDM+twCg=="
			}, {
				"type": 2,
				"height": "63138",
				"round": "0",
				"block_id": {
					"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
					"parts": {
						"total": "1",
						"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
					}
				},
				"timestamp": "2020-08-17T08:28:20.119539289Z",
				"validator_address": "96FF69D97AD48FC451345EA4AFAF2360BE3431D2",
				"validator_index": "1",
				"signature": "lbkWaa1yrqgX9ka0fZ2QSiSzEOBqVB4n6BB3WGFEy+IX2Xsl2/9jvgiIbsklgDBIbnKmlgWUPshy2NyfnxHuDg=="
			}, {
				"type": 2,
				"height": "63138",
				"round": "0",
				"block_id": {
					"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
					"parts": {
						"total": "1",
						"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
					}
				},
				"timestamp": "2020-08-17T08:28:20.11427164Z",
				"validator_address": "AFEFFB24F730A5F02E411D52778499F1D0C7F177",
				"validator_index": "2",
				"signature": "Tm+79hvzq8G9MppvuUXbOZVr1oOJ4LPmpmi0WbWQLIgbbaa8ui3VsmOJw5P7daHockKGPj4AO8/dcHxomfT/Aw=="
			}, {
				"type": 2,
				"height": "63138",
				"round": "0",
				"block_id": {
					"hash": "8DBA692D948F22B6C8BA32B0E2A03639770875979D2128EE0AE9B651CCF3571B",
					"parts": {
						"total": "1",
						"hash": "DC2B53F42E23F270C5C46563483BDDC7ABE05F4548010371CCB1CFA9305BF957"
					}
				},
				"timestamp": "2020-08-17T08:28:20.114110675Z",
				"validator_address": "F7B1127BFDA89EF7C6D28C8A9510160EA96B6DC3",
				"validator_index": "3",
				"signature": "RhAAvewl4aiVvfKLlL7Lo5ayc52glEdGjgSXRywUDRTbm8tBxNhpdkmqPlOOAzEkcg6D9dKBwYsKJ1JlnmRXBQ=="
			}]
		}
	}
}
```

## bhpcli query txs

Paging query transactions by search criteria.
```shell script
bhpcli query txs [flags]
```
Examples
```shell script
bhpcli query txs --events 'message.sender=bhp1gwmngxpzmpqd6uajaa4f45sexty6spzvyjqssa&message.action=send' --page 1 --limit 30
```
Response
```shell script
{
	"total_count": "1",
	"count": "1",
	"page_number": "1",
	"page_total": "1",
	"limit": "30",
	"txs": [{
		"height": "23218",
		"txhash": "849A4DA39A3D1C76410CA42AC0213ACC70C4B224295BC5C1AD81C4B166F95DBE",
		"raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"bhp1gwmngxpzmpqd6uajaa4f45sexty6spzvyjqssa\"},{\"key\":\"module\",\"value\":\"bank\"},{\"key\":\"action\",\"value\":\"send\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"bhp10au77t30anhmgrvnc8dxtl7lzjq76f583n4v0w\"},{\"key\":\"amount\",\"value\":\"1000000000abhp\"}]}]}]",
		"logs": [{
			"msg_index": 0,
			"success": true,
			"log": "",
			"events": [{
				"type": "message",
				"attributes": [{
					"key": "sender",
					"value": "bhp1gwmngxpzmpqd6uajaa4f45sexty6spzvyjqssa"
				}, {
					"key": "module",
					"value": "bank"
				}, {
					"key": "action",
					"value": "send"
				}]
			}, {
				"type": "transfer",
				"attributes": [{
					"key": "recipient",
					"value": "bhp10au77t30anhmgrvnc8dxtl7lzjq76f583n4v0w"
				}, {
					"key": "amount",
					"value": "1000000000abhp"
				}]
			}]
		}],
		"gas_wanted": "200000",
		"gas_used": "38775",
		"tx": {
			"type": "cosmos-sdk/StdTx",
			"value": {
				"msg": [{
					"type": "cosmos-sdk/MsgSend",
					"value": {
						"from_address": "bhp1gwmngxpzmpqd6uajaa4f45sexty6spzvyjqssa",
						"to_address": "bhp10au77t30anhmgrvnc8dxtl7lzjq76f583n4v0w",
						"amount": [{
							"denom": "abhp",
							"amount": "1000000000"
						}]
					}
				}],
				"fee": {
					"amount": [{
						"denom": "abhp",
						"amount": "60"
					}],
					"gas": "200000"
				},
				"signatures": [{
					"pub_key": {
						"type": "tendermint/PubKeySecp256k1",
						"value": "AymYTJCYWdsSa6Wgaowf5qnpyoFK1ESEh949TjQsi9Ar"
					},
					"signature": "8nyEts9dqSXA2u+VlFY0reAb92UWdagpJX0Y4kNNMMA+ZsyH4PCkaVrN+5a5s1ztkWMXNOiApe+C6ejXfyYEjA=="
				}],
				"memo": ""
			}
		},
		"timestamp": "2020-08-10T06:39:51Z",
		"events": [{
			"type": "message",
			"attributes": [{
				"key": "sender",
				"value": "bhp1gwmngxpzmpqd6uajaa4f45sexty6spzvyjqssa"
			}, {
				"key": "module",
				"value": "bank"
			}, {
				"key": "action",
				"value": "send"
			}]
		}, {
			"type": "transfer",
			"attributes": [{
				"key": "recipient",
				"value": "bhp10au77t30anhmgrvnc8dxtl7lzjq76f583n4v0w"
			}, {
				"key": "amount",
				"value": "1000000000abhp"
			}]
		}]
	}]
}
```

## bhpcli query tendermint-validator-set

Query validator details by height
```shell script
bhpcli query tendermint-validator-set [height] [flags]
```
Examples
```shell script
bhpcli query tendermint-validator-set 63139
```
Response
```shell script
{
	"block_height": "63139",
	"validators": [{
		"address": "bhpvalcons1qw50j5lxp3jzn40wz2f56gu57pcrgwyn85ggrx",
		"pub_key": "bhpvalconspub1zcjduepq64suxf6rrklnxu6e7jysqenv73k55ynqttr6s4q50aptsr7phtms9nexjs",
		"proposer_priority": "-4084850",
		"voting_power": "2000000"
	}, {
		"address": "bhpvalcons1jmlknkt66j8ug5f5t6j2ltervzlrgvwjsv0nea",
		"pub_key": "bhpvalconspub1zcjduepqsxwspe6h27k7fkk85xg46rzmg2jeksxuw6ry3jkyf7m8r38puyjq3asqef",
		"proposer_priority": "-3745450",
		"voting_power": "2000010"
	}, {
		"address": "bhpvalcons14lhlkf8hxzjlqtjpr4f80pye78gv0uth03q88d",
		"pub_key": "bhpvalconspub1zcjduepq6v05ucccwfv74d5en3lfnqya9zw6xwn50wml8rlge0kfpxk0jhksmnv7pa",
		"proposer_priority": "3915150",
		"voting_power": "2000000"
	}, {
		"address": "bhpvalcons177c3y7la4z0003kj3j9f2yqkp65kkmwrxgr5m2",
		"pub_key": "bhpvalconspub1zcjduepqwgpq9zemyvter67tltdg8a78ckwzkgxfpl3njz3fe6kk8dpgumessnt7mt",
		"proposer_priority": "3915150",
		"voting_power": "2000000"
	}]
}
```

## bhpcli query account

Query account details
```shell script
bhpcli query account [address] [flags]
```
Examples
```shell script
bhpcli query account bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq
```
Response
```shell script
{
	"type": "cosmos-sdk/Account",
	"value": {
		"address": "bhp1pplj323gwrs98pjwujvk72qvd6wkkg6fugyctq",
		"coins": [{
			"denom": "abhp",
			"amount": "9971399999610"
		}],
		"public_key": {
			"type": "tendermint/PubKeySecp256k1",
			"value": "AxL2Oz/5Mcb1X6apIUc0kguakJbSIXG5kcYS4fhSraRY"
		},
		"account_number": "14",
		"sequence": "47"
	}
}
```

## bhpcli query supply

Query the total supply of coins of the chain
```shell script
bhpcli query supply [flags]
bhpcli query supply [command]
```
Examples
```shell script
bhpcli query supply total
```
Response
```shell script
[{
	"denom": "abhp",
	"amount": "5341870000000000"
}]
```