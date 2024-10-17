package handlers

import (
	"net/http"
	"time"

	"github.com/delaneyj/datastar"
	"github.com/go-chi/chi/v5"

	"github.com/ngsalvo/datastar-test/components"
)

func SetupHome(router chi.Router) {

	homeRoute := func(w http.ResponseWriter, r *http.Request) {
		components.ClockPage().Render(r.Context(), w)
	}

	clockRoute := func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-r.Context().Done():
				ticker.Stop()
				return
			case <-ticker.C:
				datastar.RenderFragmentTempl(sse, components.ClockFragment(time.Now().Format("15:04:05")))
			}
		}
	}

	router.Get("/", homeRoute)
	router.Get("/clock", clockRoute)
}
