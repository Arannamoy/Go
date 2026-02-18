package main

import (
	"fmt"
	"net/http"
)


// home is the home page handler
func home(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"this is home page")
}

// about is the about page handler
func about(w http.ResponseWriter,r *http.Request){
	sum,a:=addValue(5,6)
	if (a!=10){
		fmt.Fprintln(w,"This is about page %d",sum)
	}else{
		fmt.Fprintln(w,a)
	}
	
}

func addValue(x int,y int) (int,int){
	var sum int;
	sum=x+y;
	return  sum,10;
}
func main(){
	http.HandleFunc("/",home)
	http.HandleFunc("/about",about)
	_=http.ListenAndServe("0.0.0.0:8000",nil)
}