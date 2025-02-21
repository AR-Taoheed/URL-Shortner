package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// URLShortner Type is a struct that holds the URLs and their time of creation
type URLEntry struct {
	OriginalURL string `json:"original_url"`

	CreatedAt time.Time `json:"created_at"`

	Clicks int `json:"clicks"`
} 

// map of string for URLEntry above
type URLShortener struct {
	urls map[string]URLEntry
}

func NewURLShortener() *URLShortener{
	return &URLShortener{
		urls: make(map[string]URLEntry),
	}
}

func (us *URLShortener) generateShortAlias() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const aliasLength = 7

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	alias := make([]byte, aliasLength)
	for i := range alias{
		alias[i] = charset[random.Intn(len(charset))]
	}
	return string(alias)
}

func (us *URLShortener) AddURL(originalURL string) string {
	alias := us.generateShortAlias()
	for us.urls[alias].OriginalURL != ""{
		alias = us.generateShortAlias()
	}

	entry := URLEntry{
		OriginalURL: originalURL,
		CreatedAt: time.Now(),
		Clicks: 0,
	}
	us.urls[alias] = entry
	return alias
}

func (us *URLShortener) GetOriginalURL(alias string) (string, error) {
	entry, exists := us.urls[alias]
	if !exists{
		return "", errors.New("alias not found")
	}

	entry.Clicks++
	us.urls[alias] = entry
	return entry.OriginalURL, nil
}


func (us *URLShortener) GetClicks(alias string) (int, error) {
	entry, exists := us.urls[alias]
	if !exists{
		return 0, errors.New("alias not found")
	}
	return entry.Clicks, nil
}

func (us *URLShortener) DeleteURL(alias string) error{
	if _, exists := us.urls[alias]; !exists{
		return errors.New("alias not found")
	}
	delete(us.urls, alias)
	return nil
}

//write for the list URLs method

func main() {
	shortener := NewURLShortener()
	alias := shortener.AddURL("https://www.github.com")
	fmt.Println("Short URL: ", alias)

	originalURL, err := shortener.GetOriginalURL(alias)
	if err != nil{
		fmt.Println("Error getting original URL: ", originalURL)
	}

	clicks, err := shortener.GetClicks(alias)
	if err != nil{fmt.Println("Error getting clicks: ", err)
	}	else{
		fmt.Println("Clicks: ", clicks)
	}

	err = shortener.DeleteURL(alias)
	if err != nil{
		fmt.Println("Error: ", err)
	}	else{
		fmt.Println("URL deleted successfully.")
	}
}