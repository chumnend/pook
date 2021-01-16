package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
	"golang.org/x/crypto/bcrypt"
)

// App struct declaration
type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Secret string
}

// Init setups up the application
func (app *App) Init(connectionString string, secret string) {
	// connect database
	var err error
	app.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	app.DB.AutoMigrate(User{})

	// set secret
	app.Secret = secret

	// setup routes
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.Use(cors)
	app.Router.HandleFunc("/", app.pingServer).Methods("GET")
	app.Router.HandleFunc("/api/v1/register", app.registerUser).Methods("POST")
	app.Router.HandleFunc("/api/v1/login", app.loginUser).Methods("POST")
}

// Run starts tthe web server
func (app *App) Run(addr string) {
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// Handlers =====================================
func (app *App) pingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}

func (app *App) registerUser(w http.ResponseWriter, r *http.Request) {
	// validate request body
	if r.Body == nil {
		response := JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		response.Send(w, 400)
		return
	}

	// read content of request body into user struct
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := JSONResponse{
			Success: false,
			Message: "Request must be JSON object",
		}

		response.Send(w, 400)
		return
	}

	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response := JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		response.Send(w, 400)
		return
	}

	// save user information to db
	user.Password = string(password)
	result := app.DB.Create(&user)
	if result.Error != nil {
		response := JSONResponse{
			Success: false,
			Message: "Email already exists",
		}

		response.Send(w, 400)
		return
	}

	// create token
	tk := &Token{
		ID:             user.ID,
		Email:          user.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(app.Secret))
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
		Message: "New user created",
		Payload: map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	}

	response.Send(w, 200)
}

func (app *App) loginUser(w http.ResponseWriter, r *http.Request) {
	// validate request body
	if r.Body == nil {
		response := JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		response.Send(w, 400)
		return
	}

	// read content of request body into user struct
	credentials := User{}

	errJSON := json.NewDecoder(r.Body).Decode(&credentials)
	if errJSON != nil {
		response := JSONResponse{
			Success: false,
			Message: "Request must be JSON object",
		}

		response.Send(w, 400)
		return
	}

	// check if email exists
	foundUser := User{}
	errDB := app.DB.Where("email = ?", credentials.Email).First(&foundUser).Error
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
	tk := Token{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(app.Secret))
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
