package main


import (
    rssd "drule.org/rssdownload"
    "fmt"
    rss "github.com/jteeuwen/go-pkg-rss"
    "os"
    "strings"
    "path"
    "bufio"
//    "time"
)

var cache rssd.Cache
var urls []string

func main() {

    urls = readUrls("rssdownload.conf")    
    
    if urls == nil {
        return
    }

    cache = *rssd.NewCache(path.Join("downloads", "rssdownloadcache.json"))
    defer cache.Save()

    timeout := 5

    for _, url := range urls {

        fmt.Println("checking URL ", url)

        feed := rss.New(timeout, true, chanHandler, itemHandler)
        if err := feed.Fetch(url, nil); err != nil {
            fmt.Fprintf(os.Stderr, "[e] %s: %s", url, err)
            continue
        }
    }
}

func readUrls(filename string) ([]string) {

    f, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
        return nil
    }
    
    out := make([]string, 0, 20)

    inScanner := bufio.NewScanner(f)
    for inScanner.Scan() {
        url := inScanner.Text()
        if url == "" {
            continue 
        }
        out = append(out, url)
    }

    return out   
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
//    fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
// no need to do anything we care about items
}


func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
    // only grab the most recent 3
    for itemIndex, item := range newitems {
        if itemIndex > 2 {
            break
        }
        // get enclosure url
        if item.Enclosures != nil && len(item.Enclosures) != 0 {
            file_url := item.Enclosures[0].Url
            start_idx := strings.LastIndex(file_url, "/")
            filename := file_url[start_idx+1:len(file_url)]
            qm_idx := strings.Index(filename, "?")
            if qm_idx != -1 {
                // strip off after qmark
                filename = filename[0:qm_idx]
            }
            fmt.Printf("found fn: %s\n", filename)
            if !cache.InCache(feed.Url, filename) {
                // do the dl,  then add to cache 
                // only add 2 cache if dl is successful
                doneChan := rssd.Download(file_url, path.Join("downloads", filename))
                fmt.Println("waiting for download to finish")
                isDone :=  <-doneChan
                if isDone {
                    fmt.Println("download success!")
                    cache.AppendCache(feed.Url, filename)
                } else {
                    fmt.Println("download failed!")
                }
            }
        }
    }

   // fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}


