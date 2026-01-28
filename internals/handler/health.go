package handler

import (
	"net/http"

	"github.com/Ankush263/http-server/internals/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
