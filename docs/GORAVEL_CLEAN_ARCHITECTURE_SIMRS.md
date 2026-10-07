# Clean Architecture Goravel — SIMRS

> **Scope:** Goravel native Go untuk backend SIMRS.
>
> Dokumen ini menggunakan pola aplikasi:
>
> ```text
> HTTP Request
>      ↓
> Request / Validation
>      ↓
> Controller
>      ↓
> Action
>      ↓
> Repository
>      ↓
> Database
>      ↓
> Repository
>      ↓
> Action
>      ↓
> Controller
>      ↓
> Response Trait
>      ↓
> HTTP Response
> ```
>
> Goravel adalah framework Go dengan struktur yang familiar dengan Laravel, tetapi implementasinya tetap menggunakan Go. Goravel menyediakan HTTP, routing, validation, ORM, authorization, testing, dan komponen framework lainnya. citeturn0view0turn0search6

---

# 1. Tujuan Arsitektur

Pada SIMRS, backend akan memiliki banyak domain:

```text
SIMRS
├── Patient
├── Registration
├── Outpatient
├── Emergency
├── Inpatient
├── MedicalRecord
├── Doctor
├── Nurse
├── Pharmacy
├── Laboratory
├── Radiology
├── Surgery
├── MCU
├── Billing
├── Insurance
└── MasterData
```

Jika semua logic ditempatkan langsung di Controller, aplikasi akan cepat menjadi sulit dirawat.

Contoh yang harus dihindari:

```text
PatientController
    ├── validation
    ├── authorization
    ├── business rules
    ├── transaction
    ├── patient query
    ├── registration query
    ├── medical record query
    ├── response formatting
    └── error handling
```

Untuk menjaga scalability, setiap layer memiliki tanggung jawab yang jelas.

---

# 2. Arsitektur Utama

```text
┌───────────────────────────────────────────────────────────────┐
│                         HTTP / API                            │
│                                                               │
│  Route → Request/Validation → Controller → Response Trait    │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                       APPLICATION                            │
│                                                               │
│                         Action                               │
│                                                               │
│  Use Case / Business Orchestration / Transaction Boundary   │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                         DATA ACCESS                           │
│                                                               │
│                       Repository                              │
│                                                               │
│                 ORM / Query Builder / SQL                    │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                         DATABASE                              │
│                                                               │
│                         PostgreSQL                            │
└───────────────────────────────────────────────────────────────┘
```

---

# 3. Request → Controller → Action → Repository → Response

Untuk sebuah endpoint SIMRS:

```text
POST /v1/patients
```

alur lengkapnya:

```text
Client
  │
  │ POST /v1/patients
  │
  ▼
┌──────────────────────┐
│ Request / Validation │
│                      │
│ Validate input       │
│ Bind JSON            │
└──────────┬───────────┘
           │
           │ PatientCreateRequest
           ▼
┌──────────────────────┐
│     Controller       │
│                      │
│ HTTP orchestration   │
└──────────┬───────────┘
           │
           │ Execute()
           ▼
┌──────────────────────┐
│       Action         │
│                      │
│ Business Use Case    │
│ Transaction          │
│ Business Rules       │
└──────────┬───────────┘
           │
           │ Repository
           ▼
┌──────────────────────┐
│     Repository       │
│                      │
│ Database Access      │
└──────────┬───────────┘
           │
           ▼
      PostgreSQL
           │
           │ Patient
           ▼
┌──────────────────────┐
│     Repository       │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│       Action         │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│     Controller       │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│   Response Trait     │
│                      │
│ Standard JSON        │
│ HTTP status          │
└──────────┬───────────┘
           │
           ▼
        Client
```

---

# 4. Layer 1 — Request

Goravel menyediakan `http.Context` yang memberikan akses terhadap HTTP request. Input dapat diambil melalui request dan JSON/form data dapat di-bind ke struct. Goravel juga menyediakan Form Request untuk validation yang lebih kompleks. citeturn0search2turn0search7

Untuk SIMRS, Request bertanggung jawab terhadap:

```text
HTTP Input
    ↓
Bind
    ↓
Validation
    ↓
Validated DTO
```

Contoh DTO:

```go
package requests

type StorePatientRequest struct {
    MedicalRecordNumber string `json:"medical_record_number"`
    Nik                 string `json:"nik"`
    Name                string `json:"name"`
    BirthDate           string `json:"birth_date"`
    Gender              string `json:"gender"`
    Phone               string `json:"phone"`
}
```

Contoh validation:

```go
func (r *StorePatientRequest) Rules() map[string]string {
    return map[string]string{
        "medical_record_number": "required",
        "name":                  "required|max_len:255",
        "birth_date":            "required",
        "gender":                "required",
    }
}
```

> Detail interface Form Request harus mengikuti versi Goravel yang digunakan oleh project. Prinsip arsitekturnya tetap sama: validation berada di HTTP/request boundary, bukan di Repository.

## Request tidak boleh

```text
Request
  ✗ query database
  ✗ create model
  ✗ update model
  ✗ menjalankan business workflow
  ✗ membuat response database
```

---

# 5. Layer 2 — Controller

Controller merupakan entry point HTTP.

Dokumentasi Goravel menempatkan controller pada:

```text
app/http/controllers
```

dan controller menerima `http.Context`. citeturn0search0turn0search3

Controller SIMRS sebaiknya tipis.

Contoh:

```go
package controllers

import (
    "goravel/app/actions/patient"
    "goravel/app/http/requests"

    "github.com/goravel/framework/contracts/http"
)

type PatientController struct {
    storePatientAction *patient.StorePatientAction
}

func NewPatientController(
    storePatientAction *patient.StorePatientAction,
) *PatientController {
    return &PatientController{
        storePatientAction: storePatientAction,
    }
}

func (c *PatientController) Store(ctx http.Context) http.Response {
    var request requests.StorePatientRequest

    if err := ctx.Request().Bind(&request); err != nil {
        return ctx.Response().Json(http.StatusBadRequest, map[string]any{
            "status":  false,
            "message": "Invalid request.",
            "errors":  err.Error(),
            "data":    nil,
        })
    }

    patient, err := c.storePatientAction.Execute(request)

    if err != nil {
        return err
    }

    return ctx.Response().Success().Json(map[string]any{
        "status":  true,
        "message": "Patient berhasil didaftarkan.",
        "errors":  nil,
        "data":    patient,
    })
}
```

> Contoh di atas menunjukkan bentuk native Goravel. Jika project memiliki `ResponseTrait`, bagian response sebaiknya dipindahkan ke trait tersebut agar seluruh endpoint konsisten.

Controller hanya mengatur:

```text
HTTP
 ↓
Input
 ↓
Action
 ↓
HTTP Response
```

Controller tidak menjadi tempat business logic.

---

# 6. Layer 3 — Action

Action adalah **use case**.

Untuk SIMRS, Action dapat merepresentasikan operasi bisnis:

```text
RegisterPatientAction
CreatePatientAction
UpdatePatientAction
GetPatientAction
ListPatientsAction

RegisterOutpatientAction
CreateMedicalRecordAction
CreatePrescriptionAction
DispensePrescriptionAction

AdmitPatientAction
DischargePatientAction

CreateMCURegistrationAction
CompleteMCUAction
```

Action bukan sekadar wrapper Repository.

Action adalah boundary business process.

---

# 7. Contoh Action SIMRS

Misalnya proses:

```text
Registrasi pasien baru
```

Business flow:

```text
RegisterPatientAction
       │
       ├── cek NIK
       │
       ├── cek nomor rekam medis
       │
       ├── create patient
       │
       └── create medical record identity
```

Contoh:

```go
package patient

import (
    "errors"

    "goravel/app/http/requests"
    "goravel/app/repositories/patient"
)

type RegisterPatientAction struct {
    patientRepository *patient.Repository
}

func NewRegisterPatientAction(
    patientRepository *patient.Repository,
) *RegisterPatientAction {
    return &RegisterPatientAction{
        patientRepository: patientRepository,
    }
}

func (a *RegisterPatientAction) Execute(
    request requests.StorePatientRequest,
) (*patient.Patient, error) {
    existingPatient, err := a.patientRepository.FindByNik(request.Nik)

    if err != nil {
        return nil, err
    }

    if existingPatient != nil {
        return nil, errors.New("patient dengan NIK tersebut sudah terdaftar")
    }

    return a.patientRepository.Create(patient.CreateData{
        MedicalRecordNumber: request.MedicalRecordNumber,
        Nik:                 request.Nik,
        Name:                request.Name,
        BirthDate:           request.BirthDate,
        Gender:              request.Gender,
        Phone:               request.Phone,
    })
}
```

Business rule:

```text
NIK sudah ada?
    ↓
YES → reject

NO
    ↓
Create Patient
```

Rule tersebut bukan tanggung jawab Repository.

---

# 8. Action dan Transaction

Transaction merupakan bagian dari business use case.

Contoh registrasi pasien:

```text
BEGIN TRANSACTION
        │
        ├── create patient
        │
        ├── create medical record
        │
        └── create patient contact
        │
        ▼
      COMMIT
```

Jika salah satu gagal:

```text
BEGIN
  ↓
Patient created
  ↓
Medical record created
  ↓
Contact failed
  ↓
ROLLBACK
```

Secara arsitektur:

```text
Controller
    ↓
RegisterPatientAction
    ↓
    ┌───────────────────────┐
    │      TRANSACTION      │
    │                       │
    │ PatientRepository     │
    │ MedicalRecordRepo     │
    │ ContactRepository     │
    └───────────────────────┘
```

Action menentukan bahwa operasi tersebut merupakan satu business transaction.

Repository hanya menjalankan persistence operation.

---

# 9. Layer 4 — Repository

Repository adalah database boundary.

Repository bertanggung jawab terhadap:

```text
SELECT
INSERT
UPDATE
DELETE
JOIN
PRELOAD / EAGER LOAD
FILTER
PAGINATION
```

Contoh:

```go
package patient

type Repository struct {
    // ORM / database dependency
}

type CreateData struct {
    MedicalRecordNumber string
    Nik                 string
    Name                string
    BirthDate           string
    Gender              string
    Phone               string
}

func (r *Repository) FindByNik(
    nik string,
) (*Patient, error) {
    // Database query
    return nil, nil
}

func (r *Repository) Create(
    data CreateData,
) (*Patient, error) {
    // Database insert
    return nil, nil
}
```

Repository:

```text
Action
  ↓
Repository
  ↓
ORM
  ↓
PostgreSQL
```

Goravel menyediakan ORM sebagai salah satu komponen framework. citeturn0view0

---

# 10. Repository Tidak Boleh Mengandung Business Rule

Jangan:

```go
func (r *Repository) Create(data CreateData) (*Patient, error) {
    if data.Gender != "M" && data.Gender != "F" {
        return nil, errors.New("gender tidak valid")
    }

    // insert
}
```

Jika aturan tersebut adalah domain/business rule, tempat yang lebih tepat adalah Action/domain validation.

Repository fokus pada:

```text
"Bagaimana data disimpan?"

bukan:

"Apakah operasi bisnis ini boleh dilakukan?"
```

---

# 11. Layer 5 — Response Trait

Response Trait digunakan untuk menyeragamkan API response.

Target format:

```json
{
    "status": true,
    "message": "Patient berhasil dibuat.",
    "errors": null,
    "data": {}
}
```

Error:

```json
{
    "status": false,
    "message": "Patient tidak ditemukan.",
    "errors": null,
    "data": null
}
```

Secara arsitektur:

```text
Action Result
     ↓
Controller
     ↓
ResponseTrait
     ↓
HTTP JSON
```

Response Trait tidak boleh:

```text
✗ query database
✗ memanggil Repository
✗ menjalankan Action
✗ menjalankan business rule
```

Fokusnya:

```text
Application Result
       ↓
HTTP Representation
```

---

# 12. Diagram Dependency

Dependency yang disarankan:

```text
┌──────────────────────────────┐
│          Controller          │
│                              │
│ HTTP boundary               │
└──────────────┬───────────────┘
               │
               ▼
┌──────────────────────────────┐
│            Action            │
│                              │
│ Use Case / Business Logic   │
└──────────────┬───────────────┘
               │
               ▼
┌──────────────────────────────┐
│         Repository           │
│                              │
│ Persistence / Database      │
└──────────────┬───────────────┘
               │
               ▼
          PostgreSQL
```

Response:

```text
Database
   ↓
Repository
   ↓
Action
   ↓
Controller
   ↓
ResponseTrait
   ↓
HTTP Response
```

---

# 13. Repository Contract

Untuk project SIMRS besar, gunakan abstraction antara Action dan implementasi database.

```text
                 ┌────────────────────────┐
                 │ PatientRepository      │
                 │ Interface / Contract   │
                 └────────────▲───────────┘
                              │
                              │ implements
                              │
                 ┌────────────┴───────────┐
                 │ PatientRepositoryImpl  │
                 │ Goravel ORM            │
                 └────────────────────────┘
                              ▲
                              │
                              │
                 ┌────────────┴───────────┐
                 │ RegisterPatientAction  │
                 └────────────────────────┘
```

Action bergantung pada contract:

```go
type PatientRepository interface {
    FindByNik(nik string) (*Patient, error)

    Create(data CreateData) (*Patient, error)
}
```

Implementasi database:

```go
type PatientRepositoryImpl struct {
    // ORM dependency
}
```

Dengan demikian:

```text
Action
  ↓
Interface
  ↓
Implementation
  ↓
PostgreSQL
```

Bukan:

```text
Action
  ↓
Eloquent/ORM detail
  ↓
PostgreSQL
```

---

# 14. Struktur Folder Goravel untuk SIMRS

Goravel menyediakan struktur default seperti `app/http`, `app/models`, `app/providers`, `routes`, dan `database`, serta mengizinkan folder tambahan untuk kebutuhan aplikasi. citeturn0search3

Untuk SIMRS yang scalable:

```text
app/
├── actions/
│   ├── patient/
│   │   ├── register_patient_action.go
│   │   ├── update_patient_action.go
│   │   └── get_patient_action.go
│   │
│   ├── registration/
│   │   ├── register_outpatient_action.go
│   │   └── cancel_registration_action.go
│   │
│   ├── medical_record/
│   │   ├── create_medical_record_action.go
│   │   └── update_medical_record_action.go
│   │
│   ├── pharmacy/
│   │   ├── create_prescription_action.go
│   │   └── dispense_prescription_action.go
│   │
│   └── mcu/
│       ├── register_mcu_action.go
│       └── complete_mcu_action.go
│
├── http/
│   ├── controllers/
│   │   ├── patient_controller.go
│   │   ├── registration_controller.go
│   │   ├── medical_record_controller.go
│   │   ├── pharmacy_controller.go
│   │   └── mcu_controller.go
│   │
│   ├── requests/
│   │   ├── patient/
│   │   ├── registration/
│   │   ├── medical_record/
│   │   ├── pharmacy/
│   │   └── mcu/
│   │
│   └── middleware/
│
├── repositories/
│   ├── patient/
│   │   ├── repository.go
│   │   └── repository_impl.go
│   │
│   ├── registration/
│   │   ├── repository.go
│   │   └── repository_impl.go
│   │
│   ├── medical_record/
│   ├── pharmacy/
│   └── mcu/
│
├── models/
│   ├── patient.go
│   ├── registration.go
│   ├── medical_record.go
│   ├── prescription.go
│   └── mcu.go
│
├── services/
│   ├── medical_record_number_service.go
│   ├── billing_service.go
│   └── queue_service.go
│
├── traits/
│   └── response_trait.go
│
└── providers/
```

> Nama folder `actions`, `repositories`, `services`, dan `traits` adalah architectural convention aplikasi, bukan berarti semuanya merupakan folder bawaan wajib Goravel. Goravel menyediakan struktur dasar dan memungkinkan aplikasi menambahkan folder sesuai kebutuhan. citeturn0search3

---

# 15. Feature / Domain Boundary

Untuk SIMRS yang semakin besar, pendekatan domain/feature dapat dibuat lebih kuat:

```text
app/
└── modules/
    ├── patient/
    │   ├── actions/
    │   ├── controllers/
    │   ├── requests/
    │   ├── repositories/
    │   ├── models/
    │   └── services/
    │
    ├── registration/
    │   ├── actions/
    │   ├── controllers/
    │   ├── requests/
    │   └── repositories/
    │
    ├── medical_record/
    ├── pharmacy/
    ├── laboratory/
    ├── radiology/
    ├── inpatient/
    ├── outpatient/
    ├── emergency/
    └── mcu/
```

Pendekatan ini cocok ketika SIMRS sudah memiliki banyak bounded context.

---

# 16. Contoh Domain Patient

```text
Patient
│
├── Create Patient
│      │
│      ├── StorePatientRequest
│      ├── CreatePatientAction
│      ├── PatientRepository
│      └── ResponseTrait
│
├── Update Patient
│      │
│      ├── UpdatePatientRequest
│      ├── UpdatePatientAction
│      ├── PatientRepository
│      └── ResponseTrait
│
├── Get Patient
│      │
│      ├── GetPatientAction
│      ├── PatientRepository
│      └── ResponseTrait
│
└── Search Patient
       │
       ├── SearchPatientRequest
       ├── SearchPatientAction
       ├── PatientRepository
       └── ResponseTrait
```

---

# 17. Contoh Domain Registrasi Rawat Jalan

Use case:

```text
Register Outpatient
```

Flow:

```text
POST /v1/registrations/outpatient
              │
              ▼
     RegisterOutpatientRequest
              │
              ▼
     RegistrationController
              │
              ▼
     RegisterOutpatientAction
              │
       ┌──────┼───────────┐
       ▼      ▼           ▼
    Patient Doctor    Schedule
    Repo     Repo       Repo
       │      │           │
       └──────┼───────────┘
              ▼
       RegistrationRepo
              │
              ▼
          PostgreSQL
              │
              ▼
     RegisterOutpatientAction
              │
              ▼
     RegistrationController
              │
              ▼
        ResponseTrait
              │
              ▼
             JSON
```

Business rule dapat mencakup:

```text
Patient harus aktif
       ↓
Doctor harus aktif
       ↓
Schedule tersedia
       ↓
Quota belum penuh
       ↓
Patient belum terdaftar pada jadwal yang sama
       ↓
Create Registration
       ↓
Generate Queue Number
```

Semua orchestration tersebut adalah tanggung jawab Action/use case.

---

# 18. Contoh Domain IGD

```text
POST /v1/emergency/registrations
```

Flow:

```text
EmergencyRequest
      │
      ▼
EmergencyController
      │
      ▼
RegisterEmergencyAction
      │
      ├── PatientRepository
      │
      ├── EmergencyRepository
      │
      ├── DoctorRepository
      │
      └── QueueRepository
      │
      ▼
PostgreSQL
      │
      ▼
ResponseTrait
```

Business rule:

```text
Patient ditemukan?
       │
       ├── YES → gunakan patient
       │
       └── NO  → buat patient temporary

Triage
  ↓
Emergency Registration
  ↓
Queue
  ↓
Doctor assignment
```

---

# 19. Contoh Domain Rawat Inap

```text
POST /v1/inpatient/admissions
```

Flow:

```text
AdmissionRequest
       │
       ▼
InpatientController
       │
       ▼
AdmitPatientAction
       │
       ├── PatientRepository
       ├── BedRepository
       ├── WardRepository
       ├── DoctorRepository
       └── AdmissionRepository
       │
       ▼
Transaction
       │
       ├── lock bed
       ├── create admission
       ├── assign bed
       └── update bed status
       │
       ▼
COMMIT
       │
       ▼
ResponseTrait
```

Penting:

```text
Bed tersedia?
       ↓
lock
       ↓
create admission
       ↓
assign bed
       ↓
ubah bed menjadi occupied
```

Ini merupakan satu business transaction.

---

# 20. Contoh Domain Farmasi

```text
POST /v1/pharmacy/prescriptions/{id}/dispense
```

Flow:

```text
DispensePrescriptionRequest
             │
             ▼
PharmacyController
             │
             ▼
DispensePrescriptionAction
             │
             ├── PrescriptionRepository
             ├── InventoryRepository
             ├── StockMovementRepository
             └── DispensingRepository
             │
             ▼
        Transaction
             │
             ├── validate prescription
             ├── check stock
             ├── deduct stock
             ├── create stock movement
             └── create dispensing record
             │
             ▼
           COMMIT
             │
             ▼
        ResponseTrait
```

Action memastikan business process konsisten.

Repository hanya menjalankan persistence.

---

# 21. Contoh Domain MCU

MCU dapat memiliki use case tersendiri:

```text
POST /v1/mcu/registrations
```

Flow:

```text
MCURegistrationRequest
          │
          ▼
MCUController
          │
          ▼
RegisterMCUAction
          │
          ├── PatientRepository
          ├── PackageRepository
          ├── DoctorRepository
          ├── MCURepository
          └── BillingRepository
          │
          ▼
      PostgreSQL
          │
          ▼
      ResponseTrait
```

Setelah registrasi:

```text
MCU Registration
      ↓
Assessment
      ↓
Laboratory
      ↓
Radiology
      ↓
Doctor Examination
      ↓
Conclusion
      ↓
MCU Report
```

Jika proses tersebut membutuhkan beberapa repository, orchestration tetap berada pada Action.

---

# 22. Service vs Action

Jangan membuat semua logic menjadi Service.

Gunakan:

```text
Action = Use Case
Service = Reusable Business Capability
Repository = Persistence
```

Contoh:

```text
RegisterPatientAction
        │
        ├── MedicalRecordNumberService
        ├── PatientRepository
        └── MedicalRecordRepository
```

Service:

```text
GenerateMedicalRecordNumberService
CalculateBillingService
GenerateQueueNumberService
StockCalculationService
```

Action:

```text
RegisterPatientAction
RegisterOutpatientAction
AdmitPatientAction
DispensePrescriptionAction
CompleteMCUAction
```

Perbedaan:

```text
"Register patient"
        ↓
Action

"Generate medical record number"
        ↓
Service
```

---

# 23. Authorization

Authentication dan authorization tidak seharusnya dicampur ke Repository.

Goravel menyediakan authorization melalui Gate dan Policy. citeturn0search9

Contoh konseptual:

```text
HTTP Request
     │
     ▼
Authentication
     │
     ▼
Authorization
     │
     ▼
Controller
     │
     ▼
Action
```

Contoh:

```text
Dokter
  → boleh membaca medical record

Perawat
  → boleh melakukan nursing documentation

Farmasi
  → boleh dispense prescription

Kasir
  → boleh melakukan billing

Admin
  → boleh melakukan master data
```

Authorization resource-level dapat menggunakan policy/gate, sedangkan business rule tetap berada di use case/domain boundary.

---

# 24. Error Flow

Error harus tetap melewati boundary dengan jelas.

```text
Repository
     │
     │ database error
     ▼
Action
     │
     │ business error
     ▼
Controller
     │
     ▼
ResponseTrait
     │
     ▼
HTTP Error
```

Contoh:

```json
{
    "status": false,
    "message": "Tempat tidur tidak tersedia.",
    "errors": null,
    "data": null
}
```

Untuk validation:

```json
{
    "status": false,
    "message": "Validation failed.",
    "errors": {
        "patient_id": [
            "patient_id wajib diisi."
        ]
    },
    "data": null
}
```

---

# 25. Jangan Membuat Controller Gemuk

## Anti-pattern

```go
func (c *PatientController) Store(ctx http.Context) http.Response {
    // bind request

    // validation

    // query NIK

    // business rule

    // generate medical record number

    // insert patient

    // insert medical record

    // transaction

    // response
}
```

Masalah:

```text
Controller
    ├── HTTP
    ├── Validation
    ├── Business Logic
    ├── Persistence
    ├── Transaction
    └── Response
```

Sulit:

```text
testing
maintenance
reuse
debugging
refactoring
```

---

# 26. Controller yang Sehat

Target:

```go
func (c *PatientController) Store(ctx http.Context) http.Response {
    request, err := c.bindRequest(ctx)

    if err != nil {
        return c.errorResponse(ctx, err)
    }

    result, err := c.registerPatientAction.Execute(request)

    if err != nil {
        return c.errorResponse(ctx, err)
    }

    return c.successResponse(
        ctx,
        result,
        "Patient berhasil didaftarkan.",
    )
}
```

Controller hanya menjadi:

```text
HTTP Adapter
```

---

# 27. Golden Rule

Gunakan pertanyaan berikut saat menentukan lokasi code:

```text
Apakah berhubungan dengan HTTP?
        │
        └── Request / Controller / Response

Apakah ini use case bisnis?
        │
        └── Action

Apakah ini reusable business capability?
        │
        └── Service

Apakah ini query / persistence?
        │
        └── Repository

Apakah ini representasi data database?
        │
        └── Model

Apakah ini authorization resource?
        │
        └── Policy / Gate
```

---

# 28. Final Architecture SIMRS

```text
                         ┌───────────────┐
                         │    CLIENT     │
                         │ Nuxt / Mobile │
                         └───────┬───────┘
                                 │
                                 │ HTTP
                                 ▼
                    ┌────────────────────────┐
                    │        ROUTES          │
                    │ routes/web.go / api.go│
                    └────────────┬───────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │ REQUEST / VALIDATION   │
                    │                        │
                    │ Bind + Validate        │
                    └────────────┬───────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │       CONTROLLER       │
                    │                        │
                    │ HTTP orchestration     │
                    └────────────┬───────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │         ACTION         │
                    │                        │
                    │ Use Case               │
                    │ Business Rules         │
                    │ Transaction            │
                    │ Orchestration          │
                    └────────────┬───────────┘
                                 │
               ┌─────────────────┼──────────────────┐
               │                 │                  │
               ▼                 ▼                  ▼
        ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
        │ Repository  │   │   Service   │   │ Policy/Gate │
        │             │   │             │   │             │
        │ Persistence │   │ Capability  │   │ Authorization│
        └──────┬──────┘   └─────────────┘   └─────────────┘
               │
               ▼
        ┌─────────────┐
        │    MODEL    │
        │             │
        │ ORM Entity  │
        └──────┬──────┘
               │
               ▼
        ┌─────────────┐
        │ PostgreSQL  │
        └──────┬──────┘
               │
               │ result
               ▼
        ┌─────────────┐
        │ Repository  │
        └──────┬──────┘
               │
               ▼
        ┌─────────────┐
        │   Action    │
        └──────┬──────┘
               │
               ▼
        ┌─────────────┐
        │ Controller  │
        └──────┬──────┘
               │
               ▼
        ┌─────────────┐
        │ResponseTrait│
        │             │
        │ JSON format │
        └──────┬──────┘
               │
               ▼
            CLIENT
```

---

# 29. Prinsip Akhir

Arsitektur SIMRS Goravel yang sehat dapat diringkas menjadi:

```text
REQUEST
    │
    │ validate / bind
    ▼
CONTROLLER
    │
    │ call use case
    ▼
ACTION
    │
    │ business rules
    │ transaction
    │ orchestration
    ▼
REPOSITORY
    │
    │ persistence
    ▼
DATABASE
```

Response:

```text
DATABASE
    ↓
REPOSITORY
    ↓
ACTION
    ↓
CONTROLLER
    ↓
RESPONSE TRAIT
    ↓
HTTP JSON
```

Dan pembagian tanggung jawab:

| Layer | Tanggung Jawab |
|---|---|
| Route | Routing |
| Request | Bind & validation |
| Controller | HTTP orchestration |
| Action | Use case & business orchestration |
| Service | Reusable business capability |
| Repository | Database persistence/query |
| Model | Database/domain representation |
| Policy/Gate | Authorization |
| Response Trait | Standard HTTP response |
| Database | Persistent data |

## Kesimpulan

Untuk SIMRS Goravel, jangan melihat:

```text
Request → Controller → Action → Repository
```

sebagai sekadar urutan folder.

Anggap masing-masing sebagai **boundary tanggung jawab**:

```text
HTTP Boundary
      ↓
Application / Use Case Boundary
      ↓
Persistence Boundary
      ↓
Database
```

Dengan demikian ketika SIMRS berkembang dari:

```text
Patient
```

menjadi:

```text
Patient
Registration
Outpatient
Inpatient
Emergency
Medical Record
Pharmacy
Laboratory
Radiology
Surgery
MCU
Billing
Insurance
```

business logic tidak menumpuk di Controller dan database access tidak menyebar ke seluruh aplikasi.

Prinsip paling penting:

> **Controller menerima dan mengorkestrasi HTTP. Action menjalankan use case. Repository mengurus persistence. Response Trait mengubah hasil aplikasi menjadi kontrak HTTP yang konsisten.**
