package main

import (
	"github.com/magiconair/properties"
)

const properyFile = "/home/narang/work/src/github.com/viveknarang/kramaAPI/api.properties"

func main() {

	p := properties.MustLoadFile(properyFile, properties.UTF8)
	url := p.GetString("db.mongo.url", "def")
	port := p.GetString("db.mongo.port", "def")

	connect(url, port)
	disconnect()

}
