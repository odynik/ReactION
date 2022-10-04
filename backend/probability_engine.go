package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type uni_distro_params struct {
	value_type  int
	total_num   int
	base        int
	min         int
	max         int
	format_type string
	rnd_type    string
}

// func newUniDistroParams(uni_rnd_sample uni_distro_params) *uni_distro_params {
// 		sample := uni_distro_params{name: name}
// 		sample.age = 42
// 		return &p
// 	}

func uniDistroSamples(num int) {
	requestURL := fmt.Sprintf("https://www.random.org/integers/?num=%s&min=1&max=6&col=1&base=10&format=plain&rnd=new", strconv.Itoa(num))
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body:\n%s\n", resBody)
}

func foo() {
	fmt.Printf("I am here!\n")
}
