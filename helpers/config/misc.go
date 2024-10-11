package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func CreateTemplate(dir string, lang string) map[string]interface{} {
	scr := utils.DefaultScripts[lang]

	for k, sc := range scr {
		if dir == "" {
			sc = fmt.Sprintf(sc.(string), ".")
		} else {
			sc = fmt.Sprintf(sc.(string), dir)
		}
		
		scr[k] = sc
	}

	return scr
}

func WriteConfig(cfg utils.PMTConfig) error {
	jsonData, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile("pmt-config.json", jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
