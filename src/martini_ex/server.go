package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	m := martini.Classic()
	m.Get("/", func() string {
		/*hmmm start here.. */
		res, err := http.Get("https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20weather.forecast%20where%20woeid%20in%20(select%20woeid%20from%20geo.places(1)%20where%20text%3D%22nome%2C%20ak%22)&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys")
		if err != nil {
			log.Fatal(err)
		}

		response, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		return string(response)

	})

	m.Run()
}

/*dhxjqsyy2zbajubwhwvptvqh*/
