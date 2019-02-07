package tester

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/ieee0824/seawall/config"
	"github.com/sclevine/agouti"
)

type ScreenShotResults map[string][]byte

type Client struct {
	driver *agouti.WebDriver
}

func NewClient() (*Client, error) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		return nil, err
	}
	return &Client{
		driver: driver,
	}, nil
}

func (c *Client) Stop() error {
	return c.driver.Stop()
}

func (_ *Client) initialize(page *agouti.Page, c *config.ClientOption) error {
	for _, o := range c.Init {
		for cmd, v := range o {
			switch v := v.(type) {
			case string:
				switch cmd {
				case "click":
					if err := page.Find(v).Click(); err != nil {
						return err
					}
				case "script":
					if err := page.RunScript(v, nil, nil); err != nil {
						return err
					}
				case "scroll":
					if err := page.Find(v).ScrollFinger(0, 0); err != nil {
						return err
					}
				}
			case map[interface{}]interface{}:
				for sel, s := range v {
					switch cmd {
					case "fill":
						if err := page.Find(sel.(string)).Fill(s.(string)); err != nil {
							return err
						}
					}
				}
			default:
				fmt.Println(reflect.TypeOf(v))
			}
		}
	}

	return nil
}

func (c *Client) ScreenShot(t *config.Target) (ScreenShotResults, error) {
	ret := ScreenShotResults{}
	for _, o := range t.ClientOptions {
		page, err := c.driver.NewPage(agouti.Desired(agouti.Capabilities{
			"chromeOptions": map[string][]string{
				"args": o.Args(),
			},
		}))
		if err != nil {
			log.Println(err)
			continue
		}
		if err := page.Size(o.WindowSize.W, o.WindowSize.H); err != nil {
			log.Println(err)
			continue
		}

		if err := page.Navigate(t.URL); err != nil {
			log.Println(err)
			continue
		}

		if err := c.initialize(page, &o); err != nil {
			log.Println(err)
			continue
		}

		time.Sleep(o.Delay)

		fileName := fmt.Sprintf("/tmp/%s-%s.png", o.Tag, time.Now().String())
		if err := page.Screenshot(fileName); err != nil {
			return nil, err
		}

		bin, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}

		ret[o.Tag] = bin
		if err := os.Remove(fileName); err != nil {
			return nil, err
		}
		if err := page.CloseWindow(); err != nil {
			return nil, err
		}
		if err := page.ClearCookies(); err != nil {
			return nil, err
		}
	}
	return ret, nil
}
