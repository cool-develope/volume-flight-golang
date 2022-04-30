# volume-flight-golang

This project is an implementatioin of a simple microservice API that can help us understand and track how a particular person's flight path may be queried.

## Architecture

- server: interface that runs the simple gRPC server
- pathctrl: controller that makes the track of flights to one
- cmd: main command line

## API

It's recommended to use `make` for building, testing the code, ... Run `make help` to get a list of the available commands.

The API is implemented using the gRPC protocol, `proto/flight/v1/path.proto`