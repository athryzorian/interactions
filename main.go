package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"

	"github.com/cenkalti/backoff/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"github.com/athryzorian/interactions/dal/operations"
)

const (
	oauthGoogleUrlAPI = "https://accounts.google.com/o/oauth2/v2/auth"
	clientID          = "171559914291-q2rsnsoeae3u90ptc7hn3pe6hud7aru2.apps.googleusercontent.com"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     clientID,
	ClientSecret: "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",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

type Message struct {
	Value string `json:"value"`
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return rootHandler(db, c)
	})

	e.GET("/auth/google/login", func(c echo.Context) error {
		return oauthGoogleLogin(c.Response().Writer, c.Request())
	})

	e.POST("/auth/signin", func(c echo.Context) error {
		return signInwithHandler(c)
	})

	e.POST("/auth/signup", func(c echo.Context) error {
		return signUpWithHandler(c)
	})

	e.POST("/auth/use", func(c echo.Context) error {
		return continueWithHandler(c)
	})

	e.GET("/auth/google/callback", func(c echo.Context) error {
		return oauthGoogleCallback(c.Response().Writer, c.Request())
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/countries", func(c echo.Context) error {
		return countriesHandler(db, c)
	})

	e.GET("/states", func(c echo.Context) error {
		return statesHandler(db, c)
	})

	e.GET("/cities", func(c echo.Context) error {
		return citiesHandler(db, c)
	})

	e.POST("/send", func(c echo.Context) error {
		return sendHandler(db, c)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func signInwithHandler(c echo.Context) error {

	fmt.Println("Sign In With Handler")
	var request = c.Request()
	request.ParseForm()
	var loginUserCredential = request.FormValue("credential")
	fmt.Println("Login User Credential:", loginUserCredential)

	// Validate the JWT token and get the user info.
	// https://developers.google.com/identity/gsi/web/guides/verify-google-id-token
	// https://pkg.go.dev/github.com/googleapis/google-api-go-client/idtoken#Validate

	idTokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		log.Fatalf("Error creating ID token validator: %s", err)
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	payload, err := idTokenValidator.Validate(context.Background(), loginUserCredential, clientID)
	if err != nil {
		log.Fatalf("Error validating ID token: %s", err)
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	issuer := payload.Issuer
	audience := payload.Audience
	expires := payload.Expires
	issuedAt := payload.IssuedAt
	userID := payload.Subject

	fmt.Printf("Issuer: %s\n", issuer)
	fmt.Printf("Audience: %s\n", audience)
	fmt.Printf("Expires At: %d\n", expires)
	fmt.Printf("Issued At: %d\n", issuedAt)
	fmt.Printf("User ID: %s\n", userID)

	// email := payload.Claims["email"]
	// emailVerified := payload.Claims["email_verified"]
	// issuer := payload.Claims["Issuer"]

	// fmt.Printf("User ID: %s\n", userID)
	// fmt.Printf("Email: %s\n", email)
	// fmt.Printf("Email Verified: %v\n", emailVerified)
	// fmt.Printf("Issuer: %v\n", issuer)

	// You can use this user information to create a session or perform other actions as needed.

	// For demonstration purposes, we just return a success message.

	return c.HTML(http.StatusOK, `<html><body>Backend -> Sign In with handler called after successful Google Login.</body></html>`)

}

func signUpWithHandler(c echo.Context) error {

	fmt.Println("Sign Up With Handler")
	return c.HTML(http.StatusOK, `<html><body>Backend -> Sign Up with handler called after successful Google Sign up.</body></html>`)

}

func continueWithHandler(c echo.Context) error {

	fmt.Println("Sign Up With Handler")
	return c.HTML(http.StatusOK, `<html><body>Backend -> Continue with handler called after successful Google login.</body></html>`)

}

func oauthGoogleLogin(w http.ResponseWriter, r *http.Request) error {

	fmt.Println("Google Login")

	oauthState := generateStateOauthCookie(w)
	u := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)

	return nil
}

func generateStateOauthCookie(w http.ResponseWriter) string {

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) error {

	fmt.Println("Auth Google Callback")

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)

	return nil
}

func getUserDataFromGoogle(code string) ([]byte, error) {

	// Use code to get token and get user info from Google.
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	//saveUser(contents)
	//saveToken(contents, token)
	return contents, nil
}

func initStore() (*sql.DB, error) {

	fmt.Println("Connecting...")

	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "myuser", "mypassword", "mydb")

	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", postgresInfo)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected:")

	/*
		_, err = db.Exec(`DROP TABLE IF EXISTS COMPANY;`)
		if err != nil {
			log.Fatalf("Error dropping table COMPANY: %s", err)
			return nil, err
		}
	*/

	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS COMPANY (ID SERIAL PRIMARY KEY, value text);`)
	// if err != nil {
	// 	log.Fatalf("Error creating table COMPANY: %s", err)
	// 	return nil, err
	// }

	// fmt.Println("Table COMPANY is created.")

	return db, nil
}

func rootHandler(db *sql.DB, c echo.Context) error {

	r, err := countRecords(db)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! (%d)\n", r))
}

func countriesHandler(db *sql.DB, c echo.Context) error {

	countries, err := operations.ListCountries(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	noOfCountries := len(countries)
	log.Println("Number of countries:", noOfCountries)

	return c.JSON(http.StatusOK, countries)

}

func statesHandler(db *sql.DB, c echo.Context) error {

	whereClause := fmt.Sprintf("parent_country='%s'", c.QueryParam("country"))
	log.Println("Where Clause:", whereClause)

	states, err := operations.ListStates(db, whereClause)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	noOfStates := len(states)
	log.Println("Number of states:", noOfStates)

	return c.JSON(http.StatusOK, states)

}

func citiesHandler(db *sql.DB, c echo.Context) error {

	whereClause := fmt.Sprintf("parent_state='%s'", c.QueryParam("state"))
	log.Println("Where Clause:", whereClause)

	cities, err := operations.ListCities(db, whereClause)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	noOfCities := len(cities)
	log.Println("Number of cities:", noOfCities)

	return c.JSON(http.StatusOK, cities)

}

func sendHandler(db *sql.DB, c echo.Context) error {

	m := &Message{}

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	tx, err := db.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Defer a rollback in case of an error
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	//stmt, err := tx.Prepare("INSERT INTO COMPANY (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value")
	stmt, err := tx.Prepare("INSERT INTO COMPANY (value) VALUES ($1)")
	if err != nil {
		log.Fatalf("Error preparing insert statement: %s", err)
		tx.Rollback()
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Value)
	if err != nil {
		log.Fatalf("Error executing insert statement : %s", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = tx.Commit()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	/*
		err := db.ExecuteTx(context.Background(), db, nil,
			func(tx *sql.Tx) error {
				_, err := tx.Exec(
					"INSERT INTO message (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value",
					m.Value,
				)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}
				return nil
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	*/

	return c.JSON(http.StatusOK, m)
}

func countRecords(db *sql.DB) (int, error) {

	rows, err := db.Query("SELECT COUNT(*) FROM COMPANY")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}

	return count, nil
}
