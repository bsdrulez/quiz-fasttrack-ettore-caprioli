/*
Copyright © 2021 Ettore Caprioli <ettore.caprioli@gmail>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package appclient

import (
	"fmt"
	"io/ioutil"
	"log"
        "strings"
	"net/http"
        "github.com/spf13/quiz-fasttrack-ettore-caprioli/game"
)

func RunClient() {
    host_and_port := "http://localhost:8080"

    // get the quiz
    quiz := http_get_call(host_and_port+"/get-quiz")
    // fmt.Println("Response:\n", quiz)

    // run the quiz
    answers := game.RunQuiz(quiz)

    // post the answers
    http_post_call(host_and_port+"/answer-quiz", answers)

}

func http_get_call(url string) string {

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal("Error getting response. ", err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response. ", err)
    }
    
    return string(body)
}

func http_post_call(url string, body string) {
    resp, err := http.Post(
        url, "application/json", 
        strings.NewReader(body))
    if err != nil {
    	print(err)
    }
    defer resp.Body.Close()
    
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    	print(err)
    }
    fmt.Println("POST RESPONSE: ", string(b))
}
