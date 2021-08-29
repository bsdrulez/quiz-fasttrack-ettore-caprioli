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
package game

import (
    "encoding/json"
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "log"
)

type question_st struct {
    Name        string  `json:"name"`
    Text        string  `json:"text"`
    Answers   []string  `json:"answers"`
}

type quiz_st struct {
    Questions   []question_st     `json:"questions"`
}


func RunQuiz(quiz_str string) string {
    
    var quiz quiz_st
    quiz = quiz_str2structure(quiz_str)
    
    var ans = make(map[string]string)
    for i, q := range quiz.Questions {
        fmt.Println("=== Question ", i, " ===\n", q.Text)
        for j, a := range q.Answers {
            fmt.Println(j, ") ", a)
        }
        fmt.Print("Insert the number of the correct answer: ")
        choice := read_int()
        ans[q.Name] = q.Answers[choice]
    }


    return answer2str(ans)
/*
    enc := json.NewEncoder(os.Stdout)
    d := map[string]string{"apple": "baluba", "lettuce": "7"}
    enc.Encode(d)

    var try map[string]string
    json.Unmarshal([]byte("{\"ciao\":\"maio\",\"bau\":\"cane\"}"), &try)
    fmt.Println(try["ciao"])
    fmt.Println(try["bau"])
*/
}

func quiz_str2structure(quiz_str string) quiz_st {
    res := quiz_st{}
    json.Unmarshal([]byte(quiz_str), &res)
    return res
}

func answer2str(ans map[string]string) string {
    ans_json, _ := json.Marshal(ans)
    return string(ans_json)
}

func read_int() int {
    reader := bufio.NewReader(os.Stdin)
    str, _ := reader.ReadString('\n')
    str = strings.Replace(str,"\n","",-1)
    num, e := strconv.Atoi(str)
    if e != nil {
        log.Fatal("Error: cannot read number")
    }
    return num
}
