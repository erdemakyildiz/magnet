package entity

/*
	MongoDB deki modeller burada toplanıyor
 */
//Kanalların tutulduğu model
type Channel struct {
	ChannelName string `bson:"channelName" json:"channelName"`
	RssUrl      string `bson:"rssUrl" json:"rssUrl"`
	Active 		bool   `bson:"active" json:"active"`
}