# Concurrent File Downloader

A high-performance command-line file downloader built in **Go** that accelerates downloads by splitting files into multiple chunks and downloading them concurrently using goroutines. The downloader uses HTTP Range requests, dynamic worker allocation, and thread-safe file writing to achieve fast, reliable, and efficient downloads.


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hiten016/Concurrent-file-downloader-in-Go.git
   cd Concurrent-file-downloader-in-Go
   ```

2. Build the executable:

   ```bash
   go build -o downloader main.go
   ```

## Usage

Download a file using:

```bash
./downloader -url <file_url> -dest <destination_file_path>
```

### Command-Line Options

- `-url` : URL of the file to download.
- `-dest` : Destination path where the downloaded file will be saved.

### Example

```bash
./downloader \
-url https://example.com/largefile.zip \
-dest ./downloads/largefile.zip
```

## How It Works

1. Retrieves the file metadata from the server.
2. Splits the file into fixed-size chunks (4 MB each).
3. Creates a dynamic pool of worker goroutines.
4. Downloads chunks concurrently using HTTP Range requests.
5. Writes downloaded chunks safely using mutex synchronization.
6. Updates the progress bar in real time.
7. Retries failed downloads if necessary.
8. Merges all chunks into the final output file.
9. Displays the total download time upon completion.

## Project Structure

```
Concurrent-file-downloader-in-Go/
├── main.go             # Application entry point
├── downloader.go       # Concurrent download logic
├── utils.go            # Helper functions
├── progress.go         # Progress bar implementation
├── go.mod
├── go.sum
└── README.md
```

