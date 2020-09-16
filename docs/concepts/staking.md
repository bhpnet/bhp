---
order: 5
---

# Staking
## Rrequirement
There is a BHP network working normally.You must installed bhpcli locally and version number is even.
## Concepts
### Voting power
Voting power is a consensus concept. BHP is a Byzantine-Fault-Tolerant POS blockchain network. During the consensus process, a set of validators will vote the proposal block. If a validator thinks the proposal block is valid, it will vote `yes`, otherwise, it will vote nil. The votes from different validator don't have the same weight. The weight of a vote is called the voting power of the corresponding validator.
### Validator
Validator is a full BHP node. If its voting power is zero, it is just a normal full node or a validator candidate. Once its voting power is positive, then it is a real validator.
### Delegator && Delegation
People that cannot, or do not want to run validator nodes, can still participate in the staking process as delegators. After delegating some bhp to validators, delegators will gain shares from corresponding validators. Delegating bhp is also called bonding bhp to validators. Later we will have detailed description on it. Besides, a validator operator is also a delegator. Usually, a validator operator not only has delegation on its own validator, but also can have delegation on other validators.
::: danger
**It is strongly NOT recommended that validator operator COMPLETELY unbind self-delegation bhp, Cause the validator will be jailed (removed out of validator set) if he do so. The delegator who bonded bhp to this validator will also suffer losses.
So, it is recommended that validator operator reserve at least 1 bhp while unbonding bhp.**
**If a validator completely unbind sele-delegation.You can use `bhpcli tx staking delegate` and `bhpcli tx slashing unjail` to restore.**
:::
### Validator Candidates
The quantity of validators can't increase infinitely. Too many validators may result in low efficient consensus which slows down the blockchain TPS. So Byzantine-Fault-Tolerant POS blockchain network will have a limiation to the validator quantity. Usually, the value is 21. If more than 21 full nodes apply to join validator set. Then only these nodes with top 21 most bonded bhp will be real validators. Others will be validator candidates and will be descending sorted according to their bonded bhp amount. Once the one or more validators are kicked out from validator set, then the top candidates will be added into validator set automatically.
### Bond && Unbond && Unbonding Period
Validator operators must bond their liquid bhp to their validators. The validator voting power is proportional to the bonded bhp including both self-bonded bhp and bhp from other delegators. Validator operators can lower their own bonded bhp by sending unbond transactions. Delegators can also lower bonded bhp by sending unbond transactions. However, these unbonded bhp won't become liquid bhp immediately. After the unbond transactions are executed, the corresponding validator operators or delegators can't sending unbond transactions on the same validators again until the unbonding period is end. Usually the unbonding period is three weeks. Once the unbonding period is end, the unbonded bhp will become liquid bhp automatically. The unbonding period mechanism makes great contribution to the security of POS blockchain network. Besides, if the self-bonded bhp equals to zero, then the corresponding validator will be removed out of validator set.
### Redelegate
Delegators can transfer their delegation from one validator to another one. Redelegation can be devided into two steps: ubond from first validator and bond to another validator. As we have talked above, ubond operation can't be completed immediately until unbonding period is end, which means delegators can't send another redelegation transactions immediately.
### Evidence && Slash
The Byzantine-Fault-Tolerant POS blockchain network assume that the Byzantine nodes possess less than 1/3 of total voting power. These Byzantine nodes must be punished. So it is necessary to collect the evidence of Byzantine behavior. According to the evidence, stake module will aotumatically slash a certain amount of bhp from corresponding validators and delegators. The slashed bhp are just burned. Besides, the Byzantine validators will be removed from the validator set and put into jail, which means their voting power is zero. During the jail period, these nodes are not event validator candidates . Once the jail period is end, they can send transactions to unjail themselves and become validator candidates again.
### Rewards
As a delegator, the more bonded bhp it has on validator, the more rewards it will earn. For a validator operator, it will have extra rewards: validator commission. The rewards come from bhp inflation and transaction fee. 

### [command details](../cli-client/staking.md)