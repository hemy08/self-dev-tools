package hdput

import (
	"net/http"
	"strings"
	"time"

	log "codehub.huawei.com/hemyzhao/lib/hemyutils/logsdk"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/logrec"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/models"
)

type Handler struct {
	debug      bool
	client     *http.Client
	preRepeat  int64
	sleepTime  int64
	testRepeat int64
	body       string
	socket     string
	url        string
}

func NewHandler() models.OptionsInterface {
	handler := &Handler{}
	return handler
}

func (h *Handler) InitData(op *models.Operation) {
	h.preRepeat = op.Opt.Pre
	h.sleepTime = op.Opt.Delay
	h.testRepeat = op.Opt.Test
	h.body = op.Opt.Body
	h.socket = op.Sock
	h.url = op.Url
	h.debug = op.Opt.DbgLog
	h.client = op.Client
}

// PreHandler 预处理消息数
func (h *Handler) PreHandler() {
	logrec.Debug(h.debug, "PUT PreHandler Entry")
	if len(h.body) == 0 || h.preRepeat == 0 {
		return
	}

	logrec.Info(h.debug, "POST PreHandler repeat time:", h.preRepeat)
	for h.preRepeat > 0 {
		h.preRepeat--
		go func() {
			_, _ = h.SendPut()
		}()
	}
}

func (h *Handler) DelayHandler() {
	logrec.Debug(h.debug, "PUT DelayHandler Entry")
	if h.sleepTime != 0 {
		logrec.Info(h.debug, "POST DelayHandler sleep time:", h.sleepTime)
		time.Sleep(time.Duration(h.sleepTime) * time.Second)
	}
}

func (h *Handler) RepeatSendHandler() {
	logrec.Debug(h.debug, "PUT RepeatSendHandler Entry")
	if len(h.body) == 0 || h.testRepeat == 0 {
		return
	}

	logrec.Info(h.debug, "POST Handler body:", h.body, ", repeat times :", h.testRepeat)
	for h.testRepeat > 0 {
		h.testRepeat--
		go func() {
			_, _ = h.SendPut()
		}()
	}
}

func (h *Handler) DoHandler() (*http.Response, error) {
	logrec.Debug(h.debug, "PUT DoHandler Entry")
	if len(h.body) == 0 {
		return nil, nil
	}

	logrec.Info(h.debug, "POST Handler body:", h.body)
	return h.SendPut()
}

func (h *Handler) AfterHandler() {
	logrec.Debug(h.debug, "PUT AfterHandler Entry")
}

func (h *Handler) SendPut() (*http.Response, error) {
	var body *strings.Reader

	if len(h.body) != 0 {
		body = strings.NewReader(h.body)
		if body == nil {
			log.Error("convert body failed")
		}
	} else {
		body = nil
	}

	req, err := http.NewRequest("PUT", h.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	return h.client.Do(req)
}
