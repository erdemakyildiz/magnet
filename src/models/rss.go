package models

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		LastBuildDate string `xml:"lastBuildDate"`
		PubDate       string `xml:"pubDate"`
		Item []struct {
			Title string `xml:"title"`
			AtomLink  struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
				Rel  string `xml:"rel,attr"`
			} `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Content struct {
				Text   string `xml:",chardata"`
				Height string `xml:"height,attr"`
				Medium string `xml:"medium,attr"`
				URL    string `xml:"url,attr"`
				Width  string `xml:"width,attr"`
			} `xml:"content"`
		} `xml:"item"`
	} `xml:"channel"`
}

