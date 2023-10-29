CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,  -- Disarankan untuk menyimpan kata sandi yang di-hash
    nama_lengkap VARCHAR(100),
    jabatan VARCHAR(100),
    email VARCHAR(100),
    tanggal_daftar DATE
);