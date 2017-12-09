package redisBus

func init() {
	err := initConnectToServer(5)
	if err != nil {
		panic(err)
	}
}
