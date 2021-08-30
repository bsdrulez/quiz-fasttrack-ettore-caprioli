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
package rank


import (
    "container/list"
)

var scores *list.List

// insert a new element in the rank and
// return the percentage of placement
func InsertScore(score int) int {

    if scores == nil {
        scores = list.New()
    }

    percentage := 100
    if scores.Len() == 0 {
        scores.PushFront(score)
    } else {
        place:=0
        for e := scores.Front(); e != nil; e = e.Next() {
            if e.Value.(int) > score {
                scores.InsertBefore(score, e)
                break
            }
            place++
        }

        percentage = place*100 / scores.Len()

        if place == scores.Len() {
            scores.PushBack(score)
        }
    }

    return percentage
}

