package main

import "text/template"

var tpl *template.Template
var g Gruopie

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	g.getArtists()
}

func main() {

}
