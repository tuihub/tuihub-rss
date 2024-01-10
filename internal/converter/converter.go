package converter

import (
	"time"

	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/mmcdole/gofeed"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen -g ignoreUnexported .

// goverter:converter
type Converter interface {
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	ToPBFeed(t *gofeed.Feed) *librarian.Feed
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	// goverter:map UpdatedParsed | ToPBTime
	// goverter:map PublishedParsed | ToPBTime
	// goverter:ignore PublishPlatform
	// goverter:ignore ReadCount
	ToPBFeedItem(t *gofeed.Item) *librarian.FeedItem

	// goverter:matchIgnoreCase
	ToPBFeedImage(t *gofeed.Image) *librarian.FeedImage

	// goverter:matchIgnoreCase
	ToPBFeedEnclosure(t *gofeed.Enclosure) *librarian.FeedEnclosure
}

func ToPBTime(t *time.Time) *timestamppb.Timestamp {
	return timestamppb.New(*t)
}
