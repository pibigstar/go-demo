package error

import "log"

func Check(err error) {
	if err != nil {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		log.Println("hava a error:", err.Error())
		panic(err)
	}
}
