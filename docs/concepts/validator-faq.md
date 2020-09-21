---
order: 2
---

# Validator FAQ

### How to Become an BHP Validator

Refer to [Join The Mainnet](../getting-start/join-bhp-mainnet.md)

It is also recommended to test your validator setup on the [Testnet](../getting-start/join-bhp-testnet.md)before launch

:::tip
To earn more delegation for your validator, you are advised to:

- Undergo a security audit
- Open-source some dev tools and workflow
- Setup your own website to build your own image
:::

### What are hardware requirements

The minimum hardware requirements are mentioned here:[Hardware Requeirment](../daemon/intro.md)

### What are the different states a validator can be in

After a validator is created with a `create-validator` transaction, they can be in three states:

- `bonded`: Validator is in the active set and participates in consensus. Validator is earning rewards and can be slashed for misbehaviour.
- `unbonding`:
  - Validator misbehaved and is in jail, i.e. outisde of the validator set. If the jailing is due to being offline for too long, the validator can send an `unjail` transaction in order to re-enter the validator set.
  - Validator is out of the top 100 bonded and become a Candidate, the validator can delegate more IRIS to himself to be in top 100, then he will get bonded automatically.
- `unbonded`: Validator is not in the active set, and therefore not signing blocks. Validator cannot be slashed, and does not earn any reward. It is still possible to delegate to this validator.

### How to backup the validator

It is really IMPORTANT to backup your validator private key carefully, it's the only way to restore your validator. Note the validator private key is a Tendermint Key.

If you are using the software sign (which is the default signing method of tendermint), your Tendermint Key is located in `<bhp-home>/config/priv_validator.json`. The easiest way is to backup the whole config folder.

### How to migrate the validator

There are many ways to migrate your validator, the most recommended way is:

1.[Run a Full Node](../getting-start/join-bhp-testnet.md) on the new server

2.ter you have caught-up, stop the Validator Node and the Full Node

3.Replace the Full Node config folder with the Validator's

4.Start the new Validator Node

### How to upload my validator's logo to the [Explorers](../getting-start/explorers.md)

1. Sign-up [Keybase](https://keybase.io/) with your validator's name

2. Upload your logo as your Keybase account's avatar

3. `Add a PGP key` to your Keybase account (I believe you will see this option after sign-up), and you will get a 16-digit string

4. [Edit your validator](../cli-client/staking.md#bhpcli-tx-staking-edit-validator) and specify `--identity=<the_16_digit_string>`