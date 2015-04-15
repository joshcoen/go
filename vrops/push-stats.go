package main



func main () {

	resname := "Number of XtremIO Deployments"
	adapterkind := "Http Post"
	reskind := "Varrow"
	resdesc := "How many have we deployed"
	metname := "XtremIO|Bricks Deployed"
	alrmlev := 0
	alrmmsg := "alarm message"
	epoc := 1
	custval := 75
	body := ("%s,%s,%s,,%s\n,%s,%s,%s,%s,%s", resname, adapterkind, reskind, resdesc, metname, alrmlev, alrmmsg, epoch, custval)
}
