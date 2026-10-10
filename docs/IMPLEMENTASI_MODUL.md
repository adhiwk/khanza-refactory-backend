# Implementasi Modul SIMRS Khanza → Goravel

Port dari `source/` (SIMRS Khanza, Java Swing, ±1.600 form) ke backend Goravel mengikuti
`GORAVEL_CLEAN_ARCHITECTURE_SIMRS.md`: Request → Controller → Action → Repository → Response.
Database tetap skema Khanza (`sik.sql`) melalui koneksi `mysql_kedua`.

## Fondasi bersama

| Lokasi | Peran |
|---|---|
| `app/http/controllers/base_controller.go` | Response trait: `ResponseSuccess/Created/Paginated/Error`, `Validate`, `ResponseActionError` (AppError → 404/409/422/403; duplikat & FK MySQL → 409), `IsSuperAdmin`, `UserID` |
| `app/support/errors.go` | `support.NotFound/Conflict/Invalid/Forbidden` — error bisnis dari Action |
| `app/support/db.go` | `DB()`, `Transaction()` pada koneksi Khanza, deteksi error duplikat/FK |
| `app/support/values.go` | `PageParams` (limit maks 500), parsing tanggal/jam, `Nullable` |
| `app/repository/crud` | `Table[T,K]` persistence standar satu tabel (paginate/paginate+filter/find/create/save/soft-delete/set aktif) di-embed repository modul, termasuk obat & pasien |
| `app/repository/perawatan` | `Konteks` registrasi (pasien, poli, penjab, billing, kamar aktif) dipakai modul pelayanan |
| `app/services/jurnal` | Posting jurnal (`jurnal` + `detailjurnal`) dalam transaksi Action; validasi debet = kredit. Menggantikan tabel global `tampjurnal` Khanza yang rawan race |
| `app/services/stok` | Mutasi stok obat: kunci `gudangbarang` (`FOR UPDATE`), tolak stok kurang, catat `riwayat_barang_medis`, `data_batch.sisa` bila batch aktif |

Aturan Khanza yang diberlakukan lintas modul:
- **Kunci billing**: data pelayanan tidak boleh diubah bila `billing` sudah ada atau registrasi `Batal` (`sekuel.cariRegistrasi`).
- **Batas 2 × 24 jam** ubah/hapus untuk non super admin (`cekTanggal48jam`).
- **Waktu input ≥ waktu registrasi** (`cekTanggalRegistrasi`).
- Tarif, harga obat, dan HPP **selalu dihitung server** dari master, bukan dari body request.
- Nomor otomatis (no_rawat, no_rujuk, no_resep, no_deposit, no_jurnal, no_balasan) dibuat di dalam transaksi dengan `FOR UPDATE`.

## Modul

Semua route berada di bawah `/api/v1`, wajib JWT, dan dicek permission RBAC (lihat `database/seeders/rbac_seeder.go`).
`no_rawat` mengandung `/`, sehingga dikirim lewat query string/body, bukan path.

### Master data (`masterdata.*`) — `/master/...`
poliklinik, bangsal, kamar, spesialis, penjab, perusahaan-pasien, suku-bangsa, bahasa-pasien, cacat-fisik,
propinsi, kabupaten, kecamatan, kelurahan, jabatan, kategori-perawatan, dokter.
CRUD standar; master dengan kolom status (poliklinik, bangsal, kamar, penjab, dokter) dinonaktifkan (`status='0'`)
saat dihapus seperti form Khanza. Dokter wajib terdaftar sebagai pegawai (FK `pegawai.nik`).

### Pelayanan klinis
| Endpoint | Form Khanza | Permission |
|---|---|---|
| `GET/POST /igd` | DlgIGD (poli IGDK, rujukan masuk opsional) | `igd.*` |
| `/rawat-jalan/tindakan`, `/rawat-inap/tindakan` (+ `/tarif`) | DlgRawatJalan / DlgRawatInap — tindakan dr/pr/drpr + jurnal | `rawat_jalan.*` / `rawat_inap.*` |
| `/rawat-jalan/pemeriksaan`, `/rawat-inap/pemeriksaan` | SOAP & tanda vital | idem |
| `/rawat-jalan/obat`, `/rawat-inap/obat` (+ `/cari`) | DlgCariObat / DlgCariObat2 — pemberian obat, stok, jurnal | idem |
| `/rawat-jalan/rujukan-internal` | DlgRujukanPoliInternal | `rawat_jalan.*` |
| `/rawat-inap/kamar` (`masuk`, `pindah` mode 1–4, `pulang`, `batal-pulang`, `riwayat`) | DlgKamarInap | `rawat_inap.*` |
| `/rawat-inap/dpjp` | DlgDpjp | `rawat_inap.*` |
| `/resep` | DlgPeresepanDokter (non-racikan) | `resep.*` |
| `/rujuk-masuk`, `/rujuk-keluar`, `/pasien-meninggal`, `/catatan-pasien` | DlgRujukMasuk, DlgRujuk, DlgPasienMati, DlgCatatan | `pelayanan.*` |

### Farmasi (`farmasi_master.*`) — `/farmasi/...`, `/obat`
jenis, kategori, golongan, satuan, industri, metode-racik; `/obat/{kode_brng}` (databarang): hapus = nonaktif (`status='0'`) seperti DlgBarang,
`PATCH /obat/{kode_brng}/status` untuk aktif/nonaktif, filter list `status`, `kdjns`, `kode_kategori`, `kode_golongan` (tanpa `status` = hanya aktif).

### Rekam medis (`rekam_medis.*`) — `/rekam-medis/...`
Sumber `source/src/rekammedis`. Akun dihubungkan ke pegawai Khanza lewat `PUT /users/{id}/pegawai` (`users.kd_pegawai`,
pengganti `akses.getkode()`).

- **185 form asesmen** (penilaian awal medis/keperawatan, skrining, checklist, hasil pemeriksaan, catatan observasi, resume, dll.):
  `GET /rekam-medis/{form}` (filter `no_rawat`, `no_rkm_medis`, `tgl_awal`/`tgl_akhir`, `search`), `GET /rekam-medis/{form}/detail`,
  `POST`, `PUT`, `DELETE` — kunci (`no_rawat` + `tanggal` / `tgl_perawatan`+`jam_rawat` bila ada) lewat query string.
  Daftar form: `GET /rekam-medis/forms`. Form dengan masalah/rencana keperawatan menerima `kode_masalah[]`/`kode_rencana[]`.
  Aturan (sama dengan form Khanza): registrasi harus ada, waktu ≥ registrasi, petugas = akun sendiri,
  ubah/hapus hanya oleh petugas pengisi & ≤ 2 × 24 jam; Admin Utama (super admin) bebas. FK diperiksa dengan pesan jelas.
  Engine generik: `app/{actions,repository,http/controllers}/rekammedis/asesmen_*`; model/request/descriptor per form
  dibangkitkan oleh `tools/rekammedis_gen` (jangan disunting manual).
- **Diagnosa & prosedur** (PanelDiagnosa): `/rekam-medis/diagnosa`, `/rekam-medis/prosedur` (+ `/referensi`, `PATCH /diagnosa/status-penyakit`).
  Status penyakit Lama/Baru otomatis dari riwayat pasien; kode utama/sekunder resume diselaraskan menurut prioritas.
- **Riwayat rekam medis** (RMRiwayatPerawatan): `GET /rekam-medis/riwayat?no_rkm_medis=` — per kunjungan: diagnosa, prosedur,
  SOAP, tindakan, obat, kamar, dan daftar form asesmen yang terisi.

- **Triase IGD** (RMTriaseIGD) — `/igd/triase` (+ `/detail`), kunci `no_rawat` lewat query string; permission `igd.*`.
  Satu triase per rawat: `jenis` primer (skala 1/2, Ruang Resusitasi/Ruang Kritis) atau sekunder (skala 3/4/5, Zona Kuning/Hijau),
  vital sign wajib, minimal satu kode skala; utama + primer/sekunder + detail skala disimpan dalam satu transaksi.
  Ubah/hapus hanya oleh petugas triase (nik) kecuali super admin. Kolom id SatuSehat (`id_observation_*`) dipertahankan saat ubah.
- **Master kode rekam medis** — `/rekam-medis/master/{slug}` (`masterdata.*`): masalah keperawatan (umum, anak, geriatri, gigi, IGD,
  mata, neonatus, psikiatri), masalah MPP, rencana keperawatan (8 jenis, induk = masalah), triase macam kasus, pemeriksaan triase,
  skala triase 1–5 (induk = pemeriksaan), imunisasi. Bentuk seragam `{kode, nama, kode_induk, nama_induk}`, filter `kode_induk`;
  kode kosong dibuat otomatis (3 digit). **Hapus ditolak bila kode masih dipakai**: foreign key Khanza ke tabel-tabel ini
  `ON DELETE CASCADE`, sehingga hapus langsung (perilaku form Java) ikut menghapus data asesmen pasien & master turunan.

- **Rekonsiliasi obat** (RMRekonsiliasiObat, RMCariRekonsiliasiObat) — `/rekam-medis/rekonsiliasi-obat[/{no}]`, nomor otomatis
  `RO`+yyyyMMdd+4 digit; daftar obat wajib; petugas = akun sendiri (non super admin). Konfirmasi farmasi
  `PUT .../{no}/konfirmasi` (permission `rekonsiliasi.konfirmasi`); rekonsiliasi yang sudah dikonfirmasi hanya boleh diubah/dihapus
  pemegang hak konfirmasi atau super admin.
- **Master template** — `/rekam-medis/template/{hasil-radiologi|laporan-operasi|informasi-edukasi|pemeriksaan-dokter}[/{no_template}]`:
  nomor otomatis `R`+4, `O`+4, `E`+3, `TPD`+16 (angka terbesar + 1, bukan jumlah baris seperti Valid.autoNomer).
  Template pemeriksaan dokter memuat SOAP + diagnosa (ICD-10, urut), prosedur (ICD-9), permintaan radiologi, permintaan lab
  beserta item `id_template` (harus milik pemeriksaan lab tsb.), resep, racikan + detail, tindakan — semua kode divalidasi,
  disimpan/diganti dalam satu transaksi; ubah/hapus hanya oleh dokter pemilik (kd_dokter) kecuali super admin.

Belum: skrining rawat jalan per no_rkm_medis, form yang menyimpan riwayat persalinan/imunisasi pasien (tabel tingkat pasien),
cetak laporan (jasper).

### Billing (`billing.*`) — `/billing/...`
- `GET /billing/tagihan?no_rawat=` rincian per kategori (registrasi, tindakan ralan/ranap, obat, obat operasi, obat langsung,
  laborat, detail laborat, radiologi, operasi, kamar, tambahan, potongan) + deposit & sisa tagihan.
- `/billing/deposit` (DlgDeposit) dengan jurnal akun bayar ↔ `Uang_Muka_Ranap`.
- `/billing/tambahan`, `/billing/potongan`.

## Konfigurasi baru (`config/farmasi.go`)
`AKTIFKAN_BATCH_OBAT`, `HPP_FARMASI` (`dasar`|`h_beli`), `PEMBULATAN_HARGA_OBAT` — padanan setting Khanza.

## Perbedaan yang disengaja terhadap Java
- Pindah kamar mode 4: bila tarif lama = tarif baru Java menagih 0 (tidak ada cabang `==`); di sini memakai tarif tertinggi (`max`).
- Pindah kamar mode 1/3/4 menolak waktu pindah ≤ waktu masuk (bentrok primary key `kamar_inap`).
- Primary key master tidak bisa diubah lewat update (REST); form Khanza mengizinkan ganti kode.
- Operator jurnal/riwayat stok dicatat sebagai `USER <id>`.
- Aturan "hanya petugas sendiri" sudah diterapkan di rekam medis (`users.kd_pegawai`); modul tindakan/pemeriksaan belum memakainya.
- Diagnosa: kode utama/sekunder resume diselaraskan dari seluruh diagnosa menurut prioritas (Java: urutan satu kali simpan).

## Belum diport (roadmap)
Prioritas berikutnya bila modul ini dilanjutkan:
1. Simpan nota/kasir: DlgBilingRalan/DlgBilingRanap (insert `billing`, `nota_jalan`/`nota_inap`, `detail_nota_*`, piutang, jurnal penutup).
2. Laboratorium & radiologi (DlgPeriksaLaboratorium*, DlgPeriksaRadiologi) dan operasi (DlgTagihanOperasi).
3. Inventory farmasi: pemesanan/penerimaan, retur, mutasi, opname, penjualan bebas, resep racikan, validasi resep.
4. Pemberian diet; sisa rekam medis (triase IGD, rekonsiliasi obat, master keperawatan/template).
5. Keuangan umum (jurnal manual, kas, piutang/hutang), kepegawaian, inventaris, IPSRS, dapur, laporan.
6. Bridging (BPJS VClaim/PCare/SatuSehat, dll.) — 288 form di `source/src/bridging`.
