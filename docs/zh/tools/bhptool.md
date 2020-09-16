# 命令行调试工具

## 介绍

`bhpdebug`是用于简单调试的工具。

接受十六进制和base64格式，并提供可读的响应。

通常，我们在日志中将字节编码为十六进制，但在JSON中将字节编码为base64

## 用法

### bhpdebug pubkey

公钥以不同的格式转换。下面得到相同的结果:

```
bhpdebug pubkey TZTQnfqOsi89SeoXVnIw+tnFJnr4X8qVC0U8AsEmFk4=
bhpdebug pubkey 4D94D09DFA8EB22F3D49EA17567230FAD9C5267AF85FCA950B453C02C126164E
```
returns
```shell script
Address: AB848E791827483D950C85F9CFC77D901FEE1E73
Hex: 4D94D09DFA8EB22F3D49EA17567230FAD9C5267AF85FCA950B453C02C126164E
JSON (base64): {"type":"tendermint/PubKeyEd25519","value":"TZTQnfqOsi89SeoXVnIw+tnFJnr4X8qVC0U8AsEmFk4="}
Bech32 Acc: bhppub1zcjduepqfk2dp80636ez702fagt4vu3sltvu2fn6lp0u49gtg57q9sfxze8qzl2yxh
Bech32 Validator Operator: bhpvaloperpub1zcjduepqfk2dp80636ez702fagt4vu3sltvu2fn6lp0u49gtg57q9sfxze8qvlg3x9
Bech32 Validator Consensus: bhpvalconspub1zcjduepqfk2dp80636ez702fagt4vu3sltvu2fn6lp0u49gtg57q9sfxze8qp40hfd
```

## bhpdebug tx

传入hex/base64编码的tx返回完整的JSON
```
bhpdebug tx <hex or base64 transaction>
bhpdebug tx wgEoKBapCj+oo2GaChRiJZkXalkgoeyuQHyuJVky9SG3URIUFA30bmXaMGXD0HKK1RADnkC+4ZgaDQoEYWJocBIFMTAwMDASDwoJCgRhYmhwEgEyEMCaDBpqCibrWumHIQPzCiGznMr8L+9qZ7DTw+o8VT6yPSfHLqoRgsHnoqibvRJAO7OK1bna1rkaBYF9dUDYneGC1s7mDiFazOqlaPS/n4ZgddD4yGHKnilfgeTx8Xvzm8M+oPRRksJVpU/ZcanqRQ==
```
returns
```shell script
{
  "type": "cosmos-sdk/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/MsgSend",
        "value": {
          "from_address": "bhp1vgjej9m2tys2rm9wgp72uf2ext6jrd63ykndxt",
          "to_address": "bhp1zsxlgmn9mgcxts7sw29d2yqrneqtacvcj2s2wn",
          "amount": [
            {
              "denom": "abhp",
              "amount": "10000"
            }
          ]
        }
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "abhp",
          "amount": "2"
        }
      ],
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "A/MKIbOcyvwv72pnsNPD6jxVPrI9J8cuqhGCweeiqJu9"
        },
        "signature": "O7OK1bna1rkaBYF9dUDYneGC1s7mDiFazOqlaPS/n4ZgddD4yGHKnilfgeTx8Xvzm8M+oPRRksJVpU/ZcanqRQ=="
      }
    ],
    "memo": ""
  }
}

```
## bhpdebug hack

生成带有密钥和时间戳（如果提供）的哈希锁

