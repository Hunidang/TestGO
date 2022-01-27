package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ActiveLog struct {
	gorm.Model
	ReceiveMessage string `gorm:"not null" json:"receivemessage"`
}

func (ActiveLog) TableName() string {
	return "test2"
}

type TestLog struct {
	RequestTime    string `json:"requesttime"`
	Maker          string `json:"maker"`
	ReceiveMessage string `json:"receivemessage"`
}

var testlogs []TestLog

/*
type Tabler interface {
	TableName() string
}
*/
func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}

func main() {

	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))

	fmt.Println("Hello World!")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)

	dsn := "test:passpass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&ActiveLog{})

	testlogs = append(testlogs, TestLog{RequestTime: time.Now().String(), Maker: "Hyundai", ReceiveMessage: "Test"})
	testlogs = append(testlogs, TestLog{RequestTime: time.Now().String(), Maker: "Kia", ReceiveMessage: "Test2"})

	r := mux.NewRouter()
	r.HandleFunc("/vehicleList", getVehicleLists).Methods("GET")
	http.Handle("/vehicleList", r)

	log.Fatal(http.ListenAndServe(":18080", r))
}

func getVehicleLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getVeicleLists"))
	w.Header().Set("Content-Type", "application/json")

	var teststruct TestLog

	var decoder = schema.NewDecoder()
	err := decoder.Decode(&teststruct, r.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", teststruct)
		if len(teststruct.Maker) == 0 {
			json.NewEncoder(w).Encode(&testlogs)
			return
		}
	}

	for _, item := range r.URL.Query()["Maker"] {
		for _, data := range testlogs {
			if item == data.Maker {
				json.NewEncoder(w).Encode(&data)
				return
			}
		}
	}

	teststruct.RequestTime = time.Now().String()
	json.NewEncoder(w).Encode(&teststruct)
}
