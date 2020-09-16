---
order: 1
---

# Validators Overview
## Introduction

The BHP is based on Cosmos SDK and Tendermint, which relies on a set of validators to secure the network. The role of validators is to run a full-node and participate in consensus by broadcasting votes which contain cryptographic signatures signed by their private keys. Validators commit new blocks to the blockchain and receive revenue in exchange for their work. They must also participate in governance by voting on proposals. Validators are weighted according to their total stake.

## Validator Responsibilities

Validators have two main responsibilities:

- Be able to constantly run a correct version of the software: Validators need to make sure that their servers are always online and their private keys are not compromised.
- Actively participate in governance: Validators are required to vote on every proposal.

Additionally, validators are expected to be active members of the community. They should always be up-to-date with the current state of the ecosystem so that they can easily adapt to any change.

## Validator Operation

### Create Validator

[bhpcli tx staking create-validator](https://docs.bhpnet.io/cli-client/staking.html#bhpcli-tx-staking-create-validator)

### Edit Validator

[bhpcli tx staking edit-validator](https://docs.bhpnet.io/cli-client/staking.html#bhpcli-tx-staking-edit-validator)

### Delegate bhp to Validator

[bhpcli tx staking delegate](https://docs.bhpnet.io/cli-client/staking.html#bhpcli-tx-staking-delegate)

### Withdraw rewards and gas

Withdraw rewards from a given delegation address,and optionally withdraw validator commission if the delegation address given is a validator operator.

[bhpcli tx distribution withdraw-rewards](https://docs.bhpnet.io/cli-client/slashing.html)

### Unjail a jailed validator

[bhpcli tx slashing unjail](https://docs.bhpnet.io/cli-client/slashing.html)

### Other Commands

More validator commands are available [view](/cli-client/staking.md)