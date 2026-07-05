package downloader

import (
    "github.com/schollz/progressbar/v3"
)

type ProgressBar struct {
    bar *progressbar.ProgressBar
}

func NewProgressBar(max int) *ProgressBar {
    bar := progressbar.New(max)
    return &ProgressBar{bar: bar}
}

func (pb *ProgressBar) Add(value int) {
    pb.bar.Add(value)
}
