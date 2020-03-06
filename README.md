# intri-type

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- [protocol buffers](http://google.github.io/proto-lens/installing-protoc.html)

```shell
brew install protobuf
```

- [gvm](https://github.com/moovweb/gvm#installing) - for multiple go version (optional)
  - If under **none** Go environment, please follow the instruction [here](https://github.com/moovweb/gvm#a-note-on-compiling-go-15).
  - When the tutor above fails you, please [install GO directly](https://golang.org/) and then you can use gvm installation for Go 1.5+

### Installing

A step by step series of examples that tell you have to get a development env running

- Set up Go modeuls and git submodules

```shell
make env
```

> This may take a while due to the git submodule is big.

### Deployment

Compile protocol buffers into `/api`

```shell
make proto
```

## Built With

- [protocol buffers](https://developers.google.com/protocol-buffers) - A language-neutral, platform-neutral extensible mechanism for serializing structured data.
- [grpc](https://grpc.io/docs/quickstart/go/) - The web framework used
- [grpc-ecosystem](https://github.com/grpc-ecosystem)
  - [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - gRPC to JSON proxy generator following the gRPC HTTP spec
- [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate) - A protoc plugin to generate polyglot message validators.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository.

## Authors

- **Martin Chang** - *Initial work* - [Intrising](https://github.com/Intrising)
