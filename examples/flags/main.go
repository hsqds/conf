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
	)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	cp := conf.NewDefaultConfigProvider()
	defer cp.Close(ctx)

	const priority = 100

	cp.AddSource(sources.NewFlagsSource(priority, "--"))

	loadErrors := cp.Load(context.Background(), serviceName1, serviceName2)
	if len(loadErrors) != 0 {
		log.Printf("%#v", loadErrors)
	}

	cfg1, err := cp.ServiceConfig(serviceName1)
	if err != nil {
		panic(err)
	}

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
