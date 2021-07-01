# BSN Service Relayer

BSN Service Relayer is intended for interchain service invocation between application chains and IRITA-HUB in BSN ecosystem.

The project is under development, with a mock implementation used to demonstrate the relayer working model and as a development reference.

## Install

Install the relayer using the following command:

```bash
make install
```

## Start

### Irita-Hub

Make sure that Irita-Hub installed and the iService service(s) deployed.

### AppChain

Make sure that the application chain is running

### Key

#### Add Irita-Hub key

```bash
relayer hub keys add [name] [passphrase]
```

### Configure

Configure the relayer according to the Irita-Hub and AppChain, default to `./config/config.yaml`

### Relayer

Start the relayer process:

```bash
relayer start
```

### Add or Update the service binding to the iService market on the application chain

```bash
relayer appchain add-svc-binding [args]
```

```bash
relayer appchain update-svc-binding [args]
```

## Development

See [Spec](./spec/design.md) for the application chain interfaces.

Add an application chain in `appchains` directory:

- Implementing the required interfaces
- Add to `appchains/factory.go` for routing
- Optionally, Add test to `tests` directory
