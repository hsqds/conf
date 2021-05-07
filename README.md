# conf
!WARNING: WIP
not recomended for production use

## main idea
application may run many services, like grpc, http or by messaging over queues.
`conf` let the user to retrieve services configs from different sources.

`instantiaite provider -> add sources -> load config data -> use config data`

## sources
each source has priority
each source may provide configs for many services

### included sources
 * environment variables
 * command line flags
 * json files

#### json files

let's try to load configuration from json file

example config `config.json`
```json
{
    "http": {
        "host": "0.0.0.0",
        "port": "8080"
    },
}
```

fallback config `defaults.json`
```json
{
    "http": {
        "host": "",
        "port": "8000"
    },
    "grpc": {
        "host": "auth.svc",
        "port": "3000"
    }
}
```

first we need to create config provider

provider uses zerolog logger
```go
	cp := conf.NewDefaultConfigProvider(&logger)
	defer cp.Close(ctx)
```

then add some sources
```go
	cp.AddSource(
        // `config.json` has higher priority than `defaults.json`
		sources.NewJSONFileSource(100, "config.json", &logger),
		sources.NewJSONFileSource(90, "defaults.json", &logger),
	)
```

then load data from sources and get needed service config
```go
	err := cp.Load(context.Background(), "http", "grpc")
    processErr(err)
    // `config.json` data will override `dafaults.json` respecting the priority
	httpConfig, err := cp.ServiceConfig("http")
    processErr(err)
    // grpc config will be loaded from `defaults.json`
    grpcConfig, err := cp.ServiceConfig("grpc")
    processErr(err)
```

now we may get config parameter using config `Get` method
```go
    value, ok := httpConfig.Get("host", "default.svc") // "0.0.0.0", true - from `configs.json`
    value, ok := grpcConfig.Get("host", "default.svc") // "auth.svc", true - from `defaults.json`
    value, ok := grpcConfig.Get("inexistingKey", "default.value") // "default.value", false - defaultValue
```

or format config parameters calling `Fmt` method
```go
    formatted, err := httpConfig.Fmt("{{.host}}:{{.port}}") // "0.0.0.0:8080"
    processErr(err)
```

## ROADMAD:
* now - remove zerolog from dependencies
* v0.1 - jsonfile source, env source, flags source
* v0.2 - ability to subscribe for source updates