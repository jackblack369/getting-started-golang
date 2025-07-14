package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vbauerster/mpb/v8"
)

type Downloader struct {
	url      string
	filePath string
}

func NewDownloader(url, filePath string) *Downloader {
	return &Downloader{
		url:      url,
		filePath: filePath,
	}
}

func (d *Downloader) GetFileSize() (int64, error) {
	resp, err := http.Head(d.url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get file info: %s", resp.Status)
	}

	return resp.ContentLength, nil
}

func (d *Downloader) Download(bar *mpb.Bar) error {
	resp, err := http.Get(d.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s", resp.Status)
	}

	file, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a proxy reader that updates the progress bar
	proxyReader := bar.ProxyReader(resp.Body)
	defer proxyReader.Close()

	// Copy data from response to file through the proxy reader
	_, err = io.Copy(file, proxyReader)
	return err
}
