package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
	"golang.org/x/crypto/bcrypt"
)

// Server struct
type Server struct {
	Addr   string
	Router *mux.Router
	DB     *gorm.DB
	Secret string
}

// NewServer creates a new server struct
func NewServer(connectionString string, secret string, port string) *Server {
	server := new(Server)

	// connect database
	var err error
	server.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	server.DB.AutoMigrate(User{})
	server.DB.AutoMigrate(Booking{})
	server.DB.AutoMigrate(Location{})

	// set secret
	if len(secret) < 0 {
		log.Fatal("Missing secret string")
	}
	server.Secret = secret

	// assign runtime address
	server.Addr = ":" + port

	// setup router
	server.Router = mux.NewRouter().StrictSlash(true)
	server.Router.Use(cors)
	server.Router.HandleFunc("/ping", server.pingHandler).Methods("GET")
	server.Router.HandleFunc("/api/v1/register", server.registerUserHandler).Methods("POST")
	server.Router.HandleFunc("/api/v1/login", server.loginUserHandler).Methods("POST")

	return server
}

// Run starts the server
func (s *Server) Run() {
	log.Println("Listening on address " + s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, s.Router))
}

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - ping")
	fmt.Fprintf(w, "Ready to serve requests")
}

func (s *Server) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - registerUser")

	// check if body exists
	if r.Body == nil {
		SendError(w, 400, "Empty request body")
		return
	}

	// read content into request body
	User := User{}
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		SendError(w, 400, "Request must be a JSON object")
		return
	}

	// validate the request body
	if len(User.FirstName) == 0 || len(User.LastName) == 0 {
		SendError(w, 400, "First name and/or last name not valid")
		return
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(User.Email) {
		SendError(w, 400, "Email not valid")
		return
	}

	if len(User.Password) < 6 {
		SendError(w, 400, "Password must contain at least 6 characters")
		return
	}

	// encrypt the password
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		SendError(w, 400, "Something went wrong. Please try again later")
		return
	}

	// save User information to db
	User.Password = string(password)
	result := s.DB.Create(&User)
	if result.Error != nil {
		SendError(w, 400, "Email already exists")
		return
	}

	// create token
	tk := &token{
		ID:             User.ID,
		Email:          User.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(s.Secret))
	if errJWT != nil {
		SendError(w, 400, "Something went wrong. Please try again later")
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

	SendJSON(w, 200, response)
}

func (s *Server) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - loginUser")

	// validate request body
	if r.Body == nil {
		SendError(w, 400, "Empty request body")
		return
	}

	// read content of request body into User struct
	credentials := User{}

	errJSON := json.NewDecoder(r.Body).Decode(&credentials)
	if errJSON != nil {
		SendError(w, 400, "Request must be a JSON object")
		return
	}

	// check if email exists
	foundUser := User{}
	errDB := s.DB.Where("email = ?", credentials.Email).First(&foundUser).Error
	if errDB != nil {
		SendError(w, 400, "Invalid email and/or password")
		return
	}

	// check if password is valid
	errPW := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password))
	if errPW != nil && errPW == bcrypt.ErrMismatchedHashAndPassword {
		SendError(w, 400, "Invalid email and/or password")
		return
	}

	// create token
	tk := token{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(s.Secret))
	if errJWT != nil {
		SendError(w, 400, "Something went wrong. Please try again later")
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

	SendJSON(w, 200, response)
}
