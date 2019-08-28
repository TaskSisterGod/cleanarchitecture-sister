package external

import "log"

type Logger struct {

}

func (logger Logger) Log(arg ...interface{}) {
	log.Println(arg...)
}