package main

import "testing"

type MockFreeUserRepository struct{}

func (r MockFreeUserRepository) Calculate(number int) int {
	// テストケースを満たすように期待する値を返却する
	return 2
}

// モックなし
func TestUserUsecase_ExecuteUserUsecase(t *testing.T) {
	// FreeUserRepositoryに依存しているため、FreeUserRepositoryがないとテストができない。
	userRepository := FreeUserRepository{}
	userUsecase := UserUsecase{userRepository: userRepository}
	actual := userUsecase.ExecuteUserUsecase(1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

// モックあり
func TestUserUsecase_ExecuteUserUsecase_Mock(t *testing.T) {
	// IFを満たす構造体を使いテストを実施
	userRepository := MockFreeUserRepository{}
	userUsecase := UserUsecase{userRepository: userRepository}
	// 内部でモックのメソッドCalculateが呼ばれる。
	actual := userUsecase.ExecuteUserUsecase(1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
