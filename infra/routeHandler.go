package infra

import "net/http"

type CustomController interface {
	Register(mux *http.ServeMux)
}

func RegisterAllRoutes(mux *http.ServeMux) func(controllers ...CustomController) *http.ServeMux {
	return func(controllers ...CustomController) *http.ServeMux {
		for _, controller := range controllers {
			controller.Register(mux)
		}

		return mux
	}
}
