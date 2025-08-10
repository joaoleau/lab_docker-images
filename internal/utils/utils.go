package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"github.com/joaoleau/golint/internal/structs"
)

const RawPrefix = "RAW_VALUE_"

func ReaderFile(path string) (*os.File, *bufio.Scanner, error) {
    fi, err := os.Open(path)
    if err != nil {
        return nil, nil, err
    }
    scanner := bufio.NewScanner(fi)
    return fi, scanner, nil
}

func WriterFiles(path string, files []*structs.File) error {
    data, err := json.MarshalIndent(files, "", "  ")
    if err != nil {
        return fmt.Errorf("erro ao serializar arquivos: %w", err)
    }

    if err := os.WriteFile(path, data, 0644); err != nil {
        return fmt.Errorf("erro ao escrever arquivo %s: %w", path, err)
    }
    return nil
}

func PrintJSONInfos(infos []*structs.JSONInfo) {
	out, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao gerar JSON: %v\n", err)
		return
	}
	fmt.Println(string(out))
}