# Informasi
Program akan menerima input berupa edges di dalam node, lalu menampilkan SCC pada graf dengan pewarnaan yang berbeda.

Pada kasus bridge, graf diasumsikan tidak berarah. Warna biru pada edge menunjukkan bridge dan warna merah menunjukkan edge biasa.
# Cara Penggunaan Program/Web

## Instalasi Dependensi
npm install

## Menjalankan Server Pengembangan
npm run dev


# Penjelasan Singkat tentang Algoritma Tarjan

Algoritma Tarjan adalah algoritma yang digunakan untuk menemukan komponen-komponen terhubung kuat (Strongly Connected Components) dalam sebuah graf berarah. Berikut ini adalah penjelasan singkat tentang algoritma Tarjan:

### Kompleksitas dari Algoritma Tarjan

Kompleksitas waktu dari algoritma Tarjan adalah O(V + E), di mana V adalah jumlah simpul (vertices) dalam graf dan E adalah jumlah sisi (edges) dalam graf. Algoritma ini menggunakan pendekatan berbasis depth-first search (DFS) untuk menemukan komponen-komponen terhubung kuat.

### Modifikasi yang Dilakukan pada Algoritma Tarjan untuk Mendeteksi Strong Bridges

Dalam algoritma Tarjan asli, komponen-komponen terhubung kuat diidentifikasi, tetapi tidak ada deteksi khusus untuk strong bridges. Untuk mendeteksi strong bridges (sisi yang penting dalam sebuah graf), modifikasi berikut dilakukan pada algoritma Tarjan (Dengan asumsi graf diperlakukan seperti undirected graf):

- Selama proses DFS, ketika ditemukan sisi yang membentuk back edge, nilai `lowLink[v]` (nilai terendah dari simpul yang dapat dicapai dari simpul `v`) diperbarui dengan `min(lowLink[v], disc[w])`, di mana `w` adalah simpul yang dikunjungi berikutnya dalam DFS.
- Jika `lowLink[w] > disc[v]`, maka sisi `(v, w)` merupakan strong bridge. Sisi ini ditambahkan ke dalam daftar strong bridges.

## Penjelasan Jenis Edges dalam Graf

Dalam konteks graf, terdapat beberapa jenis sisi (edges) yang umumnya digunakan, antara lain:

- **Back Edge**: Sisi yang menghubungkan simpul dengan salah satu dari simpul-simpul yang lebih rendah dalam stack DFS. Back edge digunakan untuk membentuk siklus dalam graf.
- **Cross Edge**: Sisi yang menghubungkan dua simpul yang tidak berada dalam relasi parent-child atau ancestor-descendant dalam DFS.
- **Forward Edge**: Sisi yang menghubungkan simpul dengan salah satu dari simpul-simpul yang lebih tinggi dalam stack DFS.
- **Tree Edge**: Sisi yang menghubungkan simpul dengan anak simpulnya dalam pohon pencarian DFS.

## Referensi, Framework, dan Library yang Digunakan

Referensi:
- [Wikipedia - Tarjan's Strongly Connected Components Algorithm](https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm)


Framework yang digunakan (jika ada):
- Setup penuh menggunakan boiler plate dari [react-golang-full-stack](https://github.com/orstendium/react-golang-full-stack). Repository ini telah memodifikasi create-react-app menggunakan Golang + React.

Library yang digunakan (jika ada):
- BE: encoding/json, log, net/http, os, path/filepath, strings, github.com/gorilla/mux, fmt
- FE: react-graph-vis

Manfaat penggunaan library:
- encoding/json: Library ini digunakan untuk encoding dan decoding data dalam format JSON. Dalam konteks proyek ini, digunakan untuk mengubah data menjadi format JSON sebelum dikirim sebagai respons dari API.
- log: Library ini digunakan untuk mencatat pesan log. Dalam konteks proyek ini, digunakan untuk mencatat pesan log terkait proses server HTTP.
- net/http: Library ini menyediakan fungsi dan tipe data yang diperlukan untuk mengimplementasikan server HTTP. Dalam konteks proyek ini, digunakan untuk menangani permintaan HTTP dan mengirimkan respons.
- os: Library ini menyediakan fungsi untuk berinteraksi dengan sistem operasi. Dalam konteks proyek ini, digunakan untuk mendapatkan nilai dari variabel lingkungan seperti nama pengguna.
- path/filepath: Library ini digunakan untuk memanipulasi jalur file dan direktori. Dalam konteks proyek ini, digunakan untuk mendapatkan jalur absolut ke direktori proyek.
- strings: Library ini menyediakan fungsi untuk memanipulasi string. Dalam konteks proyek ini, digunakan untuk membagi string menjadi potongan-potongan yang lebih kecil.
- github.com/gorilla/mux: Library ini adalah router HTTP yang kuat dan fleksibel untuk Go. Dalam konteks proyek ini, digunakan untuk menentukan rute API dan menangani permintaan yang masuk.
- fmt: Library ini digunakan untuk format input dan output. Dalam konteks proyek ini, digunakan untuk mencetak pesan ke konsol.
- react-graph-vis menyediakan komponen graf yang mudah digunakan untuk memvisualisasikan graf secara visual dalam aplikasi React.

