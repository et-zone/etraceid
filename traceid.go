package etraceid

import (
	"fmt"
	"os"
	"time"
)

var sequence int64


func init(){
	sequence =1000
}
func getSequence()int64{
	if sequence >=4095{
		sequence =1000
	}else {
		sequence++
	}
	return sequence
}

type TraceID string

func MakeTraceID() TraceID {
	ip,_:= getLocalIP()
	s:= format(int64(ip.To4()[0]))+ format(int64(ip.To4()[1]))+ format(int64(ip.To4()[2]))+ format(int64(ip.To4()[3]))
	s+= format(time.Now().UnixNano())
	s+= format(int64(os.Getpid())<<2)
	s+= format(getSequence())
	return TraceID(s)
}

func (t *TraceID)ISValid()bool{
	return len(*t)==32
}
func (t *TraceID)GetIP()string{
	s:=(*string)(t)
	return fmt.Sprintf("%v.%v.%v.%v", reduction((*s)[:2]), reduction((*s)[2:4]), reduction((*s)[4:6]), reduction((*s)[6:8]))
}

func (t *TraceID)GetTime()string{
	s:=(*string)(t)
	return time.Unix(0, reduction((*s)[8:24])).Format("2006-01-02 15:04:05")
}