package user

import "context"

type userUsecase struct {
	userRepo UserRepository
}

// NewUserUsecase creates a userUsecase struct
func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (u *userUsecase) Fetch(ctx context.Context) ([]User, error) {
	users, err := u.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	return user, nil
}
