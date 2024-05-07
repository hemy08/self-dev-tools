package models

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/logrec"
)

// ParseInputOptions 解析选项
func (op *Operation) ParseInputOptions() {
	optGet := flag.Bool("get", false, "send a GET msg to socket server")
	optPost := flag.Bool("post", false, "send a POST msg to socket server")
	optPut := flag.Bool("put", false, "send a PUT msg to socket server")
	optPatch := flag.Bool("patch", false, "send a PATCH msg to socket server")
	optDelete := flag.Bool("del", false, "send a DELETE msg to socket server")
	optBody := flag.String("b", "", "msg body")
	optHelp := flag.Bool("h", false, "usage help")
	optPre := flag.Int64("pre", 0, "repeat send message before method times")
	optDelay := flag.Int64("delay", 0, "after pre proc, delay times second")
	optTest := flag.Int64("test", 0, "send multi msg to socket server")
	optPrint := flag.Bool("print", true, "output request header")
	optDebug := flag.Bool("debug", false, "output debug log")
	optLocal := flag.Bool("local", false, "use http://localhost/xxx")
	optOutJson := flag.Bool("json", false, "output response body by json format")

	flag.Parse()
	op.Opt = &Options{
		Get:       *optGet,
		Post:      *optPost,
		Put:       *optPut,
		Patch:     *optPatch,
		Del:       *optDelete,
		Body:      *optBody,
		Help:      *optHelp,
		Pre:       *optPre,
		Delay:     *optDelay,
		Test:      *optTest,
		Print:     *optPrint,
		DbgLog:    *optDebug,
		UnixLocal: *optLocal,
		JsonOut:   *optOutJson,
	}

	logrec.Info(*optDebug, "ParseInputOptions %+v\n", *op.Opt)
}

// ParseInputArgs 解析参数
func (op *Operation) ParseInputArgs() {
	if op.Opt.Help || len(flag.Args()) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "[-options][value]... /socket_file /url")
		flag.PrintDefaults()
		os.Exit(0)
	}

	logrec.Debug(op.Opt.DbgLog, "Unix HTTP client, args: %+v", flag.Args())
	op.Sock = flag.Args()[0]
	if op.Opt.UnixLocal {
		op.Url = "http://localhost" + flag.Args()[1]
	} else {
		op.Url = "http://unix" + flag.Args()[1]
	}
}

// CreateHttpClient 创建客户端
func (op *Operation) CreateHttpClient() {
	if len(op.Sock) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "socket_file must input")
		flag.PrintDefaults()
		os.Exit(0)
	}

	op.Client = &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", op.Sock)
			},
		},
	}

	logrec.Debug(op.Opt.DbgLog, "Socket file: %+v", op.Sock)
	logrec.Debug(op.Opt.DbgLog, "op.Client: %+v", op.Client)
	logrec.Debug(op.Opt.DbgLog, "op.Client Transport: %+v", op.Client.Transport)
}
