package main


import (
//    "drule.org/rssdownload"
        "fmt"
        rss "github.com/jteeuwen/go-pkg-rss"
        "os"
//        "net/http"
//        "time"
)

const URL string = "http://feeds.twit.tv/sn_video_large"

func main() {

    err := os.Mkdir("downloads", 755)
    if err != nil && os.IsExist(err) {
        fmt.Println(err)
    }


    timeout := 5
    fmt.Println("checking URL ", URL)

    feed := rss.New(timeout, true, chanHandler, itemHandler)

    if err := feed.Fetch(URL, nil); err != nil {
        fmt.Fprintf(os.Stderr, "[e] %s: %s", URL, err)
        return
    }

}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
    fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
        fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}


