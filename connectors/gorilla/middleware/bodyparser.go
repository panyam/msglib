package middleware

import (
	"github.com/panyam/relay/connectors/gorilla/common"
	"net/http"
)

/**
 * Parses the request initially so subsequent middleware dont have to repeat
 * this.
 */
func BodyParserMiddleware(rw http.ResponseWriter, request *http.Request, context common.IRequestContext) error {
	request.ParseForm()
	return nil
}
