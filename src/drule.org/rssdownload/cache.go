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

func (c *Cache) cacheFor(name string) ([]string) {
    if v, e := c.seenfiles[name]; e {
       return v
    } else {
       v := make([]string, 0, 10)
       c.seenfiles[name] = v
       return v
    }
}

/**
 * check if a value is in the slice of a map key
 */
func (c *Cache) InCache(name string, value string ) (bool) {
    if s, e := c.seenfiles[name]; e {
        for _,v := range s {
            if v == value {
                return true
            }
        }
        return false
    } else {
        return false
    }
}

/**
 *
 */
func (c *Cache) AppendCache(name string, value string) {
    if c.InCache(name, value) {
        return
    }

    s, e := c.seenfiles[name]
    if e {
        // check if we need to resize slice
        if cap(s) == len(s) {
            old_cap := cap(s)
            new_s := make([]string, old_cap+1, old_cap + 10)
            copied := copy(new_s, s)
            if (copied != old_cap) {
                // log an error
                fmt.Printf("problem copying %d from old slice\n", old_cap)
            }
            new_s[len(s)] = value
            c.seenfiles[name] = new_s
        } else {
            s = append(s, value)
            c.seenfiles[name] = s
        }
    } else {
        s = make([]string, 1, 5)
        s[0] = value
        c.seenfiles[name] = s
    }
}

func (c *Cache) Save() {
    // write out seenfiles as json 
    f, err := os.Create(c.filename)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        defer f.Close()
        encoder := json.NewEncoder(f)
        err := encoder.Encode(c.seenfiles)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func NewCache(filename string) (*Cache) {
    f, err := os.Open(filename)
    c := Cache{filename:filename}
    if err == nil {
        defer f.Close()
        // read file
        dec := json.NewDecoder(f)
        err = dec.Decode(&c.seenfiles)
        if err != nil {
            fmt.Printf("error reading %s\n", filename)
            return nil
        }
    } else {
        c.seenfiles = make(map[string][]string)
    }

    return &c
}
