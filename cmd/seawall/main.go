package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/ieee0824/seawall/config"
	"github.com/ieee0824/seawall/tester"
)

func main() {
	client, err := tester.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()

	confPath := flag.String("c", "", "")
	outDir := flag.String("o", "", "")
	flag.Parse()
	if err := os.Mkdir(*outDir, 0777); err != nil {
		log.Fatalln(err)
	}

	c := &config.Config{}
	f, err := os.Open(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		log.Fatalln(err)
	}
	f.Close()

	for _, t := range c.Targets {
		bins, err := client.ScreenShot(t)
		fmt.Println(len(bins), err)

		for i, v := range bins {
			f, err := os.Create(fmt.Sprintf("%s/%v.png", *outDir, i))
			if err != nil {
				log.Fatalln(err)
			}
			f.Write(v)
			f.Close()
		}
	}
}