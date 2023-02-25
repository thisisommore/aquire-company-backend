package main

import (
	"template-app/app/stage/appinit"
	"template-app/app/stage/apprun"
)

func main() {
	appinit.Init()
	apprun.Run()
}
