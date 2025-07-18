package main


import (
	"fmt"
	"log"
	 "net/http"

)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found", http.StatusNotFound)
		return 
	}
	if r.Method !="GET"{
		http.Error(w,"405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"ParseForm() err:%v",err)
		return

	}
	fmt.Fprintf(w,"post request successful\n")
	name :=r.FormValue("name")
	address :=r.FormValue("email")
	fmt.Fprintf(w,"Name: %s\n",name)
	fmt.Fprintf(w,"Email: %s\n",address)
}
func main(){
    fileServer:= http.FileServer(http.Dir("./static"))

	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Server is running on port 8080")
	if err:= http.ListenAndServe(":8000",nil); err!=nil{
		log.Fatal(err)
	}
	
}