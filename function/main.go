package main

import (
    "fmt"
    "github.com/venkyGeek/carbonquest/function/pkg/carbonlogic"
)


func HelloWorld(w http.ResponseWriter, r *http.Request) {
    msg := carbonlogic.Greet()
    fmt.Fprint(w, msg)
}
