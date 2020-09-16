---
order: 2
---

# keys

Keys allows you to manage your local tendermint keystore (wallets) for BHP.

##  Available Commands

| Name                          | **Description**                                          |
| ---------------------------------- | ------------------------------------------------------------ |
| [mnemonic](#bhpcli-keys-mnemonic) | Create a bip39 mnemonic, sometimes called a seed phrase, by reading from the system entropy. To pass your own entropy, use --unsafe-entropy |
| [add](#bhpcli-keys-add)           | Create a new key and import the existing key from the mnemonic |
| [list](#bhpcli-keys-list)         | Return a list of all public keys stored by this key manager along with their associated name and address. |
| [show](#bhpcli-keys-show)         | Displays the key information based on the given name |
| [export](#bhpcli-keys-export)     | Export a private key from the local keybase in ASCII-armored encrypted format. |
| [delete](#bhpcli-keys-delete)     | Delete a key from the store.                   |
| [import](#bhpcli-keys-import) | Import a ASCII armored private key into the local keybase. |
| [update](#bhpcli-keys-update)     | Change the password used to protect private key |

### bhpcli keys mnemonic

Create a bip39 mnemonic, sometimes called a seed phrase, by reading from the system entropy. To pass your own entropy, use --unsafe-entropy

```bash
bhpcli keys mnemonic <flags>
```

**Flags**：


| **Name, shorthand** | **Default** | **Description**                                              | **Required** |
| ------------------- | ----------- | ------------------------------------------------------------ | ------------ |
| --unsafe-entropy    |             | Prompt the user to supply their own entropy, instead of relying on the system |              |

**Create mnemonic**

```bash	
bhpcli keys mnemonic
```

You will get a mnemonic consisting of 24 words, such as:

```bash
sunset print vault mail bleak mule pigeon major denial carpet machine student alpha zone lazy figure universe satoshi luxury hand bunker silver method horror
```

## bhpcli keys add

Create a new key (wallet), or recover from mnemonic/keystore.

```bash
bhpcli keys add <key-name> <flags>
```
**Flags：**

| **Name, shorthand**  | **Default** | **Description**                                              | **Required** |
| -------------------- | ----------- | ------------------------------------------------------------ | ------------ |
| --account            |             | Account number for HD derivation                             |              |
| --dry-run            |             | Perform action, but don't add key to local keystore          |              |
| --index              |             | Address index number for HD derivation                       |              |
| -i, --interactive    |             | Interactively prompt user for BIP39 passphrase and mnemonic  |              |
| --ledger             |             | Store a local reference to a private key on a Ledger device  |              |
| --no-backup          |             | Don't print out seed phrase (if others are watching the terminal) |              |
| --recover            |             | Provide seed phrase to recover existing key instead of creating |              |
| --multisig           |             | Construct and store a multisig public key (implies --pubkey) |              |
| --multisig-threshold |             | K out of N required signatures. For use in conjunction with --multisig |              |

### Create a new key

**Example：**

```bash
bhpcli keys add mykey1
```

**Response：**

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
# Please note that the following mnemonics are very important and must be backed up
affair index need gaze duty bind desert deputy pass task caution ride century law scout robust autumn across apart endorse ride harbor correct novel
```

After executing this command, enter and confirm the password to generate a new key.Password at least 8 characters.

:::warning
**important**

Write down mnemonics and keep them in a safe place!This is the only way to recover your account if you accidentally forget your password or lose your key.
:::

#### Generate keys from mnemonics

Manually generate mnemonic words with the **-i** parameter

```bash
bhpcli keys add mykey -i
```

The system will ask you to enter the secret of the key and your mnemonic

```bash
override the existing name mykey3 [y/N]: y
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
> Enter your bip39 mnemonic, or hit enter to generate one.
sunset print vault mail bleak mule pigeon major denial carpet machine student alpha zone lazy figure universe satoshi luxury hand bunker silver method horror
> Enter your bip39 passphrase. This is combined with the mnemonic to derive the seed. Most users should just hit enter to use the default, ""
```

###  Recover an existing key from seed phrase

If you forget your password or lose your key, or you wanna use your key in another place, you can recover your key by your seed phrase.

```bash
bhpcli keys add mykey --recover
```

You'll be asked to enter and repeat the new password for your key, and enter the seed phrase. Then you get your key back.

```bash
Enter a passphrase for your key:
Repeat the passphrase:
Enter your recovery seed phrase:
```

### Create a multisig key

The following example creates a multisig key with 3 sub-keys, and specify the minimum number of signatures as 2. The tx could be broadcast only when the number of signatures is greater than or equal to 2.

```shell
bhpcli keys add <multisig-keyname> --multisig-threshold=2 --multisig=<signer-keyname-1>,<signer-keyname-2>,<signer-keyname-3>
```

:::tip
`<signer-keyname> `can be the type of "local/offline/ledger", but not "multi" type.

If you don't have all the permission of sub-keys, you can ask for the pubkeys to create the offline keys first, then you will be able to create the multisig key.

Offline key can be created by "bhpcli keys add --pubkey".
:::



## bhpcli keys list

List all the keys stored by this key manager along with their associated name, type, address and pubkey.

```bash
bhpcli keys list [flags]
# Flags:
#     --indent Add indent to JSON response
```

**Example：**

```bash
bhpcli keys list
```

## bhpcli keys show

Displays the key information based on the given name

```bash	
bhpcli keys show [name [name...]] [flags]
```

**Flags：**

| **Name, shorthand**  | **Default** | **Description**                                 | **Required** |
| -------------------- | ----------- | ----------------------------------------------- | ------------ |
| -a, --address        |             | Output the address only (overrides --output)    |              |
| --bech               | acc         | The Bech32 prefix encoding for a key (acc       | val          |
| --indent             |             | Add indent to JSON response                     |              |
| --multisig-threshold | 1           | K out of N required signatures                  |              |
| -p, --pubkey         |             | Output the public key only (overrides --output) |              |

### Get details of a local key

```bash
bhpcli keys show mykey
```

Execution result：

```bash
- name: testing
  type: local
  address: bhp1rr0u7tgaz07y670tpfqma7p8nmscpk6wjtnvxr
  pubkey: bhppub1addwnpepqdag42nacaekg7234cvdpvc8n8rafazwrfd8ss3uu3kh0jes45m26dw0xu5
  mnemonic: ""
  threshold: 0
  pubkeys: []
```

### Get validator operator address

If an address has been bonded to be a validator operator (which the address you used to create a validator), then you can use `--bech val` to get the operator's address prefixed by `bhpvaloper` and the pubkey prefixed by `bhpvaloper`:

```bash
bhpcli keys show mykey --bech val
```

Execution result：

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

Export a private key from the local keybase in ASCII-armored encrypted format.

```bash
bhpcli keys export <name> [flags]
```

Example：

```bash
bhpcli keys export testing 
# Enter passphrase to decrypt your key:
# Enter passphrase to encrypt the exported key:
```

Execution result：(You need to keep the following content in the file)

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

Delete a local key by the given name.

```bash
bhpcli keys delete <name> [flags]
```

**Flags：**

| **Name, shorthand** | **Default** | **Description**                                              | **Required** |
| ------------------- | ----------- | ------------------------------------------------------------ | ------------ |
| -f, --force         | false       | Remove the key unconditionally without asking for the passphrase |              |
| -y, --yes           | false       | Skip confirmation prompt when deleting offline or ledger key references |              |

Example：

```bash
bhpcli keys delete mykey
# DANGER - enter password to permanently delete key:
# Key deleted forever (uh oh!)
```

### bhpcli keys import

Import a ASCII armored private key into the local keybase.

```bash
bhpcli keys import <name> <keyfile> [flags]
```

Example：

```bash
# Keyfile is exported with bhpcli Keys export command
bhpcli keys import mykey keyfile
```

### bhpcli keys update

Change the password used to protect private key

```bash
bhpcli keys update <name>
```

**Change the password for the local key**

```bash
bhpcli keys update mykey
```

The system will ask you to enter your current key and a new password and confirm.

```bash
# Enter the current passphrase:
# Enter the new passphrase:
# Repeat the new passphrase:
# Password successfully updated!
```

