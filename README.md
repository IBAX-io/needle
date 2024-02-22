# The Needle Language

[![Go Reference](https://pkg.go.dev/badge/flag.svg)](https://pkg.go.dev/github.com/IBAX-io/needle)

Needle is a statically typed programming language for developing smart contracts to power the IBAX blockchain. It is designed to be a platform-agnostic language to enable common libraries with execution smart contracts on the blockchain.
This repository contains the Needle compiler, virtual machines, grammar and more. then developers want to quickly write smart contracts for IBAX blockchain, it should be written in Needle.

## Usage Example

Here is a simple Needle smart contract example:

```
contract Hello {
    func hello() string {
        return "Hello, Needle!"
    }
}
```

## API Documentation

You can find the complete API documentation [here](https://pkg.go.dev/github.com/IBAX-io/needle).

## ANTLR Grammar

We are using grammar files. Link to the grammar files can be found [here](https://github.com/IBAX-io/needle/tree/main/grammar/antlr). 

## Contribution

We welcome everyone's contribution to Needle. If you want to contribute to Needle, you can submit a [pull request](https://github.com/IBAX-io/needle/pulls) or propose your suggestions in the [issue](https://github.com/IBAX-io/needle/issues) area.

## License

Needle is licensed under the GNU General Public License v3.0. For more information, please refer to the [LICENSE](https://github.com/IBAX-io/needle/blob/main/LICENSE) file.