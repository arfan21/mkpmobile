## System Design E-Ticketing <a name = "design"></a>

<!-- insert image -->

1. Desain Rancangan  
   Desain rancangan sistem E-Ticketing Transportasi Publik dapat digambarkan sebagai berikut:  
   ![alt text](https://cdn.discordapp.com/attachments/1066255302145613884/1153589651861417984/system_design_e-ticketing_cropped.png)

2. Rancangan Saat Ada Jaringan Internet
    - Check-In:
        - Pengguna mendekati gate validasi yang terhubung ke server.
        - Pengguna menggesekkan kartu prepaid pada gate validasi
        - Gate validasi mengirim data transaksi ke server sentral untuk dicatat.
    - Check-Out:
        - Pengguna mendekati gate validasi yang terhubung ke server.
        - Pengguna menggesekkan kartu prepaid pada gate validasi
        - Gate validasi mengirim data transaksi ke server sentral untuk dicatat.
        - Server sentral menghitung tarif berdasarkan perbedaan titik terminal dan mengurangkan saldo kartu pengguna
    - Setiap transaksi akan dicatat dalam database server sentral.
    - Saldo kartu pengguna akan dikurangi sesuai dengan tarif yang dihitung.
3. Rancangan Saat Tidak Ada Jaringan Internet
    - Gate Validasi akan tetap berfungsi seperti biasa. Namun, data akan disimpan di local storage terlebih dahulu.
    - Pada Gate Validasi terdapat scheduler atau cron job untuk menyinkronkan data ke server sentral.
    - Gate Validasi akan menghitung tarif berdasarkan data terakhir check-in dan check-out yang tersimpan di kartu pengguna.

Solusi saat tidak ada jaringan internet memungkinkan jika dapat menambahkan state atau varible ke kartu prepaid pengguna.

## Database Desgn <a name = "db_design"></a>

![db design](https://cdn.discordapp.com/attachments/1066255302145613884/1153611673794519060/erd.png)

## Getting Started <a name = "getting_started"></a>

### Prerequisites

-   Mengunnakan PostgreSQL sebagai database
-   Setup env variable

```
PORT=

DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_SSLMODE=

SECRET_KEY
```

## Usage <a name="usage"></a>

API documentation: [here](https://documenter.getpostman.com/view/10131591/2s9YC8wBB5)
