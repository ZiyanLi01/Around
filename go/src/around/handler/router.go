package handler

import (
	"context"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type contextKey string

const userIDKey contextKey = "UserId"

// InitRouter initializes the routes and middleware
func InitRouter() http.Handler {
	// JWT Middleware for authentication
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	router := mux.NewRouter()

	// Apply JWT middleware and WithUserID middleware to routes that require authentication
	router.Handle("/upload", jwtMiddleware.Handler(WithUserID(http.HandlerFunc(uploadHandler)))).Methods("POST")
	router.Handle("/search", jwtMiddleware.Handler(WithUserID(http.HandlerFunc(searchHandler)))).Methods("GET")
	router.Handle("/post/{id}", jwtMiddleware.Handler(WithUserID(http.HandlerFunc(deleteHandler)))).Methods("DELETE")
	router.Handle("/api/generate-image", jwtMiddleware.Handler(WithUserID(http.HandlerFunc(GenerateImageHandler)))).Methods("POST")
	router.Handle("/download-image", jwtMiddleware.Handler(WithUserID(http.HandlerFunc(DownloadImageHandler)))).Methods("GET")

	// Routes without authentication
	router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
	router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")

	// CORS configuration to allow requests from any origin
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}

// WithUserID middleware extracts the user ID from the JWT token and stores it in the request context
func WithUserID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := r.Context().Value("user")
        if user != nil {
            if token, ok := user.(*jwt.Token); ok {
                claims := token.Claims.(jwt.MapClaims)

                // Debugging logs
                log.Printf("Claims: %v\n", claims)

                userID, ok := claims["username"].(string) // Extract "username" from claims

                // More debugging logs
                log.Printf("UserID: %v, OK: %v\n", userID, ok)

                if ok && userID != "" {
                    ctx := context.WithValue(r.Context(), userIDKey, userID)
                    log.Printf("Setting UserID in context: %v\n", userID)
                    next.ServeHTTP(w, r.WithContext(ctx))
                    return
                }
            }
        }
        // If no valid UserID found, return Unauthorized error
        log.Println("Unauthorized - no valid UserID found")
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    })
}


// GetUserIDFromContext retrieves the user ID from the context
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	log.Printf("GetUserIDFromContext: %v, OK: %v\n", userID, ok)
	return userID, ok
}
