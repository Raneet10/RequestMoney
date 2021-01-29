# requestmoney

**requestmoney** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).
This repo has 2 branches :

```master``` : This uses the ```mint``` module of the sdk to mint coins and ```supply``` to transfer it to accounts. To modify the default **stake** denom in ```mint``` , the generated genesis suppposedly has to be changed.

```raneet10/supply-requestmoney``` : This simply uses ```supply``` to both mint and transfer tokens from module to the requesting account.

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
