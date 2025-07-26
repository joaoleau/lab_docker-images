package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type JSONInfo struct {
	FilePath string   `json:"filePath"`
	IsYaml   int8     `json:"isYaml"`
	Infos    []string `json:"info"`
	Errors   []string `json:"error"`
}

func printJSONInfos(infos []*JSONInfo) {
	out, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao gerar JSON: %v\n", err)
		return
	}
	fmt.Println(string(out))
}

var args []string
var infos []*JSONInfo

func init() {
	args = os.Args[1:]
}

func decodeYaml(paths []string) {
	for _, path := range paths {
		var content interface{}

		info := &JSONInfo{
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
	printJSONInfos(infos)
}
