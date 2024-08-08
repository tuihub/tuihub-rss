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
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/tuihub-rss",
			BuildVersion:      version,
			BuildDate:         "",
			Name:              "tuihub-rss",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/tuihub-rss",
		Region:     "",
		FeatureSummary: &porter.PorterFeatureSummary{ //nolint:exhaustruct // no need
			FeedSources: []*librarian.FeatureFlag{
				{
					Id:               tuihub.WellKnownToString(librarian.WellKnownFeedSource_WELL_KNOWN_FEED_SOURCE_RSS),
					Name:             "RSS",
					Description:      "",
					ConfigJsonSchema: tuihub.MustReflectJSONSchema(new(internal.PullRSSConfig)),
				},
			},
		},
		ContextJsonSchema: nil,
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
