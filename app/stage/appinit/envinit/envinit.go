package envinit

import "template-app/pkg/envconfig"

func Init() {
	envconfig.InitEnvVars()
}
