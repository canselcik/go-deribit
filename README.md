# go-deribit

## Under Development!

This is the **v2** API and I am still working on subscriptions. The deprecated, but still functioning v1 API has been tagged `v1`.

## Overview

Go library for using the [Deribit's](https://www.deribit.com/reg-3027.8327) **v2** Websocket API. 

Deribit is a modern, fast BitCoin derivatives exchange. If you are using BitMex then you are doing it wrong! Deribit does not freeze up during higher than average load. Also, it is peer-to-peer, not run by market makers on lucrative contracts who want to liquidate you.

This library is a port of the [official wrapper libraries](https://github.com/deribit) to Go.

If you wish to try it out, be kind and use my affiliate link [https://www.deribit.com/reg-3027.8327](https://www.deribit.com/reg-3027.8327)

## Example Usage

Look at `cmd/example/main.go`

## Development

The `models` directory is where all the requests and responses are stored. The contents is automatically generated from the example json snippets from `examples` which were kindly provided to me by Deribit from their Slate documentation sources.

If you need to rebuild these use `make generate-models`.

The RPC methods are also auto-generated. Use `make generate-methods` to rebuild these. They are in `rpc_methods.go` and you will need to run `goimports` afterward (or just let your IDE do it for you :)).
