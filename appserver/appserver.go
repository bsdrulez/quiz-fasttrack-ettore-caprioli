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
package appserver

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "strconv"
    "github.com/spf13/quiz-fasttrack-ettore-caprioli/rank"
)

func RunServer() {

    http.HandleFunc("/get-quiz", get_quiz)
    http.HandleFunc("/answer-quiz", answer_quiz)

    //fs := http.FileServer(http.Dir("static/"))
    //http.Handle("/static/", http.StripPrefix("/static/", fs))

    var host_and_port string = "127.0.0.1:8080";
    fmt.Println("Listening on "+host_and_port)
    http.ListenAndServe(host_and_port, nil)
}

func get_quiz(w http.ResponseWriter, r *http.Request) {

    if r.Method != "GET" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    
    quiz := read_all_file_as_string("quiz.json")
    fmt.Fprintln(w, quiz)
}

func answer_quiz(w http.ResponseWriter, r *http.Request) {


    if r.Method != "POST" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    
    var ans = make(map[string]string)
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&ans)
    if err != nil {
        log.Fatal("Error: cannot decode json")
    }

    result := check_answers(ans)
    
    fmt.Fprintln(w, result)
}

func check_answers(ans map[string]string) string {
    correct_ans_str := read_all_file_as_string("answers.json")
    var correct_ans = make(map[string]string)
    json.Unmarshal([]byte(correct_ans_str), &correct_ans)

    var corr int = 0
    var tot int = 0
    for key, val := range correct_ans {
        if val == ans[key] {
            corr++
        }
        tot++
    }
    
    var result string
    result = "Your score is "+strconv.Itoa(corr)+"/"+strconv.Itoa(tot)+"\n"
    
    
    percentage := rank.InsertScore(corr)
    result = result + "You scored higher than "+
                strconv.Itoa(percentage)+"% of all quizzers"

    return result
}

func read_all_file_as_string(filename string) string {
    content, err := ioutil.ReadFile(filename)     // the file is inside the local directory
    if err != nil {
        log.Fatal("Error: cannot read the file quiz.json")
    }
    return string(content)
}
