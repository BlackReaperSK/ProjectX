package routes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	linesAroundOccurrence = 5
	defaultLimit          = 5
	maxGoroutines         = 15 // Numero de GouRoutines
)

func SetupSearch(r *gin.Engine) {
	r.GET("/search", func(c *gin.Context) {

		query := c.Query("q")
		limitStr := c.Query("limit")

		if query == "nill" || query == "" || query == "null" {
			c.JSON(400, gin.H{"error": "Null or invalid query"})
			return
		}

		// Caminho para a pasta "Leaks"
		dirPath := "leaks"

		var results []gin.H
		occurrencesLimit := getLimit(limitStr)

		var wg sync.WaitGroup
		var mu sync.Mutex
		semaphore := make(chan struct{}, maxGoroutines) // Controlar número máximo de goroutines simultaneas

		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return nil
			}

			// Verifica se é um arquivo de texto
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".txt") {
				wg.Add(1)

				semaphore <- struct{}{}

				go func(filePath string) {
					defer wg.Done()
					defer func() { <-semaphore }()

					content, err := ioutil.ReadFile(filePath)
					if err != nil {
						fmt.Printf("Read file error %s: %v\n", filePath, err)
						return
					}

					// Divide o conteúdo em linhas
					lines := strings.Split(string(content), "\n")
					occurrences := 0

					for i, line := range lines {
						if strings.Contains(line, query) {
							// Define o início e o final para incluir as linhas ao redor da ocorrência
							start := i - linesAroundOccurrence
							if start < 0 {
								start = 0
							}
							end := i + linesAroundOccurrence + 1 // +1 para incluir a linha da ocorrência
							if end > len(lines) {
								end = len(lines)
							}

							// Adiciona o resultado à lista
							mu.Lock()
							results = append(results, gin.H{
								"path":     filePath,
								"content":  strings.Join(lines[start:end], "\n"),
								"position": fmt.Sprintf("Line %d", i+1), // Adiciona 1 para a linha de início ser 1-indexed
							})
							mu.Unlock()

							occurrences++

							// Verifica se atingiu o limite de ocorrencias
							if occurrences >= occurrencesLimit {
								return
							}
						}
					}
				}(path)
			}

			return nil
		})

		if err != nil {
			fmt.Printf("Unespected Error: %v\n", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		// Aguarde todas as goroutines concluírem
		wg.Wait()

		// Retorna os resultados da busca
		c.JSON(200, gin.H{"results": results})
	})
}

func getLimit(limitStr string) int {
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return defaultLimit
	}
	return limit
}
