---
order: 9
---

# slashing

## bhpcli tx slashing unjail

In Proof-of-Stake blockchain, validators will get block provisions by staking their token. But if they failed to keep online, they will be punished by slashing a small portion of their staked tokens. The offline validators will be removed from the validator set and put into jail, which means their voting power is zero. During the jail period, these nodes are not even validator candidates. Once the jail period ends, they can send `unjail` transactions to free themselves and become validator candidates again.

if tip `Validator still jailed, cannot yet be unjailed`,that means your validator is still in jail period, you can query the signing-info for the jail end time.

- Unjail a jailed validator

```shell script
bhpcli tx slashing unjail [flags] --from mykey
```


