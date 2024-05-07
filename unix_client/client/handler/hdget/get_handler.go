package hdget

import (
	"errors"
	"net/http"
	"strings"

	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/logrec"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/models"
)

type Handler struct {
	debug  bool
	client *http.Client
	socket string
	url    string
	body   string
}

func NewHandler() models.OptionsInterface {
	handler := &Handler{}
	return handler
}

func (h *Handler) InitData(op *models.Operation) {
	h.socket = op.Sock
	h.url = op.Url
	h.debug = op.Opt.DbgLog
	h.client = op.Client

	h.body = op.Opt.Body
	logrec.Debug(h.debug, "GET InitData &+v", *h)
}

func (h *Handler) PreHandler() {
	logrec.Debug(h.debug, "GET PreHandler Entry")
	return
}

func (h *Handler) DelayHandler() {
	logrec.Debug(h.debug, "GET DelayHandler Entry")
	return
}

func (h *Handler) RepeatSendHandler() {
	logrec.Debug(h.debug, "GET RepeatSendHandler Entry")
	return
}

func (h *Handler) DoHandler() (*http.Response, error) {
	logrec.Debug(h.debug, "GET DoHandler Entry")
	logrec.Debug(h.debug, "Url is %s", h.url)

	var req *http.Request
	var err error
	if len(h.body) != 0 {
		body := strings.NewReader(h.body)
		if body == nil {
			logrec.Warn(h.debug, "convert body failed")
			return nil, errors.New("convert body failed")
		}
		req, err = http.NewRequest("GET", h.url, body)
	} else {
		req, err = http.NewRequest("GET", h.url, nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-domain-name", "default")

	return h.client.Do(req)
}

func (h *Handler) AfterHandler() {
	logrec.Debug(h.debug, "GET AfterHandler Entry")
	return
}
