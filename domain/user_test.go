package domain

import (
	"testing"
)

var testUserA = User{
	Name: "A",
	Id:   "1",
}

var anotherTestUserA = User{
	Name: "A",
	Id:   "1",
}

func TestUserAlike(t *testing.T) {
	tests := []struct {
		name     string
		a        *User
		b        *User
		expected bool
	}{
		{
			name:     "参照先が別だがidが同じものを正しく処理できる",
			a:        &testUserA,
			b:        &anotherTestUserA,
			expected: true,
		},
		{
			name:     "参照先も同じ場合",
			a:        &testUserA,
			b:        &testUserA,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Alike(*tt.b)
			if got != tt.expected {
				t.Errorf("comparing %+v and %+v; expected %v, but got %v", tt.a, tt.b, tt.expected, got)
			}
		})
	}
}
