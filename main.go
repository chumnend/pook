package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type appType struct {
	router *mux.Router
	db     *gorm.DB
	secret string
}

var app appType

// JSONResponse struct declaration
type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// Send passes the JSONRepsone to the ResponseWriter target
func (response *JSONResponse) Send(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

type token struct {
	ID    uint
	Email string
	*jwt.StandardClaims
}

// User struct declaration
type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"`
	FirstName string `gorm:"type:varchar(100)" json:"firstName"`
	LastName  string `gorm:"type:varchar(100)" json:"lastName"`
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-token, Authorization")

		next.ServeHTTP(w, r)
	})
}

func dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dump))
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping route hit")
	fmt.Fprintf(w, "Ready to serve requests")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Register route hit")

	// validate request body
	if r.Body == nil {
		response := JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		response.Send(w, 400)
		return
	}

	// read content of request body into User struct
	User := User{}
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		response := JSONResponse{
			Success: false,
			Message: "Request must be json object",
		}

		response.Send(w, 400)
		return
	}

	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		response := JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		response.Send(w, 400)
		return
	}

	// save User information to db
	User.Password = string(password)
	result := app.db.Create(&User)
	if result.Error != nil {
		response := JSONResponse{
			Success: false,
			Message: "Email already exists",
		}

		response.Send(w, 400)
		return
	}

	// create token
	tk := &token{
		ID:             User.ID,
		Email:          User.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(app.secret))
	if errJWT != nil {
		response := JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		response.Send(w, 400)
		return
	}

	response := JSONResponse{
		Success: true,
		Message: "New User created",
		Payload: map[string]interface{}{
			"id":    User.ID,
			"email": User.Email,
			"token": tokenString,
		},
	}

	response.Send(w, 200)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Login route hit")

	// validate request body
	if r.Body == nil {
		response := JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		response.Send(w, 400)
		return
	}

	// read content of request body into User struct
	credentials := User{}

	errjson := json.NewDecoder(r.Body).Decode(&credentials)
	if errjson != nil {
		response := JSONResponse{
			Success: false,
			Message: "Request must be json object",
		}

		response.Send(w, 400)
		return
	}

	// check if email exists
	foundUser := User{}
	errDB := app.db.Where("email = ?", credentials.Email).First(&foundUser).Error
	if errDB != nil {
		response := JSONResponse{
			Success: false,
			Message: "Invalid email and/or password",
		}

		response.Send(w, 400)
		return
	}

	// check if password is valid
	errPW := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password))
	if errPW != nil && errPW == bcrypt.ErrMismatchedHashAndPassword {
		response := JSONResponse{
			Success: false,
			Message: "Invalid email and/or password",
		}

		response.Send(w, 400)
		return
	}

	// create token
	tk := token{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(app.secret))
	if errJWT != nil {
		response := JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		response.Send(w, 400)
		return
	}

	response := JSONResponse{
		Success: true,
		Message: "Successful login",
		Payload: map[string]interface{}{
			"id":    foundUser.ID,
			"email": foundUser.Email,
			"token": tokenString,
		},
	}

	response.Send(w, 200)
}

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	secret := os.Getenv("SECRET_KEY")
	port := os.Getenv("PORT")

	// configure database
	app.db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	app.db.AutoMigrate(User{})

	// set secret
	if len(secret) < 0 {
		log.Fatal("Missing secret string")
	}
	app.secret = secret

	// setup routes
	app.router = mux.NewRouter().StrictSlash(true)
	app.router.Use(cors)
	app.router.HandleFunc("/", ping).Methods("GET")
	app.router.HandleFunc("/api/v1/register", registerUser).Methods("POST")
	app.router.HandleFunc("/api/v1/login", loginUser).Methods("POST")

	// start the server
	addr := ":" + port
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, app.router))
}
