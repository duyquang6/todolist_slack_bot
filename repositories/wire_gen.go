// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repositories

import (
	"todolist-facebook-chatbot/conf"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Injectors from wire.go:

func injectTaskRepository() (taskRepository, error) {
	appConfig, err := conf.NewAppConfig()
	if err != nil {
		return taskRepository{}, err
	}
	resource, err := NewResource(appConfig)
	if err != nil {
		return taskRepository{}, err
	}
	repositoriesTaskRepository := taskRepository{
		resource:  resource,
		appConfig: appConfig,
	}
	return repositoriesTaskRepository, nil
}
