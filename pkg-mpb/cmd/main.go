package main

import (
	"log"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func main() {
	// Start the download process
	url := "https://github.com/dingodb/dingoadm/releases/download/latest/commit_id" // Replace with the actual file URL
	filePath := "commit_id"                                                         // Replace with the desired file path

	// Create a new downloader instance
	dl := NewDownloader(url, filePath)

	// Get file size first to create accurate progress bar
	size, err := dl.GetFileSize()
	if err != nil {
		log.Fatalf("Failed to get file size: %v", err)
	}

	// Create a progress bar with actual file size
	p := mpb.New()
	bar := p.AddBar(size,
		mpb.PrependDecorators(
			decor.Name("Downloading: ", decor.WC{W: 15}),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
			decor.Name(" "),
			decor.EwmaETA(decor.ET_STYLE_GO, 30),
		),
	)

	err = dl.Download(bar)
	if err != nil {
		log.Fatalf("Download failed: %v", err)
	}

	p.Wait()
	log.Println("Download completed successfully!")
}
