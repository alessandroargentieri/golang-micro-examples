package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"golang.org/x/oauth2"
)

var (
	conf     *oauth2.Config
	ctx      context.Context
	wg       sync.WaitGroup
	shutdown = make(chan struct{})
)

const successPage string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login Success</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            margin: 50px;
        }

        h1 {
            color: #007BFF;
        }

        .logo {
            font-size: 48px;
            color: #007BFF;
            margin-bottom: 20px;
        }

        p {
            color: #28A745;
        }
    </style>
</head>
<body>
    <div class="logo">
        <i class="fas fa-user"></i> <!-- Placeholder logo using Font Awesome -->
    </div>
    <h1>You logged in correctly.</h1>
    <p>You can go back to the CLI.</p>
</body>
</html>`

const errPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login Success</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            margin: 50px;
        }

        h1 {
            color: #007BFF;
        }

        .logo {
            font-size: 48px;
            color: #007BFF;
            margin-bottom: 20px;
        }

        p {
            color: #28A745;
        }
    </style>
</head>
<body>
    <div class="logo">
        <i class="fas fa-user"></i> <!-- Placeholder logo using Font Awesome -->
    </div>
    <h1>An error occurred during login.</h1>
    <p>Please retry or contact the support.</p>
</body>
</html>`

// struct representing the ~/.cli.json file
type CliConfigJsonFile struct {
	Apikeys interface{} `json:"apikeys,omitempty"`
	Bearer  struct {
		Token     string    `json:"token,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	} `json:"bearer,omitempty"`
	Meta struct {
		Admin               bool      `json:"admin,omitempty"`
		CurrentApiKey       string    `json:"current_apikey,omitempty"`
		DefaultRegion       string    `json:"default_region,omitempty"`
		LatestReleaseCheck  time.Time `json:"latest_release_check,omitempty"`
		Url                 string    `json:"url,omitempty"`
		LastCommandExecuted time.Time `json:"last_command_executed,omitempty"`
	} `json:"meta,omitempty"`
}

/*
This handler corresponds to the redirect_uri registered in the Oauth2 Client Application.
It normally involves into a Frontend URL which gets the code and sends to a Backend client application to exchange the token.
In this case, instead, by not having a Backend and not involving the browser, the callback is set directly to the function exchanging the
'code' (lasting 10 minutes) with the access token (JWT).
*/
func callbackHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/cli/oauth/redirect" || r.Method != http.MethodGet {
		close(shutdown) // Signal to shut down the server
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return
	}

	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect URI by Loukas (OIDC provider).
	code := queryParts["code"][0]
	log.Printf("code: %s\n", code)

	// The oauth2 golang library performs the call to /token to exchange the temporary code with the access token
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Println(err)
		close(shutdown) // Signal to shut down the server
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errPage, err)
		return
	}
	log.Printf("access token: %s\n", token.AccessToken)

	cliConfPath := os.Getenv("HOME") + "/.cli.json"

	// init new empty file content
	fileContent := CliConfigJsonFile{}
	// try to read the existing config file if present
	cliConfFile, err := os.ReadFile(cliConfPath)
	if err != nil {
		log.Println("~/.cli.json file not found, creating new one...")
	} else {
		// existing ~/.cli.json file correctly read, unmarshalling its content
		err = json.Unmarshal(cliConfFile, &fileContent)
		if err != nil {
			log.Printf("error while parsing the content of ~/.cli.json file: %s\n", err)
		}
	}

	// update the file content with the bearer token and the timestamp
	fileContent.Bearer.CreatedAt = time.Now().UTC()
	fileContent.Bearer.Token = token.AccessToken

	// marshal the file content into json
	newContent, err := json.Marshal(fileContent)
	if err != nil {
		log.Println(err)
		close(shutdown) // Signal to shut down the server
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errPage, err)
		return
	}

	// overwrite the old ~/.cli.json file (or create a new one if it doesn't exist) with the updated content
	if err := os.WriteFile(cliConfPath, newContent, 0777); err != nil {
		log.Println(err)
		close(shutdown) // Signal to shut down the server
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errPage, err)
		return
	}

	close(shutdown) // Signal to shut down the server

	// show succes page
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, successPage)
}

func main() {
	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     "your-app-registered-client-id",
		ClientSecret: "your-app-registered-client-secret",
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.wiremockapi.cloud/oauth/authorize",
			TokenURL: "https://oauth.wiremockapi.cloud/oauth/token",
		},
		// the callback (registered to the provider) must always be localhost
		RedirectURL: "http://127.0.0.1:9999/cli/oauth/redirect",
	}

	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Println(color.CyanString("You'll be redirected to your browser for the authentication..."))
	time.Sleep(1 * time.Second)
	exec.Command("open", url).Run()
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)

	// start local http server to work as OIDC client
	server := &http.Server{Addr: ":9999", Handler: http.HandlerFunc(callbackHandler)}

	// Use a goroutine to run the server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Error: %v\n", err)
		}
	}()

	// Set up a signal channel to capture signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("Received signal: %v\n", sig)
	case <-shutdown:
		//log.Println("Server shutdown initiated")
	}

	// Shut down the server gracefully
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Error during shutdown of the local server: %v\n", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("Login flow ended")

}
