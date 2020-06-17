package gota_common

import (
	"github.com/magiconair/properties"
	"log"
	"os/user"
	"path/filepath"
)

func LoadPropsFromUsrHome(filename string) *properties.Properties {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return properties.MustLoadFile(usr.HomeDir + string(filepath.Separator)+filename, properties.UTF8)
}
