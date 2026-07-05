package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const chunkSize = 4 * 1024 * 1024 // 4MB
const maxWorkerCount = 100        // Number of concurrent workers

func DownloadFile(url, dest string) error {
	startTime := time.Now()

	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	fileSize := resp.ContentLength

	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	chunkCount := int(fileSize / chunkSize)
	if fileSize%chunkSize != 0 {
		chunkCount++
	}

	bar := NewProgressBar(chunkCount)

	jobs := make(chan int, chunkCount)
	results := make(chan error, chunkCount)
	var wg sync.WaitGroup
	var fileMutex sync.Mutex

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: true,
		},
	}

	for i := 0; i < maxWorkerCount; i++ {
		wg.Add(1)
		go worker(url, file, jobs, results, &wg, bar, client, &fileMutex)
	}

	for i := 0; i < chunkCount; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for err := range results {
		if err != nil {
			return err
		}
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.Size() != fileSize {
		return fmt.Errorf("downloaded file size mismatch: expected %d bytes, got %d bytes", fileSize, info.Size())
	}

	endTime := time.Now()
	fmt.Printf("\nDownload completed in %v\n", endTime.Sub(startTime))

	return nil
}

func worker(url string, file *os.File, jobs <-chan int, results chan<- error, wg *sync.WaitGroup, bar *ProgressBar, client *http.Client, fileMutex *sync.Mutex) {
	defer wg.Done()
	for index := range jobs {
		err := downloadChunk(url, file, index, bar, client, fileMutex)
		if err != nil {
			results <- fmt.Errorf("error downloading chunk %d: %v", index, err)
		} else {
			results <- nil
		}
	}
}

func downloadChunk(url string, file *os.File, index int, bar *ProgressBar, client *http.Client, fileMutex *sync.Mutex) error {
	rangeHeader := fmt.Sprintf("bytes=%d-%d", index*chunkSize, (index+1)*chunkSize-1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", rangeHeader)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := make([]byte, chunkSize)
	n, err := io.ReadFull(resp.Body, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}

	fileMutex.Lock()
	_, err = file.WriteAt(buf[:n], int64(index*chunkSize))
	fileMutex.Unlock()
	if err != nil {
		return err
	}

	bar.Add(1)

	return nil
}
