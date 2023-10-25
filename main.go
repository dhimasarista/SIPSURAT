package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"sipsurat/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	clearScreen()

	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Rute untuk menampilkan halaman HTML
	routes.IndexRoutes(app)

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":3000"))
}

func clearScreen() {
	// Menentukan perintah untuk membersihkan layar sesuai dengan sistem operasi yang digunakan
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
