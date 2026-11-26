package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	WISHLIST_URL := os.Getenv("WISHLIST_URL")

	if WISHLIST_URL == "" {
		panic("WISHLIST_URL not provided")
	}
	STEAM_ID := os.Getenv("STEAMID")
	if STEAM_ID == "" {
		panic("STEAMID not provided")
	}

	getWishListURL := WISHLIST_URL + "?steamid=" + STEAM_ID
	fmt.Println("Calling Steam API with: ", getWishListURL)
	resp, err := http.Get(getWishListURL)
	if err != nil {
		log.Fatal("Error while trying to get wishlist")
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read body response")
		panic(err)
	}

	fmt.Println("Response status: ", resp.Status)
	fmt.Println("Response body: ", string(body))

}
