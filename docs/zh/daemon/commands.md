---
order: 2
---

# 命令

## 简介

BHP守护进程命令允许您初始化，启动，配置节点或者生成创世文件等。

## 用法

```bash
bhpd <command>
```

## 可用命令

|名称                       | 描述                                        |
| :----------------------------------------------- | :------------------------------------------ |
|                [init](#bhpd-init)                | 初始化验证人，p2p，创世纪和应用程序配置文件 |
| [add-genesis-account](#bhpd-add-genesis-account) | 将创世帐户添加到genesis.json                |
|               [gentx](#bhpd-gentx)               | 生成创建验证人的交易。                      |
|      [collect-gentxs](#bhpd-collect-gentxs)      | 将生成的质押交易添加到创世文件              |
| [validate-genesis](#bhpd-validate-genesis) | 在默认位置或作为参数传递的位置验证生成文件 |



### bhpd init

初始化验证人，p2p，创世纪和应用程序配置文件

```bash	
bhpd init [moniker] [flags]
```

**示例：**

```bash
bhpd init testing --chain-id=testing
```

### bhpd add-genesis-account

将创世帐户添加到genesis.json

```bash
 bhpd add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
 # address_or_key_name bhp地址或者是账户名称
 # coin 初始化token数，例如10000000abhp
```

**标志：**

| 名称，速记           | 默认     | 描述                                                         | 必须 |
| :------------------- | -------- | ------------------------------------------------------------ | ---- |
| --home-client        | /.bhpcli | 客户端主目录                                                 |      |
| --vesting-amount     |          | 需要锁定的token（数量小于等于前面初始化的数量）,锁定token无法交易，但是可以通过质押到验证者获取收益。 |      |
| --vesting-start-time |          | 指锁定开始时间，格式为unix时间戳                             |      |
| --vesting-end-time   |          | 指锁定结束时间，格式为unix时间戳                             |      |

**示例：**

```bash
# token 1000000000abhp
bhpd add-genesis-account mykey2 1000000000abhp

# 开始时间：2020-08-18 15:55:00（1597737830） 结束时间：2020-08-20 16:20:00 （1597911600）
bhpd add-genesis-account mykey 1000000000abhp --vesting-amount 800000000abhp  --vesting-start-time 1597737830  --vesting-end-time 1597911600
```

### bhpd gentx

生成创建验证人的交易。

```bash
bhpd gentx [flags]
```

**标志：**

| 名称，速记                   | 默认值          | 描述                              | 必须 |
| ---------------------------- | --------------- | --------------------------------- | ---- |
| --name                       |                 | 验证人的助记词名称                |      |
| --amount                     | 10000000000abhp | 用于质押的token的数量             |      |
| --commission-max-change-rate | 0.01            | 最高佣金变动率百分比(每日)        |      |
| --commission-max-rate        | 0.2             | 最高佣金率百分比                  |      |
| --commission-rate            | 0.1             | 初始佣金率百分比                  |      |
| --min-self-delegation        | 1               | 验证器所需的最小自委托（单位bhp） |      |

**示例：**

```bash
bhpd gentx --name mykey
```

### bhpd collect-gentxs

将生成的质押交易添加到创世文件

```bash
bhpd collect-gentxs [flags]

# Flags:
#     --gentx-dir string "gentx"目录默认为[--home]/config/gentx/ 如果多次执行，则会覆盖前一个。
```

**示例：**

```bash
bhpd collect-gentxs
```

### bhpd validate-genesis

在默认位置或作为参数传递的位置验证生成文件

```bash
bhpd validate-genesis [file]
```

**示例：**

```bash
bhpd validate-genesis
# File at \.bhpd\config\genesis.json is a valid genesis file
```

