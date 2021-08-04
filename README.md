# conf
WARNING! just for fun pet-project, not recommend for production use

`conf` simplify loading of application settings from different source. 
It may be specialized service like consul or etcd or simple json file, cli flags and environment variables.

![linter](https://github.com/hsqds/conf/actions/workflows/golangci.yml/badge.svg)
![tests](https://github.com/hsqds/conf/actions/workflows/testing.yml/badge.svg)

## Table of contents
 * [General info](#general-info)
 * [Sources](#sources)
 * [Dependencies](#dependencies)
 * [Requirements](#requirements)
 * [Setup](#setup)
 * [Examples](#examples)
 * [Roadmap](#roadmap)

## General info
While you develop, debug or testing your application, you may need to redeclare some of the settings, like database connection parameters, or some domain logic options. `conf` let you organize application configuration process flexible. 

### How does it work
Application config represented as a set of services configs provided or used by application. First you need
to instantiate `conf.Provider` then add some sources like json-file, environment variables of cli flags. 
When sources added you should `Load` configs you need. `conf.Provider` will try to load services configs
from each source you've added concurrently. Then you may take service config calling `ServiceConfig` method.
You will get service config loaded from most prioritized source.

See [usage example](./examples/alltogether/main.go).


#### Sources
 * each source has priority
 * each source may provide configs for many services

`conf` includes components to load data from 3 source types:
 * environment variables
 * command line flags
 * json files

### How to use 
`instantiaite provider -> add sources -> load config data -> use config data`

You may see detailed [example](./examples/alltogether/main.go)

## Dependencies
`conf` has only one third-party dependency
 * github.com/google/uuid

## Requirements
go version >= 1.16

## Setup
`go get github.com/hsqds/conf`

## ROADMAP
* v0.2:
  * optional merge service configs from various sources
  * dotenv source
* v0.3 - ability to subscribe for source updates
