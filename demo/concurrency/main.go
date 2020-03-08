package main

import (
	"bufio"
	"fmt"
	"github.com/golanshy/golang-microservices/demo/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/demo/src/api/services"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	failed map[string]errors.APiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.APiError
}

func getRequests() []repositories.CreateRepoRequest {

	result := make([]repositories.CreateRepoRequest, 0)
	file, err := os.Open("/Users/golanshay/Workspace/golang/src/github.com/golanshy/golang-microservices/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name:        line,
			Description: "Description " + line,
		}
		result = append(result, request)
	}
	return result
}

func main() {
	requests := getRequests()
	fmt.Println(fmt.Sprintf("About to process %d requests ", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)

	var wg = sync.WaitGroup{}

	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}

	wg.Wait()
	close(input)

	// Now we can handle success and failed results
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}

}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request:request,
		Result: result,
		Error:  err,
	}

	<-buffer
}
