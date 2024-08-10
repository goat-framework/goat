package server

import (
    "net/http"
    "path/filepath"
    "fmt"

    "github.com/goat-framework/lamb/core/template"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    viewPath := filepath.Join("views", "welcome.lamb.html")

    err := template.Compile(viewPath, "views/components")
    if err != nil {
        http.Error(w, "error compiling the view", http.StatusInternalServerError)
        fmt.Println(err.Error())
        return
    }

    w.Header().Set("Content-Type", "text/html")
    filePath := filepath.Join(".cache", "welcome.html")
    http.ServeFile(w, r , filePath)
}
