package main

type FreeUserRepository struct{}

func (r FreeUserRepository) Calculate(number int) int {
	// DBからデータを取得
	return number + 1
}
