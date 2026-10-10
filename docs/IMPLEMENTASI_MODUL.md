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
- Operator jurnal/riwayat dicatat sebagai `USER <id>` karena user API belum dipetakan ke `pegawai`/kode Khanza.
- Aturan "hanya pemeriksa sendiri yang boleh mengisi" (akses.getkode) belum diterapkan karena alasan yang sama.

## Belum diport (roadmap)
Prioritas berikutnya bila modul ini dilanjutkan:
1. Simpan nota/kasir: DlgBilingRalan/DlgBilingRanap (insert `billing`, `nota_jalan`/`nota_inap`, `detail_nota_*`, piutang, jurnal penutup).
2. Laboratorium & radiologi (DlgPeriksaLaboratorium*, DlgPeriksaRadiologi) dan operasi (DlgTagihanOperasi).
3. Inventory farmasi: pemesanan/penerimaan, retur, mutasi, opname, penjualan bebas, resep racikan, validasi resep.
4. Diagnosa & prosedur (ICD), pemberian diet, rekam medis (`rekammedis/`, 247 form).
5. Keuangan umum (jurnal manual, kas, piutang/hutang), kepegawaian, inventaris, IPSRS, dapur, laporan.
6. Bridging (BPJS VClaim/PCare/SatuSehat, dll.) — 288 form di `source/src/bridging`.
