package gota_common

import (
	"github.com/magiconair/properties"
	"log"
	"os/user"
)

var P *properties.Properties = nil

func LoadConfig() *properties.Properties {
	if P == nil {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		P = properties.MustLoadFile(usr.HomeDir+"/config/ipc.properties", properties.UTF8)
	}
	return P
}
