package hddel

import (
	"bytes"
	"net/http"

	"codehub.huawei.com/hemyzhao/lib/hemyutils/utils/stringutils"
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
}

func (h *Handler) PreHandler() {
	logrec.Debug(h.debug, "DELETE PreHandler Entry")
	return
}

func (h *Handler) DelayHandler() {
	logrec.Debug(h.debug, "DELETE DelayHandler Entry")
	return
}

func (h *Handler) RepeatSendHandler() {
	logrec.Debug(h.debug, "DELETE RepeatSendHandler Entry")
	return
}

func (h *Handler) DoHandler() (*http.Response, error) {
	logrec.Debug(h.debug, "DELETE DoHandler Entry")
	return h.SendDelete()
}

func (h *Handler) AfterHandler() {
	logrec.Debug(h.debug, "DELETE AfterHandler Entry")
	return
}

func (h *Handler) SendDelete() (*http.Response, error) {
	var body *bytes.Reader

	if len(h.body) != 0 {
		body = bytes.NewReader(stringutils.Str2bytes(h.body))
		if body == nil {
			logrec.Error(h.debug, "convert body failed")
		}
	} else {
		body = nil
	}

	req, err := http.NewRequest("DELETE", h.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	return h.client.Do(req)
}
