package routes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/search", func(c *gin.Context) {

		query := c.Query("q")

		if query == "" {
			c.JSON(400, gin.H{"error": "Null or invalid query"})
			return
		}

		// Caminho para a pasta "Leaks"
		dirPath := "leaks"

		var results []gin.H
		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return nil
			}

			// Verifica se é um arquivo de texto (extensão .txt)
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".txt") {
				// Lê o conteúdo do arquivo
				content, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Printf("Erro ao ler o arquivo %s: %v\n", path, err)
					return nil
				}

				// Divide o conteúdo em linhas
				lines := strings.Split(string(content), "\n")

				// Encontrar a posição da query nas linhas
				var start, end int
				found := false
				for i, line := range lines {
					if strings.Contains(line, query) {
						// Define o início como 30 linhas antes da ocorrência
						start = i - 30
						if start < 0 {
							start = 0
						}
						// E o final 30 linhas apos
						end = i + 30
						if end > len(lines) {
							end = len(lines)
						}
						found = true
						break
					}
				}

				// Adiciona o resultado à lista apenas se algo foi encontrado
				if found {
					results = append(results, gin.H{
						"path":     path,
						"content":  strings.Join(lines[start:end], "\n"),
						"position": fmt.Sprintf("Line %d", start+1), // Adiciona 1 para a linha de início ser 1-indexed
					})
				}
			}

			return nil
		})

		if err != nil {
			fmt.Printf("Erro ao listar arquivos: %v\n", err)
			c.JSON(500, gin.H{"error": "Erro interno do servidor"})
			return
		}

		// Retorna os resultados da busca
		c.JSON(200, gin.H{"results": results})
	})
}
