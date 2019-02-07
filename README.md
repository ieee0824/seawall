# seawall
browser test tool written in golang.

# install

```
$ go get -u github.com/ieee0824/seawall/cmd/seawall
```

# settings

save `config.yml`.
and exec `seawall -f config.yml -o outdir`

```
targets:
- url: https://google.co.jp
  client_options:
  - tag: ulefone_armor_2
    user_agent: Mozilla/5.0 (Linux; Android 7.0; Armor_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36
    window_size:
      w: 920
      h: 1080
    full_page: false
    disable_headless: true
    delay: 0s
  - tag: iPhoneX
    client_preset: iphoneX
    full_page: false
    disable_headless: true
    delay: 0s
- url: https://github.com
  client_options:
  - tag: iPhoneX
    client_preset: iphoneX
    full_page: false
    disable_headless: false
    delay: 1s
```