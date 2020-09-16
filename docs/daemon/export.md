---
order: 5
---

# Export Blockchain State

BHP can export the blockchain state and output to a json-format string which can be used as the genesis file of a new blockchain.

````shell script
bhpd export [flags]
````

Flags

| Name, shorthand   | type   | Required | Default      | Description                                                  |
| ----------------- | ------ | -------- | ------------ | ------------------------------------------------------------ |
| --for-zero-height | bool   |          | false        | Do some clean up work before exporting state. If you want to use the exported state to start a new blockchain, please add this flag. Otherwise, just leave out it |
| --height          | uint   |          | 0            | Export state from a particular height, default value is 0 which means to export the latest state |
| --home            | string |          | $HOME/.bhpd  | Specify the directory which stores node config and blockchain data |
| --jail-whitelist     | strings |        |  |  List of validators to not jail state export                        |

## Examples

Export the current blockchain state

```bash
bhpd export --home=<path-to-your-home>
```

Export blockchain state from a particular height, the height must be an existing snapshot height

```bash
bhpd export --height 10000 --home=<path-to-your-home>
```

If you want to export the blockchain state from a particular height and use the exported state as genesis state of another blockchain

```bash
bhpd export --height 10000 --for-zero-height --home=<path-to-your-home>
```
