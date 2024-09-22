package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sodexo-backend-webapp.dynamify.com/menu/listV2?storeIds=17881", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:130.0) Gecko/20100101 Firefox/130.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("X-Auth-Token", "4r2pp7gmiljhcg6atahicn7fqe1srr2p")
	req.Header.Set("Version", "2.33.42")
	req.Header.Set("Partner-Id", "9")
	req.Header.Set("Application-Id", "2")
	req.Header.Set("Device-Platform", "webUnknown")
	req.Header.Set("Device-Model", "Unknown")
	req.Header.Set("Device-Uuid", "17032e5dfa7438")
	req.Header.Set("Origin", "https://everyday.dynamify.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var a AutoGenerated

	json.Unmarshal(body, &a)

	for _, store := range a {
		for _, menu := range store.Menu {
			fmt.Println(menu.Name)
			for _, sub := range menu.MenuSubcategories {
				fmt.Println(sub.Name)
				for _, item := range sub.MenuItems {
					fmt.Println(item.Name)
				}
			}
		}
	}
}
