package functions

import "net/http"

type httpTriggerFunc = func (w http.ResponseWriter, r *http.Request)

type AzureFunctionApp struct {
  HttpFunctions []httpTriggerFunc
}
