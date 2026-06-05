package controller

import "github.com/gin-gonic/gin"

type URLUseCase interface {
}

type URLController struct {
	UseCase URLUseCase
}

func NewURLController(usecase URLUseCase) *URLController {
	return &URLController{
		UseCase: usecase,
	}
}

func (uc *URLController) ShortenURL(c *gin.Context) {

}
