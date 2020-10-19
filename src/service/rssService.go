package service

import (
	"bundle_magnet/src/client"
	"bundle_magnet/src/models/entity"
	"fmt"
	"github.com/mmcdole/gofeed/rss"
	"log"
	"strings"
)

func GetRssFeedFromUrl(item entity.Channel)  {
	fmt.Println("Data Okunuyor : " + item.ChannelName)

	data, err := client.HttGet(item.RssUrl)

	if err != nil {
		log.Fatal(err)
	}

	fp := rss.Parser{}
	rssFeed, _ := fp.Parse(strings.NewReader(data))
	fmt.Println(rssFeed.Link)

	fmt.Println("Yollandı : " + item.ChannelName)
}
