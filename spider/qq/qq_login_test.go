package qq

import (
	"go-demo/utils/env"
	"testing"
)

func TestQQLogin(t *testing.T) {
	if env.IsCI() {
		return
	}

	user, err := GetQQInfo(QQZone)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v \n", user)

	//user, err = GetQQInfo(QQFriend)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v \n", user)
}
