---
order: 6
---

# tx

## 可用命令

| 名称                               | 描述                                                         |
| ---------------------------------- | ------------------------------------------------------------ |
| [send](#bhpcli-tx-send)           | 发送令牌到另一个地址，此命令包括 generate, sign 和 broadcast 这些步骤 |
| [sign](#bhpcli-tx-sign)         | 签名生成的离线交易文件，该文件由generate-only标志生成                                                 |
| [broadcast](#bhpcli-tx-broadcast)         | 这个命令用于广播已离线签名的交易到网络                                   |

## bhpcli tx send

发送令牌到另一个地址，此命令包括 generate, sign 和 broadcast 这些步骤。

```shell script
bhpcli tx send [from_key_or_address] [to_address] [amount] [flags]
```
示例

案例中第一个参数是要转账的地址，第二个参数是接受地址，第三个参数是转账金额，第四个参数是手续费

```shell script
bhpcli tx send \
    bhp1vvknv4svw422p0tnthw5v7zgqre3yntj5srv5c \
    bhp1kqezagrup445e8mv4qpy003asfwq0acku9lex0 \
    9999500000abhp \
    --fees 500000abhp \
    --node=http://localhost:26657 \
    --chain-id testing
```
响应
```shell script
{
	"height": "0",
	"txhash": "48F2C2E7FCD3284617C16D757088FEC0E691CF8580CEE0D8A08239E4FFFE70B3",
	"raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"}]}]}]",
	"logs": [{
		"msg_index": 0,
		"success": true,
		"log": "",
		"events": [{
			"type": "message",
			"attributes": [{
				"key": "action",
				"value": "send"
			}]
		}]
	}]
}
```

## bhpcli tx sign

签名生成的离线交易文件，该文件由generate-only标志生成
```shell script
bhpcli tx sign [file] [flags]
```
[file]可以是json文件，也可以使用-在命令行输入，ctrl+d结束输入

示例
```shell script
tx send  bhp1j4kfjvsqdcpdzdnv4c9zkp5qxt49ngf2f5nmwh bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs 1000000000abhp --fees 10abhp --generate-only
bhpcli tx sign - --from=lff
```
响应
```shell script
{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"bhp1j4kfjvsqdcpdzdnv4c9zkp5qxt49ngf2f5nmwh","to_address":"bhp1t7gmv3qraqc7urcp2jqk2wv54p9jrevn9h3qhs","amount":[{"denom":"abhp","amount":"1000000000"}]}}],"fee":{"amount":[{"denom":"abhp","amount":"10"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A3GYyiZ+j9aJw1EIDZgEjUB8OitgX/U8ZqdT6t546yPS"},"signature":"50Zll4wAvph0zAail10/RmsZeE/4S5D3nNu7HeIIDddJ8BVQj0hIGi0BdDNy4kUAXHjXNSXb5q6JnzJ63wk2AQ=="}],"memo":""}}
```

## bhpcli tx broadcast
这个命令用于广播已离线签名的交易到网络。

可以将bhpcli tx sign 生成的数据进行广播
```shell script
bhpcli tx broadcast [file_path] [flags]
```
[file]可以是json文件，也可以使用-在命令行输入，ctrl+d结束输入

示例
```shell script
bhpcli tx broadcast - --from=lff
```
响应
```shell script
{"height":"0","txhash":"B939CB05BF80DD81129AFA73697B254503EE81E103A9E13F17AB0B48D9993CB0","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"}]}]}]","logs":[{"msg_index":0,"success":true,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"send"}]}]}]}
```