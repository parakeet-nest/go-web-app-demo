// Package main : a simple web app
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/rand"
)

/*
GetBytesBody returns the body of an HTTP request as a []byte.
  - It takes a pointer to an http.Request as a parameter.
  - It returns a []byte.

func GetBytesBody(request *http.Request) []byte {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	return body
}
*/

func generateFunnyName() string {
	// Seed the random number generator
	//rand.Seed(time.Now().UnixNano())

	rand.Seed(uint64(time.Now().UnixMilli()))

	adjectives := []string{"Funny", "Silly", "Crazy", "Goofy", "Wacky", "Loony", "Zany"}
	nouns := []string{"Chicken", "Goose", "Giraffe", "Koala", "Lemur", "Quokka", "Axolotl"}

	// Get random indices for both arrays
	index1 := rand.Intn(len(adjectives))
	index2 := rand.Intn(len(nouns))

	// Combine the selected elements to form the name
	randomName := adjectives[index1] + " " + nouns[index2]

	return randomName
}

func main() {
	ctx := context.Background()

	var redisServer = os.Getenv("REDIS_URL")

	client := redis.NewClient(&redis.Options{
		Addr:     redisServer, // the name is defined in the compose.yml file
		Password: "",                  // no password set
		DB:       0,                   // use default DB
	})

	appName := generateFunnyName()

	var httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	var message = os.Getenv("MESSAGE")
	if message == "" {
		message = "this is a message"
	}

	log.Println("🚀 starting web server on port: " + httpPort)
	log.Println("Hello from Docker 🐳 📝 message: " + message)

	mux := http.NewServeMux()

	fileServerHtml := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServerHtml)

	mux.HandleFunc("/info", func(response http.ResponseWriter, request *http.Request) {
		// read the content of the file info.txt
		// and return it as a response
		content, err := os.ReadFile("./public/info.txt")
		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}
		response.WriteHeader(http.StatusOK)
		response.Write(content)
	})

	mux.HandleFunc("/variables", func(response http.ResponseWriter, request *http.Request) {

		variables := map[string]interface{}{
			"message": message,
			"appName": appName,
		}

		jsonString, err := json.Marshal(variables)

		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}

		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonString)
	})

	// Use Redis data
	mux.HandleFunc("/data", func(response http.ResponseWriter, request *http.Request) {
		//greetingsValue, dbErr := client.Get(ctx, "greetings").Result()
		
		greetingsValue := greetings(ctx, client)

		data := map[string]interface{}{
			"greetings": greetingsValue,
		}
		jsonString, err := json.Marshal(data)

		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonString)
	})


	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)
}

func greetings(ctx context.Context, rdb *redis.Client) string {
	
	value, err := rdb.Get(ctx, "greetings").Result()

	if err != nil {
		log.Println("😡", err)
		return ""
	}
	return value
}
