# Clean Architecture pada Goravel

Dokumen ini menjelaskan pola arsitektur backend Goravel dengan alur utama:

```text
HTTP Request
     │
     ▼
┌──────────────┐
│   Request    │  Validasi & normalisasi input
└──────┬───────┘
       │ validated data
       ▼
┌──────────────┐
│  Controller  │  HTTP orchestration
└──────┬───────┘
       │
       ▼
┌──────────────┐
│    Action    │  Business/use-case logic
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Repository   │  Database persistence/query
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ ResponseTrait│  Standardisasi HTTP response
└──────┬───────┘
       │
       ▼
   HTTP Response
```

> Prinsip utamanya: setiap layer mempunyai satu tanggung jawab yang jelas. Controller tidak menjadi tempat business logic, Repository tidak menjadi tempat business rule, dan Response Trait tidak melakukan query database.

---

## 1. Gambaran Besar

Untuk aplikasi Goravel, request dapat diproses menggunakan pipeline berikut:

```text
Client
  │
  │ HTTP Request
  ▼
Request
  │
  │ validated input
  ▼
Controller
  │
  │ call use case
  ▼
Action
  │
  │ business operation
  ▼
Repository
  │
  │ query / persistence
  ▼
Database
  │
  │ result
  ▼
Repository
  │
  │ entity / model / DTO
  ▼
Action
  │
  │ business result
  ▼
Controller
  │
  │ response payload
  ▼
ResponseTrait
  │
  │ standardized JSON
  ▼
Client
```

Secara konseptual:

```text
Request
   ↓
Controller
   ↓
Action
   ↓
Repository
   ↓
Database

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

# 2. Tanggung Jawab Setiap Layer

## 2.1 Request

Request bertanggung jawab terhadap **input HTTP**.

Tanggung jawab:

- menerima parameter request;
- validasi input;
- authorization level request jika diperlukan;
- normalisasi input sederhana;
- menghasilkan data yang aman untuk diteruskan ke Controller.

Contoh:

```text
POST /api/v1/projects
```

Request:

```json
{
  "nama_proyek": "Perumahan A",
  "alamat": "Jl. Contoh No. 10",
  "luas_lahan": 5000
}
```

Request melakukan:

```text
Input
  ↓
Validation
  ↓
Validated Data
```

Contoh:

```php
class StoreProjectRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'nama_proyek' => ['required', 'string', 'max:255'],
            'alamat' => ['nullable', 'string'],
            'luas_lahan' => ['required', 'numeric', 'min:0'],
        ];
    }
}
```

### Request tidak boleh

Request sebaiknya tidak:

- melakukan query database;
- memanggil Repository;
- menjalankan business logic;
- membuat response JSON;
- melakukan transaksi database.

---

# 3. Controller

Controller merupakan **HTTP entry point**.

Controller menerima Request dan meneruskan data ke Action.

Controller idealnya tipis.

```text
Request
   │
   ▼
Controller
   │
   └──> Action
```

Contoh:

```php
class ProjectController extends Controller
{
    public function store(StoreProjectRequest $request)
    {
        $project = $this->storeProjectAction->execute(
            $request->validated()
        );

        return $this->success(
            $project,
            'Project berhasil dibuat.'
        );
    }
}
```

Controller bertugas:

- menerima HTTP request;
- mengambil validated data;
- memanggil Action;
- menentukan HTTP response;
- menggunakan Response Trait untuk response standar.

### Controller tidak boleh

Controller sebaiknya tidak berisi:

```php
DB::table(...)->insert(...);
```

atau:

```php
$project = Project::create(...);
```

atau business rule seperti:

```php
if ($project->status === 'finished') {
    // business logic
}
```

Business logic tersebut berada di Action.

---

# 4. Action

Action adalah **use case / business operation**.

Action merupakan layer terpenting dalam pola ini karena business process berada di sini.

Contoh use case:

```text
StoreProjectAction
UpdateProjectAction
DeleteProjectAction
GetProjectAction
ListProjectsAction
ApproveProjectAction
FinishProjectAction
```

Alur:

```text
Controller
    │
    ▼
Action
    │
    ├── validation/business rule lanjutan
    ├── transaction
    ├── orchestration
    ├── authorization business
    └── Repository
```

Contoh:

```php
class StoreProjectAction
{
    public function __construct(
        protected ProjectRepository $repository,
    ) {
    }

    public function execute(array $data): Project
    {
        return $this->repository->create($data);
    }
}
```

Untuk business process yang lebih kompleks:

```php
class ApproveProjectAction
{
    public function __construct(
        protected ProjectRepository $projectRepository,
        protected BudgetRepository $budgetRepository,
    ) {
    }

    public function execute(string $projectId): Project
    {
        $project = $this->projectRepository->findOrFail($projectId);

        if ($project->status !== 'draft') {
            throw new DomainException(
                'Project hanya dapat disetujui dari status draft.'
            );
        }

        $project = $this->projectRepository->approve($project);

        $this->budgetRepository->initialize($project);

        return $project;
    }
}
```

### Action boleh

Action boleh:

- menjalankan business rule;
- memanggil beberapa Repository;
- menjalankan transaction;
- melakukan orchestration;
- menentukan urutan proses;
- mengubah state domain;
- memanggil service lain jika diperlukan.

### Action tidak boleh

Action sebaiknya tidak:

- mengetahui detail HTTP;
- mengakses `$request`;
- membuat `JsonResponse`;
- menentukan HTTP status code;
- melakukan formatting response API.

Contoh yang sebaiknya dihindari:

```php
public function execute(Request $request)
{
    return response()->json(...);
}
```

Action seharusnya menerima data aplikasi, bukan object HTTP Request.

---

# 5. Repository

Repository bertanggung jawab terhadap **persistence dan database access**.

Alurnya:

```text
Action
   │
   ▼
Repository
   │
   ├── Query
   ├── Insert
   ├── Update
   ├── Delete
   └── Fetch
   │
   ▼
Database
```

Contoh interface:

```php
interface ProjectRepository
{
    public function create(array $data): Project;

    public function find(string $id): ?Project;

    public function findOrFail(string $id): Project;

    public function update(Project $project, array $data): Project;

    public function delete(Project $project): bool;
}
```

Implementasi:

```php
class ProjectRepositoryImpl implements ProjectRepository
{
    public function create(array $data): Project
    {
        return Project::create($data);
    }

    public function find(string $id): ?Project
    {
        return Project::query()
            ->where('id', $id)
            ->first();
    }

    public function findOrFail(string $id): Project
    {
        return Project::query()
            ->where('id', $id)
            ->firstOrFail();
    }

    public function update(Project $project, array $data): Project
    {
        $project->update($data);

        return $project->refresh();
    }

    public function delete(Project $project): bool
    {
        return $project->delete();
    }
}
```

### Repository boleh

Repository boleh:

- menggunakan Eloquent;
- menggunakan Query Builder;
- membuat query kompleks;
- eager loading;
- filtering;
- pagination;
- insert;
- update;
- delete;
- mengambil data dari database.

### Repository tidak boleh

Repository tidak boleh:

```php
if ($project->status === 'finished') {
    throw new Exception(...);
}
```

jika aturan tersebut merupakan business rule.

Business rule berada di Action.

Repository juga tidak seharusnya membuat HTTP response:

```php
return response()->json(...);
```

---

# 6. Response Trait

Response Trait bertanggung jawab untuk **standardisasi HTTP response**.

Tujuannya supaya semua endpoint menghasilkan format response yang konsisten.

Contoh:

```php
trait ResponseTrait
{
    protected function success(
        mixed $data = null,
        string $message = 'Success',
        int $status = 200,
    ) {
        return response()->json([
            'status' => true,
            'message' => $message,
            'errors' => null,
            'data' => $data,
        ], $status);
    }

    protected function error(
        string $message,
        mixed $errors = null,
        int $status = 400,
    ) {
        return response()->json([
            'status' => false,
            'message' => $message,
            'errors' => $errors,
            'data' => null,
        ], $status);
    }
}
```

Response:

```json
{
  "status": true,
  "message": "Project berhasil dibuat.",
  "errors": null,
  "data": {
    "id": "01K...",
    "nama_proyek": "Perumahan A"
  }
}
```

### Response Trait tidak boleh

Response Trait tidak boleh:

- melakukan query database;
- menjalankan business rule;
- memanggil Action;
- memanggil Repository.

Fungsinya hanya:

```text
Application Result
       ↓
HTTP Representation
```

---

# 7. Full Request Lifecycle

Contoh:

```text
POST /v1/projects
```

## Step 1 — HTTP Request

Client mengirim:

```json
{
  "nama_proyek": "Perumahan A",
  "alamat": "Jl. Contoh",
  "luas_lahan": 5000
}
```

↓

## Step 2 — Request

```text
StoreProjectRequest
```

Melakukan:

```text
required?
string?
numeric?
max length?
```

Hasil:

```php
$request->validated()
```

↓

## Step 3 — Controller

```php
public function store(StoreProjectRequest $request)
{
    $project = $this->storeProjectAction->execute(
        $request->validated()
    );

    return $this->success(
        $project,
        'Project berhasil dibuat.',
        201
    );
}
```

↓

## Step 4 — Action

```php
StoreProjectAction::execute()
```

Action menjalankan business process.

```text
Business Rule
      ↓
Repository
```

↓

## Step 5 — Repository

```php
$this->repository->create($data);
```

Repository melakukan:

```text
Eloquent
   ↓
PostgreSQL
```

↓

## Step 6 — Database

Database menyimpan:

```text
projects
```

↓

## Step 7 — Repository

Repository mengembalikan:

```php
Project
```

↓

## Step 8 — Action

Action mengembalikan result:

```php
Project
```

↓

## Step 9 — Controller

Controller menerima result:

```php
$project
```

↓

## Step 10 — Response Trait

Response distandarkan:

```json
{
  "status": true,
  "message": "Project berhasil dibuat.",
  "errors": null,
  "data": {}
}
```

---

# 8. Diagram Dependency

Dependency sebaiknya mengalir ke arah yang benar:

```text
┌──────────────────────┐
│      HTTP Layer      │
│                      │
│ Request              │
│ Controller            │
│ Response Trait        │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│   Application Layer  │
│                      │
│ Action               │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│   Infrastructure     │
│                      │
│ Repository           │
│ Eloquent / Query     │
└──────────┬───────────┘
           │
           ▼
       PostgreSQL
```

Secara dependency:

```text
Controller
    ↓
Action
    ↓
Repository
    ↓
Database
```

Bukan:

```text
Controller
    ├──→ Database
    ├──→ Repository
    ├──→ Business Logic
    └──→ Response
```

Controller harus tetap tipis.

---

# 9. Dependency Inversion

Untuk arsitektur yang lebih clean, Action sebaiknya bergantung pada **Repository Contract**, bukan implementasi database secara langsung.

```text
             ┌──────────────────────┐
             │   ProjectRepository  │
             │      Interface       │
             └──────────▲───────────┘
                        │
              implements│
                        │
             ┌──────────┴───────────┐
             │ ProjectRepositoryImpl│
             │   Eloquent / DB      │
             └──────────────────────┘
                        ▲
                        │
                        │
             ┌──────────┴───────────┐
             │ StoreProjectAction   │
             └──────────────────────┘
```

Dengan pola ini:

```php
class StoreProjectAction
{
    public function __construct(
        protected ProjectRepository $repository,
    ) {
    }
}
```

Action tidak peduli apakah Repository menggunakan:

```text
Eloquent
Query Builder
Raw SQL
External API
Mock Repository
```

Yang penting contract terpenuhi.

---

# 10. Struktur Folder

Struktur sederhana:

```text
app/
├── Actions/
│   └── Project/
│       ├── StoreProjectAction.php
│       ├── UpdateProjectAction.php
│       ├── DeleteProjectAction.php
│       └── ApproveProjectAction.php
│
├── Http/
│   ├── Controllers/
│   │   └── ProjectController.php
│   │
│   └── Requests/
│       └── Project/
│           ├── StoreProjectRequest.php
│           └── UpdateProjectRequest.php
│
├── Repositories/
│   └── Project/
│       ├── ProjectRepository.php
│       └── ProjectRepositoryImpl.php
│
├── Models/
│   └── Project.php
│
└── Traits/
    └── ResponseTrait.php
```

Untuk project besar, dapat dipisahkan berdasarkan domain/feature:

```text
app/
└── Modules/
    └── Project/
        ├── Actions/
        ├── Controllers/
        ├── Requests/
        ├── Repositories/
        ├── Models/
        ├── DTOs/
        └── Services/
```

Struktur feature-based biasanya lebih mudah diskalakan ketika jumlah module bertambah.

---

# 11. Contoh Full Flow

```php
class ProjectController extends Controller
{
    use ResponseTrait;

    public function __construct(
        protected StoreProjectAction $storeProjectAction,
    ) {
    }

    public function store(StoreProjectRequest $request)
    {
        $project = $this->storeProjectAction->execute(
            $request->validated()
        );

        return $this->success(
            $project,
            'Project berhasil dibuat.',
            201
        );
    }
}
```

Action:

```php
class StoreProjectAction
{
    public function __construct(
        protected ProjectRepository $repository,
    ) {
    }

    public function execute(array $data): Project
    {
        return $this->repository->create($data);
    }
}
```

Repository:

```php
class ProjectRepositoryImpl implements ProjectRepository
{
    public function create(array $data): Project
    {
        return Project::query()->create($data);
    }
}
```

Response:

```json
{
  "status": true,
  "message": "Project berhasil dibuat.",
  "errors": null,
  "data": {
    "id": "01K...",
    "nama_proyek": "Perumahan A"
  }
}
```

---

# 12. Transaction Berada di Mana?

Untuk transaction yang merupakan bagian dari satu business use case, transaction idealnya dikelola oleh **Action/Application layer**.

```text
Controller
    ↓
Action
    ↓
BEGIN TRANSACTION
    │
    ├── Repository A
    │
    ├── Repository B
    │
    └── Repository C
    │
COMMIT
```

Jika gagal:

```text
BEGIN
  ↓
Repository A
  ↓
Repository B
  ↓
ERROR
  ↓
ROLLBACK
```

Contoh:

```php
public function execute(array $data): Project
{
    return DB::transaction(function () use ($data) {
        $project = $this->projectRepository->create($data);

        $this->budgetRepository->initialize($project);

        $this->phaseRepository->initialize($project);

        return $project;
    });
}
```

Repository tetap fokus pada database operation.

Action menentukan bahwa beberapa operasi tersebut merupakan **satu business transaction**.

---

# 13. Kapan Membuat Service?

Tidak semua business logic harus dipaksa masuk ke Action.

Gunakan Action sebagai **use case boundary**.

Contoh:

```text
CreateProjectAction
       │
       ├── ProjectRepository
       ├── BudgetService
       └── PhaseService
```

Service cocok untuk logic yang:

- reusable;
- cukup kompleks;
- tidak merepresentasikan satu HTTP use case;
- digunakan oleh beberapa Action.

Contoh:

```text
CalculateBudgetService
GenerateProjectCodeService
CalculateMaterialPriceService
```

Namun hindari membuat:

```text
ProjectService
```

yang akhirnya menjadi tempat semua business logic.

Jika setiap operasi mempunyai use case yang jelas, lebih baik:

```text
CreateProjectAction
UpdateProjectAction
ApproveProjectAction
FinishProjectAction
```

daripada satu God Service:

```text
ProjectService
```

---

# 14. Error Flow

Error juga mengikuti boundary layer.

```text
Repository
    │
    │ exception / not found
    ▼
Action
    │
    │ business exception
    ▼
Controller / Exception Handler
    │
    ▼
ResponseTrait
    │
    ▼
HTTP Error
```

Contoh response:

```json
{
  "status": false,
  "message": "Project tidak dapat disetujui.",
  "errors": {
    "status": [
      "Project harus berada pada status draft."
    ]
  },
  "data": null
}
```

Jangan membuat Repository mengembalikan response seperti:

```php
return response()->json([
    'status' => false,
]);
```

Repository harus tetap independen dari HTTP.

---

# 15. Prinsip Utama

## Single Responsibility

```text
Request
  = validation

Controller
  = HTTP orchestration

Action
  = business/use-case

Repository
  = persistence

ResponseTrait
  = HTTP response formatting
```

## Separation of Concerns

Jangan mencampur:

```text
Validation
Business Logic
Database Query
HTTP Response
```

dalam satu class.

Hindari:

```php
public function store(Request $request)
{
    // validation
    // business rule
    // query
    // transaction
    // response
}
```

Gunakan:

```text
Request
   ↓
Controller
   ↓
Action
   ↓
Repository
   ↓
Database
```

---

# 16. Diagram Final

```text
                           CLIENT
                             │
                             │ HTTP
                             ▼
                    ┌─────────────────┐
                    │     REQUEST     │
                    │                 │
                    │ Validation      │
                    │ Normalization   │
                    └────────┬────────┘
                             │
                             │ validated data
                             ▼
                    ┌─────────────────┐
                    │   CONTROLLER    │
                    │                 │
                    │ HTTP orchestration
                    └────────┬────────┘
                             │
                             │ execute()
                             ▼
                    ┌─────────────────┐
                    │     ACTION      │
                    │                 │
                    │ Use Case        │
                    │ Business Rule   │
                    │ Transaction     │
                    │ Orchestration   │
                    └────────┬────────┘
                             │
                             │ repository contract
                             ▼
                    ┌─────────────────┐
                    │   REPOSITORY    │
                    │                 │
                    │ Query           │
                    │ Insert          │
                    │ Update          │
                    │ Delete          │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │    DATABASE     │
                    │   PostgreSQL    │
                    └────────┬────────┘
                             │
                             │ result
                             ▼
                    ┌─────────────────┐
                    │   REPOSITORY    │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │     ACTION      │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   CONTROLLER    │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ RESPONSE TRAIT  │
                    │                 │
                    │ Standard JSON   │
                    │ HTTP Status     │
                    └────────┬────────┘
                             │
                             ▼
                           CLIENT
```

---

# 17. Rule of Thumb

Gunakan aturan sederhana berikut saat mengembangkan endpoint baru:

```text
Apakah ini validasi input?
        │
        └── YES → Request

Apakah ini HTTP orchestration?
        │
        └── YES → Controller

Apakah ini business rule / use case?
        │
        └── YES → Action

Apakah ini database operation?
        │
        └── YES → Repository

Apakah ini format HTTP response?
        │
        └── YES → ResponseTrait
```

Dengan demikian, endpoint Goravel dapat dipertahankan dalam bentuk:

```text
Request
   ↓
Controller
   ↓
Action
   ↓
Repository
   ↓
Database

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

## Kesimpulan

Arsitektur ini menjaga setiap layer tetap fokus:

> **Request validates, Controller orchestrates HTTP, Action executes business rules, Repository handles persistence, dan Response Trait standardizes HTTP output.**

Tujuan akhirnya bukan sekadar membuat lebih banyak folder atau class, tetapi memastikan perubahan pada satu concern tidak menyebar ke seluruh aplikasi.

Contohnya:

```text
Perubahan database
    → Repository

Perubahan business rule
    → Action

Perubahan validasi
    → Request

Perubahan format API response
    → ResponseTrait

Perubahan endpoint / HTTP behavior
    → Controller
```

Itulah boundary yang harus dijaga ketika project Goravel berkembang menjadi aplikasi besar seperti SIMRS atau ERP.
