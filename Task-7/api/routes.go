package api

import (
	"log"
	"net/http"
	"time"

	"github.com/Tridipchavda/middleware/models"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// Api struct to hold DB
type api struct {
	db *gorm.DB
}

// Function to get API with DB
func NewAPI(db *gorm.DB) *api {
	return &api{db: db}
}

// JWT secret key
var secretKey = []byte("1524360798")

// Middleware function to Authenticate with Token
func (a *api) AuthenticateToken(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Token Cookie
		storedToken, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Parse the token to Verify its Valid or Not with Cliams
		tkn, err := jwt.ParseWithClaims(storedToken.Value, &jwt.StandardClaims{}, func(tkn *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		// Handling if the Token is invalid
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Move to The Main Route Handler
		next(w, r)
	})
}

// Function to get the JWT Token
func (a *api) Sign(w http.ResponseWriter, r *http.Request) {
	// Parse the form and get the value
	r.ParseForm()
	user := r.Form.Get("user")
	pass := r.Form.Get("password")

	log.Println(user, pass)
	// Store details in User model and check if user 's credentials are valid
	var u models.User
	a.db.Find(&u, "name = ?", user)
	if user != u.Name || pass != u.Pass {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// make Unsigned Token with Header , Claims Details
	unsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    user,
		ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
	})

	// Sign the Unsigned Token with secret key
	signed, err := unsigned.SignedString(secretKey)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the JWT Token in the Cookie
	http.SetCookie(w, &http.Cookie{Name: "token", Value: signed})
	w.Write([]byte(signed))
}

// Function which is Provided after Middleware
func (a *api) LogIn(w http.ResponseWriter, r *http.Request) {
	// Login Page
	w.Write([]byte("Login done "))
}
