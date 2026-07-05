# Concurrent File Downloader

Concurrent File Downloader is a command-line tool written in Go for efficiently downloading large files from a given URL. It utilizes concurrent downloading, mutexes for safe file handling, and various optimizations to reduce overall download time.

## Features

- **Concurrent Downloading**: Downloads file chunks concurrently using multiple goroutines, leveraging available network bandwidth.
- **Progress Bar**: Displays a real-time progress bar with percentage completion and estimated time remaining.
- **Robust Error Handling**: Manages network errors, retries failed downloads, and ensures data integrity through checksum verification.
- **Total Download Time**: Calculates and displays the total time taken to download the file upon completion.
- **Chunked Download**: Breaks down large files into smaller, manageable chunks (4MB each) to optimize network usage and memory consumption.
- **Dynamic Worker Pool**: Adjusts the number of concurrent workers dynamically based on file size and system resources, maximizing download efficiency.
- **HTTP Client Optimization**: Uses a custom HTTP client with increased idle connections and optimized timeout settings for reliable and efficient file retrieval.
- **Safe File Handling with Mutexes**: Ensures thread-safe file writes using mutexes (`sync.Mutex`), preventing concurrent writes from causing data corruption.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/4d4r5h/concurrent-file-downloader.git
   cd concurrent-file-downloader
   ```

2. Build the executable:

   ```bash
   go build -o downloader main.go
   ```

## Usage

To download a file:

```bash
./downloader -url <file_url> -dest <destination_file_path>
```

### Options

- `-url`: Specifies the URL of the file to download.
- `-dest`: Specifies the destination path where the downloaded file will be saved.

### Example

```bash
./downloader -url https://example.com/largefile.zip -dest ./downloads/largefile.zip
```