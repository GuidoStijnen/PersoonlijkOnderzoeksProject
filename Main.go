package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	OS   string `json:"os"`
	IPv4 string `json:"ipv4"`
}

var (
	tpl    *template.Template // Template-instantie voor het renderen van HTML
	config Config             // Configuratieobject voor databaseverbinding
)

func init() {
	var err error
	tpl, err = template.ParseGlob("*.html") // Parseer alle HTML-templates in de huidige map
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadConfig()

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true",
		config.Username, config.Password, config.Hostname, config.Port, config.Database)

	db, err := sql.Open("mysql", connString) // Open verbinding met de MySQL-database
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping() // Ping database voor confermatie dat er een verbidning is
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the MySQL server")

	//Met deze code worden de juiste functies opgeroepen op basis van het aangevraagde URL-pad, waardoor de juiste respons wordt gegeven aan de client.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRoot(w, r, db, connString)
	})
	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		handleDevices(w, r, db, connString)
	})
	http.HandleFunc("/devices/delete", func(w http.ResponseWriter, r *http.Request) {
		handleDeleteDevice(w, r, db, connString)
	})
	http.HandleFunc("/devices/add", func(w http.ResponseWriter, r *http.Request) {
		handleCreateDevice(w, r, db, connString)
	})

	serverAddress := fmt.Sprintf(":%d", config.Port)
	log.Printf("Server is running on http://localhost%s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

// De functie laadt de inhoud van een YAML-configuratiebestand, decodeert het YAML-formaat en slaat de resulterende configuratie op in de config-variabele.
func loadConfig() {
	configFile := "config.yaml"

	content, err := os.ReadFile(configFile) // lees de configuration file
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	err = yaml.Unmarshal(content, &config) // Deserialize de configuration file
	if err != nil {
		log.Fatalf("Failed to parse configuration file: %v", err)
	}
}

// De functie neemt een templatebestand en gegevens als invoer, rendert het template met de gegeven gegevens en schrijft de gegenereerde HTML naar de http.ResponseWriter
func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	err := tpl.ExecuteTemplate(w, templateFile, data) // Render template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Deze functie controleert het pad van het verzoek en handelt het verzoek af op basis van de methode van het verzoek.
// Als het een POST-methode is wordt het verzoek doorgestuurd naar handleCreateDevice
// anders wordt de startpagina gerenderd "index.html"
func handleRoot(w http.ResponseWriter, r *http.Request, db *sql.DB, connString string) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		handleCreateDevice(w, r, db, connString)
	} else {
		renderTemplate(w, "index.html", nil)
	}
}

// Deze functie haalt alle apparaten op uit de database
// slaat ze op in een lijst van Device objecten en renderen de "devices.html" template met de lijst van apparaten om deze aan de gebruiker weer te geven.
func handleDevices(w http.ResponseWriter, r *http.Request, db *sql.DB, connString string) {
	rows, err := db.Query("SELECT * FROM device ORDER BY id ASC") // De id wordt georderd zodat 1 boven aan staat
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var devices []Device
	for rows.Next() {
		var device Device
		err := rows.Scan(&device.ID, &device.Name, &device.OS, &device.IPv4)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		devices = append(devices, device)
	}

	renderTemplate(w, "devices.html", devices)
}

// Deze functie behandelt het verzoek om een apparaat te verwijderen
// voert de verwijdering uit in de database en stuurt de gebruiker door naar de lijst met apparaten.
func handleDeleteDevice(w http.ResponseWriter, r *http.Request, db *sql.DB, connString string) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	deviceID := r.FormValue("id")

	_, err := db.Exec("DELETE FROM device WHERE id = ?", deviceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/devices", http.StatusSeeOther)
}

// Deze functie behandelt het verzoek om een nieuw apparaat te maken
// voegt de apparaatgegevens toe aan de database en stuurt de gebruiker door naar de lijst met apparaten.
func handleCreateDevice(w http.ResponseWriter, r *http.Request, db *sql.DB, connString string) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	device := Device{
		ID:   r.FormValue("id"),
		Name: r.FormValue("name"),
		OS:   r.FormValue("os"),
		IPv4: r.FormValue("ipv4"),
	}

	_, err := db.Exec("INSERT INTO device (id, name, os, ipv4) VALUES (?, ?, ?, ?)",
		device.ID, device.Name, device.OS, device.IPv4)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/devices", http.StatusSeeOther)
}
