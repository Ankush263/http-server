package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ankush263/http-server/internals/utils"
)

type EchoRequest struct {
	Message string `json:"message"`
}

// This function will return the message, that has been passed into it.
func Echo(w http.ResponseWriter, r *http.Request) {
	var req EchoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.Message == "" {
		utils.Error(w, http.StatusBadRequest, "message is required")
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"echo": req.Message,
	})
}
