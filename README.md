# centrifuge-protobufs
[![Build Status](https://travis-ci.com/centrifuge/centrifuge-protobufs.svg?token=Sbf68xBZUZLMB3kGTKcX&branch=master)](https://travis-ci.com/centrifuge/centrifuge-protobufs)

Protobuf files &amp; go bindings for Centrifuge Documents. These define the document schema exchanged by nodes on the P2P network.

When editing these specs please note the following:
- Never reuse the protobuf field numbers. As per the protobuf standard a field number should only be used once and marked as [reserved](https://developers.google.com/protocol-buffers/docs/proto#reserved) when they are not in use anymore.
- Never have your field start with XXX_, those are used for internal fields in Go and will be ignored by precise proofs.
- While technically this is not a requirement, it's not recommended to use the same field names across the CoreDocument and other messages.
- Fields that are essential and shared among all documents belong into CoreDocument, fields specific to the document type belong into invoice, po etc.


**Getting help:** Head over to our developer documentation at [developer.centrifuge.io](http://developer.centrifuge.io) to learn how to setup a node and interact with it. If you have any questions, feel free to join our [slack channel](https://join.slack.com/t/centrifuge-io/shared_invite/enQtNDYwMzQ5ODA3ODc0LTU4ZjU0NDNkOTNhMmUwNjI2NmQ2MjRiNzA4MGIwYWViNTkxYzljODU2OTk4NzM4MjhlOTNjMDAwNWZkNzY2YWY) 

**DISCLAIMER:** The code released here presents a very early alpha version that should not be used in production and has not been audited. Use this at your own risk.

# How to commit
Make sure you run the `make` before committing to make sure the generated files are up to date:

```bash,
# To install dependencies and generate files
make all
# To just generate files
make gen_proto
```

# Installation

```
$ make help
usage: make [target] ...

targets:
help                 Show this help message.
install              Install dependencies required to generate bindings & documentation
vendorinstall        Installs all protobuf dependencies with go-vendorinstall
lint                 runs prototool lint
proto-gen-go         generates the go bindings
proto-all            runs prototool all
```
## Making sure all dependencies are installed
### Install dependencies
```
make install
```

# Quick intro to prototool
We are using [prototool](https://github.com/uber/prototool) to lint or protobuf
files and build our go stubs.

## Helpful commands

Below a few helpful commands:

```
# To lint your files:
prototool lint

# To build the go stubs:
prototool gen-go

```

