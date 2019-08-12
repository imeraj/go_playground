package middlewares

import (
	"net/http"

	"github.com/imeraj/go_playground/lenslocked/context"
	"github.com/imeraj/go_playground/lenslocked/controllers"
)

type Auth struct {
	sc *controllers.Sessions
}

func NewAuth(sc *controllers.Sessions) *Auth {
	return &Auth{
		sc: sc,
	}
}

func (mw *Auth) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := mw.sc.GetCookie(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		user, err := mw.sc.GetSessionService().ByRemember(cookie["remember_token"])
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)

		next(w, r)
	})
}

func (mw *Auth) Apply(next http.Handler) http.Handler {
	return mw.ApplyFn(next.ServeHTTP)
}
