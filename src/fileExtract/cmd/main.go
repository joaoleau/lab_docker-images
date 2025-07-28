package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/joaoleau/golint/structs"
	"github.com/joaoleau/golint/utils"
)

var args []string
var infos []*structs.JSONInfo
var files []*structs.File

func init() {
	args = os.Args[1:]
}

func yamlExtract(path string, file *structs.File) {
	fi, sc, err := utils.ReaderFile(path)
	info := &structs.JSONInfo{
		FilePath: path,
		IsYaml:   1,
		Infos:    make([]string, 0),
		Errors:   make([]string, 0),
	}

	if err != nil {
		info.Errors = append(info.Errors, fmt.Sprintf("Erro ao abrir %s: %v\n", path, err))
		return
	}
	
	inSource := false
	inHelm := false
	inValue := false
	newRawLines := make([]string, 0)
	inValuesFrom := false
	for sc.Scan() {
		line := sc.Text()
		trimmedLine := strings.TrimSpace(sc.Text())

		//SOURCE
		if strings.TrimSuffix(trimmedLine, ":") == "source" { inSource = true }
		
		if inSource && strings.Split(trimmedLine, ":")[0] == "targetRevision" { 
			linesRaw := strings.Split(trimmedLine, ":")
			line := strings.TrimSpace(linesRaw[len(linesRaw) - 1])
			file.TargetRevision = line
		}
		if inSource && strings.Split(trimmedLine, ":")[0] == "path" { 
			linesRaw := strings.Split(trimmedLine, ":")
			line := strings.TrimSpace(linesRaw[len(linesRaw) - 1])
			file.ChartPath = line
		}

		//HELM
		if strings.TrimSuffix(trimmedLine, ":") == "helm" { inHelm = true; inSource = false }
		if inHelm && strings.TrimSuffix(trimmedLine, ":") == "valuesFrom" { 
			inValuesFrom = true
			continue
		}


		if inValuesFrom {
			if strings.HasPrefix(trimmedLine, "- ") {
				trimmed := strings.TrimPrefix(trimmedLine, "- ")
				file.ValuesFrom = append(file.ValuesFrom, file.ChartPath+"/"+trimmed)
			}
			
			inValuesFrom = false
		}
		
		if inHelm && strings.Split(trimmedLine, ":")[0] == "values" { inValue = true; file.ValuesFrom = append(file.ValuesFrom, path); continue }

		//VALUE
		if inValue {
			newRawLines = append(newRawLines, line)
		}
	}

	if len(newRawLines) > 0 {
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		outputFilename := utils.RawPrefix + base
		outputPath := filepath.Join(dir, outputFilename)
		err := os.WriteFile(outputPath, []byte(strings.Join(newRawLines, "\n")), 0644)
		if err != nil {
			info.Errors = append(info.Errors, fmt.Sprintf("Erro ao escrever %s: %v", outputPath, err))
		} else {
			info.Infos = append(info.Infos, fmt.Sprintf("Arquivo %s criado com conteúdo valuesFrom", outputPath))
		}
	}

	if err := fi.Close(); err != nil {
		info.Errors = append(info.Errors, fmt.Sprintf("Erro ao fechar %s: %v\n", path, err))
	}

	infos = append(infos, info)
}

func IsYAMLFile(path string, file *structs.File) {
	ext := strings.ToLower(path)
	if strings.HasSuffix(ext, ".yaml") || strings.HasSuffix(ext, ".yml") { 
		file.IsYaml = 1 
	}
}

func main() {
	for _, path := range args {
		file := &structs.File{
			FilePath: path,
			IsYaml:   0,
			TargetRevision: "",
			ChartPath: "",
			ValuesFrom: make([]string, 0),
		}
		IsYAMLFile(path, file)
		if file.IsYaml > 0 {yamlExtract(path, file)}
		files = append(files, file)
	}
	utils.WriterFiles("files.json", files)
	utils.PrintJSONInfos(infos)
}
