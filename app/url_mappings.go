package app

import (
	"github.com/zoharngo/golang-microservices/src/api/controllers/repositories"
	"github.com/zoharngo/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
