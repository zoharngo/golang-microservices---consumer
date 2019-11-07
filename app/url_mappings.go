package app

import (
	"github.com/zoharngo/golang-microservices---consumer/controllers/repositories"
	"github.com/zoharngo/golang-microservices---consumer/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}