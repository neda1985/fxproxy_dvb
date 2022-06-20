package apis

import (
	"fmt"
	"fxproxy/pkg/logger"
	"fxproxy/pkg/validator"
	"io"
	"net/http"
	"net/url"
	"os"
)

func CoreHandler(rw http.ResponseWriter, request *http.Request) {
	if !validator.ValidatePath(request.URL.Path) {
		returnNotFound(rw)
		return
	}
	u := url.URL{
		Scheme: os.Getenv("SCHEMA"),
		Host:   os.Getenv("DOWNSTREAM"),
		Path:   request.URL.Path,
	}
	req, err := http.NewRequest(request.Method, u.String(), request.Body)
	if err != nil {
		logger.LogError(err)
		returnInternalServerError(rw)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogError(err)
		returnInternalServerError(rw)
		return
	}
	defer resp.Body.Close()
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}
func returnNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "not found")
}
func returnInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "internal server error")
}
