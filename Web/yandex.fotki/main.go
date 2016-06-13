package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/SlyMarbo/rss"
)

const fetchUrl = "http://api-fotki.yandex.ru/api/podhistory/poddate;2015-04-01T12:00:00Z/?limit=100"

const targetDir = "./images/"

func main() {

	feed, err := rss.Fetch(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Items {
		log.Print("Processing: ", item.Title)
		imgUrl, err := ImgUrl(item.Enclosures)
		if err != nil {
			log.Print(err)
			continue
		}

		u, err := url.Parse(imgUrl)
		if err != nil {
			log.Print(err)
			continue
		}
		splitedPath := strings.Split(u.Path, "/")
		fileName := targetDir + splitedPath[len(splitedPath)-1] + ".jpg"
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			continue
		}

		img, format, err := LoadImage(imgUrl)
		if err != nil {
			log.Print(err)
			continue
		}
		if format == "" {
			log.Print("Empty format")
			continue
		}

		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Print(err)
			continue
		}

		if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 85}); err != nil {
			log.Print(err)
			continue
		}
	}

}

func ImgUrl(enclosures []*rss.Enclosure) (string, error) {
	for _, encl := range enclosures {
		if encl.Rel == "edit-media" {
			return encl.Url, nil
		}
	}
	return "", fmt.Errorf("Rel:'edit-media' not found")
}

func LoadImage(imgUrl string) (img image.Image, format string, err error) {

	resp, err := http.Get(imgUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("Status code: %d (%s)", resp.StatusCode, resp.Status)
	}
	log.Print("Content-Type: ", resp.Header.Get("Content-Type"))

	img, format, err = image.Decode(resp.Body)
	if err != nil {
		return nil, "", err
	}
	log.Printf("Image format: %s [%dx%d]", format, img.Bounds().Dx(), img.Bounds().Dy())
	return img, format, nil
}
