package storage

import "tantan-demo/config"

func init() {
	c := "../config/config.json"
	if err := config.Load(c); err != nil {
		panic(err)
	}
	Connect()
}
