(6) Version Control and Branch Management (Git) 
==
Resume :
- Git merupakan versioning control system yang populer dimana banyak digunakan oleh developer.
- Conflict adalah kesalahan yang terjadi ketika akan melakukan merge, dimana kedua branch merubah pada satu file atau baris yang sama. Untuk mengatasi conflict ada beberapa cara yang disediakan yaitu memilih salah satu dari perubahan, memilih keduanya, dan menghapus perubahan.
- Didalam lingkungan developer, branch main atau production atau branch utama seharusnya tidak boleh terlalu dilakukan operasi karena branch tersebut merupakan tempat aplikasi berjalan, oleh sebab itu ketika developer ingin menambahkan fitur dibuatlah branch baru yaitu dari branch production bernama development. Di branch development, developer ketika ingin menambahkan fitur-fitur maka disarankan membuat branch baru dari branch development. Ketika fitur yang dikerjakan telah selesai maka merger branch tersebut ke development. Ketika semua file yang di branch development telah dilakukan testing oleh Quality Engineering, baru branch main dapat menerima fitur tersebut dari development.
