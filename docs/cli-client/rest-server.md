---
order: 11
---

# rest-server

Start LCD (light-client daemon), a local REST server

```shell script
nohup bhpcli rest-server --laddr "tcp://0.0.0.0:26690" --home /root/bhp/build/node0/bhpcli/ --chain-id=testing > bhpcli.log &
```