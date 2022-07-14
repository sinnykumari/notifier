package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Notify whishlist items")
	resp, err := http.Get("https://www.spicevillage.eu/products/fortune-mustard-oil-edible-500ml")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	searchItem := "Fortune Mustard Oil"
	fmt.Println("Search for", searchItem)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
