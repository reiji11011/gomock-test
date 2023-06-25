package main

import "testing"

func Test_Calculate(t *testing.T) {

}

// Calculateを実際に呼び出してテストする
func TestFreeUser_Calculate(t *testing.T) {
	freeUser := FreeUserRepository{}
	actual := freeUser.Calculate(1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

// Calculateをモックに差し替えてテストする
// モックを用意
type MockUser struct{}

// モックの振る舞いを定義
func (m *MockUser) Calculate(number int) int {
	return 2
}

func TestFreeUser_Calculate_Mock(t *testing.T) {
	mockUser := MockUser{}
	actual := mockUser.Calculate(1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

}
