package rssdownload

import (
    "os"
    "fmt"
    "encoding/json"
)


type Cache struct {
    filename string
    seenfiles map[string][]string
}

func (c *Cache) JsonOut() (string) {
    b, _ := json.Marshal(c.seenfiles)
    return string(b)
}


func NewCache(filename string) (*Cache) {
    f, err := os.Open(filename)
    c := Cache{filename:filename}
    if os.IsExist(err) {
        // read file
        dec := json.NewDecoder(f)        
        err = dec.Decode(&c.seenfiles)
        if err != nil {
            fmt.Printf("error reading %s\n", filename)
            return nil
        }
    } else {
        c.seenfiles = map[string][]string{}
    }

 
    return &c
}
