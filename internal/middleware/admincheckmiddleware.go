package middleware

import "net/http"

type AdminCheckMiddleware struct {
}

func NewAdminCheckMiddleware() *AdminCheckMiddleware {
	return &AdminCheckMiddleware{}
}

func (m *AdminCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if user role is admin
		// Note: Requires JWT middleware to populate context with "userRole" claim
		userRole := r.Context().Value("userRole")
		if userRole == nil || userRole != "admin" {
			http.Error(w, "Prohibited: Admin Access Required", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
