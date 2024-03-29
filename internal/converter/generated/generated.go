// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	gofeed "github.com/mmcdole/gofeed"
	v1 "github.com/tuihub/protos/pkg/librarian/v1"
	converter "github.com/tuihub/tuihub-rss/internal/converter"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ToPBFeed(source *gofeed.Feed) *v1.Feed {
	var pV1Feed *v1.Feed
	if source != nil {
		var v1Feed v1.Feed
		v1Feed.Title = (*source).Title
		v1Feed.Link = (*source).Link
		v1Feed.Description = (*source).Description
		var pV1FeedItemList []*v1.FeedItem
		if (*source).Items != nil {
			pV1FeedItemList = make([]*v1.FeedItem, len((*source).Items))
			for i := 0; i < len((*source).Items); i++ {
				pV1FeedItemList[i] = c.ToPBFeedItem((*source).Items[i])
			}
		}
		v1Feed.Items = pV1FeedItemList
		v1Feed.Language = (*source).Language
		v1Feed.Image = c.ToPBFeedImage((*source).Image)
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for j := 0; j < len((*source).Authors); j++ {
				pV1FeedPersonList[j] = c.pGofeedPersonToPV1FeedPerson((*source).Authors[j])
			}
		}
		v1Feed.Authors = pV1FeedPersonList
		pV1Feed = &v1Feed
	}
	return pV1Feed
}
func (c *ConverterImpl) ToPBFeedEnclosure(source *gofeed.Enclosure) *v1.FeedEnclosure {
	var pV1FeedEnclosure *v1.FeedEnclosure
	if source != nil {
		var v1FeedEnclosure v1.FeedEnclosure
		v1FeedEnclosure.Url = (*source).URL
		v1FeedEnclosure.Length = (*source).Length
		v1FeedEnclosure.Type = (*source).Type
		pV1FeedEnclosure = &v1FeedEnclosure
	}
	return pV1FeedEnclosure
}
func (c *ConverterImpl) ToPBFeedImage(source *gofeed.Image) *v1.FeedImage {
	var pV1FeedImage *v1.FeedImage
	if source != nil {
		var v1FeedImage v1.FeedImage
		v1FeedImage.Url = (*source).URL
		v1FeedImage.Title = (*source).Title
		pV1FeedImage = &v1FeedImage
	}
	return pV1FeedImage
}
func (c *ConverterImpl) ToPBFeedItem(source *gofeed.Item) *v1.FeedItem {
	var pV1FeedItem *v1.FeedItem
	if source != nil {
		var v1FeedItem v1.FeedItem
		v1FeedItem.Title = (*source).Title
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pV1FeedPersonList[i] = c.pGofeedPersonToPV1FeedPerson((*source).Authors[i])
			}
		}
		v1FeedItem.Authors = pV1FeedPersonList
		v1FeedItem.Description = (*source).Description
		v1FeedItem.Content = (*source).Content
		v1FeedItem.Guid = (*source).GUID
		v1FeedItem.Link = (*source).Link
		v1FeedItem.Image = c.ToPBFeedImage((*source).Image)
		v1FeedItem.Published = (*source).Published
		v1FeedItem.PublishedParsed = converter.ToPBTime((*source).PublishedParsed)
		v1FeedItem.Updated = (*source).Updated
		v1FeedItem.UpdatedParsed = converter.ToPBTime((*source).UpdatedParsed)
		var pV1FeedEnclosureList []*v1.FeedEnclosure
		if (*source).Enclosures != nil {
			pV1FeedEnclosureList = make([]*v1.FeedEnclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pV1FeedEnclosureList[j] = c.ToPBFeedEnclosure((*source).Enclosures[j])
			}
		}
		v1FeedItem.Enclosures = pV1FeedEnclosureList
		pV1FeedItem = &v1FeedItem
	}
	return pV1FeedItem
}
func (c *ConverterImpl) pGofeedPersonToPV1FeedPerson(source *gofeed.Person) *v1.FeedPerson {
	var pV1FeedPerson *v1.FeedPerson
	if source != nil {
		var v1FeedPerson v1.FeedPerson
		v1FeedPerson.Name = (*source).Name
		v1FeedPerson.Email = (*source).Email
		pV1FeedPerson = &v1FeedPerson
	}
	return pV1FeedPerson
}
