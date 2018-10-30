package adaptor

import "testing"

func TestAdaptor(t *testing.T)  {
	player := PlayerAdaptor{}
	player.play("mp3","死了都要爱")
	player.play("wma","滴滴")
	player.play("mp4","复仇者联盟")
}
