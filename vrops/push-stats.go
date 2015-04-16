package main

import "fmt"
import "net/http"

func main () {

	resname := "Number of XtremIO Deployments"
	adapterkind := "Http Post"
	reskind := "Varrow"
	resdesc := "How many have we deployed"
	metname := "XtremIO|Bricks Deployed"
	alrmlev := "0"
	alrmmsg := "alarm message"
	epoch := "1429185278"
	custval := "75"
	body := resname+","+adapterkind+","+reskind+",,"+resdesc+",\n"+metname+","+alrmlev+","+alrmmsg+","+epoch+","+custval

	fmt.Println(body)

	vcurl := "https://clt-vrops02-e.labclt.local/HttpPostAdapter/OpenAPIServlet"

	client := &http.Client{}
	req := http.Post("https://clt-vrops02-e.labclt.local/HttpPostAdapter/OpenAPIServlet", body, nil)
	req.SetBasicAuth("admin","V@rrow1!")
	client.Do(req)

}

