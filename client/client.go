package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetObject(object, output, apikey, proxy string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", proxy+"/get", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Object", object)
	req.Header.Add("Api-Key", apikey)

	resp, err := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 401 {
		fmt.Printf("Error: %v\n", resp.Status)
		os.Exit(1)
	}

	ioutil.WriteFile(output, b, 0644)
}
