---
order: 4
---

# config

bhpcli config 命令可以交互式地配置一些公共参数的默认值，例如chain-id，home，fee 和 node。完成配置后，后续的bhpcli命令可以省略对这些flag参数的指定。
```shell script
bhpcli config <key> [value] [flags]
```
示例
```shell script
bhpcli config chain-id testing
```
响应
```shell script
configuration saved to /root/.bhpcli/config/config.toml
```                                             