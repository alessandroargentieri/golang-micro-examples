package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
 
func main() {
 
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
 
	req.Header.Set("Cache-Control", "no-cache")
 
	client := &http.Client{Timeout: time.Second * 10}
 
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
 
	fmt.Printf("%s\n", body)
}
