package etraceid

import (
	"context"
)

const key = "ectx_k"

func ToContext(ctx context.Context,t TraceID)context.Context{
	if string(t)==""{
		return context.WithValue(ctx,key,MakeTraceID())
	}
	return context.WithValue(ctx,key,t)
}

func ContexFormat(ctx context.Context)string{
	val:=ctx.Value(key)

	if val==nil{
		return ""
	}
	return string(val.(TraceID))
}
func ContextToTrace(ctx context.Context)TraceID{
	val:=ctx.Value(key)
	if val==nil{
		return ""
	}else {
		return val.(TraceID)
	}
}