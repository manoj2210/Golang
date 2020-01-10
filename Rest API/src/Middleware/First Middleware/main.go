package main

import(
  "net/http"
  "fmt"
)

func middleware(h http.Handler) http.Handler{
  return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){ 
    fmt.Fprintf(w,"Hi this from the middleware I'm executing you\n")
    h.ServeHTTP(w,r) //Once Transferred control to Server can't back
    fmt.Fprintf(w,"Don't worry Middleware got Executed\n") //This is not executed
    fmt.Println("OVER")
  })
}

func mainServer(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Hi this is from Server\n")
}
func main(){
  newHandler:=http.HandlerFunc(mainServer)
  http.Handle("/", middleware(newHandler))
  http.ListenAndServe(":8080",nil)
}