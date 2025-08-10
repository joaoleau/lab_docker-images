package main

import (
	"fmt"
	"os"
	"github.com/joaoleau/golint/internal/structs"
	"github.com/joaoleau/golint/internal/utils"
	"gopkg.in/yaml.v3"
)


var args []string
var infos []*structs.JSONInfo

func init() {
	args = os.Args[1:]
}

func decodeYaml(paths []string) {
	for _, path := range paths {
		var content interface{}

		info := &structs.JSONInfo{
			FilePath: path,
			IsYaml:   0,
			Infos:    make([]string, 0),
			Errors:   make([]string, 0),
		}

		data, err := os.ReadFile(path)
		if err != nil {
			info.Errors = append(info.Errors, fmt.Sprintf("Arquivo %s não foi encontrado!", path))
			infos = append(infos, info)
			continue
		}
		info.Infos = append(info.Infos, fmt.Sprintf("Arquivo %s foi encontrado!", path))

		err = yaml.Unmarshal(data, &content)
		if err != nil {
			info.Errors = append(info.Errors, fmt.Sprintf("Erro ao decodificar YAML: %v", err))
		} else {
			info.IsYaml = 1
			info.Infos = append(info.Infos, "YAML válido!")
		}

		infos = append(infos, info)
	}
}

func main() {
	decodeYaml(args)
	utils.PrintJSONInfos(infos)
}
