package main

// import (
// 	"fmt"
// 	"os"
// 	"helm.sh/helm/v3/pkg/chart/loader"
// 	"helm.sh/helm/v3/pkg/chartutil"
// 	"helm.sh/helm/v3/pkg/engine"
// )

// func main() {
// 	// Caminho para o chart Helm
// 	chartPath := "./meu-chart"

// 	// Carrega o chart
// 	chart, err := loader.Load(chartPath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao carregar chart: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Valores para o template
// 	vals := map[string]interface{}{
// 		"nome": "valor",
// 	}

// 	// Prepara os valores
// 	valsYaml, err := chartutil.ToRenderValues(chart, chartutil.Values(vals), chartutil.ReleaseOptions{
// 		Name:      "meu-release",
// 		Namespace: "default",
// 	}, nil)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao preparar valores: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Renderiza os templates
// 	rendered, err := engine.Render(chart, valsYaml)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao renderizar: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Exibe os arquivos renderizados
// 	for filename, content := range rendered {
// 		fmt.Printf("--- %s ---\n%s\n", filename, content)
// 	}
// }