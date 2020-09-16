---
order: 2
---

# keys

Keys模块用于管理BHP的Tendermint本地密钥库（钱包）。

## 可用命令

| 名称                               | 描述                                                         |
| ---------------------------------- | ------------------------------------------------------------ |
| [mnemonic](#bhpcli-keys-mnemonic) | 通过读取系统熵来创建bip39助记词，也可以称为种子短语 |
| [add](#bhpcli-keys-add)           | 创建新密钥、从助记词导入已有密钥 |
| [list](#bhpcli-keys-list)         | 列出所有密钥                                                 |
| [show](#bhpcli-keys-show)         | 根据给定的名称显示密钥信息                                   |
| [export](#bhpcli-keys-export)     | 以Ascii加密格式从本地密钥库导出私钥。                   |
| [delete](#bhpcli-keys-delete)     | 删除指定的密钥                                               |
| [import](#bhpcli-keys-import) | 将ASCII加密的私钥导入到本地密钥库中。 |
| [update](#bhpcli-keys-update)     | 更改用于保护私钥的密码                                       |

### bhpcli keys mnemonic

通过读取系统熵来创建24个单词组成的bip39助记词（也称为种子短语）。如果需要传递自定义的熵，请使用`unsafe-entropy`模式。

```bash
bhpcli keys mnemonic <flags>
```

**标志：**


| 名称，速记       | 默认 | 描述                                     | 必须 |
| ---------------- | ---- | ---------------------------------------- | ---- |
| --unsafe-entropy |      | 提示用户提供自定义熵，而不是通过系统生成 |      |

**创建助记词**

```bash	
bhpcli keys mnemonic
```

执行上述命令后你将得到24个单词组成的助记词，例如：

```bash
sunset print vault mail bleak mule pigeon major denial carpet machine student alpha zone lazy figure universe satoshi luxury hand bunker silver method horror
```

## bhpcli keys add

创建一个新的密钥（钱包），或通过助记词/密钥库导入已有密钥。

```bash
bhpcli keys add <key-name> <flags>
```
**标志：**

| 名称，速记           | 默认值 | 描述                                               | 必须 |
| -------------------- | ------ | -------------------------------------------------- | ---- |
| --account            |        | HD推导的账号                                       |      |
| --dry-run            |        | 本地模拟交易，不会向本地密钥库中添加密钥           |      |
| --index              |        | HD推导的索引号                                     |      |
| -i, --interactive    |        | 交互式提示用户使用BIP39短语和助记符                |      |
| --ledger             |        | 在Ledger设备上存储私钥的本地引用                   |      |
| --no-backup          |        | 不输出助记词（如果有其他人正在观察你在终端的操作） |      |
| --recover            |        | 提供助记词以恢复现有密钥而不是新建                 |      |
| --multisig           |        | 创建多签密钥                                       |      |
| --multisig-threshold |        | 指定多签密钥最少签名数                             |      |

#### 创建密钥

**示例：**

```bash
bhpcli keys add mykey1
```

**响应：**

```bash
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:

- name: mykey1
  type: local
  address: bhp1uzeawqdv6wpw9qhv6wlh50vtwz2877vy9wn4tk
  pubkey: bhppub1addwnpepqfexthpf6fgc0296rqd9fvldsuh2eyjf6v3q2cyvlhs0ggn60nj9qw7gkdh
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.
# 请注意下面的助记词非常重要，一定要备份
affair index need gaze duty bind desert deputy pass task caution ride century law scout robust autumn across apart endorse ride harbor correct novel
```

执行该命令后输入并确认密码，将生成一个新的密钥。密码至少8个字符。

:::warning
**重要**

写下助记词并保存在安全的地方！如果你不慎忘记密码或丢失了密钥，这是恢复账户的唯一方法。
:::

#### 通过助记词生成密钥

如果你想要自己手动生成助记词后，生成对应的密钥，可以通过**-i**参数指定。

```bash
bhpcli keys add mykey -i
```

系统会要求你输入该密钥的秘密和你的助记词

```bash
override the existing name mykey3 [y/N]: y
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
> Enter your bip39 mnemonic, or hit enter to generate one.
sunset print vault mail bleak mule pigeon major denial carpet machine student alpha zone lazy figure universe satoshi luxury hand bunker silver method horror
> Enter your bip39 passphrase. This is combined with the mnemonic to derive the seed. Most users should just hit enter to use the default, ""
```

####  通过助记词恢复密钥

如果你忘记了密码或丢失了密钥，或者你想在其他地方使用密钥，可以通过助记词短语来恢复密钥。

```bash
bhpcli keys add mykey --recover
```

系统会要求你输入密钥的新密码并确认，然后输入助记词。然后你将得到恢复的密钥。

```bash
Enter a passphrase for your key:
Repeat the passphrase:
Enter your recovery seed phrase:
```

#### 创建多签名密钥

以下例子为创建一个包含3个子密钥的多签密钥，且指定最小签名数为2。只有交易签名数大于等于2时，该交易才会被广播。

```shell
bhpcli keys add <multisig-keyname> --multisig-threshold=2 --multisig=<signer-keyname-1>,<signer-keyname-2>,<signer-keyname-3>
```

:::tip
`<signer-keyname>` 可以为“local/offline/ledger”类型，但不允许为“multi”类型。

如果你没有子密钥的所有许可，则可先请求获取pubkeys并以此创建“offline”密钥，然后你将可以创建多签密钥。

其中“offline”类型密钥可以通过“bhpcli keys add --pubkey”命令创建。
:::



## bhpcli keys list

返回此密钥管理器存储的所有密钥的名称、类型、地址和公钥列表。

```bash
bhpcli keys list [flags]
# Flags:
#     --indent 格式化输出的JSON数据
```

**示例：**

```bash
bhpcli keys list
```

## bhpcli keys show

查询本地密钥的详细信息。

```bash	
bhpcli keys show [name [name...]] [flags]
```

**标志：**

| 名称，速记           | 默认值 | 描述                               | 必须 |
| -------------------- | ------ | ---------------------------------- | ---- |
| -a, --address        |        | 只输出地址                         |      |
| --bech               | acc    | Bech32编码的前缀（acc\|val\|cons） |      |
| --indent             |        | 添加json格式缩进                   |      |
| --multisig-threshold | 1      | K out of N required signatures     |      |
| -p, --pubkey         |        | 仅输出公钥                         |      |

#### 查看指定密钥

```bash
bhpcli keys show mykey
```

执行结果：

```bash
- name: testing
  type: local
  address: bhp1rr0u7tgaz07y670tpfqma7p8nmscpk6wjtnvxr
  pubkey: bhppub1addwnpepqdag42nacaekg7234cvdpvc8n8rafazwrfd8ss3uu3kh0jes45m26dw0xu5
  mnemonic: ""
  threshold: 0
  pubkeys: []
```

#### 查询验证人operator地址

如果某个地址已绑定成为验证人operator（用于创建验证人的地址），则可以使用`--bech val`获取其` bhpvaloper`前缀的operator地址和` bhpvaloper`前缀的公钥地址：

```bash
bhpcli keys show mykey --bech val
```

执行结果：

```bash
- name: testing
  type: local
  address: bhpvaloper1rr0u7tgaz07y670tpfqma7p8nmscpk6wcghkfr
  pubkey: bhpvaloperpub1addwnpepqdag42nacaekg7234cvdpvc8n8rafazwrfd8ss3uu3kh0jes45m260ha0v7
  mnemonic: ""
  threshold: 0
  pubkeys: []
```

### bhpcli keys export

以Ascii加密格式从本地密钥库导出私钥。

```bash
bhpcli keys export <name> [flags]
```

示例：

```bash
bhpcli keys export testing 
# Enter passphrase to decrypt your key:
# Enter passphrase to encrypt the exported key:
```

执行结果：(你需要将下面的内容保持到文件中)

```bash
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: AD3373B0EA5DF8BB295CB6A15C96A6EE

BG5oP7KCqNkqTpHOFrKq5UZAvQTtnEDAv9J2+BE64sq1TlUl9wkUGUfzE41/8f4A
R4G+0A/exgjEZh7OAbWAVTxGzZ54XtYvj3yMue0=
=ZvB0
-----END TENDERMINT PRIVATE KEY-----
```

### bhpcli keys delete

根据名称删除本地密钥

```bash
bhpcli keys delete <name> [flags]
```

**标志：**

| 名称，速记  | 默认  | 描述                                   | 必须 |
| ----------- | ----- | -------------------------------------- | ---- |
| -f, --force | false | 无密码强制删除秘钥                     |      |
| -y, --yes   | false | 删除离线和ledger引用密钥时跳过确认提示 |      |

示例：

```bash
bhpcli keys delete mykey
# DANGER - enter password to permanently delete key:
# Key deleted forever (uh oh!)
```

### bhpcli keys import

将ASCII加密的私钥导入到本地密钥库中。

```bash
bhpcli keys import <name> <keyfile> [flags]
```

示例：

```bash
# keyfile 内容为bhpcli keys export 命令导出
bhpcli keys import mykey keyfile
```

### bhpcli keys update

更改用于保护私钥的密码。

```bash
bhpcli keys update <name>
```

**修改本地密钥的密码**

```bash
bhpcli keys update mykey
```

系统会要求你输入当前密钥和新密码并确认。

```bash
# Enter the current passphrase:
# Enter the new passphrase:
# Repeat the new passphrase:
# Password successfully updated!
```

