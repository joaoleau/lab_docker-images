package structs

type JSONInfo struct {
	FilePath string   `json:"filePath"`
	IsYaml   int8     `json:"isYaml"`
	Infos    []string `json:"info"`
	Errors   []string `json:"error"`
}

type File struct {
	FilePath string `json:"filePath"`
	IsYaml int8 `json:"isYaml"`
	TargetRevision string `json:"targetRevision"`
	ChartPath string `json:"chartPath"`
	ValuesFrom []string `json:"valuesFrom"`
}