package main

import "fmt"

type UserRepository interface {
	Calculate(number int) int
}

type UserUsecase struct {
	userRepository UserRepository
}

func main() {
	userRepository := FreeUserRepository{}
	userUsecase := UserUsecase{userRepository: userRepository}
	result := userUsecase.ExecuteUserUsecase(1)
	fmt.Println(result)
}

func (u UserUsecase) ExecuteUserUsecase(number int) int {
	return u.userRepository.Calculate(number)
}
