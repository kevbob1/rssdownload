package rssdownload

import "testing"
//import "fmt"

func TestNew(t *testing.T) {
    c := *NewCache("somefile.json")
    c.seenfiles["http://www.google.com/rss"] = []string{ "podcast3.mp3", "podcast4.mp3"}
    c.seenfiles["http://www.google.com/rss2"] = []string{ "podcast6.mp3", "podcast7.mp3"}
    if &c == nil {
        t.Errorf("pointer should not be nil, %s", c)
    } else {
        t.Error(c.JsonOut())
    }
}

