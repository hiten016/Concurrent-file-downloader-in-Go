package cmd

import (
    "flag"
    "fmt"
    "os"
    "concurrent-file-downloader/downloader"
)

func Execute() {
    url := flag.String("url", "", "URL of the file to download")
    dest := flag.String("dest", "", "Destination path")
    flag.Parse()

    if *url == "" || *dest == "" {
        fmt.Println("Usage: -url <file_url> -dest <destination_path>")
        os.Exit(1)
    }

    err := downloader.DownloadFile(*url, *dest)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
