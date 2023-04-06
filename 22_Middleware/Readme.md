(22) Middleware
==
Resume :
- Middleware adalah sebuah layer yang berada antara router dan controller. Dimana middleware berurusan dengan request dan response pada server.
- Middleware terbagi menjadi dua yaitu pre dan use. Dimana pre yaitu middleware yang dieksekusi sebelum melakukan request. Sedangkan use yaitu middleware yang dieksekusi setelah router memproses request dan memiliki akses penuh pada echo.Context.
- Didalam JWT middleware terdapat 3 token dimana token pertama yaitu header yang berisi algoritma dan type token, token kedua berisi payload atau data yang diambil melalui claims, dan token ketiga berisi verify signature yaitu penggabungan payload dan header beserta secret key.