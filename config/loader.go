package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

//Conf is the global configure instance
var Conf config

type config struct {
	Process process
	Path    path
}

type process struct {
	Backup bool
}

type path struct {
	LogPath string
	OutPath string
}

/*LoadConfig Loads the configurations parameters stored on
* the configuration file (config.toml)
 */
func LoadConfig(cf string) error {

	if _, err := toml.DecodeFile(cf, &Conf); err != nil {
		log.Fatal("Load config failed", err)
		return err
	}
	return nil
}
