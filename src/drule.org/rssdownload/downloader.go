package rssdownload

import (
    "net/http"
    "os"
    "fmt"
    "io"
)

const READ_SIZE int64 = 1024 * 1024 * 5 // 5MB

func Download(file_url, destination_file string) (chan bool) {
    channel := make(chan bool, 5)
    f, err := os.Create(destination_file)
    if err != nil {
        fmt.Printf("error downloading %s, %s\n", destination_file, err)
        if f != nil {
            f.Close()
        }
        os.Remove(destination_file)
        channel <- false
    }

    //start download in an anonymous goroutine
    //because well it could be slow or something

    go func() {
        fmt.Printf("start download of %s\n", file_url)
        res, err := http.Get(file_url)
        switch {
        case err != nil:
            fmt.Println(err)
            defer f.Close()
            channel <- false
        case res.StatusCode != 200:
            fmt.Printf("http status code %d", res.StatusCode )
            channel <- false
        default:
            var read_bytes int64
            var err error
            for {
                read_bytes, err = io.CopyN(f, res.Body, READ_SIZE)

                if err != nil && err != io.EOF {
                    fmt.Printf("read %d, error %s\n", read_bytes, err)
                    channel <- false
                    break
                }
                if err != nil && err == io.EOF {
                    channel <- true
                    break
                }
            }
        }
    }()

    return channel
}

