package main

import (
	"bufio"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
)

var roleID = os.Getenv("ROLE_ID")
var secretID = os.Getenv("SECRET_ID")
var vaultAddr = os.Getenv("VAULT_ADDR")
var transitPath = "transit/encrypt/cckey"
var transformPath = "transform/encode/payments"
var appRolePath = "auth/approle/login"
var dataFilePath = "/mnt/data/numbers.data"

/// Login function to log into Vault via App Role
/// This function recieves a token from Vault and sets it to the client
func login(client *api.Client) {
	log.Printf("Logging into Vault with role_id: %s", roleID)

	auth, err := client.Logical().Write(appRolePath, map[string]interface{}{
		"role_id":   roleID,
		"secret_id": secretID,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Set token for future client requests
	client.SetToken(auth.Auth.ClientToken)
	log.Printf("Successfully logged into Vault.")
}

// Transform reads in a plaintext creditcard with a Vault client.
// It does a write to the Transform path where it gets back an encrypted Secret type.
// It writes the encrypted value to a file via the writeData function.
func transform(creditCard string, client *api.Client) {
	log.Printf("Writing to: %s", vaultAddr+transformPath)

	// write creditcard value to Vault transform path
	secret, err := client.Logical().Write(transformPath, map[string]interface{}{
		"value": creditCard,
	})

	if err != nil {
		log.Fatal(err)
	}

	// write encrypted data to file
	writeData(secret.Data["encoded_value"].(string))
}

// Transit reads in a plaintext creditcard with a Vault client.
// It does a write to the Transit path where it gets back an encrypted Secret type.
// It writes the encrypted value to a file via the writeData function.
func transit(creditCard string, client *api.Client) {
	log.Printf("Writing to: %s", vaultAddr+transitPath)

	// base64 encode creditcard value and write to Vault transit path
	secret, err := client.Logical().Write(transitPath, map[string]interface{}{
		"plaintext": base64.StdEncoding.EncodeToString([]byte(creditCard)),
	})

	if err != nil {
		log.Fatal(err)
	}

	// write encrypted data to file
	writeData(secret.Data["ciphertext"].(string))
}

/// WriteData reads in a string and writes it to a file.
func writeData(cc string) {
	// open data file
	file, err := os.OpenFile(dataFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// write encrypted credit card data to the file
	if _, err := file.WriteString(cc + " \n"); err != nil {
		log.Println(err)
	}
}

/// ReadData opens a file and scans each line, storing it in a string array.
/// It returns the string array for consumption.
func readData() []string {
	var lines []string

	// open data file
	file, err := os.OpenFile(dataFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read all lines of file into an array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

/// LandingPage takes in a ResponseWriter and an HTTP request r
/// This function creates a Vault client and authenticates
/// It parses each request and executes the HTML template
func landingPage(w http.ResponseWriter, r *http.Request, client *api.Client) {
	log.Printf("%s %s %s %s", r.Proto, r.Method, r.RemoteAddr, r.URL.Path)

	if r.Method == http.MethodPost {
		// login to Vault
		login(client)

		if r.FormValue("submit") == "transform" {
			transform(r.FormValue("cnumber"), client)
		} else if r.FormValue("submit") == "transit" {
			transit(r.FormValue("cnumber"), client)
		}

		log.Printf("Revoking token...")
		client.Auth().Token().RevokeSelf("")

	}

	// read in all entries from data file
	data := readData()

	// template out index file
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func main() {
	log.Printf("Starting server...")
	// declare Vault config
	config := &api.Config{
		Address: vaultAddr,
	}

	// declare new Vault client
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		landingPage(w, r, client)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
