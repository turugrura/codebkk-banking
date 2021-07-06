package handler

import (
	"fmt"
	"net/http"

	"github.com/turugrura/codebkk-banking/errs"
)

func handleError(w http.ResponseWriter, err error) {
	switch v := err.(type) {
	case errs.AppError:
		w.WriteHeader(v.Code)
		fmt.Fprintln(w, v.Message)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, v)
	}
}
