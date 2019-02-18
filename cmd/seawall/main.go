package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/ieee0824/seawall/config"
	"github.com/ieee0824/seawall/tester"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	log.SetFlags(log.Llongfile)
	client, err := tester.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()

	confPath := flag.String("c", "", "")
	outDir := flag.String("o", "", "")
	autoOpenMode := flag.Bool("a", false, "")
	flag.Parse()
	if err := os.Mkdir(*outDir, 0777); err != nil && !strings.Contains(err.Error(), "file exists") {
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
		if err != nil {
			continue
		}

		for i, v := range bins {
			writeFileName := fmt.Sprintf("%s/%v.png", *outDir, i)
			f, err := os.Create(writeFileName)
			if err != nil {
				log.Fatalln(err)
			}
			f.Write(v)
			f.Close()

			if *autoOpenMode {
				if err := open.Run(writeFileName); err != nil {
					log.Println(err)
				}
			} else {
				fmt.Println(writeFileName)
			}
		}
	}
}
