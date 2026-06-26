package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mrdyuke/shorty/internal/domain"
)

type URLUseCase interface {
	ShortenURL(ctx context.Context, urlPack *domain.URLPack) (string, error)
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
	requestCtx := c.Request.Context()

	urlPack := new(domain.URLPack)
	if err := c.BindJSON(urlPack); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := uc.UseCase.ShortenURL(requestCtx, urlPack)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"shortURL": shortURL})
}
