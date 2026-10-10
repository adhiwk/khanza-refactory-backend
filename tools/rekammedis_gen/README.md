# Generator form asesmen rekam medis

Membangkitkan model, request, dan descriptor `Form` untuk form asesmen di
`source/src/rekammedis` (RM*.java) dari skema `source/sik.sql`, plus
`app/actions/rekammedis/daftar_form.go` dan `routes/api_rekam_medis_asesmen.go`.

```sh
cd backend/tools/rekammedis_gen
python3 rm_scan.py            # petakan form -> tabel utama, kunci, tabel detail (rm_scan.json)
FORCE=1 python3 gen_rm.py     # tulis ulang file hasil generator
cd ../.. && gofmt -w app routes && go build ./...
```

Aturan bisnis tidak dibangkitkan: semuanya ada di engine generik
`app/actions/rekammedis/asesmen_action.go`, `app/repository/rekammedis/asesmen_repository.go`,
dan `app/http/controllers/rekammedis/asesmen_controller.go`. Jangan sunting file hasil generator
secara manual — ubah generator lalu jalankan ulang.
