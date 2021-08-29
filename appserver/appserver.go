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
)

func RunServer() {

    http.HandleFunc("/get-quiz", get_quiz)

    //fs := http.FileServer(http.Dir("static/"))
    //http.Handle("/static/", http.StripPrefix("/static/", fs))

    var host_and_port string = "127.0.0.1:8080";
    fmt.Println("Listening on "+host_and_port)
    http.ListenAndServe(host_and_port, nil)
}

func get_quiz(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintln(w, "What number am I thinking right now?\n- 5\n- 8")
    content, err := ioutil.ReadFile("quiz.json")     // the file is inside the local directory
    if err != nil {
        log.Fatal("Error: cannot read the file quiz.json")
    }
    fmt.Fprintln(w, string(content))
}

