package main

import (
	"encoding/json"
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
	var result = map[string]interface{}{}

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
			result[writeFileName] = false
			f, err := os.Create(writeFileName)
			if err != nil {
				log.Println(err)
				continue
			}
			if _, err := f.Write(v); err != nil {
				f.Close()
				log.Println(err)
				continue
			}
			f.Close()

			if *autoOpenMode {
				if err := open.Run(writeFileName); err != nil {
					log.Println(err)
				}
			} else {
				result[writeFileName] = true
			}
		}
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(resultJson))
}
