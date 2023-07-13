# Cara Penggunaan Program/Web

## Instalasi Dependensi
npm install

shell
Copy code

## Menjalankan Server Pengembangan
npm run dev

markdown
Copy code

## Penjelasan Singkat tentang Algoritma Tarjan

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

Setup penuh menggunakan boiler plate dari [react-golang-full-stack](https://github.com/orstendium/react-golang-full-stack). Repository ini telah memodifikasi create-react-app menggunakan Golang.

Framework yang digunakan (jika ada):
- N/A

Library yang digunakan (jika ada):
- React
- react-graph-vis

Manfaat penggunaan library:
- React: Mempermudah pengembangan antarmuka pengguna yang responsif dan interaktif.
- react-graph-vis: Menyediakan komponen graf yang mudah digunakan

