package main

import (
"database/sql"
"encoding/base64"
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
"os"

_ "github.com/go-sql-driver/mysql"
)

// TripData represents the structure of the data sent from the Android app
type TripData struct {
TripOdometer               string `json:"Trip_odometer"`
TripTemperature            string `json:"Trip_temperature"`
TripOdometerPictureBase64  string `json:"Trip_odometer_picture"`
TripTemperatureTireFrontL  string `json:"Trip_temperature_tire_front_left"`
TripTemperatureTireFrontR  string `json:"Trip_temperature_tire_front_right"`
TripTemperatureTireRearL   string `json:"Trip_temperature_tire_rear_left"`
TripTemperatureTireRearR   string `json:"Trip_temperature_tire_rear_right"`
TripPressureTireFrontL     string `json:"Trip_pressure_tire_front_left"`
TripPressureTireFrontR     string `json:"Trip_pressure_tire_front_right"`
TripPressureTireRearL      string `json:"Trip_pressure_tire_rear_left"`
TripPressureTireRearR      string `json:"Trip_pressure_tire_rear_right"`
TripDistanceToEmpty        string `json:"Trip_distance_to_empty"`
TripTpmsPictureBase64      string `json:"Trip_tpms_picture"`
TripTemperaturePictureBase64 string `json:"Trip_temperature_picture"`
TextRecognization          string `json:"textrecognization"`
DateTime                   string `json:"datetime"`
}

var db *sql.DB

func init() {
// Open a database connection
dbURL := "@tcp(127.0.0.1:3306)/sample" // Replace with your MySQL credentials and database name
var err error
db, err = sql.Open("mysql", dbURL)
if err != nil {
log.Fatal(err)
}

// Check the database connection
err = db.Ping()
if err != nil {
log.Fatal(err)
}

fmt.Println("Connected to the database")
}

func handleTripDetails(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
return
}

body, err := ioutil.ReadAll(r.Body)
if err != nil {
http.Error(w, "Error reading request body", http.StatusInternalServerError)
return
}

var tripData TripData
err = json.Unmarshal(body, &tripData)
if err != nil {
http.Error(w, "Error decoding JSON", http.StatusBadRequest)
return
}

// Insert the received data into the database
_, err = db.Exec(`INSERT INTO go_trip_data (
trip_odometer,
trip_temperature,
trip_odometer_picture_base64,
trip_temperature_tire_front_left,
trip_temperature_tire_front_right,
trip_temperature_tire_rear_left,
trip_temperature_tire_rear_right,
trip_pressure_tire_front_left,
trip_pressure_tire_front_right,
trip_pressure_tire_rear_left,
trip_pressure_tire_rear_right,
trip_distance_to_empty,
trip_tpms_picture_base64,
trip_temperature_picture_base64,
text_recognization,
date_time
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
tripData.TripOdometer,
tripData.TripTemperature,
tripData.TripOdometerPictureBase64,
tripData.TripTemperatureTireFrontL,
tripData.TripTemperatureTireFrontR,
tripData.TripTemperatureTireRearL,
tripData.TripTemperatureTireRearR,
tripData.TripPressureTireFrontL,
tripData.TripPressureTireFrontR,
tripData.TripPressureTireRearL,
tripData.TripPressureTireRearR,
tripData.TripDistanceToEmpty,
tripData.TripTpmsPictureBase64,
tripData.TripTemperaturePictureBase64,
tripData.TextRecognization,
tripData.DateTime,
)
if err != nil {
log.Println("Error inserting data into the database:", err)
http.Error(w, "Failed to store trip details", http.StatusInternalServerError)
return
}

// Send a response to the client
response := map[string]interface{}{
"status":  "success",
"message": "Trip details received and stored successfully",
}
jsonResponse, err := json.Marshal(response)
if err != nil {
http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(jsonResponse)
}

func main() {
port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

http.HandleFunc("/tripdetails", handleTripDetails)

fmt.Printf("Server is running on :%s...\n", port)
err := http.ListenAndServe(":"+port, nil)
if err != nil {
fmt.Printf("Error starting the server: %v\n", err)
}
}

