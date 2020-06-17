package gota_common

import (
	"github.com/magiconair/properties"
	"log"
	"os"
	"os/user"
)

func LoadPropsFromUsrHome(filename string) *properties.Properties {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return properties.MustLoadFile(usr.HomeDir + string(os.PathSeparator)+filename, properties.UTF8)
}

func LoadPrjProps(filename string) *properties.Properties {
	prjhome, err := os.Getwd()
	if err != nil {
		log.Fatal( err )
	}
	return properties.MustLoadFile(prjhome + string(os.PathSeparator)+filename, properties.UTF8)
}
