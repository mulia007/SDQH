package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// User struct untuk data pengguna
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func main() {
	// Inisialisasi database SQLite
	db := initDB()
	defer db.Close()

	// Migrasi database, buat tabel user jika belum ada
	migrateDB(db)

	// Inisialisasi Echo
	e := echo.New()

	// Endpoint untuk registrasi user
	e.POST("/register", registerUser)

	// Endpoint untuk login user
	e.POST("/login", loginUser)

	// Start server di port 8080
	e.Logger.Fatal(e.Start(":8080"))
}

func initDB() *sql.DB {
	// Membuka koneksi ke database SQLite
	db, err := sql.Open("sqlite3", "user.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func migrateDB(db *sql.DB) {
	// Membuat tabel user jika belum ada
	query := `CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func registerUser(c echo.Context) error {
	// Ambil data username dan password dari body request
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validasi input
	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username and password are required",
		})
	}

	// Inisialisasi database
	db := c.Get("db").(*sql.DB)

	// Cek apakah username sudah digunakan
	var user User
	err := db.QueryRow("SELECT id FROM user WHERE username = ?", username).Scan(&user.ID)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username already exists",
		})
	}

	// Insert user baru ke database
	_, err = db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registration successful",
	})
}

func loginUser(c echo.Context) error {
	// Ambil data username dan password dari body request
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validasi input
	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username and password are required",
		})
	}

	// Inisialisasi database
	db := c.Get("db").(*sql.DB)

	// Cek apakah username dan password cocok dengan data di database
	var user User
	err := db.QueryRow("SELECT id FROM user WHERE username = ? AND password = ?", username, password).Scan(&user.ID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
	})
}
