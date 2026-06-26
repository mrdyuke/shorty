package usecase

import (
	"context"
	"crypto/md5"

	"github.com/mrdyuke/shorty/internal/domain"
)

type URLRepo interface {
	SaveURL(ctx context.Context, urlPack *domain.URLPack) error
}

type URLUseCase struct {
	Repo URLRepo
}

func NewURLUseCase(repo URLRepo) *URLUseCase {
	return &URLUseCase{
		Repo: repo,
	}
}

func (uc *URLUseCase) ShortenURL(ctx context.Context, urlPack *domain.URLPack) (string, error) {
	hash := md5.New()

	_, err := hash.Write([]byte(urlPack.OriginalURL))
	if err != nil {
		return "", err
	}

	shortURL := string(hash.Sum(nil)[:6])
	urlPack.ShortURL = shortURL

	err = uc.Repo.SaveURL(ctx, urlPack)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
