package models

import "net/http"

// Operation 处理逻辑
type Operation struct {
	Client *http.Client
	Opt    *Options
	OptionsInterface
	Sock string
	Url  string
}

// Options 入参选项
type Options struct {
	Get       bool
	Post      bool
	Put       bool
	Patch     bool
	Del       bool
	Body      string
	Pre       int64
	Delay     int64
	Test      int64
	Help      bool
	Print     bool
	DbgLog    bool
	UnixLocal bool
	JsonOut   bool
}

// OperationInterface 参数处理
type OperationInterface interface {
	ParseInputOptions()
	ParseInputArgs()
	CreateHttpClient()
}

// OptionsInterface 参数处理
type OptionsInterface interface {
	InitData(op *Operation)
	PreHandler()
	DelayHandler()
	RepeatSendHandler()
	DoHandler() (*http.Response, error)
	AfterHandler()
}

type Controller struct {
	Options
	OptionsInterface
}
