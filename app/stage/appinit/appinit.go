// Package appinit provides method to Init all stages of app
package appinit

import (
	"template-app/app/stage/appinit/dbconinit"
	"template-app/app/stage/appinit/dbmigrate"
	"template-app/app/stage/appinit/envinit"
	"template-app/app/stage/appinit/logoinit"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
