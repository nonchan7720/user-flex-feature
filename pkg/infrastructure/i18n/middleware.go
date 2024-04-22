package i18n

import "net/http"

func HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, AcceptLanguageMiddlewareFunc(w, r))
	})
}

func acceptLanguageMiddlewareFunc(r *http.Request) *http.Request {
	ctx := r.Context()
	acceptLanguage := r.Header.Get("Accept-Language")
	return r.WithContext(ToLanguage(ctx, acceptLanguage))
}

func AcceptLanguageMiddlewareFunc(_ http.ResponseWriter, r *http.Request) *http.Request {
	return acceptLanguageMiddlewareFunc(r)
}
