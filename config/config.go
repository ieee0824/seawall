package config

import (
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"
)

const (
	DefaultUserAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	DefaultPCUserAgent     = DefaultUserAgent
	DefaultMobileUserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36"
)

const (
	SizeFullHD         = "FullHD"
	SizeWUXGA          = "WUXGA"
	Size4K             = "4k"
	SizeiPhone5s       = "iPhone5s"
	SizeiPhone6        = "iPhone6"
	SizeiPhone6Plus    = "iPhone6Plus"
	SizeiPhone6s       = "iPhone6s"
	SizeiPhone6sPlus   = "iPhone6sPlus"
	SizeiPhoneSE       = "iPhoneSE"
	SizeiPhone7        = "iPhone7"
	SizeiPhone7Plus    = "iPhone7Plus"
	SizeiPhone8        = "iPhone8"
	SizeiPhone8Plus    = "iPhone8Plus"
	SizeiPhoneX        = "iPhoneX"
	SizeiPhoneXR       = "iPhoneXR"
	SizeiPhoneXS       = "iPhoneXS"
	SizeiPhoneXSMax    = "iPhoneXSMax"
	SizeUlefoneArmor2  = "Ulefone Armor 2"
	SizeUlefoneArmor3  = "Ulefone Armor 3"
	SizeUlefoneArmor3T = "Ulefone Armor 3T"
	SizeUlefoneArmor5  = "Ulefone Armor 5"
	SizeUlefoneArmor6  = "Ulefone Armor 6"
	SizePixel2         = "Pixel 2"
	SizePixel2XL       = "Pixel 2 XL"
)

var WindowSizeList = map[string]*Bounds{
	SizeFullHD:         {1920, 1080},
	SizeWUXGA:          {1920, 1200},
	Size4K:             {4096, 2160},
	SizeiPhone5s:       {640, 1136},
	SizeiPhone6:        {750, 1334},
	SizeiPhone6Plus:    {1080, 1920},
	SizeiPhone6s:       {750, 1334},
	SizeiPhone6sPlus:   {1080, 1920},
	SizeiPhoneSE:       {640, 1136},
	SizeiPhone7:        {750, 1334},
	SizeiPhone7Plus:    {1080, 1920},
	SizeiPhone8:        {750, 1334},
	SizeiPhone8Plus:    {1080, 1920},
	SizeiPhoneX:        {1125, 2436},
	SizeiPhoneXR:       {828, 1792},
	SizeiPhoneXS:       {1125, 2436},
	SizeiPhoneXSMax:    {1242, 2688},
	SizeUlefoneArmor2:  {920, 1080},
	SizeUlefoneArmor3:  {1080, 2160},
	SizeUlefoneArmor3T: {1080, 2160},
	SizeUlefoneArmor5:  {720, 1512},
	SizeUlefoneArmor6:  {1080, 2246},
	SizePixel2:         {1080, 1920},
	SizePixel2XL:       {1440, 2880},
}

type Bounds struct {
	W int `yaml:"w"`
	H int `yaml:"h"`
}

func (b *Bounds) Arg() []string {
	if b == nil || b.W == 0 || b.H == 0 {
		return []string{"--window-size=1920,4320"}
	}
	return []string{fmt.Sprintf("--window-size=%d,%d", b.W, b.H)}
}

type UA string

func (u UA) Arg() []string {
	if u == "" {
		return []string{"--user-agent=\"" + DefaultUserAgent + `"`}
	}
	return []string{"--user-agent=\"" + string(u) + `"`}
}

type DisableHeadlessOption bool

func (h DisableHeadlessOption) Arg() []string {
	if !h {
		return []string{"--headless"}
	}
	return []string{}
}

type FullPageOption bool

func (f FullPageOption) Arg() []string {
	if f {
		return []string{"--fullPage"}
	}
	return []string{}
}

type Preset string

var presets = map[Preset]struct {
	Memo       string
	UserAgent  UA
	WindowSize *Bounds
}{
	"iphone5s_safari": {
		"ios iOS 11.2.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_1 like Mac OS X) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0 Mobile/15C153 Safari/604.1",
		WindowSizeList[SizeiPhone5s],
	},
	"iphone6s_safari": {
		"ios iOS 11.2.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_1 like Mac OS X) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0 Mobile/15C153 Safari/604.1",
		WindowSizeList[SizeiPhone6s],
	},
	"iphone6s_plus_safari": {
		"ios iOS 11.2.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_1 like Mac OS X) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0 Mobile/15C153 Safari/604.1",
		WindowSizeList[SizeiPhone6sPlus],
	},
	"iphoneX": {
		"ios iOS 11.2.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_1 like Mac OS X) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0 Mobile/15C153 Safari/604.1",
		WindowSizeList[SizeiPhoneX],
	},
}

type ClientOption struct {
	Init            []map[Cmd]interface{} `yaml:"init"`
	Tag             string                `yaml:"tag"`
	ClientPreset    Preset                `yaml:"client_preset"`
	UserAgent       UA                    `yaml:"user_agent"`
	WindowSize      *Bounds               `yaml:"window_size"`
	FullPage        FullPageOption        `yaml:"full_page"`
	DisableHeadless DisableHeadlessOption `yaml:"disable_headless"`
	Delay           time.Duration         `yaml:"delay"`
}

func (c *ClientOption) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(out)
}

func (c *ClientOption) Args() []string {
	p, ok := presets[c.ClientPreset]
	if ok {
		c.WindowSize = p.WindowSize
		c.UserAgent = p.UserAgent
	}
	ret := []string{
		"--disable-gpu",
		"--log-level=3",
		"--no-sandbox",
		"--ignore-certificate-errors",
		"--disable-application-cache",
	}

	ret = append(ret, c.UserAgent.Arg()...)
	ret = append(ret, c.DisableHeadless.Arg()...)
	ret = append(ret, c.FullPage.Arg()...)

	if c.FullPage {
		var b *Bounds
		if c.WindowSize == nil {
			b = &Bounds{
				WindowSizeList[SizeFullHD].W,
				4320,
			}
		} else {
			b = &Bounds{
				c.WindowSize.W,
				4320,
			}
		}
		c.WindowSize = b
		ret = append(ret, b.Arg()...)
	} else {
		ret = append(ret, c.WindowSize.Arg()...)
	}
	fmt.Println(ret)
	return ret
}

type Cmd string

type Target struct {
	URL           string         `yaml:"url"`
	ClientOptions []ClientOption `yaml:"client_options"`
}

func (t *Target) String() string {
	out, err := yaml.Marshal(t)
	if err != nil {
		return ""
	}
	return string(out)
}

type Config struct {
	Targets []*Target `yaml:"targets"`
}

func (c *Config) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(out)
}
