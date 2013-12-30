package rssdownload

import (
    "testing"
    "os"
)

func TestNew(t *testing.T) {
    c := *NewCache("somefile.json")
    c.seenfiles["http://www.google.com/rss"] = []string{ "podcast3.mp3", "podcast4.mp3"}
    c.seenfiles["http://www.google.com/rss2"] = []string{ "podcast6.mp3", "podcast7.mp3"}
    if &c == nil {
        t.Errorf("pointer should not be nil, %s", c)
    } else {
//        t.Error(c.JsonOut())
    }
}

func TestSave(t *testing.T) {
    c := *NewCache("somefile.json")
    c.seenfiles["http://www.google.com/rss"] = []string{ "podcast3.mp3", "podcast4.mp3"}
    c.seenfiles["http://www.google.com/rss2"] = []string{ "podcast6.mp3", "podcast7.mp3"}
    c.Save()
    _, err := os.Open("somefile.json")
    if os.IsNotExist(err) {
        t.Error("somefile.json needs to exist after save")
    }
}

