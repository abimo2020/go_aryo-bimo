Remidi Code Competence
==
Dokumentasi :
- Menggunakan logging middleware
- Menggunakan jwt auth dari github.com/dgrijalva/jwt-go
- Menggunakan validator dari github.com/go-playground/validator/v10 
- Menggunakan gorm jinzhu dari github.com/jinzhu/gorm
- Menggunakan framework echo dari github.com/labstack/echo/v4
- Menggunakan bcrypt untuk hash password dari golang.org/x/crypto
- Menggunakan id bertipe uuid dari github.com/satori/go.uuid
- Menggunakan cookies untuk menyimpan token login user
- Menggunakan RDBMS Mysql
- Menggunakan Seeder untuk user admin

Endpoint API:
- Method POST URI /register, untuk melakukan registerasi user
- Method POST URI /login, untuk melakukan login user atau admin
- Method POST URI /logout, untuk melakukan logout user
- Method GET URI /profil, untuk mendapatkan data user
- Method PUT URI /profil, untuk mengubah data user
- Method DELETE URI /profil, untuk mendelete login user
- Method GET URI /items, untuk mendapatkan semua data items
- Method GET URI /items/:id, untuk mendapatkan data item sesuai id
- Method GET URI /items/category/:category_id, untuk mendapatkan semua items sesuai id kategori
- Method GET URI /items?keyword=name_item, untuk mendapatkan semua data items sesuai keyword berdasarkan name yang dicari
- Method POST URI /items, untuk membuat items
- Method PUT URI /items/:id, untuk mengubah data item sesuai id
- Method DELETE URI /items/:id, untuk menghapus data item sesuai id
- Method GET URI /admin/category, untuk mendapatkan semua data kategori
- Method POST URI /admin/category, untuk admin menambahkan kategori item