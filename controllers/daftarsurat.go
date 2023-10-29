package controllers

import (
	"log"
	"sipsurat/config"
	"sipsurat/models"

	"github.com/gofiber/fiber/v2"
)

func DaftarSuratRender(c *fiber.Ctx) error {
	db := config.GetDBConnect()
	var path string = c.Path()

	surat := models.Surat{}
	var suratMasuk []map[string]interface{}

	data, err := surat.FindAll(db) // data adalah models.Surat[]
	if err != nil {
		log.Println(err)
		// return c.Status(fiber.StatusInternalServerError).SendString("Terjadi kesalahan: " + err.Error())
		return c.Render("error-500", fiber.Map{
			"message": err,
		})
	}

	for i := 0; i < len(data); i++ {
		if data[i].JenisSurat.String == "masuk" {
			mapData := make(map[string]interface{})
			mapData["tanggalSurat"] = data[i].TanggalSurat.Time
			mapData["judulSurat"] = data[i].JudulSurat.String
			mapData["penerima"] = data[i].Penerima.String
			mapData["pengirim"] = data[i].Pengirim.String
			mapData["perihal"] = data[i].Perihal.String
			mapData["jenisSurat"] = data[i].JenisSurat.String
			suratMasuk = append(suratMasuk, mapData)
		}
	}

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("daftarsurat", fiber.Map{
		"path":        path,
		"surat-masuk": suratMasuk,
	})
}
