// routes/routes.go
package routes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas no servidor Gin
func SetupRoutes(r *gin.Engine) {
	// Define uma rota que aceita a query "q"
	r.GET("/", func(c *gin.Context) {
		// Obtém o valor da query "q" a partir da URL
		query := c.Query("q")

		// Verifica se a query está presente
		if query == "" {
			c.JSON(400, gin.H{"error": "A query 'q' é obrigatória"})
			return
		}

		// Caminho para a pasta "Leaks"
		dirPath := "leaks"

		// Slice para armazenar os resultados da busca
		var results []gin.H

		// Lista os arquivos na pasta
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
				for i, line := range lines {
					if strings.Contains(line, query) {
						// Define o início como 30 linhas antes da ocorrência
						start = i - 30
						if start < 0 {
							start = 0
						}

						// Define o final como 30 linhas após a ocorrência
						end = i + 30
						if end > len(lines) {
							end = len(lines)
						}
						break
					}
				}

				// Adiciona o resultado à lista
				results = append(results, gin.H{
					"path":     path,
					"content":  strings.Join(lines[start:end], "\n"),
					"position": fmt.Sprintf("Linha %d", start+1), // Adiciona 1 para a linha de início ser 1-indexed
				})
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
