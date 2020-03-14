package helper

import (
	"os"
	"strconv"
)

func Env(key string, value string) (data string){
	data = os.Getenv(key)
	if data == "" {
		data = value
	}

	return data
}

func EnvNumber(key string, value int) int  {
	var (
		err error
		data int
		)
	tStr := os.Getenv(key)
	if data, err= strconv.Atoi(tStr);err!=nil{
		data = value
	}

	return data
}
