package usecase

type URLRepo interface {
}

type URLUseCase struct {
	Repo URLRepo
}

func NewURLUseCase(repo URLRepo) *URLUseCase {
	return &URLUseCase{
		Repo: repo,
	}
}
