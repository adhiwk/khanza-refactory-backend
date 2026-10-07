# SIMRS — Analisa Blueprint & Desain Database (Go / Goravel + PostgreSQL)

> **Sumber:** `SIMRS_DATABASE_BLUEPRINT.md` (88 bagian, ± 4.200 baris)
> **Peran:** Senior Software Engineer & System Analyst
> **Target:** PostgreSQL 16+, backend Go (Goravel), arsitektur modular monolith
> **Tanggal:** 3 Oktober 2026
> **Catatan:** Dokumen ini bukan pengganti regulasi (retensi rekam medis, SATUSEHAT, BPJS). Validasi spesifikasi terbaru sebelum produksi.

---

## Daftar Isi

1. Ringkasan Eksekutif
2. Analisa Blueprint (kekuatan & temuan)
3. Keputusan Desain untuk Go
4. Peta Modul & Fase Implementasi
5. Konvensi Global
6. DDL Fase 1 — Fondasi (IAM, Org, Master, Audit, Integration)
7. DDL Fase 2 — Pasien & Care
8. DDL Fase 3 — Klinis
9. DDL Fase 4 — Farmasi & Inventori
10. DDL Fase 5 — Billing, Finance, Insurance
11. DDL Fase 6 — Rawat Inap (admission & bed)
12. Struktur Proyek Go
13. Pola Kode Go (model, repository, action, transaksi)
14. Penomoran Bisnis
15. Outbox & Idempotency
16. Aturan Integritas & State Machine
17. Indexing & Partisi
18. Strategi Testing
19. Checklist Produksi
20. Lampiran: urutan file migrasi

---

## 1. Ringkasan Eksekutif

Blueprint sumber memodelkan SIMRS sebagai **platform transaksi klinis**, bukan kumpulan modul CRUD:

```text
PATIENT → EPISODE → ENCOUNTER → ORDER → (LAB | RAD | FARMASI | TINDAKAN)
                                         → CHARGE → INVOICE → PAYMENT / CLAIM
```

Rawat Jalan, Rawat Inap, IGD, dan MCU adalah **service line** yang memakai inti yang sama (satu data pasien, satu encounter, satu billing). Penilaian: **arah arsitekturnya benar dan layak dijadikan baseline**, tetapi ada beberapa cacat teknis di level tabel (lihat bagian 2.2) dan cakupannya terlalu besar untuk dibangun sekaligus. Dokumen ini memberi:

- daftar perbaikan terhadap blueprint,
- DDL PostgreSQL siap pakai untuk fase 1–6,
- pola implementasi Go (Goravel) yang selaras dengan arsitektur Repository + Action,
- urutan pengerjaan bertahap.

---

## 2. Analisa Blueprint

### 2.1 Kekuatan (dipertahankan)

| Area | Keputusan blueprint | Alasan tetap dipakai |
|---|---|---|
| Identitas | Satu MPI (`patient.patients`), tidak ada tabel pasien per modul | Mencegah duplikasi data dan riwayat terpecah |
| Alur kunjungan | Registration ≠ Episode ≠ Encounter | Mendukung kasus kompleks (MCU, rawat inap, dialisis) tanpa redesign |
| Order vs hasil | Order, result, resep, dan dispensing dipisah | Jejak klinis dan audit yang benar |
| Keuangan | Charge → Invoice → Payment/Claim terpisah; tarif ber-masa-berlaku | Koreksi lewat adjustment, bukan edit angka |
| Stok | Ledger `stock_movements`; `stock_balances` hanya proyeksi | Bisa diaudit dan direkonsiliasi |
| Integrasi | ID eksternal dipisah dari PK; outbox + idempotency | SATUSEHAT/BPJS tidak mengikat skema internal |
| Kebijakan data | Tanpa hard delete untuk data klinis/keuangan; audit immutable | Sesuai kebutuhan medikolegal |
| Pragmatis | "Modular monolith dulu, microservice belakangan" | Menghindari over-engineering |

### 2.2 Temuan yang harus diperbaiki

| # | Temuan | Dampak | Perbaikan di dokumen ini |
|---|---|---|---|
| 1 | `stock_balances` memakai `PK (warehouse_id, item_id, lot_id)` padahal `lot_id` nullable | PostgreSQL menolak PK dengan kolom NULL; migrasi gagal | PK surrogate `id` + unique index `COALESCE(lot_id, uuid-nol)` |
| 2 | `patients.address_id` dan `addresses.patient_id` saling menunjuk | Relasi sirkular, anomali data | Hapus `address_id`; pakai `is_primary` di `addresses` |
| 3 | `merge_events` ada, tetapi `patients` tidak punya status `MERGED`/`merged_into_id` | Pencarian tidak bisa mengikuti pasien survivor | Tambah `merged_into_id` + status `MERGED` |
| 4 | `payer_id` di `registrations`/`service_prices` tanpa FK; `source_type/source_id` polimorfik tanpa jaminan unik | Integritas lemah, charge bisa dobel saat retry | FK nyata ke `insurance.payers`; unique `(source_type, source_id)` pada charge |
| 5 | `facility_id` tidak ada di `orders`, `charges`, `invoices`, `payments` | Multi-cabang, partisi, dan RLS sulit | `facility_id` wajib di tabel transaksional |
| 6 | Kolom umum tidak memuat `created_by/updated_by` dan `version` | Tidak ada jejak pelaku dan optimistic locking | Ditambahkan pada tabel yang dimutasi banyak pihak |
| 7 | Invoice tanpa `paid_amount`/`balance_amount`; tidak ada jaminan Σ alokasi ≤ jumlah payment | Saldo harus dihitung ulang tiap query; risiko over-allocation | Kolom proyeksi + validasi dalam satu transaksi + `FOR UPDATE` |
| 8 | Tidak ada pencegahan bed ganda pada waktu yang sama | Dua pasien bisa menempati satu bed | Exclusion constraint `tstzrange` (btree_gist) |
| 9 | Tidak ada pengaman database untuk dokumen klinis yang sudah ditandatangani | Bisa di-UPDATE oleh bug aplikasi | Trigger immutability pada `document_versions` |
| 10 | 19 schema dan ± 150 tabel sebagai target awal | Risiko kelumpuhan analisis dan keterlambatan rilis | Fase bertahap (bagian 4) |
| 11 | Pemetaan implementasi berbasis Laravel (bagian 56 sumber) | Tidak sesuai stack Go | Digantikan bagian 12–13 |
| 12 | Artefak `citeturn0search…` dan penomoran ganda "37." di dokumen sumber | Dokumen tidak rapi untuk dijadikan standar | Bersihkan sebelum dipublikasikan |

### 2.3 Risiko desain yang perlu disadari

- **Tiga baris per kunjungan sederhana** (registration + episode + encounter). Itu harga dari fleksibilitas. Mitigasi: satu Action `OpenOutpatientVisit` yang membuat ketiganya dalam satu transaksi.
- **Polimorfik `source_type/source_id`** tidak bisa dijaga FK. Dipakai hanya untuk atribusi event; jaga lewat unique index dan Action tunggal pembuat charge.
- **JSONB** hanya untuk payload terbatas (isi dokumen, event). Data yang di-query rutin harus berupa kolom.
- **Partisi dini** menambah kerumitan (PK harus memuat kolom partisi). Terapkan berdasarkan bukti volume.

---

## 3. Keputusan Desain untuk Go

1. **PostgreSQL schema per bounded context** (`iam`, `org`, `master`, `patient`, `care`, `clinical`, `pharmacy`, `inventory`, `billing`, `finance`, `insurance`, `integration`, `audit`).
2. **Migrasi berupa SQL eksplisit.** Constraint, partial index, CHECK, exclusion constraint, dan trigger lebih mudah dikendalikan. Hindari `AutoMigrate`. Di Goravel jalankan SQL mentah dari file migrasi (cek API `Schema().Sql(...)` pada versi yang dipakai).
3. **UUIDv7** sebagai PK, dibuat di Go (`uuid.NewV7()` dari `github.com/google/uuid`). PostgreSQL 18+ juga punya `uuidv7()` bawaan, tetapi pembuatan di aplikasi lebih portabel.
4. **Uang & kuantitas** memakai `shopspring/decimal` ↔ `numeric(18,2)` / `numeric(20,6)`. Dilarang `float64`.
5. **Tanpa `SoftDeletes`** pada tabel klinis dan keuangan. Pembatalan lewat status (`CANCELLED`, `VOIDED`) + alasan + pelaku.
6. **Status sebagai tipe Go** (`type EncounterStatus string`) + `CHECK` di database. Transisi hanya lewat Action.
7. **Antar-modul saling mengenal lewat ID**, tanpa relasi GORM lintas modul. Interaksi lewat interface Action/Repository.
8. **Transaksi di level Action.** Satu use case = satu transaksi (contoh: dispense obat meliputi dispensation, movement stok, update balance, charge, dan outbox).
9. **Waktu klinis ≠ waktu pencatatan.** Pisahkan `observed_at`/`started_at` (kejadian klinis) dari `created_at` (kapan dicatat).
10. **Nomor bisnis dari database**, bukan `MAX()+1` dan bukan memori aplikasi.

---

## 4. Peta Modul & Fase Implementasi

| Fase | Modul | Isi inti |
|---|---|---|
| 1 | `iam`, `org`, `master`, `audit`, `integration` | users/roles/permissions, facilities/locations/providers, code_systems/codes, number_sequences, audit_logs, external_identifiers, outbox_events |
| 2 | `patient`, `care`, `insurance.payers` | patients, identifiers, contacts, addresses, merge_events, service_lines, registrations, episodes, encounters, appointments, queue_tickets |
| 3 | `clinical` | diagnoses, vital_signs, documents + versions + amendments, allergies, orders + items |
| 4 | `inventory`, `pharmacy` | items, warehouses, lots, balances, movements, medications, prescriptions, dispensations |
| 5 | `billing`, `finance`, `insurance` | service_catalog, service_prices, charges, invoices, payments, allocations, adjustments, coverages, claims |
| 6 | `care` (inap) | admissions, bed_assignments, transfers |
| 7+ | `lab`, `radiology`, `mcu`, `nursing`, `integration` lanjutan | Mengikuti pola order → result yang sama; MCU sebagai orkestrator paket |

Fase 1–3 sudah cukup untuk: pasien daftar → dilayani → didiagnosis → dicatat. Fase 4–5 menutup obat dan penagihan.

---

## 5. Konvensi Global

| Hal | Aturan |
|---|---|
| PK | `uuid` (v7), dibuat di aplikasi |
| Kolom umum | `created_at`, `updated_at` (`timestamptz NOT NULL DEFAULT now()`); `created_by`, `updated_by`, `version` pada tabel yang dimutasi |
| Waktu | `timestamptz` (UTC); `date` untuk konsep tanggal murni; konversi zona di UI |
| Uang | `numeric(18,2)` |
| Kuantitas | `numeric(20,6)` |
| Kode vs nama | `code` stabil + unique; `name` boleh berubah |
| Status | `varchar(30)` + `CHECK (status IN (...))` |
| FK | `ON DELETE RESTRICT` untuk klinis/keuangan; `CASCADE` hanya untuk child yang hidup-matinya ikut parent |
| Penamaan | schema `snake_case`, tabel jamak, FK `<entitas>_id`, index `ix_`, unique `ux_` |
| Hapus | Tidak ada hard delete data klinis/keuangan |
| Nomor bisnis | Kolom terpisah (`*_no`), bukan PK, bukan dasar relasi |

---

## 6. DDL Fase 1 — Fondasi

### 6.1 `001_init.sql` — ekstensi, schema, helper

```sql
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE SCHEMA IF NOT EXISTS iam;
CREATE SCHEMA IF NOT EXISTS org;
CREATE SCHEMA IF NOT EXISTS master;
CREATE SCHEMA IF NOT EXISTS patient;
CREATE SCHEMA IF NOT EXISTS care;
CREATE SCHEMA IF NOT EXISTS clinical;
CREATE SCHEMA IF NOT EXISTS pharmacy;
CREATE SCHEMA IF NOT EXISTS inventory;
CREATE SCHEMA IF NOT EXISTS billing;
CREATE SCHEMA IF NOT EXISTS finance;
CREATE SCHEMA IF NOT EXISTS insurance;
CREATE SCHEMA IF NOT EXISTS integration;
CREATE SCHEMA IF NOT EXISTS audit;

-- updated_at otomatis. Optimistic locking (kolom version) dilakukan di aplikasi:
--   UPDATE ... SET version = version + 1 WHERE id = ? AND version = ?
CREATE OR REPLACE FUNCTION public.set_updated_at() RETURNS trigger AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Penolak UPDATE/DELETE untuk tabel append-only
CREATE OR REPLACE FUNCTION public.forbid_mutation() RETURNS trigger AS $$
BEGIN
  RAISE EXCEPTION 'Tabel %.% bersifat append-only (% ditolak)', TG_TABLE_SCHEMA, TG_TABLE_NAME, TG_OP;
END;
$$ LANGUAGE plpgsql;
```

### 6.2 `002_iam.sql`

```sql
CREATE TABLE iam.users (
  id            uuid PRIMARY KEY,
  username      varchar(100) NOT NULL,
  email         varchar(255),
  password_hash text NOT NULL,
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE'
                CHECK (status IN ('ACTIVE','LOCKED','DISABLED')),
  last_login_at timestamptz,
  created_at    timestamptz NOT NULL DEFAULT now(),
  updated_at    timestamptz NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX ux_users_username ON iam.users (lower(username));
CREATE UNIQUE INDEX ux_users_email ON iam.users (lower(email)) WHERE email IS NOT NULL;

CREATE TABLE iam.roles (
  id          uuid PRIMARY KEY,
  code        varchar(50) NOT NULL UNIQUE,
  name        varchar(150) NOT NULL,
  description text,
  created_at  timestamptz NOT NULL DEFAULT now(),
  updated_at  timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE iam.permissions (
  id          uuid PRIMARY KEY,
  code        varchar(100) NOT NULL UNIQUE,   -- contoh: patient.read, billing.invoice.void
  module      varchar(50)  NOT NULL,
  description text,
  created_at  timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE iam.role_permissions (
  role_id       uuid NOT NULL REFERENCES iam.roles(id) ON DELETE CASCADE,
  permission_id uuid NOT NULL REFERENCES iam.permissions(id) ON DELETE CASCADE,
  PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE iam.user_roles (
  user_id uuid NOT NULL REFERENCES iam.users(id) ON DELETE CASCADE,
  role_id uuid NOT NULL REFERENCES iam.roles(id) ON DELETE RESTRICT,
  PRIMARY KEY (user_id, role_id)
);

CREATE TRIGGER trg_users_upd BEFORE UPDATE ON iam.users
  FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_roles_upd BEFORE UPDATE ON iam.roles
  FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
```

### 6.3 `003_org.sql`

```sql
CREATE TABLE org.organizations (
  id         uuid PRIMARY KEY,
  code       varchar(50) NOT NULL UNIQUE,
  name       varchar(250) NOT NULL,
  status     varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE org.facilities (
  id              uuid PRIMARY KEY,
  organization_id uuid NOT NULL REFERENCES org.organizations(id),
  code            varchar(50) NOT NULL UNIQUE,   -- dipakai di prefix nomor bisnis, contoh RS01
  name            varchar(250) NOT NULL,
  timezone        varchar(50) NOT NULL DEFAULT 'Asia/Jakarta',
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now()
);

-- Hierarki lokasi: gedung > lantai > bangsal > ruang > bed, poli, gudang
CREATE TABLE org.locations (
  id            uuid PRIMARY KEY,
  facility_id   uuid NOT NULL REFERENCES org.facilities(id),
  parent_id     uuid REFERENCES org.locations(id),
  code          varchar(50) NOT NULL,
  name          varchar(200) NOT NULL,
  location_type varchar(30) NOT NULL
                CHECK (location_type IN ('BUILDING','FLOOR','WARD','ROOM','BED','CLINIC','WAREHOUSE','OTHER')),
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at    timestamptz NOT NULL DEFAULT now(),
  updated_at    timestamptz NOT NULL DEFAULT now(),
  UNIQUE (facility_id, code)
);
CREATE INDEX ix_locations_parent ON org.locations (parent_id);

CREATE TABLE org.providers (
  id            uuid PRIMARY KEY,
  user_id       uuid UNIQUE REFERENCES iam.users(id),   -- tidak semua provider punya akun
  full_name     varchar(200) NOT NULL,
  provider_type varchar(30) NOT NULL
                CHECK (provider_type IN ('DOCTOR','NURSE','MIDWIFE','PHARMACIST','ANALYST','RADIOGRAPHER','OTHER')),
  license_no    varchar(100),                           -- STR/SIP
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at    timestamptz NOT NULL DEFAULT now(),
  updated_at    timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE org.provider_locations (
  provider_id uuid NOT NULL REFERENCES org.providers(id) ON DELETE CASCADE,
  location_id uuid NOT NULL REFERENCES org.locations(id) ON DELETE CASCADE,
  PRIMARY KEY (provider_id, location_id)
);

-- Cakupan akses user per fasilitas (dipakai untuk filter multi-cabang)
CREATE TABLE iam.user_facilities (
  user_id     uuid NOT NULL REFERENCES iam.users(id) ON DELETE CASCADE,
  facility_id uuid NOT NULL REFERENCES org.facilities(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, facility_id)
);
```

### 6.4 `004_master.sql` — terminologi & penomoran

```sql
CREATE TABLE master.code_systems (
  id      uuid PRIMARY KEY,
  code    varchar(50) NOT NULL UNIQUE,    -- ICD10, ICD9CM, LOINC, KFA, SNOMED, INTERNAL_SEX ...
  name    varchar(200) NOT NULL,
  version varchar(50),
  status  varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

CREATE TABLE master.codes (
  id             uuid PRIMARY KEY,
  code_system_id uuid NOT NULL REFERENCES master.code_systems(id),
  code           varchar(100) NOT NULL,
  display        varchar(500) NOT NULL,
  status         varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE','RETIRED')),
  valid_from     date,
  valid_until    date,
  UNIQUE (code_system_id, code)
);
CREATE INDEX ix_codes_display_trgm ON master.codes USING gin (display gin_trgm_ops);

CREATE TABLE master.code_mappings (
  id             uuid PRIMARY KEY,
  source_code_id uuid NOT NULL REFERENCES master.codes(id),
  target_code_id uuid NOT NULL REFERENCES master.codes(id),
  relationship   varchar(30) NOT NULL DEFAULT 'EQUIVALENT'
                 CHECK (relationship IN ('EQUIVALENT','BROADER','NARROWER','RELATED')),
  UNIQUE (source_code_id, target_code_id)
);

-- Generator nomor bisnis (lihat bagian 14)
CREATE TABLE master.number_sequences (
  scope      varchar(150) PRIMARY KEY,    -- contoh: RS01:REG:20261003
  last_value bigint NOT NULL DEFAULT 0,
  updated_at timestamptz NOT NULL DEFAULT now()
);
```

### 6.5 `005_audit_integration.sql`

```sql
-- Audit: append-only, dipartisi per bulan
CREATE TABLE audit.audit_logs (
  id          uuid NOT NULL,
  occurred_at timestamptz NOT NULL DEFAULT now(),
  actor_id    uuid,                         -- iam.users.id (tanpa FK agar log tetap utuh)
  facility_id uuid,
  action      varchar(50) NOT NULL,         -- CREATE | UPDATE | READ | SIGN | VOID | LOGIN ...
  entity_type varchar(100) NOT NULL,
  entity_id   uuid,
  patient_id  uuid,                         -- memudahkan "siapa membuka rekam medis pasien X"
  before_data jsonb,
  after_data  jsonb,
  ip_address  inet,
  request_id  varchar(100),
  PRIMARY KEY (id, occurred_at)
) PARTITION BY RANGE (occurred_at);

CREATE TABLE audit.audit_logs_default PARTITION OF audit.audit_logs DEFAULT;
CREATE TABLE audit.audit_logs_2026_10 PARTITION OF audit.audit_logs
  FOR VALUES FROM ('2026-10-01') TO ('2026-11-01');

CREATE INDEX ix_audit_entity  ON audit.audit_logs (entity_type, entity_id, occurred_at DESC);
CREATE INDEX ix_audit_patient ON audit.audit_logs (patient_id, occurred_at DESC) WHERE patient_id IS NOT NULL;
CREATE INDEX ix_audit_actor   ON audit.audit_logs (actor_id, occurred_at DESC);

CREATE TRIGGER trg_audit_immutable BEFORE UPDATE OR DELETE ON audit.audit_logs
  FOR EACH ROW EXECUTE FUNCTION public.forbid_mutation();

CREATE TABLE integration.systems (
  id     uuid PRIMARY KEY,
  code   varchar(50) NOT NULL UNIQUE,   -- SATUSEHAT, BPJS_VCLAIM, PACS, LIS
  name   varchar(150) NOT NULL,
  status varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

-- Pemetaan ID internal ↔ ID eksternal
CREATE TABLE integration.external_identifiers (
  id          uuid PRIMARY KEY,
  system_id   uuid NOT NULL REFERENCES integration.systems(id),
  entity_type varchar(100) NOT NULL,    -- patient, encounter, practitioner ...
  entity_id   uuid NOT NULL,
  external_id varchar(200) NOT NULL,
  created_at  timestamptz NOT NULL DEFAULT now(),
  UNIQUE (system_id, entity_type, external_id),
  UNIQUE (system_id, entity_type, entity_id)
);

CREATE TABLE integration.outbox_events (
  id              uuid PRIMARY KEY,
  aggregate_type  varchar(100) NOT NULL,
  aggregate_id    uuid NOT NULL,
  event_type      varchar(100) NOT NULL,
  payload         jsonb NOT NULL,
  idempotency_key varchar(200) NOT NULL UNIQUE,
  status          varchar(20) NOT NULL DEFAULT 'PENDING'
                  CHECK (status IN ('PENDING','PROCESSING','DONE','FAILED','DEAD')),
  attempts        integer NOT NULL DEFAULT 0,
  next_attempt_at timestamptz NOT NULL DEFAULT now(),
  last_error      text,
  created_at      timestamptz NOT NULL DEFAULT now(),
  processed_at    timestamptz
);
CREATE INDEX ix_outbox_pending ON integration.outbox_events (next_attempt_at)
  WHERE status IN ('PENDING','FAILED');
```

---

## 7. DDL Fase 2 — Pasien & Care

### 7.1 `006_payers_patient.sql`

```sql
-- Dipindah ke fase 2 karena registrations membutuhkan FK ke payer
CREATE TABLE insurance.payers (
  id         uuid PRIMARY KEY,
  code       varchar(50) NOT NULL UNIQUE,
  name       varchar(250) NOT NULL,
  payer_type varchar(30) NOT NULL CHECK (payer_type IN ('SELF','BPJS','INSURANCE','CORPORATE')),
  status     varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE patient.patients (
  id                uuid PRIMARY KEY,
  medical_record_no varchar(50) NOT NULL UNIQUE,
  status            varchar(30) NOT NULL DEFAULT 'ACTIVE'
                    CHECK (status IN ('ACTIVE','INACTIVE','DECEASED','MERGED')),
  merged_into_id    uuid REFERENCES patient.patients(id),
  full_name         varchar(200) NOT NULL,
  normalized_name   varchar(200) NOT NULL,        -- lower + trim + tanpa gelar/tanda baca
  birth_date        date,
  birth_place       varchar(150),
  sex_code          varchar(30),
  nationality_code  varchar(30),
  marital_status    varchar(30),
  deceased_at       timestamptz,
  created_by        uuid REFERENCES iam.users(id),
  created_at        timestamptz NOT NULL DEFAULT now(),
  updated_at        timestamptz NOT NULL DEFAULT now(),
  version           integer NOT NULL DEFAULT 1,
  CHECK ((status = 'MERGED') = (merged_into_id IS NOT NULL)),
  CHECK (merged_into_id IS NULL OR merged_into_id <> id)
);
CREATE INDEX ix_patients_name_trgm ON patient.patients USING gin (normalized_name gin_trgm_ops);
CREATE INDEX ix_patients_birth ON patient.patients (birth_date);

CREATE TABLE patient.identifiers (
  id               uuid PRIMARY KEY,
  patient_id       uuid NOT NULL REFERENCES patient.patients(id),
  identifier_type  varchar(50) NOT NULL,           -- NIK, BPJS, PASSPORT, SATUSEHAT, OTHER
  identifier_value varchar(150) NOT NULL,
  normalized_value varchar(150) NOT NULL,
  issuer           varchar(200),
  is_primary       boolean NOT NULL DEFAULT false,
  valid_from       date,
  valid_until      date,
  created_at       timestamptz NOT NULL DEFAULT now()
);
-- Unik per (tipe, penerbit, nilai) — bukan global
CREATE UNIQUE INDEX ux_identifiers ON patient.identifiers
  (identifier_type, COALESCE(issuer, ''), normalized_value);
CREATE INDEX ix_identifiers_patient ON patient.identifiers (patient_id);

CREATE TABLE patient.contacts (
  id                uuid PRIMARY KEY,
  patient_id        uuid NOT NULL REFERENCES patient.patients(id),
  contact_type      varchar(30) NOT NULL,
  name              varchar(200) NOT NULL,
  relationship_code varchar(50),
  phone             varchar(50),
  email             varchar(255),
  is_emergency      boolean NOT NULL DEFAULT false
);
CREATE INDEX ix_contacts_patient ON patient.contacts (patient_id);

CREATE TABLE patient.addresses (
  id            uuid PRIMARY KEY,
  patient_id    uuid NOT NULL REFERENCES patient.patients(id),
  address_type  varchar(30) NOT NULL DEFAULT 'HOME',
  is_primary    boolean NOT NULL DEFAULT false,
  line1         text NOT NULL,
  line2         text,
  village_code  varchar(50),
  district_code varchar(50),
  city_code     varchar(50),
  province_code varchar(50),
  postal_code   varchar(20),
  country_code  varchar(10) NOT NULL DEFAULT 'ID',
  valid_from    date,
  valid_until   date
);
CREATE UNIQUE INDEX ux_addresses_primary ON patient.addresses (patient_id) WHERE is_primary;

CREATE TABLE patient.photo_refs (
  id           uuid PRIMARY KEY,
  patient_id   uuid NOT NULL REFERENCES patient.patients(id),
  object_key   text NOT NULL,
  content_type varchar(100) NOT NULL,
  checksum     varchar(128),
  created_at   timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE patient.merge_events (
  id                  uuid PRIMARY KEY,
  survivor_patient_id uuid NOT NULL REFERENCES patient.patients(id),
  merged_patient_id   uuid NOT NULL REFERENCES patient.patients(id),
  reason              text NOT NULL,
  performed_by        uuid NOT NULL REFERENCES iam.users(id),
  performed_at        timestamptz NOT NULL DEFAULT now(),
  CHECK (survivor_patient_id <> merged_patient_id)
);

CREATE TRIGGER trg_patients_upd BEFORE UPDATE ON patient.patients
  FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
```

### 7.2 `007_care.sql`

```sql
CREATE TABLE care.service_lines (
  id            uuid PRIMARY KEY,
  code          varchar(50) NOT NULL UNIQUE,   -- OUTPATIENT, INPATIENT, EMERGENCY, MCU, DAYCARE ...
  name          varchar(150) NOT NULL,
  service_type  varchar(50) NOT NULL,
  workflow_code varchar(100),
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

CREATE TABLE care.registrations (
  id                uuid PRIMARY KEY,
  registration_no   varchar(50) NOT NULL UNIQUE,
  patient_id        uuid NOT NULL REFERENCES patient.patients(id),
  facility_id       uuid NOT NULL REFERENCES org.facilities(id),
  service_line_id   uuid NOT NULL REFERENCES care.service_lines(id),
  location_id       uuid REFERENCES org.locations(id),
  payer_id          uuid REFERENCES insurance.payers(id),
  referral_id       uuid,                      -- FK ditambahkan saat care.referrals dibuat
  priority_code     varchar(50),
  registered_at     timestamptz NOT NULL DEFAULT now(),
  registration_date date NOT NULL,
  status            varchar(30) NOT NULL DEFAULT 'REGISTERED'
                    CHECK (status IN ('REGISTERED','CHECKED_IN','COMPLETED','CANCELLED')),
  cancel_reason     text,
  created_by        uuid NOT NULL REFERENCES iam.users(id),
  created_at        timestamptz NOT NULL DEFAULT now(),
  updated_at        timestamptz NOT NULL DEFAULT now(),
  version           integer NOT NULL DEFAULT 1
);
CREATE INDEX ix_reg_patient ON care.registrations (patient_id, registered_at DESC);
CREATE INDEX ix_reg_facility_date ON care.registrations (facility_id, registration_date, service_line_id);

CREATE TABLE care.episodes (
  id              uuid PRIMARY KEY,
  episode_no      varchar(50) NOT NULL UNIQUE,
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  service_line_id uuid NOT NULL REFERENCES care.service_lines(id),
  registration_id uuid REFERENCES care.registrations(id),
  started_at      timestamptz NOT NULL,
  ended_at        timestamptz,
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE'
                  CHECK (status IN ('ACTIVE','ON_HOLD','FINISHED','CANCELLED')),
  priority_code   varchar(50),
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now(),
  version         integer NOT NULL DEFAULT 1,
  CHECK (ended_at IS NULL OR ended_at >= started_at)
);
CREATE INDEX ix_episodes_patient ON care.episodes (patient_id, started_at DESC);

CREATE TABLE care.encounters (
  id              uuid PRIMARY KEY,
  encounter_no    varchar(50) NOT NULL UNIQUE,
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  episode_id      uuid NOT NULL REFERENCES care.episodes(id),
  registration_id uuid REFERENCES care.registrations(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  location_id     uuid REFERENCES org.locations(id),
  provider_id     uuid REFERENCES org.providers(id),
  encounter_type  varchar(50) NOT NULL,        -- CONSULTATION, TRIAGE, WARD_ROUND, PROCEDURE ...
  class_code      varchar(50) NOT NULL,        -- AMB, EMER, IMP (selaras HL7 ActEncounterCode)
  started_at      timestamptz NOT NULL,
  ended_at        timestamptz,
  status          varchar(30) NOT NULL DEFAULT 'PLANNED'
                  CHECK (status IN ('PLANNED','IN_PROGRESS','ON_HOLD','FINISHED','CANCELLED')),
  priority_code   varchar(50),
  created_by      uuid REFERENCES iam.users(id),
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now(),
  version         integer NOT NULL DEFAULT 1,
  CHECK (ended_at IS NULL OR ended_at >= started_at)
);
CREATE INDEX ix_enc_patient ON care.encounters (patient_id, started_at DESC);
CREATE INDEX ix_enc_episode ON care.encounters (episode_id);
CREATE INDEX ix_enc_active ON care.encounters (facility_id, location_id, started_at)
  WHERE status IN ('PLANNED','IN_PROGRESS');

CREATE TABLE care.encounter_participants (
  encounter_id     uuid NOT NULL REFERENCES care.encounters(id) ON DELETE CASCADE,
  provider_id      uuid NOT NULL REFERENCES org.providers(id),
  participant_role varchar(50) NOT NULL,
  started_at       timestamptz,
  ended_at         timestamptz,
  PRIMARY KEY (encounter_id, provider_id, participant_role)
);

CREATE TABLE care.appointments (
  id              uuid PRIMARY KEY,
  appointment_no  varchar(50) NOT NULL UNIQUE,
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  location_id     uuid NOT NULL REFERENCES org.locations(id),
  provider_id     uuid REFERENCES org.providers(id),
  service_line_id uuid NOT NULL REFERENCES care.service_lines(id),
  scheduled_start timestamptz NOT NULL,
  scheduled_end   timestamptz,
  status          varchar(30) NOT NULL DEFAULT 'BOOKED'
                  CHECK (status IN ('BOOKED','ARRIVED','FULFILLED','NO_SHOW','CANCELLED')),
  source          varchar(50) NOT NULL DEFAULT 'FRONT_DESK',
  registration_id uuid REFERENCES care.registrations(id),
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX ix_appt_provider_time ON care.appointments (provider_id, scheduled_start);
CREATE INDEX ix_appt_patient ON care.appointments (patient_id, scheduled_start DESC);

-- Nomor antrean bukan identitas registrasi
CREATE TABLE care.queue_tickets (
  id              uuid PRIMARY KEY,
  queue_no        varchar(30) NOT NULL,
  registration_id uuid NOT NULL REFERENCES care.registrations(id),
  location_id     uuid NOT NULL REFERENCES org.locations(id),
  service_code    varchar(50) NOT NULL,
  issued_at       timestamptz NOT NULL DEFAULT now(),
  called_at       timestamptz,
  served_at       timestamptz,
  status          varchar(30) NOT NULL DEFAULT 'WAITING'
                  CHECK (status IN ('WAITING','CALLED','SERVING','DONE','SKIPPED','CANCELLED')),
  priority_score  integer NOT NULL DEFAULT 0
);
CREATE INDEX ix_queue_active ON care.queue_tickets (location_id, priority_score DESC, issued_at)
  WHERE status IN ('WAITING','CALLED');

CREATE TRIGGER trg_reg_upd  BEFORE UPDATE ON care.registrations FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_ep_upd   BEFORE UPDATE ON care.episodes      FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_enc_upd  BEFORE UPDATE ON care.encounters    FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_appt_upd BEFORE UPDATE ON care.appointments  FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
```

---

## 8. DDL Fase 3 — Klinis

### 8.1 `008_clinical.sql`

```sql
CREATE TABLE clinical.diagnoses (
  id             uuid PRIMARY KEY,
  patient_id     uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id   uuid NOT NULL REFERENCES care.encounters(id),
  code_id        uuid NOT NULL REFERENCES master.codes(id),     -- ICD-10
  diagnosis_type varchar(30) NOT NULL CHECK (diagnosis_type IN ('PRIMARY','SECONDARY','COMPLICATION','DIFFERENTIAL')),
  clinical_notes text,
  diagnosed_at   timestamptz NOT NULL,
  diagnosed_by   uuid NOT NULL REFERENCES org.providers(id),
  status         varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','RESOLVED','ENTERED_IN_ERROR')),
  created_at     timestamptz NOT NULL DEFAULT now(),
  updated_at     timestamptz NOT NULL DEFAULT now()
);
-- Satu diagnosis utama aktif per encounter
CREATE UNIQUE INDEX ux_dx_primary ON clinical.diagnoses (encounter_id)
  WHERE diagnosis_type = 'PRIMARY' AND status = 'ACTIVE';
CREATE INDEX ix_dx_patient ON clinical.diagnoses (patient_id, diagnosed_at DESC);

CREATE TABLE clinical.vital_signs (
  id             uuid PRIMARY KEY,
  patient_id     uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id   uuid NOT NULL REFERENCES care.encounters(id),
  observed_at    timestamptz NOT NULL,                 -- waktu klinis
  systolic       smallint,
  diastolic      smallint,
  heart_rate     smallint,
  resp_rate      smallint,
  temperature_c  numeric(4,1),
  spo2           smallint CHECK (spo2 BETWEEN 0 AND 100),
  weight_kg      numeric(6,2),
  height_cm      numeric(5,1),
  recorded_by    uuid NOT NULL REFERENCES iam.users(id),
  status         varchar(30) NOT NULL DEFAULT 'VALID' CHECK (status IN ('VALID','ENTERED_IN_ERROR')),
  created_at     timestamptz NOT NULL DEFAULT now()   -- waktu pencatatan
);
CREATE INDEX ix_vs_enc ON clinical.vital_signs (encounter_id, observed_at DESC);
CREATE INDEX ix_vs_patient ON clinical.vital_signs (patient_id, observed_at DESC);

-- Dokumen klinis (SOAP, resume medis, asesmen) — versi dan koreksi dikontrol
CREATE TABLE clinical.documents (
  id                 uuid PRIMARY KEY,
  document_no        varchar(50) NOT NULL UNIQUE,
  patient_id         uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id       uuid NOT NULL REFERENCES care.encounters(id),
  document_type      varchar(50) NOT NULL,             -- SOAP, DISCHARGE_SUMMARY, ASSESSMENT ...
  author_id          uuid NOT NULL REFERENCES org.providers(id),
  status             varchar(30) NOT NULL DEFAULT 'DRAFT'
                     CHECK (status IN ('DRAFT','SIGNED','AMENDED','VOIDED')),
  current_version_id uuid,                              -- FK ditambahkan setelah document_versions ada
  signed_at          timestamptz,
  signed_by          uuid REFERENCES iam.users(id),
  created_at         timestamptz NOT NULL DEFAULT now(),
  updated_at         timestamptz NOT NULL DEFAULT now(),
  version            integer NOT NULL DEFAULT 1,
  CHECK (status NOT IN ('SIGNED','AMENDED') OR signed_at IS NOT NULL)
);
CREATE INDEX ix_doc_enc ON clinical.documents (encounter_id, document_type);

CREATE TABLE clinical.document_versions (
  id           uuid PRIMARY KEY,
  document_id  uuid NOT NULL REFERENCES clinical.documents(id),
  version_no   integer NOT NULL,
  content      jsonb NOT NULL,
  content_hash varchar(128) NOT NULL,
  created_by   uuid NOT NULL REFERENCES iam.users(id),
  created_at   timestamptz NOT NULL DEFAULT now(),
  UNIQUE (document_id, version_no)
);
-- Versi adalah append-only: koreksi = versi baru, bukan UPDATE
CREATE TRIGGER trg_docver_immutable BEFORE UPDATE OR DELETE ON clinical.document_versions
  FOR EACH ROW EXECUTE FUNCTION public.forbid_mutation();

ALTER TABLE clinical.documents
  ADD CONSTRAINT fk_doc_current_version
  FOREIGN KEY (current_version_id) REFERENCES clinical.document_versions(id)
  DEFERRABLE INITIALLY DEFERRED;

CREATE TABLE clinical.amendments (
  id              uuid PRIMARY KEY,
  document_id     uuid NOT NULL REFERENCES clinical.documents(id),
  from_version_id uuid NOT NULL REFERENCES clinical.document_versions(id),
  to_version_id   uuid NOT NULL REFERENCES clinical.document_versions(id),
  reason          text NOT NULL,
  requested_by    uuid NOT NULL REFERENCES iam.users(id),
  approved_by     uuid REFERENCES iam.users(id),
  created_at      timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE clinical.allergies (
  id           uuid PRIMARY KEY,
  patient_id   uuid NOT NULL REFERENCES patient.patients(id),
  substance_code_id uuid REFERENCES master.codes(id),
  substance_text varchar(300) NOT NULL,
  reaction     text,
  severity     varchar(30) CHECK (severity IN ('MILD','MODERATE','SEVERE')),
  status       varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE','ENTERED_IN_ERROR')),
  recorded_by  uuid NOT NULL REFERENCES iam.users(id),
  recorded_at  timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX ix_allergy_patient ON clinical.allergies (patient_id) WHERE status = 'ACTIVE';

-- Header order generik; detail spesifik (lab, radiologi, resep) di modul masing-masing
CREATE TABLE clinical.orders (
  id           uuid PRIMARY KEY,
  order_no     varchar(50) NOT NULL UNIQUE,
  patient_id   uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id uuid NOT NULL REFERENCES care.encounters(id),
  facility_id  uuid NOT NULL REFERENCES org.facilities(id),
  ordered_by   uuid NOT NULL REFERENCES org.providers(id),
  order_type   varchar(50) NOT NULL
               CHECK (order_type IN ('LAB','RADIOLOGY','MEDICATION','PROCEDURE','NURSING','DIET','REFERRAL')),
  priority_code varchar(50) NOT NULL DEFAULT 'ROUTINE',
  ordered_at   timestamptz NOT NULL DEFAULT now(),
  status       varchar(30) NOT NULL DEFAULT 'DRAFT'
               CHECK (status IN ('DRAFT','ACTIVE','IN_PROGRESS','COMPLETED','CANCELLED')),
  cancel_reason text,
  notes        text,
  created_at   timestamptz NOT NULL DEFAULT now(),
  updated_at   timestamptz NOT NULL DEFAULT now(),
  version      integer NOT NULL DEFAULT 1
);
CREATE INDEX ix_orders_enc ON clinical.orders (encounter_id, order_type);
CREATE INDEX ix_orders_worklist ON clinical.orders (facility_id, order_type, ordered_at)
  WHERE status IN ('ACTIVE','IN_PROGRESS');

CREATE TABLE clinical.order_items (
  id              uuid PRIMARY KEY,
  order_id        uuid NOT NULL REFERENCES clinical.orders(id),
  service_code_id uuid REFERENCES master.codes(id),
  catalog_item_id uuid,                              -- FK ke master.service_catalog (fase 5)
  sequence_no     integer NOT NULL,
  qty             numeric(20,6),
  unit_code       varchar(50),
  instructions    text,
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE'
                  CHECK (status IN ('ACTIVE','IN_PROGRESS','COMPLETED','CANCELLED')),
  UNIQUE (order_id, sequence_no)
);

CREATE TRIGGER trg_doc_upd   BEFORE UPDATE ON clinical.documents FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_order_upd BEFORE UPDATE ON clinical.orders    FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
CREATE TRIGGER trg_dx_upd    BEFORE UPDATE ON clinical.diagnoses FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
```

---

## 9. DDL Fase 4 — Farmasi & Inventori

### 9.1 `009_inventory_pharmacy.sql`

```sql
CREATE TABLE inventory.items (
  id             uuid PRIMARY KEY,
  code           varchar(100) NOT NULL UNIQUE,
  name           varchar(300) NOT NULL,
  item_type      varchar(50) NOT NULL CHECK (item_type IN ('MEDICATION','CONSUMABLE','DEVICE','OTHER')),
  base_unit_code varchar(50) NOT NULL,
  status         varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  created_at     timestamptz NOT NULL DEFAULT now(),
  updated_at     timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE inventory.warehouses (
  id          uuid PRIMARY KEY,
  facility_id uuid NOT NULL REFERENCES org.facilities(id),
  location_id uuid REFERENCES org.locations(id),
  code        varchar(50) NOT NULL UNIQUE,
  name        varchar(200) NOT NULL,
  status      varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

CREATE TABLE inventory.item_lots (
  id           uuid PRIMARY KEY,
  item_id      uuid NOT NULL REFERENCES inventory.items(id),
  lot_no       varchar(100) NOT NULL,
  expiry_date  date,
  manufacturer varchar(200),
  status       varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','QUARANTINE','EXPIRED')),
  UNIQUE (item_id, lot_no)
);

-- PERBAIKAN blueprint: PK surrogate karena lot_id nullable
CREATE TABLE inventory.stock_balances (
  id           uuid PRIMARY KEY,
  warehouse_id uuid NOT NULL REFERENCES inventory.warehouses(id),
  item_id      uuid NOT NULL REFERENCES inventory.items(id),
  lot_id       uuid REFERENCES inventory.item_lots(id),
  qty_on_hand  numeric(20,6) NOT NULL DEFAULT 0 CHECK (qty_on_hand >= 0),
  qty_reserved numeric(20,6) NOT NULL DEFAULT 0 CHECK (qty_reserved >= 0),
  updated_at   timestamptz NOT NULL DEFAULT now(),
  CHECK (qty_reserved <= qty_on_hand)
);
CREATE UNIQUE INDEX ux_stock_bal ON inventory.stock_balances
  (warehouse_id, item_id, COALESCE(lot_id, '00000000-0000-0000-0000-000000000000'::uuid));

-- Ledger: sumber kebenaran historis stok
CREATE TABLE inventory.stock_movements (
  id             uuid PRIMARY KEY,
  warehouse_id   uuid NOT NULL REFERENCES inventory.warehouses(id),
  item_id        uuid NOT NULL REFERENCES inventory.items(id),
  lot_id         uuid REFERENCES inventory.item_lots(id),
  movement_type  varchar(50) NOT NULL
                 CHECK (movement_type IN ('PURCHASE_IN','TRANSFER_IN','TRANSFER_OUT','DISPENSE_OUT',
                                          'RETURN_IN','ADJUSTMENT_IN','ADJUSTMENT_OUT','EXPIRED_OUT','DAMAGE_OUT')),
  direction      smallint NOT NULL CHECK (direction IN (1, -1)),   -- +1 masuk, -1 keluar
  quantity       numeric(20,6) NOT NULL CHECK (quantity > 0),
  unit_cost      numeric(18,2),
  reference_type varchar(100),
  reference_id   uuid,
  occurred_at    timestamptz NOT NULL DEFAULT now(),
  created_by     uuid NOT NULL REFERENCES iam.users(id),
  notes          text,
  created_at     timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX ix_mv_item ON inventory.stock_movements (warehouse_id, item_id, occurred_at DESC);
CREATE INDEX ix_mv_ref ON inventory.stock_movements (reference_type, reference_id);
CREATE TRIGGER trg_mv_immutable BEFORE UPDATE OR DELETE ON inventory.stock_movements
  FOR EACH ROW EXECUTE FUNCTION public.forbid_mutation();

CREATE TABLE pharmacy.medications (
  id            uuid PRIMARY KEY,
  item_id       uuid NOT NULL UNIQUE REFERENCES inventory.items(id),
  code_id       uuid REFERENCES master.codes(id),          -- KFA
  generic_name  varchar(300) NOT NULL,
  dosage_form   varchar(100),
  strength      varchar(100),
  is_narcotic   boolean NOT NULL DEFAULT false,
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

CREATE TABLE pharmacy.prescriptions (
  id              uuid PRIMARY KEY,
  prescription_no varchar(50) NOT NULL UNIQUE,
  order_id        uuid UNIQUE REFERENCES clinical.orders(id),
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id    uuid NOT NULL REFERENCES care.encounters(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  prescriber_id   uuid NOT NULL REFERENCES org.providers(id),
  prescribed_at   timestamptz NOT NULL DEFAULT now(),
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE'
                  CHECK (status IN ('DRAFT','ACTIVE','PARTIALLY_DISPENSED','DISPENSED','CANCELLED')),
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE pharmacy.prescription_items (
  id              uuid PRIMARY KEY,
  prescription_id uuid NOT NULL REFERENCES pharmacy.prescriptions(id),
  medication_id   uuid NOT NULL REFERENCES pharmacy.medications(id),
  quantity        numeric(20,6) NOT NULL CHECK (quantity > 0),
  unit_code       varchar(50) NOT NULL,
  dose_text       varchar(300),
  frequency_text  varchar(200),
  route_code      varchar(50),
  duration_days   integer,
  instructions    text,
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE'
                  CHECK (status IN ('ACTIVE','DISPENSED','CANCELLED'))
);

-- Dispensing = pemenuhan resep (bisa parsial)
CREATE TABLE pharmacy.dispensations (
  id              uuid PRIMARY KEY,
  dispensation_no varchar(50) NOT NULL UNIQUE,
  prescription_id uuid NOT NULL REFERENCES pharmacy.prescriptions(id),
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id    uuid NOT NULL REFERENCES care.encounters(id),
  warehouse_id    uuid NOT NULL REFERENCES inventory.warehouses(id),
  dispensed_by    uuid NOT NULL REFERENCES org.providers(id),
  dispensed_at    timestamptz NOT NULL DEFAULT now(),
  status          varchar(30) NOT NULL DEFAULT 'COMPLETED' CHECK (status IN ('COMPLETED','RETURNED','VOIDED'))
);

CREATE TABLE pharmacy.dispensation_items (
  id                   uuid PRIMARY KEY,
  dispensation_id      uuid NOT NULL REFERENCES pharmacy.dispensations(id),
  prescription_item_id uuid NOT NULL REFERENCES pharmacy.prescription_items(id),
  lot_id               uuid REFERENCES inventory.item_lots(id),
  quantity             numeric(20,6) NOT NULL CHECK (quantity > 0),
  stock_movement_id    uuid NOT NULL REFERENCES inventory.stock_movements(id)
);
```

Pengurangan stok yang aman dari race condition (tanpa read-then-write):

```sql
UPDATE inventory.stock_balances
   SET qty_on_hand = qty_on_hand - $1, updated_at = now()
 WHERE id = $2 AND qty_on_hand - qty_reserved >= $1;
-- RowsAffected = 0  →  stok tidak cukup, batalkan transaksi
```

---

## 10. DDL Fase 5 — Billing, Finance, Insurance

### 10.1 `010_billing_finance.sql`

```sql
CREATE TABLE master.service_catalog (
  id              uuid PRIMARY KEY,
  code            varchar(100) NOT NULL UNIQUE,
  name            varchar(250) NOT NULL,
  service_type    varchar(50) NOT NULL,
  clinical_domain varchar(50),
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

ALTER TABLE clinical.order_items
  ADD CONSTRAINT fk_oi_catalog FOREIGN KEY (catalog_item_id) REFERENCES master.service_catalog(id);

-- Tarif ber-masa-berlaku; jangan menimpa tarif yang sudah dipakai transaksi
CREATE TABLE master.service_prices (
  id              uuid PRIMARY KEY,
  service_id      uuid NOT NULL REFERENCES master.service_catalog(id),
  payer_id        uuid REFERENCES insurance.payers(id),   -- NULL = tarif umum
  price           numeric(18,2) NOT NULL CHECK (price >= 0),
  currency        varchar(3) NOT NULL DEFAULT 'IDR',
  effective_from  date NOT NULL,
  effective_until date,
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE')),
  CHECK (effective_until IS NULL OR effective_until >= effective_from)
);
CREATE UNIQUE INDEX ux_price_start ON master.service_prices
  (service_id, COALESCE(payer_id, '00000000-0000-0000-0000-000000000000'::uuid), effective_from);
-- Pemeriksaan tumpang tindih periode dilakukan di Action saat menambah tarif.

CREATE TABLE billing.charges (
  id              uuid PRIMARY KEY,
  charge_no       varchar(50) NOT NULL UNIQUE,
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  encounter_id    uuid REFERENCES care.encounters(id),
  episode_id      uuid REFERENCES care.episodes(id),
  service_id      uuid REFERENCES master.service_catalog(id),
  source_type     varchar(100),
  source_id       uuid,
  quantity        numeric(20,6) NOT NULL CHECK (quantity > 0),
  unit_price      numeric(18,2) NOT NULL,
  gross_amount    numeric(18,2) NOT NULL,
  discount_amount numeric(18,2) NOT NULL DEFAULT 0,
  net_amount      numeric(18,2) NOT NULL,
  payer_id        uuid REFERENCES insurance.payers(id),
  occurred_at     timestamptz NOT NULL,
  status          varchar(30) NOT NULL DEFAULT 'PENDING'
                  CHECK (status IN ('PENDING','BILLED','VOIDED')),
  void_reason     text,
  created_at      timestamptz NOT NULL DEFAULT now(),
  CHECK (net_amount = gross_amount - discount_amount),
  CHECK ((source_type IS NULL) = (source_id IS NULL))
);
-- Kunci idempotensi: satu event sumber → satu charge aktif
CREATE UNIQUE INDEX ux_charge_source ON billing.charges (source_type, source_id)
  WHERE source_id IS NOT NULL AND status <> 'VOIDED';
CREATE INDEX ix_charge_enc ON billing.charges (encounter_id) WHERE status = 'PENDING';

CREATE TABLE billing.invoices (
  id              uuid PRIMARY KEY,
  invoice_no      varchar(50) NOT NULL UNIQUE,
  patient_id      uuid NOT NULL REFERENCES patient.patients(id),
  facility_id     uuid NOT NULL REFERENCES org.facilities(id),
  encounter_id    uuid REFERENCES care.encounters(id),
  episode_id      uuid REFERENCES care.episodes(id),
  payer_id        uuid REFERENCES insurance.payers(id),
  issued_at       timestamptz NOT NULL DEFAULT now(),
  subtotal        numeric(18,2) NOT NULL DEFAULT 0,
  discount_amount numeric(18,2) NOT NULL DEFAULT 0,
  tax_amount      numeric(18,2) NOT NULL DEFAULT 0,
  total_amount    numeric(18,2) NOT NULL DEFAULT 0,
  paid_amount     numeric(18,2) NOT NULL DEFAULT 0 CHECK (paid_amount >= 0),   -- proyeksi
  status          varchar(30) NOT NULL DEFAULT 'DRAFT'
                  CHECK (status IN ('DRAFT','ISSUED','PARTIALLY_PAID','PAID','VOIDED')),
  finalized_at    timestamptz,
  void_reason     text,
  created_by      uuid REFERENCES iam.users(id),
  updated_at      timestamptz NOT NULL DEFAULT now(),
  version         integer NOT NULL DEFAULT 1,
  CHECK (total_amount = subtotal - discount_amount + tax_amount),
  CHECK (paid_amount <= total_amount OR status = 'VOIDED')
);
CREATE INDEX ix_inv_patient ON billing.invoices (patient_id, issued_at DESC);
CREATE INDEX ix_inv_open ON billing.invoices (facility_id, issued_at) WHERE status IN ('ISSUED','PARTIALLY_PAID');

CREATE TABLE billing.invoice_items (
  id         uuid PRIMARY KEY,
  invoice_id uuid NOT NULL REFERENCES billing.invoices(id),
  charge_id  uuid NOT NULL UNIQUE REFERENCES billing.charges(id),   -- satu charge hanya di satu invoice
  quantity   numeric(20,6) NOT NULL,
  unit_price numeric(18,2) NOT NULL,
  amount     numeric(18,2) NOT NULL
);

CREATE TABLE finance.payments (
  id             uuid PRIMARY KEY,
  payment_no     varchar(50) NOT NULL UNIQUE,
  facility_id    uuid NOT NULL REFERENCES org.facilities(id),
  patient_id     uuid REFERENCES patient.patients(id),
  payer_id       uuid REFERENCES insurance.payers(id),
  payment_method varchar(50) NOT NULL,
  amount         numeric(18,2) NOT NULL CHECK (amount > 0),
  allocated_amount numeric(18,2) NOT NULL DEFAULT 0,                 -- proyeksi
  currency       varchar(3) NOT NULL DEFAULT 'IDR',
  paid_at        timestamptz NOT NULL DEFAULT now(),
  reference_no   varchar(150),
  status         varchar(30) NOT NULL DEFAULT 'POSTED' CHECK (status IN ('POSTED','VOIDED')),
  received_by    uuid NOT NULL REFERENCES iam.users(id),
  CHECK (allocated_amount <= amount)
);

CREATE TABLE finance.payment_allocations (
  payment_id       uuid NOT NULL REFERENCES finance.payments(id),
  invoice_id       uuid NOT NULL REFERENCES billing.invoices(id),
  allocated_amount numeric(18,2) NOT NULL CHECK (allocated_amount > 0),
  PRIMARY KEY (payment_id, invoice_id)
);

CREATE TABLE finance.adjustments (
  id              uuid PRIMARY KEY,
  adjustment_no   varchar(50) NOT NULL UNIQUE,
  invoice_id      uuid NOT NULL REFERENCES billing.invoices(id),
  adjustment_type varchar(50) NOT NULL CHECK (adjustment_type IN ('CREDIT_NOTE','DEBIT_NOTE','WRITE_OFF','DISCOUNT')),
  amount          numeric(18,2) NOT NULL,
  reason          text NOT NULL,
  status          varchar(30) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING','APPROVED','REJECTED')),
  approved_by     uuid REFERENCES iam.users(id),
  created_by      uuid NOT NULL REFERENCES iam.users(id),
  created_at      timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE insurance.coverages (
  id            uuid PRIMARY KEY,
  patient_id    uuid NOT NULL REFERENCES patient.patients(id),
  payer_id      uuid NOT NULL REFERENCES insurance.payers(id),
  member_no     varchar(100) NOT NULL,
  class_code    varchar(30),
  valid_from    date,
  valid_until   date,
  is_primary    boolean NOT NULL DEFAULT false,
  status        varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','INACTIVE'))
);

CREATE TABLE insurance.claims (
  id                 uuid PRIMARY KEY,
  claim_no           varchar(100) NOT NULL UNIQUE,
  patient_id         uuid NOT NULL REFERENCES patient.patients(id),
  encounter_id       uuid REFERENCES care.encounters(id),
  payer_id           uuid NOT NULL REFERENCES insurance.payers(id),
  invoice_id         uuid REFERENCES billing.invoices(id),
  claimed_amount     numeric(18,2) NOT NULL,
  approved_amount    numeric(18,2),
  submitted_at       timestamptz,
  accepted_at        timestamptz,
  rejected_at        timestamptz,
  status             varchar(30) NOT NULL DEFAULT 'DRAFT'
                     CHECK (status IN ('DRAFT','SUBMITTED','ACCEPTED','REJECTED','PAID')),
  external_reference varchar(150)
);

CREATE TABLE insurance.claim_events (
  id          uuid PRIMARY KEY,
  claim_id    uuid NOT NULL REFERENCES insurance.claims(id),
  event_type  varchar(50) NOT NULL,
  status      varchar(30) NOT NULL,
  message     text,
  payload     jsonb,
  occurred_at timestamptz NOT NULL DEFAULT now(),
  created_by  uuid REFERENCES iam.users(id)
);

CREATE TRIGGER trg_inv_upd BEFORE UPDATE ON billing.invoices FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();
```

Aturan transaksi alokasi pembayaran: `SELECT ... FROM finance.payments WHERE id = $1 FOR UPDATE`, lalu `SELECT ... FROM billing.invoices WHERE id = ANY($2) FOR UPDATE`, validasi `allocated_amount + baru ≤ amount` dan `paid_amount + baru ≤ total_amount`, insert alokasi, update proyeksi, ubah status invoice. Semua dalam satu transaksi.

---

## 11. DDL Fase 6 — Rawat Inap

### 11.1 `011_inpatient.sql`

```sql
CREATE TABLE care.admissions (
  id                    uuid PRIMARY KEY,
  episode_id            uuid NOT NULL UNIQUE REFERENCES care.episodes(id),
  admission_no          varchar(50) NOT NULL UNIQUE,
  admitted_at           timestamptz NOT NULL,
  admission_source      varchar(50) NOT NULL,                -- IGD, POLI, RUJUKAN, LANGSUNG
  admitting_provider_id uuid REFERENCES org.providers(id),
  status                varchar(30) NOT NULL DEFAULT 'ADMITTED'
                        CHECK (status IN ('ADMITTED','DISCHARGE_PLANNED','DISCHARGED','CANCELLED')),
  discharge_at          timestamptz,
  discharge_disposition varchar(50),                         -- PULANG, DIRUJUK, MENINGGAL, APS
  CHECK (discharge_at IS NULL OR discharge_at >= admitted_at)
);

CREATE TABLE care.bed_assignments (
  id              uuid PRIMARY KEY,
  admission_id    uuid NOT NULL REFERENCES care.admissions(id),
  bed_location_id uuid NOT NULL REFERENCES org.locations(id),
  assigned_at     timestamptz NOT NULL,
  released_at     timestamptz,
  assignment_type varchar(50) NOT NULL DEFAULT 'PRIMARY',
  status          varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','RELEASED','CANCELLED')),
  CHECK (released_at IS NULL OR released_at > assigned_at),
  -- Satu bed tidak boleh ditempati dua pasien pada periode yang beririsan
  CONSTRAINT ex_bed_no_overlap EXCLUDE USING gist (
    bed_location_id WITH =,
    tstzrange(assigned_at, COALESCE(released_at, 'infinity'::timestamptz)) WITH &&
  ) WHERE (status <> 'CANCELLED')
);
-- Satu admission hanya punya satu bed aktif pada satu waktu
CREATE UNIQUE INDEX ux_bed_active_per_admission ON care.bed_assignments (admission_id)
  WHERE released_at IS NULL AND status = 'ACTIVE';

CREATE TABLE care.transfers (
  id               uuid PRIMARY KEY,
  admission_id     uuid NOT NULL REFERENCES care.admissions(id),
  from_location_id uuid NOT NULL REFERENCES org.locations(id),
  from_bed_id      uuid REFERENCES org.locations(id),
  to_location_id   uuid NOT NULL REFERENCES org.locations(id),
  to_bed_id        uuid REFERENCES org.locations(id),
  requested_at     timestamptz NOT NULL DEFAULT now(),
  completed_at     timestamptz,
  reason           text,
  status           varchar(30) NOT NULL DEFAULT 'REQUESTED'
                   CHECK (status IN ('REQUESTED','IN_TRANSIT','COMPLETED','CANCELLED'))
);

CREATE TABLE care.referrals (
  id               uuid PRIMARY KEY,
  referral_no      varchar(50) NOT NULL UNIQUE,
  patient_id       uuid NOT NULL REFERENCES patient.patients(id),
  direction        varchar(10) NOT NULL CHECK (direction IN ('IN','OUT')),
  from_facility    varchar(250),
  to_facility      varchar(250),
  reason           text,
  referred_at      timestamptz NOT NULL,
  status           varchar(30) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','ACCEPTED','CLOSED','CANCELLED'))
);

ALTER TABLE care.registrations
  ADD CONSTRAINT fk_reg_referral FOREIGN KEY (referral_id) REFERENCES care.referrals(id);
```

Lab, radiologi, MCU, dan nursing (fase 7+) mengikuti pola yang sama: header order domain → item → spesimen/studi → hasil/laporan, dengan hasil berstatus (`PRELIMINARY`, `FINAL`, `AMENDED`) dan koreksi lewat versi baru. Paket MCU wajib punya versi yang tidak berubah untuk episode yang sudah berjalan.

---

## 12. Struktur Proyek Go

Selaras dengan pola yang sudah dipakai (`app/modules/<modul>` berisi model + Repository, dan Action untuk use case).

```text
app/
  modules/
    shared/                  # base types, uuid, decimal helper, error domain
      ids.go                 # NewID() → uuid.NewV7()
      status.go
    patient/
      model.go               # Patient, Identifier, Address ...
      repository.go          # interface + implementasi GORM
      actions/
        register_patient.go
        merge_patient.go
        search_patient.go
    care/
      model.go
      repository.go
      actions/
        open_outpatient_visit.go   # registration + episode + encounter, 1 transaksi
        close_encounter.go
        admit_patient.go
        transfer_bed.go
    clinical/
      actions/ {record_vitals, record_diagnosis, sign_document, amend_document, place_order}.go
    pharmacy/
      actions/ {create_prescription, dispense}.go
    inventory/
      actions/ {stock_in, stock_out, adjust_stock}.go
    billing/
      actions/ {generate_charge, build_invoice, finalize_invoice, void_invoice}.go
    finance/
      actions/ {receive_payment, allocate_payment}.go
    integration/
      outbox/ {publisher.go, worker.go}
    numbering/
      generator.go
  http/controllers/...
database/
  migrations/ 001_init.sql … 011_inpatient.sql
```

Prinsip batas modul:

- Modul **tidak** mengimpor model modul lain; hanya memakai ID dan interface.
- Action lintas modul menerima dependensi lewat konstruktor (contoh: `Dispense` menerima `inventory.StockOut`, `billing.GenerateCharge`, `outbox.Publisher`).
- Controller tipis: validasi request → panggil satu Action → bentuk response.

---

## 13. Pola Kode Go

### 13.1 Dasar: ID, status, uang

```go
package shared

import "github.com/google/uuid"

func NewID() uuid.UUID { return uuid.Must(uuid.NewV7()) }
```

```go
package care

type EncounterStatus string

const (
    EncounterPlanned    EncounterStatus = "PLANNED"
    EncounterInProgress EncounterStatus = "IN_PROGRESS"
    EncounterOnHold     EncounterStatus = "ON_HOLD"
    EncounterFinished   EncounterStatus = "FINISHED"
    EncounterCancelled  EncounterStatus = "CANCELLED"
)

// Tabel transisi yang diizinkan — satu-satunya tempat aturan status didefinisikan
var encounterTransitions = map[EncounterStatus][]EncounterStatus{
    EncounterPlanned:    {EncounterInProgress, EncounterCancelled},
    EncounterInProgress: {EncounterOnHold, EncounterFinished, EncounterCancelled},
    EncounterOnHold:     {EncounterInProgress, EncounterCancelled},
}

func (s EncounterStatus) CanMoveTo(to EncounterStatus) bool {
    for _, t := range encounterTransitions[s] {
        if t == to {
            return true
        }
    }
    return false
}
```

### 13.2 Model (tabel ber-schema, tanpa soft delete)

```go
package patient

import (
    "time"

    "github.com/google/uuid"
)

type Patient struct {
    ID              uuid.UUID  `gorm:"type:uuid;primaryKey"`
    MedicalRecordNo string
    Status          string
    MergedIntoID    *uuid.UUID `gorm:"type:uuid"`
    FullName        string
    NormalizedName  string
    BirthDate       *time.Time
    BirthPlace      *string
    SexCode         *string
    Version         int
    CreatedAt       time.Time
    UpdatedAt       time.Time
}

func (Patient) TableName() string { return "patient.patients" }
```

```go
package billing

import (
    "time"

    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type Charge struct {
    ID             uuid.UUID       `gorm:"type:uuid;primaryKey"`
    ChargeNo       string
    PatientID      uuid.UUID       `gorm:"type:uuid"`
    FacilityID     uuid.UUID       `gorm:"type:uuid"`
    EncounterID    *uuid.UUID      `gorm:"type:uuid"`
    SourceType     *string
    SourceID       *uuid.UUID      `gorm:"type:uuid"`
    Quantity       decimal.Decimal `gorm:"type:numeric(20,6)"`
    UnitPrice      decimal.Decimal `gorm:"type:numeric(18,2)"`
    GrossAmount    decimal.Decimal `gorm:"type:numeric(18,2)"`
    DiscountAmount decimal.Decimal `gorm:"type:numeric(18,2)"`
    NetAmount      decimal.Decimal `gorm:"type:numeric(18,2)"`
    Status         string
    OccurredAt     time.Time
    CreatedAt      time.Time
}

func (Charge) TableName() string { return "billing.charges" }
```

> Catatan Goravel: `orm.Model` bawaan memakai `ID uint` dan `SoftDeletes`. Untuk tabel klinis/keuangan, definisikan struct sendiri seperti di atas dan isi `ID` secara eksplisit di Action/Repository (`shared.NewID()`) sebelum `Create`, sehingga tidak bergantung pada hook ORM.

### 13.3 Optimistic locking

```go
func (r *encounterRepo) UpdateStatus(tx orm.Query, e *Encounter, to EncounterStatus) error {
    res, err := tx.Model(&Encounter{}).
        Where("id = ? AND version = ?", e.ID, e.Version).
        Update(map[string]any{"status": string(to), "version": e.Version + 1})
    if err != nil {
        return err
    }
    if res.RowsAffected == 0 {
        return ErrConcurrentModification // 409 Conflict di controller
    }
    return nil
}
```

### 13.4 Action dengan transaksi tunggal: `OpenOutpatientVisit`

```go
type OpenOutpatientVisit struct {
    numbering numbering.Generator
    regRepo   RegistrationRepository
    epRepo    EpisodeRepository
    encRepo   EncounterRepository
    outbox    outbox.Publisher
}

type OpenVisitInput struct {
    PatientID  uuid.UUID
    FacilityID uuid.UUID
    LocationID uuid.UUID
    ProviderID *uuid.UUID
    PayerID    *uuid.UUID
    ActorID    uuid.UUID
}

func (a *OpenOutpatientVisit) Execute(in OpenVisitInput) (*Encounter, error) {
    var enc *Encounter
    err := facades.Orm().Transaction(func(tx orm.Query) error {
        now := time.Now().UTC()

        regNo, err := a.numbering.Next(tx, in.FacilityID, "REG", now)
        if err != nil { return err }
        epNo, err := a.numbering.Next(tx, in.FacilityID, "EPS", now)
        if err != nil { return err }
        encNo, err := a.numbering.Next(tx, in.FacilityID, "ENC", now)
        if err != nil { return err }

        reg := &Registration{ID: shared.NewID(), RegistrationNo: regNo /* ... */}
        if err := a.regRepo.Create(tx, reg); err != nil { return err }

        ep := &Episode{ID: shared.NewID(), EpisodeNo: epNo, RegistrationID: &reg.ID /* ... */}
        if err := a.epRepo.Create(tx, ep); err != nil { return err }

        enc = &Encounter{ID: shared.NewID(), EncounterNo: encNo, EpisodeID: ep.ID /* ... */}
        if err := a.encRepo.Create(tx, enc); err != nil { return err }

        return a.outbox.Publish(tx, outbox.Event{
            AggregateType:  "encounter",
            AggregateID:    enc.ID,
            EventType:      "encounter.created",
            IdempotencyKey: "encounter.created:" + enc.ID.String(),
            Payload:        map[string]any{"encounter_id": enc.ID},
        })
    })
    return enc, err
}
```

Semua repository menerima `tx orm.Query` supaya berjalan di transaksi yang sama.

### 13.5 Action lintas modul: `Dispense`

Urutan dalam **satu** transaksi:

1. Kunci dan validasi resep (`SELECT ... FOR UPDATE`, status harus `ACTIVE`/`PARTIALLY_DISPENSED`).
2. Untuk tiap item: kurangi `stock_balances` dengan `UPDATE ... WHERE qty_on_hand - qty_reserved >= ?`; jika `RowsAffected = 0`, kembalikan error stok kurang.
3. Insert `stock_movements` (`DISPENSE_OUT`, `direction = -1`).
4. Insert `dispensations` + `dispensation_items` (menautkan `stock_movement_id`).
5. Perbarui status `prescription_items` dan `prescriptions`.
6. Panggil `billing.GenerateCharge` dengan `source_type = 'dispensation_item'`, `source_id = <id>` (idempoten lewat unique index).
7. Publish outbox (`MedicationDispense` untuk SATUSEHAT).

### 13.6 Dokumen klinis: tanda tangan & koreksi

- `SignDocument`: pastikan status `DRAFT`, set `signed_at/signed_by`, status `SIGNED`, tulis audit.
- `AmendDocument`: buat baris baru di `document_versions` (`version_no + 1`), buat `amendments` dengan alasan, ubah `documents.current_version_id` dan status `AMENDED`. **Tidak ada UPDATE ke versi lama**; trigger `forbid_mutation` menjadi lapisan pengaman kedua.

---

## 14. Penomoran Bisnis

Format: `<FASILITAS>-<JENIS>-<YYYYMMDD>-<URUT 6 digit>`, contoh `RS01-REG-20261003-000001`.

```sql
INSERT INTO master.number_sequences (scope, last_value)
VALUES ($1, 1)
ON CONFLICT (scope)
DO UPDATE SET last_value = master.number_sequences.last_value + 1, updated_at = now()
RETURNING last_value;
```

```go
func (g *generator) Next(tx orm.Query, facilityID uuid.UUID, kind string, at time.Time) (string, error) {
    code := g.facilityCode(facilityID)            // cache, mis. "RS01"
    day := at.In(g.loc).Format("20060102")
    scope := fmt.Sprintf("%s:%s:%s", code, kind, day)

    var n int64
    if err := tx.Raw(upsertSQL, scope).Scan(&n); err != nil {
        return "", err
    }
    return fmt.Sprintf("%s-%s-%s-%06d", code, kind, day, n), nil
}
```

Catatan:

- Penomoran berjalan di dalam transaksi pemanggil; baris sequence terkunci sampai commit, sehingga **tidak ada nomor ganda**. Konsekuensinya, transaksi yang rollback tidak menghasilkan lubang nomor (berbeda dengan `SEQUENCE` PostgreSQL). Jika regulasi/akuntansi mengizinkan lubang, `SEQUENCE` lebih cepat.
- Nomor rekam medis (`medical_record_no`) memakai scope tanpa tanggal (`RS01:MR`).
- Format nomor tidak boleh menjadi dasar relasi; relasi selalu lewat UUID.

---

## 15. Outbox & Idempotency

**Publisher (di dalam transaksi bisnis):** insert ke `integration.outbox_events` dengan `idempotency_key` unik; `ON CONFLICT DO NOTHING` membuat publish ulang aman.

**Worker (proses terpisah/scheduler):**

```sql
WITH picked AS (
  SELECT id FROM integration.outbox_events
   WHERE status IN ('PENDING','FAILED') AND next_attempt_at <= now()
   ORDER BY next_attempt_at
   LIMIT 50
   FOR UPDATE SKIP LOCKED
)
UPDATE integration.outbox_events e
   SET status = 'PROCESSING', attempts = attempts + 1
  FROM picked WHERE e.id = picked.id
RETURNING e.*;
```

Aturan:

- Sukses → `DONE` + `processed_at`.
- Gagal → `FAILED`, `next_attempt_at = now() + backoff(attempts)`; lewat batas percobaan → `DEAD` (perlu intervensi manual, tampilkan di dashboard).
- Kirim ke sistem eksternal dengan idempotency key yang sama pada setiap retry.
- Simpan hasil pemetaan di `integration.external_identifiers`.
- Pekerjaan yang macet di `PROCESSING` terlalu lama (worker mati) dikembalikan ke `FAILED` oleh job pembersih.

---

## 16. Aturan Integritas & State Machine

### 16.1 Invarian yang harus dijaga (di DB bila mungkin, di Action bila tidak)

| Area | Invarian | Penjaga |
|---|---|---|
| Pasien | NIK/BPJS tidak ganda per penerbit; pasien `MERGED` wajib punya `merged_into_id` | unique index + CHECK |
| Encounter | `ended_at ≥ started_at`; encounter selesai tidak menerima order baru | CHECK + Action |
| Diagnosis | Hanya satu diagnosis utama aktif per encounter | partial unique index |
| Dokumen | Dokumen `SIGNED` tidak diubah; koreksi = versi baru | trigger + Action |
| Bed | Satu bed tidak dihuni dua pasien pada waktu beririsan | exclusion constraint |
| Stok | Saldo tidak negatif; setiap perubahan saldo punya movement | CHECK + satu Action |
| Charge | Satu event sumber → satu charge aktif | partial unique index |
| Invoice | `total = subtotal − diskon + pajak`; invoice final tidak diubah angkanya | CHECK + status |
| Payment | Σ alokasi ≤ jumlah payment; Σ alokasi per invoice ≤ total | `FOR UPDATE` + CHECK |
| Audit | Tidak bisa diubah/dihapus | trigger `forbid_mutation` |

### 16.2 State machine ringkas

```text
Encounter : PLANNED → IN_PROGRESS ⇄ ON_HOLD → FINISHED
                 └──────────→ CANCELLED (dari status mana pun sebelum FINISHED)

Document  : DRAFT → SIGNED → AMENDED (→ AMENDED …)        VOIDED (khusus, butuh otorisasi)

Order     : DRAFT → ACTIVE → IN_PROGRESS → COMPLETED        CANCELLED

Prescription : DRAFT → ACTIVE → PARTIALLY_DISPENSED → DISPENSED   CANCELLED

Invoice   : DRAFT → ISSUED → PARTIALLY_PAID → PAID          VOIDED (hanya jika belum ada pembayaran,
                                                              selain itu pakai credit note)

Claim     : DRAFT → SUBMITTED → ACCEPTED | REJECTED → PAID
```

Dokumentasikan setiap state machine di kode (tabel transisi seperti `encounterTransitions`) dan uji sebagai unit test.

---

## 17. Indexing & Partisi

**Prinsip:** index berdasarkan query nyata, bukan tebakan. Mulai dari yang sudah ada di DDL, lalu pantau dengan `pg_stat_statements` dan `EXPLAIN (ANALYZE, BUFFERS)`.

| Kebutuhan | Index |
|---|---|
| Cari pasien by nama (fuzzy) | `gin (normalized_name gin_trgm_ops)` |
| Cari pasien by NIK/BPJS | `ux_identifiers` |
| Riwayat kunjungan pasien | `(patient_id, started_at DESC)` |
| Worklist order aktif | partial index `WHERE status IN ('ACTIVE','IN_PROGRESS')` |
| Antrean aktif | partial index `WHERE status IN ('WAITING','CALLED')` |
| Invoice belum lunas | partial index `WHERE status IN ('ISSUED','PARTIALLY_PAID')` |
| Outbox | partial index status pending |
| Audit per pasien/entitas | `(patient_id, occurred_at DESC)`, `(entity_type, entity_id, ...)` |

**Partisi (berdasarkan bukti volume, bukan di awal):**

| Tabel | Kunci | Catatan |
|---|---|---|
| `audit.audit_logs` | bulan | sudah dipartisi di DDL; buat partisi bulan berikutnya lewat job terjadwal |
| `inventory.stock_movements` | bulan | PK menjadi `(id, occurred_at)` |
| `clinical.vital_signs` | bulan | bila volume IGD/ICU besar |
| `integration.outbox_events` | bulan atau purge `DONE` lama | |

Setiap tabel yang dipartisi: PK harus memuat kolom partisi, dan FK *ke* tabel partisi sebaiknya dihindari (itulah alasan `audit_logs` tidak punya FK).

**Skala bertahap:** satu database → read replica untuk laporan → materialized view/read model → CDC ke penyimpanan analitik. Jangan menjadikan OLTP sebagai mesin BI.

---

## 18. Strategi Testing

Gunakan database PostgreSQL nyata (container) untuk test integrasi; mock tidak cukup untuk constraint.

| Area | Test wajib |
|---|---|
| Penomoran | 100 goroutine meminta nomor bersamaan → tidak ada duplikat |
| Stok | Pengurangan konkuren melebihi saldo → salah satu gagal, saldo tidak negatif; saldo = Σ movement |
| Billing | Retry `GenerateCharge` dengan sumber sama → tetap satu charge |
| Payment | Alokasi melebihi jumlah → ditolak; dua alokasi konkuren ke invoice yang sama → konsisten |
| Dokumen | UPDATE langsung ke `document_versions` → ditolak trigger; amend → versi bertambah |
| Bed | Dua assignment beririsan pada satu bed → ditolak |
| Pasien | Merge pasien → pencarian pasien lama mengarah ke survivor; riwayat tidak hilang |
| State machine | Semua transisi valid/invalid per entitas (table-driven test) |
| Outbox | Worker mati di tengah proses → event tidak hilang dan tidak terkirim ganda |
| Migrasi | Naik dari kosong ke terbaru, dan dari snapshot versi sebelumnya |

---

## 19. Checklist Produksi

- [ ] Satu MPI; alur duplikat & merge pasien tersedia
- [ ] Registration, Episode, Encounter terpisah; `OpenOutpatientVisit` membuat ketiganya atomik
- [ ] Riwayat bed & transfer tersimpan, tidak ada overwrite
- [ ] Dokumen klinis berversi; tanda tangan & koreksi diaudit
- [ ] Diagnosis terpetakan ke ICD-10 via `master.codes`
- [ ] Order ≠ hasil; resep ≠ dispensing
- [ ] Stok memakai ledger; rekonsiliasi saldo vs movement terjadwal
- [ ] Charge ≠ invoice ≠ payment ≠ claim; tarif ber-masa-berlaku
- [ ] ID eksternal terpisah dari PK; outbox + idempotency aktif
- [ ] Audit append-only; akses baca data sensitif dapat diaudit
- [ ] Partisi audit dibuat otomatis untuk bulan berikutnya
- [ ] Backup terjadwal **dan uji restore** (RPO/RTO disepakati)
- [ ] Prosedur downtime (formulir manual & input susulan) tersedia
- [ ] Kebijakan retensi data ditetapkan sesuai regulasi yang berlaku
- [ ] Data dictionary dan state machine terdokumentasi
- [ ] Spesifikasi SATUSEHAT/BPJS terbaru diverifikasi sebelum integrasi
- [ ] Artefak dokumen sumber (`citeturn…`, penomoran ganda) dibersihkan

---

## 20. Lampiran — Urutan File Migrasi

| File | Isi |
|---|---|
| `001_init.sql` | ekstensi, schema, fungsi helper |
| `002_iam.sql` | users, roles, permissions |
| `003_org.sql` | organizations, facilities, locations, providers, user_facilities |
| `004_master.sql` | code_systems, codes, mappings, number_sequences |
| `005_audit_integration.sql` | audit_logs (partisi), systems, external_identifiers, outbox |
| `006_payers_patient.sql` | payers, patients, identifiers, contacts, addresses, merge_events |
| `007_care.sql` | service_lines, registrations, episodes, encounters, appointments, queue |
| `008_clinical.sql` | diagnoses, vital_signs, documents, versions, amendments, allergies, orders |
| `009_inventory_pharmacy.sql` | inventory, medications, prescriptions, dispensations |
| `010_billing_finance.sql` | service_catalog, prices, charges, invoices, payments, adjustments, claims |
| `011_inpatient.sql` | admissions, bed_assignments, transfers, referrals |
| `012+` | lab, radiology, mcu, nursing (mengikuti pola order → result) |

**Catatan integrasi dengan backend yang sudah berjalan:** tabel `users` yang kini dipakai modul `user`/`auth` dapat dipertahankan terlebih dahulu. Pindah ke `iam.users` cukup dengan mengubah `TableName()` model dan migrasi data sekali jalan. Mulailah dari fase 1–3, selesaikan hingga alur "daftar → periksa → diagnosis" berjalan end-to-end beserta test invarian, baru lanjut ke fase berikutnya.

---

*Dokumen ini adalah baseline desain. Sesuaikan tipe kolom, daftar status, dan aturan retensi dengan kebijakan klinis dan regulasi rumah sakit sebelum migrasi dijalankan di produksi.*
