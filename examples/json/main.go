package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
)

func main() {
	const (
		serviceName1 = "serviceName1"
		serviceName2 = "serviceName2"
		confFilename = "config.json"
	)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	cp := conf.NewDefaultConfigProvider()
	defer cp.Close(ctx)

	f, err := os.Open(confFilename)
	if err != nil {
		panic(fmt.Errorf("could not open config file (%q): %w", confFilename, err))
	}
	defer f.Close()

	cp.AddSource(sources.NewJSONSource(100, f))

	loadErrors := cp.Load(context.Background(), serviceName1, serviceName2)
	if len(loadErrors) != 0 {
		log.Printf("%#v", loadErrors)
	}

	cfg1, err := cp.ServiceConfig(serviceName1)
	if err != nil {
		panic(err)
	}

	log.Printf("got %q cfg: %v\n", serviceName1, cfg1)

	cfg2, err := cp.ServiceConfig(serviceName2)
	if err != nil {
		panic(err)
	}

	log.Printf("%s config %#v", serviceName1, cfg1)
	log.Printf("%s config %#v", serviceName2, cfg2)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '-', tabwriter.TabIndent)
	pattern := "%s\t%s\t%s\n"
	fmt.Fprintf(w, pattern, "KEY", "VALUE", "DEFAULT")
	fmt.Fprintf(w, pattern, "key1", cfg1.Get("key11", "default11"), "default11")
	fmt.Fprintf(w, pattern, "key2", cfg1.Get("key12", "default12"), "default12")
	fmt.Fprintf(w, pattern, "key3", cfg1.Get("key13", "default13"), "default13")

	err = w.Flush()
	if err != nil {
		log.Printf("table flush error: %s", err)
	}

	formatted, err := cfg1.Fmt("{{.host}}:{{.port}}")
	if err != nil {
		panic(err)
	}

	log.Printf("fmt: %q", formatted)
}
