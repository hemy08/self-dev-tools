package manager

import (
	"fmt"
	"net/http"

	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/hddel"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/hdget"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/hdpatch"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/hdpost"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/handler/hdput"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client/models"
)

var OpHandler = make(map[string]func() models.OptionsInterface, 0)

// InitializationOptionHandler 初始化
func InitializationOptionHandler() {
	OpHandler[http.MethodGet] = func() models.OptionsInterface { return hdget.NewHandler() }
	OpHandler[http.MethodPut] = func() models.OptionsInterface { return hdput.NewHandler() }
	OpHandler[http.MethodPatch] = func() models.OptionsInterface { return hdpatch.NewHandler() }
	OpHandler[http.MethodDelete] = func() models.OptionsInterface { return hddel.NewHandler() }
	OpHandler[http.MethodPost] = func() models.OptionsInterface { return hdpost.NewHandler() }

}

// GetOptionHandler Get Service Plugin
func GetOptionHandler(name string) (models.OptionsInterface, error) {
	var service models.OptionsInterface
	if f, ok := OpHandler[name]; ok {
		service = f()
		return service, nil
	}

	return nil, fmt.Errorf("the service [%s] do not exist", name)
}
