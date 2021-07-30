package main

import (
    "log"
    "net/http"
    "os"
    "fmt"
)

func main() {
    dir_to_serve, _ := os.Getwd()
    web_server_bind_addr := ""
    web_server_bind_port := 80

    fmt.Println("WebServerForFolder utility")
    fmt.Println("https://github.com/Stanislav-Povolotsky/WebServerForFolder")
    fmt.Println("")

    if (len(os.Args) == 2) && (os.Args[1] == "/?" || os.Args[1] == "--help") {
        fmt.Println("Usage:", os.Args[0], "[<port> [<folder>]]")
        return
    }

    if (len(os.Args) >= 2) {
        var port int
        if _, err := fmt.Sscanf(os.Args[1], "%d", &port); err == nil {
            web_server_bind_port = port
        }
    }
    if (len(os.Args) >= 3) {
        dir_to_serve = os.Args[2]
    }

    web_server_bind := fmt.Sprintf("%s:%d", web_server_bind_addr, web_server_bind_port)
    fmt.Println(fmt.Sprintf("Starting webserver at %s to serve directory %s", web_server_bind, dir_to_serve))
    fs := http.FileServer( http.Dir( dir_to_serve ) )
    log.Fatal(http.ListenAndServe( web_server_bind, fs ))
}
