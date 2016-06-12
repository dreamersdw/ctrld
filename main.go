package main

import "fmt"
import "github.com/BurntSushi/toml"

type Config struct {
	Etcdaddr string
}

func main() {
	cfg := Config{}
	toml.DecodeFile("ctrld.toml", &cfg)
	fmt.Println(cfg.Etcdaddr)
}
