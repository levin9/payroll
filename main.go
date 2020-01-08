package main

import (
	_ "payroll/config"
	"payroll/handlers"
	"payroll/tasks"

	"github.com/spf13/viper"
)

func main() {

	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
