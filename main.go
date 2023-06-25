package main

import "fmt"

type UserRepository interface {
	Calculate(number int) int
}

type UserUsecase struct {
	userRepository UserRepository
}

func main() {
	var userRepository UserRepository
	userRepository = FreeUserRepository{}
	userUsecase := UserUsecase{userRepository: userRepository}
	result := userUsecase.ExecuteUserUsecase()
	fmt.Println(result)
}

func (u UserUsecase) ExecuteUserUsecase() int {
	return u.userRepository.Calculate(1)
}

// 別ファイル（user_repository.go）
type FreeUserRepository struct{}

func (r FreeUserRepository) Calculate(number int) int {
	// DBからデータを取得
	return number + 1
}
