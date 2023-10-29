package routes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func SetupExport(r *gin.Engine) {
	r.GET("/export", func(c *gin.Context) {
		query := c.Query("q")
		limitStr := c.Query("limit")

		if query == "nill" || query == "" || query == "null" {
			c.JSON(400, gin.H{"error": "Null or invalid query"})
			return
		}

		// Caminho para a pasta "Leaks"
		dirPath := "leaks"

		var results []string
		occurrencesLimit := getLimit(limitStr)

		var wg sync.WaitGroup
		var mu sync.Mutex
		semaphore := make(chan struct{}, maxGoroutines)

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

					// Lê o conteúdo do arquivo
					content, err := ioutil.ReadFile(filePath)
					if err != nil {
						fmt.Printf("Read file error %s: %v\n", filePath, err)
						return
					}

					// Divide o conteúdo em linhas
					lines := strings.Split(string(content), "\n")
					occurrences := 0

					for _, line := range lines {
						if strings.Contains(line, query) {
							// Adiciona o resultado à lista
							mu.Lock()
							results = append(results, line)
							mu.Unlock()

							occurrences++
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
			fmt.Printf("Unexpected Error: %v\n", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		// Aguarde todas as goroutines concluírem
		wg.Wait()

		// Concatenate the lines and return the raw text
		rawText := strings.Join(results, "\n")
		c.String(200, rawText)
	})
}
