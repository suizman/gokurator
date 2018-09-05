package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetObject(object, output, proxy string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", proxy+"/get", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("object", object)
	resp, err := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(output, b, 755)
}
