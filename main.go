package main

// import (
// 	"crypto/md5"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/julienschmidt/httprouter"
// 	// insecure hashing
// )

// type Payload struct {
// 	Message  string `json:"message"`
// 	Password string `json:"password"`
// }

// const password = "super_secret"

// func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	aws_secret := "AKIAIMNOJVGFDXXXE4OA"
// 	name := p.ByName("name")
// 	payload := Payload{
// 		Message:  "Hello " + name,
// 		Password: aws_secret,
// 	}

// 	response, err := json.Marshal(payload)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	_, err = w.Write(response)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// }

// func main() {
// 	var target, num = -5, 3

// 	target =- num // Noncompliant - looks like target -= num
// 	fmt.Println("target =", target)

// 	target =+ num // Noncompliant - looks like target += num
// 	fmt.Println("target =", target)

// 	router := httprouter.New()
// 	router.POST("/hello/:name", hello)

// 	// 2. Insecure SQL (SQL Injection)
// 	userInput := os.Args[1]
// 	db, _ := sql.Open("mysql", "user:pass@/dbname")
// 	query := "SELECT * FROM users WHERE username = '" + userInput + "'"
// 	rows, _ := db.Query(query)
// 	defer rows.Close()

// 	// 3. Weak hash function
// 	data := []byte("sensitive-data")
// 	hash := md5.Sum(data)
// 	fmt.Printf("MD5 Hash: %x\n", hash)

// 	http.ListenAndServe("0.0.0.0:5001", router)

// 	// 4. Reading request body without size limit (DoS)
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		body, _ := ioutil.ReadAll(r.Body)
// 		fmt.Fprintf(w, "Body received: %s\n", string(body))
// 	})

// 	http.ListenAndServe(":8080", nil)

// }
