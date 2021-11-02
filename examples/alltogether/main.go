package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
)

func main() {
	// For example let's try to configure application that
	// will read data from mongodb and provide http and grpc APIs.
	// What kind of data do we need to set up application?
	// mongodb: username, password, hostname, port, db name
	// http endpoint: host and port to listen connections
	// grpc endpoint: host and port to listen
	//
	// for clarity we will store each service config in different source
	// mongodb and grpc services configs will be stored in json file
	// to configure http endpoint we will use defaults
	log.Println("Starting")

	// first we need to create provider instance
	// also you may use conf.NewConfigProvider() method
	// in case if you want to use alternative realizations
	// of service storage, config storage or loader
	p := conf.NewDefaultConfigProvider()
	// don't forget to close connections to shutdown gracefully
	defer p.Close(context.TODO())

	// let's add sources
	// say we want to get our app settings from json file
	// but for debug we want to use flags also we will use native
	// map to store defaults
	// we got flags > json > map
	// ---------------------
	// | SOURCE | PRIORITY |
	// ---------------------
	// | flags  | 200      |
	// | json   | 100      |
	// | map    | 0        |
	// ---------------------

	const (
		flagsPriority = 200
		jsonPriority  = 100
		mapPriority   = 0
	)

	// adding flags source is simple
	// just pass priority and prefix
	p.AddSource(sources.NewFlagsSource(flagsPriority, "--"))

	// now lets add json source
	// constructor function accept priority and reader
	// so we need to open config file first
	jsonFile, err := os.Open("./examples/alltogether/config.json")
	if err != nil {
		panic(fmt.Errorf("could not open json config file: %w", err))
	}

	p.AddSource(sources.NewJSONSource(jsonPriority, jsonFile))

	// finally add defaults
	p.AddSource(sources.NewMapSource(mapPriority, map[string]conf.Config{
		"http": conf.NewMapConfig(map[string]string{
			"hostname": "0.0.0.0",
			"port":     "80",
		}),
	}))

	// ok.
	// Loading configuration data.
	// Passing context (WithTimeout recommended, but to keep it simple,
	// will use TODO) and list of services, which settings we need to load.
	// Config provider will try to load each service configuration from each
	// source we've added earlier. But we have no defaults for grpc and mongo,
	// as we have no http settings in config.json.
	errs := p.Load(context.TODO(), "grpc", "http", "mongodb")

	// So we get a list of errors. Actually a slice of `conf.LoadError`.
	// Is it critical? No!
	// We may just log it and continue
	log.Printf("config loading errors: %#v", errs)

	// Requesting http endpoint config
	httpConfig, err := p.ServiceConfig("http")
	if err != nil {
		// In this case err may be critical - we got no config from any source!
		panic(err)
	}

	// Now when we get config instance, let's see how we can use it's data.
	// Consider http config.
	// `Get` method takes key and default value
	hostname := httpConfig.Get("hostname", "localhost")
	_ = hostname
	// hostname == "0.0.0.0"
	// But! Do you remember about flags source we've added earlier?
	// If you will run you application with flags
	// `app --http-hostname=any-hostname.here --http-port=8080`
	// config will contain `hostname == "any-hostname.here"`
	// and `port == "8080"`

	// Also we may use `Fmt` method and format parameters using standard go
	// templating syntax.
	httpListen, err := httpConfig.Fmt("{{.hostname}}:{{.port}}")
	// err may be critical here
	_ = err
	// If default settings were not overriden by json or flags, we'll get
	// `httpListen == "hostname:80"``
	_ = httpListen
}
