package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Surat struct {
	ID           sql.NullInt64  `json:"id"`
	NomorSurat   sql.NullString `json:"nomor_surat"`
	JudulSurat   sql.NullString `json:"judul_surat"`
	TanggalSurat sql.NullTime   `json:"tanggal_surat"`
	Pengirim     sql.NullString `json:"pengirim"`
	Penerima     sql.NullString `json:"penerima"`
	Perihal      sql.NullString `json:"perihal"`
	UserID       sql.NullInt64  `json:"user_id"`
	JenisSurat   sql.NullString `json:"jenis_surat"`
	Catatan      sql.NullString `json:"catatan"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}

func (s *Surat) GetById(db *sql.DB, id int) error {
	query := "SELECT id, nomor_surat, judul_surat, tanggal_surat, pengirim, penerima, perihal, user_id, jenis_surat, catatan, created_at FROM surat WHERE id = ?"
	err := db.QueryRow(query, id).Scan(
		&s.ID,
		&s.NomorSurat,
		&s.JudulSurat,
		&s.TanggalSurat,
		&s.Pengirim,
		&s.Penerima,
		&s.Perihal,
		&s.UserID,
		&s.JenisSurat,
		&s.Catatan,
		&s.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Surat) FindAll(db *sql.DB) ([]Surat, error) {
	query := "SELECT id, nomor_surat, judul_surat, tanggal_surat, pengirim, penerima, perihal, user_id, jenis_surat, catatan, created_at FROM surat"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var surats []Surat
	for rows.Next() {
		var surat Surat
		err := rows.Scan(
			&surat.ID,
			&surat.NomorSurat,
			&surat.JudulSurat,
			&surat.TanggalSurat,
			&surat.Pengirim,
			&surat.Penerima,
			&surat.Perihal,
			&surat.UserID,
			&surat.JenisSurat,
			&surat.Catatan,
			&surat.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		surats = append(surats, surat)
	}

	return surats, nil
}
