package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
)

func main() {
	const (
		serviceName1 = "serviceName1"
		serviceName2 = "serviceName2"
	)

	cp := conf.NewDefaultConfigProvider()
	cp.AddSource(
		sources.NewMapSource(100, map[string]conf.Config{
			serviceName1: conf.NewMapConfig(map[string]string{
				"key1": "value1",
			}),
			serviceName2: conf.NewMapConfig(map[string]string{
				"key2": "value2",
			}),
		}),
	)

	cp.AddSource(
		sources.NewMapSource(101, map[string]conf.Config{
			serviceName1: conf.NewMapConfig(map[string]string{
				"key1": "boom!",
			}),
		}),
	)

	loadErrors := cp.Load(context.Background(), serviceName1, serviceName2)
	if len(loadErrors) != 0 {
		// no need to panic
		// not every source may provide config data for every service
		log.Printf("%#v", loadErrors)
	}

	cfg1, err := cp.ServiceConfig(serviceName1)
	if err != nil {
		panic(err)
	}

	log.Printf("%q config %#v", serviceName1, cfg1)

	cfg2, err := cp.ServiceConfig(serviceName2)
	if err != nil {
		panic(err)
	}

	log.Printf("%q config %#v", serviceName2, cfg2)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '-', tabwriter.TabIndent)
	pattern := "%s\t%s\t%s\n"
	fmt.Fprintf(w, pattern, "KEY", "VALUE", "DEFAULT")
	fmt.Fprintf(w, pattern, "key1", cfg1.Get("key1", "default11"), "default11")
	fmt.Fprintf(w, pattern, "key2", cfg1.Get("key2", "default12"), "default12")
	fmt.Fprintf(w, pattern, "key3", cfg1.Get("key3", "default13"), "default13")

	err = w.Flush()
	if err != nil {
		log.Printf("table flush error: %s", err)
	}
}
