package internal

import (
	"context"

	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	"github.com/tuihub/tuihub-rss/internal/converter/generated"

	"github.com/mmcdole/gofeed"
	"github.com/muzhou233/go-favicon"
)

type Handler struct {
	porter.UnimplementedLibrarianPorterServiceServer
	rss     RSS
	favicon *favicon.Finder
}

func NewHandler() *Handler {
	return &Handler{
		porter.UnimplementedLibrarianPorterServiceServer{},
		NewRSS(),
		favicon.New(favicon.IgnoreManifest),
	}
}

func (h Handler) PullFeed(ctx context.Context, req *porter.PullFeedRequest) (
	*porter.PullFeedResponse, error) {
	data, err := h.rss.Get(req.GetChannelId())
	if err != nil {
		return nil, err
	}
	feed, err := h.rss.Parse(data)
	if err != nil {
		return nil, err
	}
	if len(feed.Link) > 0 {
		if icons, err1 := h.favicon.Find(feed.Link); err1 == nil && len(icons) > 0 {
			for _, icon := range icons {
				if icon.Height > 0 && icon.Width > 0 {
					feed.Image = &gofeed.Image{
						URL:   icons[0].URL,
						Title: "",
					}
					break
				}
			}
		}
	}
	converter := &generated.ConverterImpl{}
	res := converter.ToPBFeed(feed)
	return &porter.PullFeedResponse{Data: res}, nil
}
