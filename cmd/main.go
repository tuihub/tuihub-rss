package main

import (
	"context"
	"os"

	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
	"github.com/tuihub/tuihub-go"
	"github.com/tuihub/tuihub-go/logger"
	"github.com/tuihub/tuihub-rss/internal"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	config := tuihub.PorterConfig{
		Name:       "tuihub-rss",
		Version:    version,
		GlobalName: "github.com/tuihub/tuihub-rss",
		FeatureSummary: &porter.PorterFeatureSummary{
			SupportedAccounts:   nil,
			SupportedAppSources: nil,
			SupportedFeedSources: []string{
				tuihub.WellKnownToString(librarian.WellKnownFeedSource_WELL_KNOWN_FEED_SOURCE_RSS),
			},
			SupportedNotifyDestinations: nil,
		},
	}
	server, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if err = server.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
