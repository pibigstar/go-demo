package copy

import "testing"

func TestDeepCopy(t *testing.T) {
	type User struct {
		Name string
	}
	user1 := &User{Name: "pibigstar"}

	var user2 User
	err := DeepCopy(user1, &user2)
	if err != nil {
		t.Error(err)
	}
	t.Log(user2)

	var user3 User
	err = Copy(user1, &user3)
	if err != nil {
		t.Error(err)
	}
	t.Log(user3)
}
