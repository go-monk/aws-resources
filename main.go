package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/config"
)

var usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "List AWS resources returned by the AWS Resource Groups Tagging API.\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Output can include previously tagged or deleted resources.\n\n")
	flag.PrintDefaults()
}

var tagsFlag Tags

func init() {
	flag.Var(&tagsFlag, "tag", "filter by `key=value` tag (repeatable)")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("aws-resources: ")

	flag.Usage = usage
	profile := flag.String("profile", "", "AWS profile to use")
	region := flag.String("region", "", "AWS region to query")
	withTags := flag.Bool("tags", false, "include tags in the output")
	flag.Parse()

	// Create a context that gets canceled on Ctrl-C or kill PID.
	ctx, cancel := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Load AWS config with optional profile and region override.
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile(*profile), config.WithRegion(*region))
	if err != nil {
		log.Fatal(err)
	}

	resources, err := GetResources(ctx, cfg, tagsFlag)
	if err != nil {
		log.Fatal(err)
	}

	if *withTags {
		resources.PrintWithTags()
	} else {
		resources.Print()
	}
}
