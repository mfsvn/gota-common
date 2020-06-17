package gota_common

import (
	"github.com/magiconair/properties"
	"log"
	"os/user"
)

func LoadPropsFromUsrHome(filename string) *properties.Properties {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return properties.MustLoadFile(usr.HomeDir+filename, properties.UTF8)
}
