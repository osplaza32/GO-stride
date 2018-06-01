package main

import "github.com/dghubble/sling"
import (
	"net/http"
	"time"
	"fmt"
)

type ApiParms struct{
	PPU string "ppu,omitempty"
}

func main() {
	params := &ApiParms{"ppu:xl3525"}

	req,_ := sling.New().Get("http://porsilapongo.cl/API_MOD/robo.php").QueryStruct(params).Request()
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp,_ :=client.Do(req)
	fmt.Printf("%#v\n", resp)
}
