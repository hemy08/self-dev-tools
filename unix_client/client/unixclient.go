package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "codehub.huawei.com/hemyzhao/lib/hemyutils/logsdk"

	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/logrec"

	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/manager"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/models"
)

var op = models.Operation{}

func StartUnixDemoClient() {
	op.ParseInputOptions()
	op.ParseInputArgs()
	op.CreateHttpClient()
	manager.InitializationOptionHandler()

	response, err := Handler(&op)
	if err != nil {
		logrec.Error(true, "response failed, %v", err)
		panic(err)
	}

	logrec.Debug(op.Opt.DbgLog, "Response Header: \n %+v", response.Header)
	logrec.Debug(op.Opt.DbgLog, "Response StatusCode: %d", response.StatusCode)

	resBody, err := ioutil.ReadAll(response.Body)
	defer func() {
		if response == nil {
			return
		}
		err := response.Body.Close()
		if err != nil {
			logrec.Warn(true, "close file fail: %v", err)
		}
	}()
	if err != nil {
		logrec.Error(true, "read resp error %v", err.Error())
		panic(err)
	}

	FormatPrintResponseString(resBody)
	FormatPrintResponseJson(resBody)
}

func Handler(op *models.Operation) (*http.Response, error) {
	logrec.Info(op.Opt.DbgLog, "Handler Op:%+v\n", *op.Opt)
	switch {
	case op.Opt.Get == true:
		logrec.Debug(op.Opt.DbgLog, "Handler get")
		return DoHandler(http.MethodGet, op)
	case op.Opt.Post == true:
		logrec.Debug(op.Opt.DbgLog, "Handler Post")
		return DoHandler(http.MethodPost, op)
	case op.Opt.Patch == true:
		logrec.Debug(op.Opt.DbgLog, "Handler SendPatch")
		return DoHandler(http.MethodPatch, op)
	case op.Opt.Put == true:
		logrec.Debug(op.Opt.DbgLog, "Handler Put")
		return DoHandler(http.MethodPut, op)
	case op.Opt.Del == true:
		logrec.Debug(op.Opt.DbgLog, "Handler Del")
		return DoHandler(http.MethodDelete, op)
	}

	return nil, errors.New("op not support")
}

func DoHandler(method string, op *models.Operation) (*http.Response, error) {
	logrec.Debug(op.Opt.DbgLog, "DoHandler %s op：%+v", method, op)

	hd, err := manager.GetOptionHandler(method)
	if err != nil {
		logrec.Error(op.Opt.DbgLog, "DoHandler get handler by %s failed, op: %+v", method, op)
		return nil, err
	}

	hd.InitData(op)
	hd.PreHandler()
	hd.RepeatSendHandler()
	hd.DelayHandler()
	resp, err := hd.DoHandler()
	hd.AfterHandler()
	return resp, err
}

func FormatPrintResponseJson(resBody []byte) {
	if op.Opt.JsonOut {
		var out bytes.Buffer
		err := json.Indent(&out, resBody, "", "\t")
		if err != nil {
			log.Error("json.Indent failed : ", err.Error())
			return
		}

		_, _ = out.WriteTo(os.Stdout)
	}
}

func FormatPrintResponseString(resBody []byte) {
	// 优先json输出
	if op.Opt.JsonOut {
		return
	}

	if op.Opt.Print {
		fmt.Println(string(resBody))
	}
}
