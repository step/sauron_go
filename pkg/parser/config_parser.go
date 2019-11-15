package parser

import (
	"fmt"
	"reflect"

	"github.com/step/saurontypes"
	"github.com/spf13/viper"
)

func GetTasks(assignmentName string)  []saurontypes.Task {
	ParseConfig(assignmentName, "config", "../../")

	query := fmt.Sprintf("assignments.%s.tasks", assignmentName)
	result := viper.Get(query)

	fmt.Printf("result %s\n",result)

	return []saurontypes.Task{}
}

func ParseConfig(assignmentName string, configName string, configPath string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	
	// query:= fmt.Sprintf("assignments.%s.tests.job_name", assignmentName)

	// fmt.Println("assignment test job", viper.GetString(query))
}