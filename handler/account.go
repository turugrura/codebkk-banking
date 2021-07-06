package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/turugrura/codebkk-banking/errs"
	"github.com/turugrura/codebkk-banking/logs"
	"github.com/turugrura/codebkk-banking/service"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{accService: accService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logs.Error(err)
		handleError(w, errs.NewValidationError("request body incorect format"))
		return
	}

	customerId, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	accRes, err := h.accService.NewAccount(customerId, request)
	if err != nil {
		logs.Error(err)
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accRes)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	accs, err := h.accService.GetAccounts(customerId)
	if err != nil {
		logs.Error(err)
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accs)
}
