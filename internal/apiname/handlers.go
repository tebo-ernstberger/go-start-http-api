package apiname

import (
	"math/rand"
	"net/http"

	"github.com/ernst01/common/pkg/response"
)

const rmax = 100000

func (s *Server) handleReadApiname() http.HandlerFunc {
	type APINameResponse struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &APINameResponse{
			ID:      rand.Intn(rmax),
			Message: "Well done Sparky!",
		}

		if resp.ID > (rmax / 2) {
			response.SendError(w, http.StatusInternalServerError, "Something went wrong")
		} else {
			response.SendSuccess(w, http.StatusOK, resp)
		}
	}
}