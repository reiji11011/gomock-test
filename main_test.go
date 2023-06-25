package main

import "testing"

type MockFreeUserRepository struct{}

func (r MockFreeUserRepository) Calculate(number int) int {
	// テストケースを満たすように期待する値を返却する
	return 2
}

// 新しいモックの追加
type MockFreeUserRepository2 struct{}

func (r MockFreeUserRepository2) Calculate(number int) int {
	// テストケースを満たすように期待する値を返却する
	return 3
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

// モックあり
// 複数パターンのテストを実施
func TestUserUsecase_ExecuteUserUsecase_Mock2(t *testing.T) {
	// IFを満たす構造体を使いテストを実施
	userRepository := MockFreeUserRepository{}
	userUsecase := UserUsecase{userRepository: userRepository}
	// テストケース1
	actual := userUsecase.ExecuteUserUsecase(1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	// テストケース2
	// テストケースに合わせてモックを追加
	userRepository2 := MockFreeUserRepository2{}
	userUsecase2 := UserUsecase{userRepository: userRepository2}
	actual = userUsecase2.ExecuteUserUsecase(2)
	expected = 3

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

}
