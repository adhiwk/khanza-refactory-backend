# SIMRS Database Blueprint

> **Status:** Architecture Blueprint / Design Baseline
>
> **Audience:** Senior System Analyst, Solution Architect, Backend Engineer, DBA, Clinical Informatics, Integration Engineer
>
> **Target:** Scalable Hospital Information System (SIMRS) with Electronic Medical Record (RME), clinical services, pharmacy, inventory, finance, interoperability, audit, and analytics.
>
> **Recommended database:** PostgreSQL 16+ (the model is portable to other relational databases with minor adaptations).
>
> **Design principle:** Domain-driven modular relational design. The database is an operational source of truth (OLTP), while reporting/analytics should be separated progressively as volume grows.

---

## 1. Executive Design Decision

This blueprint does **not** model SIMRS as four independent systems called Rawat Jalan, Rawat Inap, IGD, and MCU.

Instead, the core is organized as:

```text
PATIENT / MPI
      |
      v
CARE EPISODE
      |
      v
ENCOUNTER
      |
      +--------------------+
      |                    |
      v                    v
CLINICAL ORDERS       DIRECT CLINICAL DATA
      |
      +-----------+-----------+-------------+
      |           |           |             |
      v           v           v             v
    LAB        RADIOLOGY    PHARMACY     PROCEDURE
      |           |           |             |
      v           v           v             v
   RESULTS     REPORTS     DISPENSE      RESULTS
      \           |           |             /
       +----------+-----------+------------+
                  |
                  v
                CHARGE
                  |
                  v
               INVOICE
                  |
                  v
               PAYMENT

Parallel cross-cutting concerns:
- Identity / authorization
- Terminology / master data
- Audit / provenance
- Integration / outbox / external identifiers
- Documents / object storage metadata
- Reporting / analytics
```

### 1.1 Core architectural decision

The following distinction is mandatory:

| Concept | Meaning |
|---|---|
| Patient | The person receiving care |
| Registration | Administrative registration for a visit/service request |
| Care Episode | A business/clinical episode grouping related care |
| Encounter | A concrete interaction with a provider/service |
| Order | A request for a clinical or operational service |
| Result | Output of an ordered service |
| Charge | Billable financial consequence of a service/event |
| Invoice | Financial document grouping charges |
| Payment | Settlement of an invoice/receivable |

This prevents the common mistake of making `rawat_jalan`, `igd`, `rawat_inap`, and `mcu` separate databases with duplicated patient, doctor, registration, billing, and clinical data.

---

# 2. Design Goals

The database must satisfy the following goals.

1. **Single patient identity** across all services.
2. **Encounter-centric clinical history**.
3. **Care delivery extensibility**: outpatient, inpatient, emergency, MCU, daycare, homecare, dialysis, surgery, rehabilitation, corporate health, etc.
4. **Immutable clinical history** where appropriate.
5. **Explicit state transitions**, not unexplained integer statuses.
6. **Strong auditability** for sensitive clinical and financial data.
7. **Transactional consistency** for clinical, inventory, and financial operations.
8. **Terminology normalization** and versioning.
9. **Interoperability-ready identifiers** without coupling internal schema to external APIs.
10. **Scalable indexing and partitioning strategy**.
11. **Operational database optimized for OLTP**.
12. **Reporting architecture that can evolve to read replicas, materialized views, CDC, or a warehouse**.
13. **No hard delete for finalized clinical/financial records**.
14. **No binary DICOM storage in the SIMRS relational database**.
15. **No external system identifier used as the internal primary key**.

---

# 3. Non-Goals

This blueprint deliberately does not prescribe:

- A specific frontend framework.
- A specific Laravel package.
- Microservices from day one.
- Kubernetes from day one.
- A full FHIR server implementation inside SIMRS.
- PACS implementation.
- LIS analyzer implementation.
- A data warehouse schema for every KPI.

The internal database remains the operational source of truth. External standards are integration contracts.

---

# 4. Recommended Technology Baseline

## 4.1 PostgreSQL

Recommended because SIMRS benefits from:

- ACID transactions.
- Foreign keys.
- Partial indexes.
- Expression indexes.
- JSONB for bounded/extensible payloads.
- Generated columns where appropriate.
- Range/date partitioning.
- Materialized views.
- Row-level security when justified.
- Strong concurrency behavior.
- Logical replication / CDC ecosystem.

## 4.2 Identifier strategy

Use UUIDv7 or ULID-like sortable identifiers for internal IDs.

Recommended pattern:

```text
id                -> internal immutable identifier
business_number   -> human-readable business identifier
external_*        -> external-system identifier
```

Example:

```text
id                  = 0199... sortable UUID
medical_record_no   = RM-00012345
satusehat_patient_id = 10000001
bpjs_member_no       = 0001234567890
```

Never use medical record number, NIK, BPJS number, accession number, invoice number, or SATUSEHAT ID as the primary key.

---

# 5. Database Schemas / Bounded Contexts

Recommended PostgreSQL schema separation:

```text
iam
org
master
patient
care
clinical
nursing
lab
radiology
pharmacy
inventory
billing
finance
insurance
mcu
integration
document
audit
reporting
```

A smaller deployment can initially use one PostgreSQL schema while retaining the same logical module boundaries. Physical schema separation is optional; domain separation is not.

---

# 6. Global Conventions

## 6.1 Common columns

Operational tables should normally use:

```text
id                UUID/UUIDv7       PK
created_at        timestamptz       NOT NULL
updated_at        timestamptz       NOT NULL
```

For business records, use explicit lifecycle fields where needed:

```text
status
cancelled_at
cancelled_by
cancel_reason
voided_at
voided_by
void_reason
```

Do not blindly put `deleted_at` on clinical records and call that auditability.

## 6.2 Foreign keys

Foreign keys should exist for internal relational integrity.

Recommended:

```text
ON DELETE RESTRICT
```

for clinical and financial relationships.

Use `ON DELETE CASCADE` only for truly owned child records such as draft configuration or line items whose lifecycle is entirely controlled by the parent.

## 6.3 Money

Use:

```text
numeric(18,2)
```

or another explicitly defined fixed precision.

Never use floating point for money.

## 6.4 Quantity

Use `numeric(20,6)` when quantities can require fractions.

## 6.5 Time

Use `timestamptz` for event timestamps.

Use `date` for date-only concepts.

Internally prefer UTC-aware timestamps. The UI converts to the facility timezone.

## 6.6 Code vs name

Do not use mutable display names as business keys.

Use:

```text
code
name
status
```

with unique constraints on stable codes.

---

# 7. IAM Domain

## 7.1 `iam.users`

```text
id                  UUID PK
username            varchar(100) unique
email               varchar(255) nullable
password_hash       text
status              varchar(30)
last_login_at       timestamptz nullable
created_at          timestamptz
updated_at          timestamptz
```

## 7.2 `iam.roles`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(150)
status              varchar(30)
created_at          timestamptz
updated_at          timestamptz
```

## 7.3 `iam.permissions`

```text
id                  UUID PK
code                varchar(150) unique
name                varchar(200)
module              varchar(100)
created_at          timestamptz
updated_at          timestamptz
```

Examples:

```text
clinical.note.create
clinical.note.sign
clinical.note.amend
lab.result.verify
radiology.report.verify
pharmacy.dispense
inventory.adjust
billing.invoice.finalize
finance.payment.post
```

## 7.4 `iam.user_roles`

```text
user_id             UUID FK iam.users
role_id             UUID FK iam.roles
organization_id     UUID FK org.organizations nullable
facility_id         UUID FK org.facilities nullable
valid_from          timestamptz nullable
valid_until         timestamptz nullable
PK (user_id, role_id, organization_id, facility_id)
```

## 7.5 `iam.role_permissions`

```text
role_id             UUID FK iam.roles
permission_id       UUID FK iam.permissions
PK (role_id, permission_id)
```

Authorization must be evaluated server-side. UI visibility is not security.

---

# 8. Organization / Facility Domain

## 8.1 `org.organizations`

```text
id                  UUID PK
code                varchar(50) unique
name                varchar(200)
organization_type   varchar(50)
status              varchar(30)
parent_id           UUID FK org.organizations nullable
```

Supports hospital groups or parent organizations.

## 8.2 `org.facilities`

```text
id                  UUID PK
organization_id     UUID FK org.organizations
code                varchar(50)
name                varchar(200)
facility_type       varchar(50)
timezone             varchar(64)
status              varchar(30)
```

## 8.3 `org.locations`

```text
id                  UUID PK
facility_id         UUID FK org.facilities
parent_id           UUID FK org.locations nullable
code                varchar(50)
name                varchar(200)
location_type       varchar(50)
status              varchar(30)
```

Examples:

```text
BUILDING
FLOOR
POLYCLINIC
WARD
ROOM
BED
LAB
RADIOLOGY
PHARMACY
OPERATING_ROOM
```

## 8.4 `org.providers`

Represents a healthcare professional identity.

```text
id                  UUID PK
user_id             UUID FK iam.users nullable
code                varchar(50) unique
full_name           varchar(200)
provider_type       varchar(50)
license_no          varchar(100) nullable
specialty_id        UUID nullable
status              varchar(30)
```

## 8.5 `org.provider_locations`

```text
provider_id         UUID FK org.providers
location_id         UUID FK org.locations
valid_from          date
valid_until         date nullable
PK (provider_id, location_id, valid_from)
```

---

# 9. Master / Terminology Domain

A SIMRS must not scatter terminology values throughout business tables.

## 9.1 `master.code_systems`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(200)
uri                 text nullable
version             varchar(100) nullable
status              varchar(30)
```

## 9.2 `master.codes`

```text
id                  UUID PK
code_system_id      UUID FK master.code_systems
code                varchar(150)
display              varchar(300)
definition           text nullable
status               varchar(30)
valid_from           date nullable
valid_until          date nullable
```

Unique:

```text
(code_system_id, code)
```

This table can represent local codes and mapped standards.

## 9.3 `master.code_mappings`

```text
id                  UUID PK
source_code_id      UUID FK master.codes
target_system_id    UUID FK master.code_systems
target_code         varchar(150)
target_display      varchar(300) nullable
mapping_status      varchar(30)
valid_from          date nullable
valid_until         date nullable
```

Important for mapping local services/diagnoses/medications to external standards.

SATUSEHAT currently validates terminology including ICD-10, ICD-9-CM, LOINC, and SNOMED CT for relevant interoperability flows, so terminology mapping should be a first-class concern rather than hardcoded strings. citeturn0search4

---

# 10. Patient / MPI Domain

## 10.1 `patient.patients`

```text
id                  UUID PK
medical_record_no   varchar(50) unique
status              varchar(30)
full_name           varchar(200)
normalized_name     varchar(200)
birth_date          date nullable
birth_place         varchar(150) nullable
sex_code            varchar(30) nullable
nationality_code    varchar(30) nullable
marital_status      varchar(30) nullable
address_id          UUID nullable
created_at          timestamptz
updated_at          timestamptz
```

Do not store calculated age.

## 10.2 `patient.identifiers`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
identifier_type     varchar(50)
identifier_value    varchar(150)
normalized_value    varchar(150)
is_primary           boolean
issuer               varchar(200) nullable
valid_from          date nullable
valid_until         date nullable
created_at          timestamptz
```

Examples:

```text
NIK
BPJS
PASSPORT
SATUSEHAT
LOCAL_MR
OTHER
```

Sensitive identifiers should be protected and access-controlled.

Do not assume every identifier is globally unique without defining the issuer/system.

## 10.3 `patient.contacts`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
contact_type        varchar(30)
name                varchar(200)
relationship_code   varchar(50) nullable
phone               varchar(50) nullable
email               varchar(255) nullable
is_emergency        boolean
```

## 10.4 `patient.addresses`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
address_type        varchar(30)
line1               text
line2               text nullable
city_code           varchar(50) nullable
province_code       varchar(50) nullable
postal_code         varchar(20) nullable
country_code        varchar(10)
valid_from          date nullable
valid_until         date nullable
```

## 10.5 `patient.photo_refs`

Store only object-storage metadata, not uncontrolled binary blobs in the core patient table.

```text
id                  UUID PK
patient_id          UUID FK patient.patients
object_key          text
content_type        varchar(100)
checksum            varchar(128) nullable
created_at          timestamptz
```

## 10.6 `patient.merge_events`

For MPI duplicate resolution.

```text
id                  UUID PK
survivor_patient_id UUID FK patient.patients
merged_patient_id   UUID FK patient.patients
reason              text
performed_by        UUID FK iam.users
performed_at        timestamptz
```

Never physically destroy the historical identity trail simply because two patient records were merged.

---

# 11. Care Delivery Domain

This is where Rawat Jalan, Rawat Inap, IGD, and MCU belong conceptually.

## 11.1 `care.service_lines`

```text
id                  UUID PK
code                varchar(50) unique
name                varchar(150)
service_type        varchar(50)
workflow_code       varchar(100)
status              varchar(30)
```

Examples:

```text
OUTPATIENT
INPATIENT
EMERGENCY
MCU
DAYCARE
HOMECARE
DIALYSIS
REHABILITATION
SURGERY
CORPORATE_HEALTH
```

A new service line must not require redesigning the Patient or Encounter model.

## 11.2 `care.registrations`

```text
id                  UUID PK
registration_no     varchar(50) unique
patient_id          UUID FK patient.patients
facility_id         UUID FK org.facilities
service_line_id     UUID FK care.service_lines
registered_at       timestamptz
registration_date   date
location_id         UUID FK org.locations nullable
payer_id            UUID nullable
referral_id         UUID nullable
priority_code       varchar(50) nullable
status              varchar(30)
created_by          UUID FK iam.users
created_at          timestamptz
updated_at          timestamptz
```

## 11.3 `care.episodes`

```text
id                  UUID PK
episode_no          varchar(50) unique
patient_id          UUID FK patient.patients
facility_id         UUID FK org.facilities
service_line_id     UUID FK care.service_lines
registration_id     UUID FK care.registrations nullable
started_at          timestamptz
ended_at            timestamptz nullable
status              varchar(30)
priority_code       varchar(50) nullable
created_at          timestamptz
updated_at          timestamptz
```

This is the grouping layer for complex services.

Examples:

```text
MCU Episode
IGD Episode
Inpatient Episode
Outpatient Episode
Dialysis Episode
```

## 11.4 `care.encounters`

```text
id                  UUID PK
encounter_no        varchar(50) unique
patient_id          UUID FK patient.patients
episode_id          UUID FK care.episodes
registration_id     UUID FK care.registrations nullable
facility_id         UUID FK org.facilities
location_id         UUID FK org.locations nullable
provider_id         UUID FK org.providers nullable
encounter_type      varchar(50)
class_code          varchar(50)
started_at          timestamptz
ended_at            timestamptz nullable
status              varchar(30)
priority_code       varchar(50) nullable
created_at          timestamptz
updated_at          timestamptz
```

The internal encounter model is intentionally richer than a visit table.

## 11.5 `care.encounter_participants`

```text
encounter_id        UUID FK care.encounters
provider_id         UUID FK org.providers
participant_role    varchar(50)
started_at          timestamptz nullable
ended_at            timestamptz nullable
PK (encounter_id, provider_id, participant_role)
```

---

# 12. Appointment and Queue

## 12.1 `care.appointments`

```text
id                  UUID PK
appointment_no      varchar(50) unique
patient_id          UUID FK patient.patients
facility_id         UUID FK org.facilities
location_id         UUID FK org.locations
provider_id         UUID FK org.providers nullable
service_line_id     UUID FK care.service_lines
scheduled_start     timestamptz
scheduled_end       timestamptz nullable
status              varchar(30)
source              varchar(50)
created_at          timestamptz
updated_at          timestamptz
```

## 12.2 `care.queue_tickets`

```text
id                  UUID PK
queue_no            varchar(30)
registration_id     UUID FK care.registrations
location_id         UUID FK org.locations
service_code        varchar(50)
issued_at           timestamptz
called_at            timestamptz nullable
served_at            timestamptz nullable
status              varchar(30)
priority_score       integer default 0
```

Queue numbering must not be the identity of the registration.

---

# 13. Inpatient Domain

## 13.1 `care.admissions`

```text
id                  UUID PK
episode_id          UUID FK care.episodes unique
admission_no        varchar(50) unique
admitted_at         timestamptz
admission_source    varchar(50)
admitting_provider_id UUID FK org.providers nullable
status              varchar(30)
discharge_at        timestamptz nullable
discharge_disposition varchar(50) nullable
```

## 13.2 `care.bed_assignments`

```text
id                  UUID PK
admission_id        UUID FK care.admissions
bed_location_id     UUID FK org.locations
assigned_at         timestamptz
released_at         timestamptz nullable
assignment_type     varchar(50)
status              varchar(30)
```

Do not overwrite the patient's previous bed assignment.

This supports bed history and transfers.

## 13.3 `care.transfers`

```text
id                  UUID PK
admission_id        UUID FK care.admissions
from_location_id    UUID FK org.locations
from_bed_id         UUID FK org.locations nullable
to_location_id      UUID FK org.locations
to_bed_id           UUID FK org.locations nullable
requested_at        timestamptz
completed_at        timestamptz nullable
reason              text nullable
status              varchar(30)
```

---

# 14. Emergency / IGD Domain

IGD requires additional triage and acuity data.

## 14.1 `clinical.triage_assessments`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
triage_time         timestamptz
triage_category     varchar(50)
chief_complaint     text
consciousness_code  varchar(50) nullable
assessed_by         UUID FK org.providers nullable
status              varchar(30)
```

## 14.2 `clinical.emergency_events`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
event_type          varchar(50)
event_time          timestamptz
performed_by        UUID FK org.providers nullable
notes               text nullable
```

Do not force all emergency workflow into generic outpatient fields.

---

# 15. Clinical Core

## 15.1 `clinical.problems`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters nullable
code_system_id      UUID FK master.code_systems
code                varchar(100)
display              varchar(300)
problem_type        varchar(50)
clinical_status     varchar(50)
verification_status varchar(50)
onset_at             timestamptz nullable
resolved_at          timestamptz nullable
recorded_by          UUID FK org.providers
created_at           timestamptz
```

This supports diagnoses/conditions while retaining local and external terminology.

## 15.2 `clinical.diagnoses`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
patient_id          UUID FK patient.patients
code_system_id      UUID FK master.code_systems
code                varchar(100)
display              varchar(300)
diagnosis_type       varchar(50)
sequence_no          integer
certainty_code       varchar(50) nullable
present_on_admission boolean nullable
recorded_by          UUID FK org.providers
recorded_at          timestamptz
status               varchar(30)
```

Examples:

```text
PRIMARY
SECONDARY
WORKING
FINAL
COMORBID
```

## 15.3 `clinical.procedures`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
patient_id          UUID FK patient.patients
procedure_code_id   UUID FK master.codes
performed_by        UUID FK org.providers nullable
performed_at        timestamptz
status              varchar(30)
body_site_code      varchar(100) nullable
notes               text nullable
```

## 15.4 `clinical.vital_signs`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
patient_id          UUID FK patient.patients
observed_at         timestamptz
recorded_by         UUID FK org.providers
systolic_bp         numeric(6,2) nullable
diastolic_bp        numeric(6,2) nullable
heart_rate          numeric(6,2) nullable
respiratory_rate    numeric(6,2) nullable
temperature         numeric(6,2) nullable
spo2                numeric(6,2) nullable
weight_kg            numeric(8,3) nullable
height_cm           numeric(8,3) nullable
bmi                 numeric(8,3) nullable
notes               text nullable
```

For highly extensible observations, use a separate observation model rather than adding hundreds of columns.

## 15.5 `clinical.observations`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters nullable
observation_code_id UUID FK master.codes
value_type          varchar(30)
value_numeric       numeric(20,8) nullable
value_text          text nullable
value_boolean       boolean nullable
value_code_id       UUID FK master.codes nullable
unit_code           varchar(50) nullable
observed_at         timestamptz
issued_at            timestamptz nullable
status               varchar(30)
recorded_by         UUID FK org.providers
```

Use JSONB only for bounded structured observations that do not deserve relational modeling; do not use JSONB as an excuse to avoid modeling core clinical data.

---

# 16. Clinical Documentation / EMR

## 16.1 `clinical.documents`

```text
id                  UUID PK
document_no         varchar(50) unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
document_type       varchar(50)
author_id            UUID FK org.providers
status               varchar(30)
version_no           integer
signed_at            timestamptz nullable
signed_by            UUID FK org.providers nullable
created_at           timestamptz
updated_at           timestamptz
```

## 16.2 `clinical.document_versions`

```text
id                  UUID PK
document_id         UUID FK clinical.documents
version_no          integer
content_json        jsonb
content_hash        varchar(128)
created_by          UUID FK iam.users
created_at          timestamptz
reason              text nullable
```

Signed/finalized clinical documents should be versioned rather than silently overwritten.

## 16.3 `clinical.amendments`

```text
id                  UUID PK
document_id         UUID FK clinical.documents
original_version_no integer
new_version_no      integer
reason              text
requested_by        UUID FK iam.users
approved_by         UUID FK iam.users nullable
created_at          timestamptz
```

---

# 17. Allergy Domain

## 17.1 `clinical.allergies`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
substance_code_id   UUID FK master.codes
reaction_code_id    UUID FK master.codes nullable
severity_code       varchar(50) nullable
verification_status varchar(50)
status              varchar(30)
onset_at             timestamptz nullable
recorded_by         UUID FK org.providers
recorded_at         timestamptz
```

Never model allergy as a free-text-only field if it will drive medication safety.

---

# 18. Clinical Orders

## 18.1 `clinical.orders`

Generic order header.

```text
id                  UUID PK
order_no            varchar(50) unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
ordered_by          UUID FK org.providers
order_type          varchar(50)
priority_code       varchar(50)
ordered_at          timestamptz
status              varchar(30)
notes               text nullable
```

Examples:

```text
LAB
RADIOLOGY
MEDICATION
PROCEDURE
NURSING
DIET
REFERRAL
```

## 18.2 `clinical.order_items`

```text
id                  UUID PK
order_id            UUID FK clinical.orders
service_code_id     UUID FK master.codes nullable
catalog_item_id     UUID nullable
sequence_no         integer
qty                  numeric(20,6) nullable
unit_code            varchar(50) nullable
instructions        text nullable
status              varchar(30)
```

Do not make all order domains share every possible field. Domain-specific tables should hold their specialized information.

---

# 19. Service Catalog

## 19.1 `master.service_catalog`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(250)
service_type        varchar(50)
clinical_domain     varchar(50)
status              varchar(30)
```

Examples:

```text
CONSULTATION
CBC
URINALYSIS
CHEST_XRAY
ECG
AUDIOMETRY
DRUG_DISPENSING
BED_DAILY_RATE
```

## 19.2 `master.service_prices`

```text
id                  UUID PK
service_id          UUID FK master.service_catalog
payer_id            UUID nullable
price               numeric(18,2)
currency             varchar(3)
effective_from      date
effective_until     date nullable
status              varchar(30)
```

Tariffs must be effective-dated. Never overwrite historical pricing that has already generated financial transactions.

---

# 20. Laboratory Domain

## 20.1 `lab.test_catalog`

```text
id                  UUID PK
service_id          UUID FK master.service_catalog
code                varchar(100) unique
name                varchar(250)
specimen_type_id    UUID nullable
result_type         varchar(50)
status              varchar(30)
```

## 20.2 `lab.orders`

```text
id                  UUID PK
clinical_order_id   UUID FK clinical.orders unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
status              varchar(30)
priority_code       varchar(50)
collected_at        timestamptz nullable
completed_at        timestamptz nullable
```

## 20.3 `lab.order_items`

```text
id                  UUID PK
lab_order_id        UUID FK lab.orders
test_catalog_id     UUID FK lab.test_catalog
sequence_no         integer
status              varchar(30)
```

## 20.4 `lab.specimens`

```text
id                  UUID PK
lab_order_id        UUID FK lab.orders
specimen_no         varchar(50) unique
specimen_type_code  varchar(100)
collected_at        timestamptz nullable
received_at         timestamptz nullable
collected_by        UUID FK org.providers nullable
received_by         UUID FK iam.users nullable
status              varchar(30)
rejection_reason    text nullable
```

SATUSEHAT models laboratory specimen data explicitly, including specimen type, collection time, receiving time, patient, encounter, and related request; therefore specimen should not be reduced to a boolean flag on the lab order. citeturn0search9

## 20.5 `lab.results`

```text
id                  UUID PK
lab_order_item_id   UUID FK lab.order_items
specimen_id         UUID FK lab.specimens nullable
result_code_id      UUID FK master.codes
value_type          varchar(30)
value_numeric       numeric(20,8) nullable
value_text          text nullable
value_code_id       UUID FK master.codes nullable
unit_code           varchar(50) nullable
reference_low       numeric(20,8) nullable
reference_high      numeric(20,8) nullable
abnormal_flag       varchar(30) nullable
observed_at         timestamptz nullable
verified_at         timestamptz nullable
verified_by         UUID FK org.providers nullable
status              varchar(30)
```

## 20.6 `lab.result_comments`

```text
id                  UUID PK
result_id           UUID FK lab.results
comment_text        text
created_by          UUID FK iam.users
created_at          timestamptz
```

Result lifecycle:

```text
ORDERED -> COLLECTED -> RECEIVED -> PROCESSING -> RESULTED -> VERIFIED -> FINAL
```

Cancellation and rejection are explicit branches, not magic status numbers.

---

# 21. Radiology / RIS / PACS Domain

The SIMRS relational database should manage the order, scheduling, procedure, report, and PACS references. DICOM objects themselves belong in PACS/object storage.

## 21.1 `radiology.orders`

```text
id                  UUID PK
clinical_order_id   UUID FK clinical.orders unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
accession_no        varchar(50) unique
priority_code       varchar(50)
status              varchar(30)
ordered_at          timestamptz
scheduled_at        timestamptz nullable
performed_at        timestamptz nullable
```

## 21.2 `radiology.order_items`

```text
id                  UUID PK
radiology_order_id  UUID FK radiology.orders
procedure_code_id   UUID FK master.codes
body_site_code      varchar(100) nullable
laterality_code     varchar(50) nullable
status              varchar(30)
```

## 21.3 `radiology.studies`

```text
id                  UUID PK
radiology_order_id  UUID FK radiology.orders
study_instance_uid  varchar(128) unique
accession_no        varchar(50)
modality            varchar(20)
study_description   varchar(300) nullable
performed_at        timestamptz nullable
pacs_system         varchar(50) nullable
pacs_reference      text nullable
status              varchar(30)
```

## 21.4 `radiology.series`

```text
id                  UUID PK
study_id            UUID FK radiology.studies
series_instance_uid varchar(128) unique
modality            varchar(20)
series_number       integer nullable
series_description  varchar(300) nullable
```

## 21.5 `radiology.reports`

```text
id                  UUID PK
study_id            UUID FK radiology.studies
report_no           varchar(50) unique
status              varchar(30)
findings            text
impression          text
reported_by         UUID FK org.providers
verified_by         UUID FK org.providers nullable
reported_at         timestamptz nullable
verified_at         timestamptz nullable
```

DICOM's imaging workflow includes concepts such as patient, service episode, imaging service request, procedure type, requested procedure, scheduled procedure step, and procedure plan. Keep those concepts in the RIS domain while storing only the DICOM identifiers/references in the SIMRS database. citeturn0search24

---

# 22. Pharmacy Domain

## 22.1 `pharmacy.medications`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(300)
generic_name        varchar(300) nullable
strength            varchar(100) nullable
dosage_form         varchar(100) nullable
route_code          varchar(50) nullable
status              varchar(30)
```

## 22.2 `pharmacy.prescriptions`

```text
id                  UUID PK
clinical_order_id   UUID FK clinical.orders unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
prescription_no     varchar(50) unique
prescribed_by       UUID FK org.providers
ordered_at          timestamptz
status              varchar(30)
```

## 22.3 `pharmacy.prescription_items`

```text
id                  UUID PK
prescription_id     UUID FK pharmacy.prescriptions
medication_id       UUID FK pharmacy.medications
qty                 numeric(20,6)
dose                varchar(100)
frequency           varchar(100)
duration            varchar(100)
route_code          varchar(50) nullable
instructions        text nullable
sequence_no         integer
status              varchar(30)
```

## 22.4 `pharmacy.dispensations`

```text
id                  UUID PK
dispensing_no       varchar(50) unique
prescription_id     UUID FK pharmacy.prescriptions
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
dispensed_by        UUID FK iam.users
dispensed_at        timestamptz
status              varchar(30)
```

## 22.5 `pharmacy.dispensation_items`

```text
id                  UUID PK
dispensation_id     UUID FK pharmacy.dispensations
prescription_item_id UUID FK pharmacy.prescription_items
inventory_item_id   UUID FK inventory.items
qty_prescribed      numeric(20,6)
qty_dispensed       numeric(20,6)
status              varchar(30)
```

Prescription and dispensing are intentionally separate. A prescription is an order; dispensing is an actual fulfillment event.

---

# 23. Inventory / Warehouse Domain

## 23.1 `inventory.items`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(300)
item_type           varchar(50)
base_unit_code      varchar(50)
status              varchar(30)
```

## 23.2 `inventory.warehouses`

```text
id                  UUID PK
facility_id         UUID FK org.facilities
location_id         UUID FK org.locations
code                varchar(50) unique
name                varchar(200)
status              varchar(30)
```

## 23.3 `inventory.item_lots`

```text
id                  UUID PK
item_id             UUID FK inventory.items
lot_no              varchar(100)
expiry_date         date nullable
manufacturer        varchar(200) nullable
status              varchar(30)
```

Unique:

```text
(item_id, lot_no)
```

## 23.4 `inventory.stock_balances`

```text
warehouse_id        UUID FK inventory.warehouses
item_id             UUID FK inventory.items
lot_id              UUID FK inventory.item_lots nullable
qty_on_hand         numeric(20,6)
qty_reserved        numeric(20,6)
updated_at          timestamptz
PK (warehouse_id, item_id, lot_id)
```

## 23.5 `inventory.stock_movements`

This is the inventory ledger.

```text
id                  UUID PK
warehouse_id        UUID FK inventory.warehouses
item_id             UUID FK inventory.items
lot_id              UUID FK inventory.item_lots nullable
movement_type       varchar(50)
quantity            numeric(20,6)
unit_cost           numeric(18,2) nullable
reference_type      varchar(100) nullable
reference_id        UUID nullable
occurred_at         timestamptz
created_by          UUID FK iam.users
notes               text nullable
```

Examples:

```text
PURCHASE_IN
TRANSFER_IN
TRANSFER_OUT
DISPENSE_OUT
RETURN_IN
ADJUSTMENT_IN
ADJUSTMENT_OUT
EXPIRED_OUT
DAMAGE_OUT
```

`stock_balances` is a performance projection/cache of the ledger. The movement ledger is the audit source for stock changes.

Do not rely on one mutable `qty` field as the only historical truth.

---

# 24. MCU Domain

MCU is a **service line / care delivery orchestration domain**, not a special kind of patient and not simply a polyclinic.

## 24.1 `mcu.packages`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(250)
description         text nullable
status              varchar(30)
valid_from          date nullable
valid_until         date nullable
```

## 24.2 `mcu.package_versions`

```text
id                  UUID PK
package_id          UUID FK mcu.packages
version_no          integer
price               numeric(18,2) nullable
effective_from      date
effective_until      date nullable
status              varchar(30)
```

Package composition must be versioned. Never change a historical package definition used by an existing MCU episode.

## 24.3 `mcu.package_items`

```text
id                  UUID PK
package_version_id  UUID FK mcu.package_versions
service_id          UUID FK master.service_catalog
sequence_no         integer
required             boolean
qty                  numeric(20,6)
```

## 24.4 `mcu.episodes`

```text
id                  UUID PK
care_episode_id     UUID FK care.episodes unique
package_version_id  UUID FK mcu.package_versions
company_id          UUID nullable
purpose_code        varchar(50)
status              varchar(30)
started_at          timestamptz
completed_at        timestamptz nullable
```

Examples of purpose:

```text
PRE_EMPLOYMENT
PERIODIC
CORPORATE
EXECUTIVE
GENERAL
DRIVER
FITNESS
```

## 24.5 `mcu.item_statuses`

```text
id                  UUID PK
mcu_episode_id      UUID FK mcu.episodes
package_item_id     UUID FK mcu.package_items
encounter_id        UUID FK care.encounters nullable
status               varchar(30)
started_at           timestamptz nullable
completed_at         timestamptz nullable
```

MCU workflow:

```text
MCU Registration
    -> MCU Episode
    -> Package Version
    -> Package Items
    -> Multiple Encounters / Clinical Services
    -> Results
    -> Physician Assessment
    -> Final MCU Conclusion
    -> Final Report
```

This lets one MCU package use laboratory, radiology, ECG, audiometry, spirometry, doctor assessment, and other services without duplicating those domains.

---

# 25. Nursing Domain

## 25.1 `nursing.assessments`

```text
id                  UUID PK
encounter_id        UUID FK care.encounters
patient_id          UUID FK patient.patients
assessment_type     varchar(50)
assessed_at         timestamptz
assessed_by         UUID FK org.providers
status              varchar(30)
content_json        jsonb
```

## 25.2 `nursing.care_plans`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters
created_by          UUID FK org.providers
status              varchar(30)
created_at          timestamptz
```

## 25.3 `nursing.care_plan_items`

```text
id                  UUID PK
care_plan_id        UUID FK nursing.care_plans
problem_code_id     UUID FK master.codes nullable
goal_text            text
intervention_text    text
status               varchar(30)
```

---

# 26. Referral Domain

## 26.1 `care.referrals`

```text
id                  UUID PK
referral_no         varchar(50) unique
patient_id          UUID FK patient.patients
source_encounter_id UUID FK care.encounters nullable
source_facility_id  UUID FK org.facilities
source_provider_id  UUID FK org.providers nullable
target_facility_id  UUID FK org.facilities nullable
target_provider_id  UUID FK org.providers nullable
reason               text
status               varchar(30)
referred_at          timestamptz
completed_at         timestamptz nullable
```

---

# 27. Insurance / Payer Domain

## 27.1 `insurance.payers`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(250)
payer_type          varchar(50)
status              varchar(30)
```

## 27.2 `insurance.coverages`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
payer_id            UUID FK insurance.payers
member_no           varchar(150)
plan_code           varchar(100) nullable
valid_from          date nullable
valid_until         date nullable
status              varchar(30)
```

## 27.3 `insurance.eligibility_checks`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
coverage_id         UUID FK insurance.coverages
encounter_id        UUID FK care.encounters nullable
checked_at           timestamptz
status               varchar(30)
external_reference  varchar(150) nullable
response_json       jsonb nullable
```

External responses may be retained in JSONB, but the core financial/clinical state should remain relational.

---

# 28. Billing Domain

Billing should be event/charge based.

```text
Clinical Event
     |
     v
   Charge
     |
     v
 Invoice
     |
     v
 Payment / Claim
```

## 28.1 `billing.charges`

```text
id                  UUID PK
charge_no           varchar(50) unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters nullable
episode_id          UUID FK care.episodes nullable
service_id          UUID FK master.service_catalog nullable
source_type         varchar(100) nullable
source_id           UUID nullable
quantity            numeric(20,6)
unit_price          numeric(18,2)
gross_amount        numeric(18,2)
discount_amount     numeric(18,2)
net_amount          numeric(18,2)
payer_id            UUID FK insurance.payers nullable
occurred_at         timestamptz
status              varchar(30)
```

The `source_type/source_id` pair should be supplemented by explicit foreign keys when the relationship is sufficiently stable. Polymorphic references are useful for event attribution but should not replace relational integrity everywhere.

## 28.2 `billing.invoices`

```text
id                  UUID PK
invoice_no          varchar(50) unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters nullable
episode_id          UUID FK care.episodes nullable
payer_id            UUID FK insurance.payers nullable
issued_at            timestamptz
subtotal             numeric(18,2)
discount_amount      numeric(18,2)
tax_amount           numeric(18,2)
total_amount         numeric(18,2)
status               varchar(30)
finalized_at         timestamptz nullable
```

## 28.3 `billing.invoice_items`

```text
id                  UUID PK
invoice_id          UUID FK billing.invoices
charge_id           UUID FK billing.charges
quantity            numeric(20,6)
unit_price          numeric(18,2)
amount              numeric(18,2)
```

Never mutate historical invoice totals casually. Corrections should use adjustment/credit/debit documents according to finance policy.

---

# 29. Finance / Payment Domain

## 29.1 `finance.payments`

```text
id                  UUID PK
payment_no          varchar(50) unique
patient_id          UUID FK patient.patients nullable
payer_id            UUID FK insurance.payers nullable
payment_method      varchar(50)
amount              numeric(18,2)
currency             varchar(3)
paid_at              timestamptz
reference_no        varchar(150) nullable
status              varchar(30)
received_by         UUID FK iam.users
```

## 29.2 `finance.payment_allocations`

```text
payment_id          UUID FK finance.payments
invoice_id          UUID FK billing.invoices
allocated_amount    numeric(18,2)
PK (payment_id, invoice_id)
```

A payment can be allocated to multiple invoices, and an invoice may be settled by multiple payments.

## 29.3 `finance.adjustments`

```text
id                  UUID PK
adjustment_no       varchar(50) unique
invoice_id          UUID FK billing.invoices
adjustment_type     varchar(50)
amount              numeric(18,2)
reason              text
approved_by         UUID FK iam.users nullable
created_by          UUID FK iam.users
created_at          timestamptz
status              varchar(30)
```

---

# 30. Claim Domain

## 30.1 `insurance.claims`

```text
id                  UUID PK
claim_no            varchar(100) unique
patient_id          UUID FK patient.patients
encounter_id        UUID FK care.encounters nullable
payer_id            UUID FK insurance.payers
invoice_id          UUID FK billing.invoices nullable
submitted_at        timestamptz nullable
accepted_at         timestamptz nullable
rejected_at         timestamptz nullable
claimed_amount      numeric(18,2)
approved_amount     numeric(18,2) nullable
status              varchar(30)
external_reference  varchar(150) nullable
```

## 30.2 `insurance.claim_events`

```text
id                  UUID PK
claim_id            UUID FK insurance.claims
event_type          varchar(50)
status              varchar(30)
message             text nullable
payload_json        jsonb nullable
occurred_at         timestamptz
created_by          UUID FK iam.users nullable
```

---

# 31. Document / File Metadata

Do not store large documents, DICOM objects, or arbitrary binary files directly in OLTP tables unless there is a clear reason.

## 31.1 `document.files`

```text
id                  UUID PK
patient_id          UUID FK patient.patients nullable
encounter_id        UUID FK care.encounters nullable
document_type       varchar(50)
object_provider     varchar(50)
object_key          text
content_type        varchar(150)
size_bytes          bigint
checksum            varchar(128) nullable
created_by          UUID FK iam.users
created_at          timestamptz
status              varchar(30)
```

Examples:

```text
Consent PDF
Referral Letter
Discharge Summary PDF
MCU Report PDF
Radiology Report PDF
Scanned Document
```

DICOM remains under PACS and is referenced by study/series identifiers.

---

# 32. Integration Domain

The integration database model is deliberately separated from clinical source tables.

## 32.1 `integration.systems`

```text
id                  UUID PK
code                varchar(100) unique
name                varchar(200)
system_type         varchar(50)
base_url             text nullable
status              varchar(30)
```

Examples:

```text
SATUSEHAT
BPJS
ORTHANC
LIS
PAYMENT_GATEWAY
SMS_GATEWAY
```

## 32.2 `integration.external_identifiers`

```text
id                  UUID PK
system_id           UUID FK integration.systems
entity_type         varchar(100)
entity_id           UUID
external_id         varchar(300)
external_type       varchar(100) nullable
created_at          timestamptz
updated_at          timestamptz
```

Unique:

```text
(system_id, entity_type, entity_id)
(system_id, entity_type, external_id)
```

This allows local IDs and external IDs to coexist without polluting every table with dozens of integration-specific columns.

## 32.3 `integration.outbox_events`

```text
id                  UUID PK
aggregate_type      varchar(100)
aggregate_id        UUID
event_type          varchar(150)
payload_json        jsonb
occurred_at         timestamptz
published_at        timestamptz nullable
attempt_count       integer default 0
last_error          text nullable
status              varchar(30)
```

The application writes business data and the outbox event in the same database transaction.

Then a worker publishes the event.

## 32.4 `integration.messages`

```text
id                  UUID PK
system_id           UUID FK integration.systems
message_type        varchar(100)
direction            varchar(20)
correlation_id      varchar(150) nullable
idempotency_key     varchar(200) nullable
request_payload     jsonb nullable
response_payload    jsonb nullable
status               varchar(30)
attempt_count        integer
sent_at              timestamptz nullable
completed_at         timestamptz nullable
last_error           text nullable
created_at          timestamptz
```

Unique idempotency keys should be used where external systems may retry requests.

SATUSEHAT is explicitly an interoperability ecosystem using HL7 FHIR and HTTPS REST APIs; therefore external resource IDs should be integration mappings, not replacements for internal database identity. citeturn0search0turn0search6

---

# 33. Audit Domain

Audit is not the same as application logs.

## 33.1 `audit.audit_logs`

```text
id                  UUID PK
occurred_at         timestamptz
user_id             UUID FK iam.users nullable
facility_id         UUID FK org.facilities nullable
request_id          varchar(100) nullable
session_id          varchar(150) nullable
action              varchar(50)
entity_type         varchar(100)
entity_id           UUID nullable
before_json         jsonb nullable
after_json          jsonb nullable
reason              text nullable
ip_address          inet nullable
user_agent          text nullable
```

Actions:

```text
CREATE
UPDATE
SIGN
AMEND
VOID
CANCEL
LOGIN
LOGOUT
EXPORT
VIEW_SENSITIVE
```

For high-risk patient data, consider dedicated access logging for read events as well as write events.

## 33.2 `audit.security_events`

```text
id                  UUID PK
occurred_at         timestamptz
user_id             UUID FK iam.users nullable
event_type          varchar(100)
severity             varchar(30)
ip_address           inet nullable
details_json        jsonb nullable
```

---

# 34. Reference Data for Units and Measurements

## 34.1 `master.units`

```text
id                  UUID PK
code                varchar(50) unique
name                varchar(100)
unit_type           varchar(50)
status              varchar(30)
```

## 34.2 `master.conversion_factors`

```text
id                  UUID PK
from_unit_id        UUID FK master.units
to_unit_id          UUID FK master.units
factor              numeric(30,12)
valid_from          date nullable
valid_until         date nullable
```

This is important for pharmacy and inventory.

---

# 35. Scheduling Domain

## 35.1 `care.schedules`

```text
id                  UUID PK
facility_id         UUID FK org.facilities
location_id         UUID FK org.locations
provider_id         UUID FK org.providers
service_line_id     UUID FK care.service_lines
weekday              smallint
start_time          time
end_time            time
capacity             integer
valid_from          date
valid_until         date nullable
status              varchar(30)
```

## 35.2 `care.schedule_exceptions`

```text
id                  UUID PK
schedule_id         UUID FK care.schedules
exception_date      date
exception_type      varchar(50)
start_time          time nullable
end_time            time nullable
capacity             integer nullable
reason              text nullable
```

---

# 36. Consent / Privacy Domain

## 36.1 `clinical.consents`

```text
id                  UUID PK
patient_id          UUID FK patient.patients
consent_type        varchar(100)
status              varchar(30)
given_at             timestamptz
withdrawn_at         timestamptz nullable
obtained_by         UUID FK iam.users
method              varchar(50)
document_id         UUID FK document.files nullable
```

Sensitive workflows should explicitly model consent where legally/operationally required.

---

# 37. Clinical State Machine Principles

Each important aggregate should have a documented state machine.

## 37.1 Encounter

```text
PLANNED
   |
   v
ARRIVED
   |
   v
IN_PROGRESS
   |
   v
COMPLETED
```

Possible terminal branches:

```text
CANCELLED
NO_SHOW
```

## 37. Clinical document

```text
DRAFT
  |
  v
SIGNED
  |
  v
FINAL
  |
  v
AMENDED
```

## 37. Lab

```text
ORDERED
 -> COLLECTED
 -> RECEIVED
 -> PROCESSING
 -> RESULTED
 -> VERIFIED
 -> FINAL
```

## 37. Radiology

```text
ORDERED
 -> SCHEDULED
 -> PERFORMED
 -> STUDY_AVAILABLE
 -> REPORTED
 -> VERIFIED
 -> FINAL
```

## 37. Prescription

```text
DRAFT
 -> SIGNED
 -> DISPENSING
 -> PARTIALLY_DISPENSED
 -> DISPENSED
```

## 37. Invoice

```text
DRAFT
 -> FINALIZED
 -> PARTIALLY_PAID
 -> PAID
```

Cancellation/voiding must be governed by business rules, not generic CRUD.

---

# 38. Why `Patient -> Episode -> Encounter` is the Core

This is the most important structural decision in this blueprint.

Example outpatient:

```text
Patient
  |
  +-- Registration
       |
       +-- Outpatient Episode
            |
            +-- Doctor Encounter
                 |
                 +-- Lab Order
                 +-- Radiology Order
                 +-- Prescription
```

Example IGD:

```text
Patient
  |
  +-- IGD Registration
       |
       +-- Emergency Episode
            |
            +-- Triage Encounter
            +-- Doctor Encounter
            +-- Procedure Encounter
            +-- Admission
```

Example inpatient:

```text
Patient
  |
  +-- Inpatient Episode
       |
       +-- Admission
       +-- Daily/Clinical Encounters
       +-- Nursing Care
       +-- Procedures
       +-- Pharmacy
       +-- Lab
       +-- Radiology
       +-- Discharge
```

Example MCU:

```text
Patient
  |
  +-- MCU Registration
       |
       +-- MCU Episode
            |
            +-- Doctor Encounter
            +-- Lab Encounter
            +-- Radiology Encounter
            +-- ECG Encounter
            +-- Audiometry Encounter
            +-- Final MCU Assessment
```

The same clinical services can therefore be reused across multiple care delivery models.

---

# 39. Master Data Hierarchy

A robust SIMRS should distinguish:

```text
Reference / Terminology
        |
        +-- ICD
        +-- SNOMED CT
        +-- LOINC
        +-- local codes

Clinical Catalog
        |
        +-- Services
        +-- Lab Tests
        +-- Radiology Procedures
        +-- Procedures
        +-- Medications

Commercial Catalog
        |
        +-- Tariffs
        +-- Payer prices
        +-- Package prices
```

Do not put price directly into the clinical master service record.

---

# 40. FHIR / SATUSEHAT Mapping Strategy

The internal model should be normalized for hospital operations, while integration mapping translates it into FHIR resources.

Conceptual mapping:

| Internal | FHIR-oriented resource |
|---|---|
| `patient.patients` | Patient |
| `org.providers` | Practitioner |
| `org.facilities` | Organization |
| `org.locations` | Location |
| `care.encounters` | Encounter |
| `care.episodes` | EpisodeOfCare |
| `clinical.observations` | Observation |
| `clinical.diagnoses` | Condition |
| `clinical.procedures` | Procedure |
| `clinical.orders` | ServiceRequest / domain resource |
| `lab.specimens` | Specimen |
| `lab.results` | Observation / DiagnosticReport |
| `radiology.studies` | ImagingStudy |
| `pharmacy.prescriptions` | MedicationRequest |
| `pharmacy.dispensations` | MedicationDispense |
| `billing.charges` | ChargeItem / related financial resources |
| `billing.invoices` | Invoice |
| `insurance.claims` | Claim |
| `document.files` | DocumentReference |

SATUSEHAT's published resource list includes Patient, Encounter, EpisodeOfCare, Observation, DiagnosticReport, ImagingStudy, MedicationRequest, MedicationDispense, ServiceRequest, Specimen, ChargeItem, Invoice, Claim, and related resources. citeturn0search1

This mapping is deliberately kept outside the transactional schema so changes in interoperability profiles do not force a redesign of the hospital core.

---

# 41. Indexing Strategy

Indexes must be designed from real access patterns.

## 41.1 Patient

```sql
CREATE UNIQUE INDEX ux_patients_mr
ON patient.patients (medical_record_no);

CREATE INDEX ix_patient_identifiers_lookup
ON patient.identifiers (identifier_type, normalized_value);
```

For name search, use PostgreSQL `pg_trgm` or a dedicated search strategy rather than `%name%` over millions of rows.

## 41.2 Encounter

Common indexes:

```text
(patient_id, started_at DESC)
(facility_id, started_at DESC)
(provider_id, started_at DESC)
(status, started_at DESC)
(episode_id)
```

## 41.3 Orders

```text
(patient_id, ordered_at DESC)
(encounter_id, ordered_at DESC)
(status, ordered_at)
```

## 41.4 Audit

```text
(occurred_at DESC)
(user_id, occurred_at DESC)
(entity_type, entity_id, occurred_at DESC)
(request_id)
```

## 41.5 Integration

```text
(status, occurred_at)
(system_id, status, created_at)
(idempotency_key)
```

Avoid creating an index for every foreign key automatically without checking workload; however, high-volume foreign-key columns used in joins/filtering generally need indexes.

---

# 42. Partitioning Strategy

Do not partition every table.

Partition only high-volume append-heavy tables where operational experience shows benefit.

Likely candidates:

```text
clinical.observations
clinical.vital_signs
clinical.document_versions
lab.results
integration.messages
integration.outbox_events
 audit.audit_logs
billing.charges
inventory.stock_movements
```

Recommended initial partitioning key:

```text
occurred_at / created_at by month
```

Example:

```text
audit_logs_2026_10
audit_logs_2026_11
audit_logs_2026_12
```

Keep partition management automated.

Do not partition simply because the database is expected to become large. Partitioning adds operational complexity.

---

# 43. Concurrency and Transaction Rules

## 43.1 Registration

Registration number generation must be concurrency-safe.

Never:

```text
SELECT MAX(number) + 1
```

Use database sequence, allocation table, or transactional number generator.

## 43.2 Inventory

For stock deduction:

```text
BEGIN

lock/select balance row
validate available quantity
insert stock movement
update stock balance

COMMIT
```

Use row locking or an equivalent concurrency control mechanism.

## 43.3 Payment

Payment posting and allocation must be atomic.

## 43.4 Clinical finalization

Finalization must validate:

- actor authorization;
- current state;
- required data;
- timestamps;
- related orders/results;
- audit event.

---

# 44. No Hard Delete Policy

## Can delete

Mostly configuration/draft data that has never entered a business transaction.

## Should not hard-delete

```text
Patient
Encounter
Diagnosis
Procedure
Lab Result
Radiology Report
Prescription
Dispensing
Charge
Invoice
Payment
Stock Movement
Audit Log
Integration Message
Signed Document
```

Use lifecycle states and correction/amendment events.

---

# 45. Data Integrity Rules

Examples:

### Patient

```text
medical_record_no must be unique
```

### Identifier

```text
identifier uniqueness depends on issuer/system
```

### Encounter

```text
patient_id must match episode.patient_id
```

### Order

```text
patient_id and encounter_id must refer to the same patient
```

### Result

```text
result patient must match order patient
```

### Dispensing

```text
dispensing patient must match prescription patient
```

### Charge

```text
charge patient must match source transaction patient
```

These are domain invariants and should be enforced in application/domain logic plus database constraints where practical.

---

# 46. Multi-Facility / Multi-Branch Scalability

If the hospital group may grow into multiple facilities, do not add `hospital_id` randomly to every table.

Define:

```text
organization
   |
   +-- facility
         |
         +-- location
```

Transactions should reference `facility_id` where operationally relevant.

This supports:

```text
Hospital Group
  |
  +-- Hospital A
  +-- Hospital B
  +-- Clinic C
  +-- Diagnostic Center D
```

A patient can exist at group level while encounters belong to a facility.

---

# 47. Multi-Tenant Considerations

If the system will become SaaS/multi-tenant, add a clear tenant boundary early:

```text
tenant_id
```

to every tenant-owned aggregate.

For a single hospital deployment, do not add unnecessary tenant complexity.

If multi-tenancy is a real requirement, PostgreSQL Row Level Security can be evaluated, but application authorization must remain explicit.

---

# 48. Reporting Architecture

Do not build large management reports by joining dozens of transactional tables on every request.

Evolution path:

```text
Phase 1
OLTP
  |
  +-- indexed reporting queries
  +-- materialized views

Phase 2
OLTP
  |
  +-- read replica
  |
  +-- reporting database

Phase 3
OLTP
  |
  +-- CDC / event pipeline
          |
          v
      Data Warehouse
          |
          v
         BI
```

Typical reporting facts:

```text
Fact Encounter
Fact Admission
Fact Procedure
Fact Lab Result
Fact Pharmacy Dispense
Fact Charge
Fact Payment
Fact Stock Movement
```

Dimensions:

```text
Dim Date
Dim Patient (privacy-aware)
Dim Provider
Dim Facility
Dim Service
Dim Payer
Dim Diagnosis
Dim Location
```

Do not expose the full clinical OLTP schema directly to BI users.

---

# 49. Read Model / Materialized Views

For high-traffic screens, create read models such as:

```text
reporting.patient_current_summary
reporting.daily_service_volume
reporting.daily_revenue
reporting.bed_occupancy
reporting.pharmacy_stock_status
reporting.lab_turnaround_time
reporting.radiology_turnaround_time
```

These are derived models, not sources of clinical truth.

---

# 50. Caching Strategy

Redis/cache may be used for:

```text
Reference data
Permission lookup
Queue display
Short-lived availability
Dashboard aggregates
Session data
```

Do not use cache as the authoritative source for:

```text
Patient identity
Diagnosis
Clinical notes
Lab results
Stock ledger
Payment ledger
```

---

# 51. Security Classification

Recommended data classes:

```text
PUBLIC
INTERNAL
CONFIDENTIAL
SENSITIVE_HEALTH
FINANCIAL
SECURITY_AUDIT
```

At minimum, protect:

- NIK
- insurance/member identifiers
- clinical notes
- diagnoses
- lab results
- imaging reports
- prescriptions
- payment information
- audit logs

Use least privilege and explicit access logging for sensitive records.

---

# 52. Backup / Recovery

Database design must be accompanied by operational recovery design.

Minimum:

```text
Full Backup
WAL / PITR
Off-site Backup
Backup Encryption
Restore Test
Retention Policy
RPO
RTO
```

Example target, to be agreed with hospital management:

```text
RPO <= 15 minutes
RTO <= 1 hour
```

These are examples, not universal requirements.

---

# 53. Downtime Architecture

SIMRS needs an operational downtime plan.

During database/application outage:

```text
Downtime Registration
Downtime Clinical Notes
Downtime Medication
Downtime Lab
Downtime Radiology
Downtime Billing
```

After recovery:

```text
Downtime Records
    -> Back Entry
    -> Validation
    -> Reconciliation
    -> Audit
```

Do not let downtime recovery overwrite original timestamps without preserving the actual event time and entry time.

---

# 54. Event Model

Important business events can be represented as:

```text
PatientRegistered
EncounterStarted
EncounterCompleted
DiagnosisRecorded
ClinicalDocumentSigned
LabOrdered
SpecimenCollected
LabResultVerified
RadiologyStudyPerformed
RadiologyReportVerified
PrescriptionSigned
MedicationDispensed
StockMoved
ChargeCreated
InvoiceFinalized
PaymentPosted
ClaimSubmitted
```

Events should be immutable facts.

Do not confuse event logs with audit logs:

```text
Business event = something happened in the business domain.
Audit event   = who did what to data/security state.
Application log = technical execution information.
```

Keep these concepts separate.

---

# 55. Recommended Aggregate Boundaries

The application should treat these as important aggregates:

```text
Patient
Registration
Care Episode
Encounter
Clinical Document
Order
Lab Order
Radiology Order
Prescription
Dispensation
Inventory Balance
Stock Movement
Charge
Invoice
Payment
Claim
MCU Episode
```

A single HTTP request should not casually modify ten unrelated aggregates without an explicit transaction/use case.

---

# 56. Laravel Mapping

For a Laravel backend, map modules approximately as:

```text
app/Domain/
├── Patient/
├── Organization/
├── Care/
├── Clinical/
├── Nursing/
├── Laboratory/
├── Radiology/
├── Pharmacy/
├── Inventory/
├── Billing/
├── Finance/
├── Insurance/
├── MCU/
├── Integration/
└── Audit/
```

A request flow can be:

```text
FormRequest
    |
    v
DTO
    |
    v
Application Action / Use Case
    |
    +---- Domain Service
    |
    +---- Repository
    |
    v
Transaction
    |
    +---- Domain changes
    +---- Audit
    +---- Outbox event
```

Do not place business rules in Eloquent models simply because they are convenient.

---

# 57. API Design Implication

Do not expose database tables directly as API resources.

Bad:

```text
GET /patients/{id}/tables/clinical_diagnoses
```

Better:

```text
GET /patients/{id}
GET /patients/{id}/encounters
GET /encounters/{id}
GET /encounters/{id}/clinical-summary
POST /encounters/{id}/orders
POST /lab/orders/{id}/collect
POST /lab/results/{id}/verify
POST /prescriptions/{id}/dispense
```

The API should express business operations and resource boundaries.

---

# 58. Naming Convention

Recommended PostgreSQL naming:

```text
snake_case
plural table names
singular FK concept + _id
created_at / updated_at
```

Examples:

```text
patients
care_episodes
encounters
lab_orders
stock_movements
```

If using PostgreSQL schemas:

```text
patient.patients
care.episodes
care.encounters
lab.orders
inventory.stock_movements
```

Do not mix:

```text
pasien
patients
tbl_patient
m_patient
```

within the same architecture.

---

# 59. Suggested Core ERD

```text
patient.patients
      |
      +----< care.registrations
      |             |
      |             +----< care.episodes
      |                       |
      |                       +----< care.encounters
      |                                  |
      |                                  +----< clinical.diagnoses
      |                                  +----< clinical.procedures
      |                                  +----< clinical.observations
      |                                  +----< clinical.orders
      |                                  |           |
      |                                  |           +----< lab.orders
      |                                  |           +----< radiology.orders
      |                                  |           +----< pharmacy.prescriptions
      |                                  |
      |                                  +----< clinical.documents
      |
      +----< clinical.allergies
      +----< insurance.coverages

clinical.orders
      |
      +---- lab.orders
      |       |
      |       +---- lab.order_items
      |       +---- lab.specimens
      |       +---- lab.results
      |
      +---- radiology.orders
      |       |
      |       +---- radiology.order_items
      |       +---- radiology.studies
      |       +---- radiology.reports
      |
      +---- pharmacy.prescriptions
              |
              +---- pharmacy.prescription_items
                       |
                       +---- pharmacy.dispensations
                                |
                                +---- inventory.stock_movements

clinical/procedure events
      |
      v
billing.charges
      |
      v
billing.invoices
      |
      v
finance.payments
```

---

# 60. MCU ERD

```text
care.service_lines
        |
        +---- MCU
              |
              v
        mcu.packages
              |
              v
        mcu.package_versions
              |
              v
        mcu.package_items
              |
              v
        mcu.episodes
              |
              +----< mcu.item_statuses
              |          |
              |          +---- encounter_id
              |
              +---- care.episodes
                       |
                       +---- care.encounters
                                  |
                                  +---- Lab
                                  +---- Radiology
                                  +---- ECG
                                  +---- Doctor Assessment
```

This makes MCU scalable because package composition is configuration/versioning, while clinical execution remains in shared clinical domains.

---

# 61. Rawat Jalan ERD

```text
Patient
  |
Registration
  |
Outpatient Episode
  |
Encounter
  |
+---- Diagnosis
+---- Vital Signs
+---- Clinical Note
+---- Orders
       |
       +---- Lab
       +---- Radiology
       +---- Prescription
       +---- Procedure
  |
Charges
  |
Invoice
  |
Payment / Claim
```

---

# 62. IGD ERD

```text
Patient
  |
IGD Registration
  |
Emergency Episode
  |
+---- Triage
+---- Emergency Encounter
+---- Doctor Encounter
+---- Procedure
+---- Lab/Radiology
+---- Disposition
          |
          +---- Discharge
          +---- Referral
          +---- Admission
```

---

# 63. Rawat Inap ERD

```text
Patient
  |
Inpatient Episode
  |
Admission
  |
Bed Assignment
  |
+---- Clinical Encounters
+---- Nursing Assessments
+---- Doctor Notes
+---- Medication
+---- Lab
+---- Radiology
+---- Procedures
+---- Transfers
  |
Discharge
  |
Discharge Summary
  |
Billing / Claim
```

---

# 64. Laboratory Lifecycle

```text
Order
  |
  v
Specimen Collection
  |
  v
Specimen Received
  |
  v
Processing
  |
  v
Resulted
  |
  v
Verified
  |
  v
Final
```

If an analyzer integration is introduced:

```text
Lab Order
   |
   v
Analyzer Queue
   |
   v
Analyzer Result
   |
   v
Result Validation
   |
   v
Verified Result
```

Analyzer raw messages belong in the integration layer; normalized clinical results belong in `lab.results`.

---

# 65. Radiology Lifecycle

```text
Radiology Order
     |
     v
Scheduled
     |
     v
Modality Worklist
     |
     v
Study Performed
     |
     v
PACS / Orthanc
     |
     v
Radiologist Report
     |
     v
Verified Report
```

The SIMRS should keep:

```text
patient
encounter
accession_no
study_uid
series_uid
report
pacs_reference
```

It should not become the PACS binary repository.

---

# 66. Pharmacy Lifecycle

```text
Prescription
     |
     v
Verification
     |
     v
Picking
     |
     v
Dispensing
     |
     v
Stock Movement
     |
     v
Charge
```

For controlled/high-risk medications, introduce additional authorization and witness rules according to hospital policy.

---

# 67. Financial Lifecycle

```text
Clinical/Operational Event
        |
        v
Charge
        |
        v
Invoice
        |
        +---- Cash/Bank Payment
        |
        +---- Insurance Claim
        |
        +---- Adjustment
```

The billing model must not mutate clinical history simply to make financial reconciliation easier.

---

# 68. Business Number Strategy

Use separate human-readable numbers for different documents.

Examples:

```text
MR-00000123
REG-20261003-000001
EP-20261003-000001
ENC-20261003-000001
LAB-20261003-000001
RAD-20261003-000001
RX-20261003-000001
DISP-20261003-000001
INV-20261003-000001
PAY-20261003-000001
CLM-20261003-000001
MCU-20261003-000001
```

The format can be configured by facility and document type, but the internal UUID remains immutable.

---

# 69. Sequence / Numbering Rules

Never generate numbers using application memory.

Bad:

```text
last_number + 1
```

Good:

```text
DB sequence
or
transaction-safe number allocation
```

For multi-facility systems, numbering may include facility code:

```text
RS01-REG-20261003-000001
RS02-REG-20261003-000001
```

The business number format should not become a relational dependency.

---

# 70. Data Retention

Retention must be defined by hospital policy and applicable Indonesian regulations.

Separate:

```text
Clinical retention
Financial retention
Audit retention
Integration retention
Application log retention
```

Do not apply one generic deletion job to all tables.

For very large audit/integration data, archive by time period.

---

# 71. Data Quality Rules

Implement data quality checks for:

```text
Duplicate patient
Invalid NIK
Duplicate encounter
Impossible timestamps
Result without specimen
Dispense without prescription
Stock negative without approved policy
Payment above invoice balance
Invoice without charges
Diagnosis without encounter
Final document without signer
External ID mapped to multiple local entities
```

These should become automated validation jobs/alerts.

---

# 72. Observability

Every API transaction should carry a `request_id`.

Recommended technical observability fields:

```text
request_id
correlation_id
user_id
facility_id
aggregate_type
aggregate_id
operation
latency_ms
status_code
error_code
```

This enables tracing:

```text
UI
 -> API
 -> DB
 -> Queue
 -> External API
```

without exposing technical logs as clinical records.

---

# 73. Outbox Pattern Example

When finalizing a lab result:

```text
BEGIN

UPDATE lab.results
SET status = 'VERIFIED'

INSERT INTO audit.audit_logs (...)

INSERT INTO integration.outbox_events (...)

COMMIT
```

Then:

```text
Worker
  -> reads outbox
  -> calls external system
  -> records integration message
  -> marks outbox published
```

If SATUSEHAT is unavailable, the lab result remains finalized locally and the integration worker retries.

SATUSEHAT documents its interoperability APIs as REST/FHIR and validates submitted resources, making asynchronous retry and integration observability particularly important for a production SIMRS. citeturn0search2turn0search4

---

# 74. Idempotency Rules

Every external operation that can be retried should have an idempotency strategy.

Examples:

```text
payment:create
claim:submit
satusehat:patient:create
satusehat:encounter:create
notification:send
```

Possible key:

```text
facility_code + business_number + operation
```

Store the key in `integration.messages` and reject/reuse the previous result on duplicate requests.

---

# 75. Anti-Patterns Explicitly Prohibited

## 75.1 One giant patient table

Avoid hundreds of nullable clinical columns.

## 75.2 Separate patient table per module

Never:

```text
lab_patients
radiology_patients
pharmacy_patients
```

## 75.3 Status integers without reference

Avoid:

```text
status = 1
status = 2
```

without a documented state model.

## 75.4 Hard deleting clinical records

Avoid.

## 75.5 Storing DICOM in MySQL/PostgreSQL

Avoid for normal PACS architecture.

## 75.6 One mutable stock quantity as the only truth

Avoid.

## 75.7 Billing directly from UI

All charges must be generated from trusted business events.

## 75.8 SATUSEHAT IDs as primary keys

Avoid.

## 75.9 JSONB for everything

Avoid.

## 75.10 Microservices before domain boundaries are stable

Avoid.

## 75.11 `SELECT MAX(number)+1`

Never use it for document numbering.

---

# 76. Recommended First Implementation Scope

If the project starts from zero, implement in this order.

## Phase 1 — Core identity

```text
IAM
Organization
Facility
Location
Provider
Master Terminology
Patient MPI
```

## Phase 2 — Core care

```text
Registration
Appointment
Queue
Care Episode
Encounter
```

## Phase 3 — Outpatient / IGD / Inpatient

```text
Clinical Notes
Diagnosis
Procedure
Vital Signs
Orders
Admission
Bed
Discharge
```

## Phase 4 — Diagnostic

```text
Laboratory
Radiology
PACS integration
```

## Phase 5 — Pharmacy / Inventory

```text
Medication
Prescription
Dispensing
Warehouse
Stock Ledger
```

## Phase 6 — Finance

```text
Charge
Invoice
Payment
Claim
```

## Phase 7 — MCU

```text
Package
Package Version
MCU Episode
Package Execution
MCU Final Assessment
```

## Phase 8 — Interoperability

```text
Outbox
External IDs
Integration Messages
SATUSEHAT
BPJS
```

## Phase 9 — Analytics

```text
Reporting DB
Read Models
Warehouse
BI
```

---

# 77. Recommended Development Rule

Do not implement every table as CRUD.

Implement business use cases.

Examples:

```text
RegisterPatient
RegisterVisit
StartEncounter
CompleteEncounter
RecordDiagnosis
SignClinicalNote
CreateLabOrder
CollectSpecimen
VerifyLabResult
CreateRadiologyOrder
FinalizeRadiologyReport
SignPrescription
DispenseMedication
TransferStock
CreateCharge
FinalizeInvoice
PostPayment
SubmitClaim
CompleteMCUEpisode
```

Each use case should own:

```text
validation
authorization
transaction boundary
state transition
audit
outbox/event
```

---

# 78. Testing Strategy

Database design is incomplete without test scenarios.

## Patient

- Duplicate patient detection.
- Identifier collision.
- Merge patient.
- Patient search at scale.

## Encounter

- Duplicate registration.
- Concurrent registration.
- Wrong patient/encounter relation.
- Encounter state transition.

## Lab

- Duplicate specimen collection.
- Result without specimen.
- Verify twice.
- Result amendment.

## Radiology

- Duplicate accession number.
- Duplicate study UID.
- PACS unavailable.
- Report amendment.

## Pharmacy

- Prescription quantity mismatch.
- Partial dispensing.
- Concurrent dispensing.
- Insufficient stock.
- Lot/expiry handling.

## Billing

- Duplicate charge.
- Invoice finalization.
- Partial payment.
- Overpayment.
- Reversal/adjustment.

## Integration

- External timeout.
- Retry.
- Duplicate event.
- Invalid response.
- External system unavailable.

---

# 79. Performance Targets

Actual targets must be established from workload tests, but useful starting engineering objectives are:

```text
Simple patient lookup             < 200 ms p95
Encounter detail                  < 300 ms p95
Order creation                    < 500 ms p95
Clinical finalization             < 500 ms p95
Dashboard                         asynchronous/read-model based
External integration              non-blocking where possible
```

These are engineering targets, not guaranteed performance numbers.

Performance must be tested with realistic data volumes.

---

# 80. Scale Model

The system should be able to evolve through these stages.

### Stage A — Single database

```text
Nuxt / Web
    |
Laravel API
    |
PostgreSQL
Redis
Queue
PACS
```

### Stage B — Read scaling

```text
Application
   |
Primary PostgreSQL
   |
Read Replica
```

### Stage C — Reporting separation

```text
PostgreSQL Primary
       |
       +---- Reporting DB
       +---- Read Replica
```

### Stage D — Event / analytics

```text
PostgreSQL
   |
CDC / Events
   |
Data Platform
   |
Warehouse / BI
```

The database model remains stable while infrastructure evolves.

---

# 81. Recommended PostgreSQL Extensions

Potentially useful:

```text
pgcrypto
pg_trgm
btree_gin
```

Use extensions intentionally and document them.

`pg_trgm` is particularly useful for patient name/identifier search patterns that cannot be handled efficiently by ordinary B-tree indexes.

---

# 82. Migration Strategy

Migrations should be incremental.

Recommended sequence:

```text
001 IAM
002 Organization
003 Master
004 Patient
005 Care
006 Clinical
007 Lab
008 Radiology
009 Pharmacy
010 Inventory
011 Billing
012 Finance
013 Insurance
014 MCU
015 Integration
016 Audit
017 Reporting
```

Do not create one 10,000-line migration containing the entire hospital database.

However, production deployment should still treat schema migration as a controlled release artifact.

---

# 83. Data Dictionary Requirement

For every production table, maintain a data dictionary containing:

```text
Field
Type
Nullable
Default
Description
Business Meaning
Source
Allowed Values
Validation
PII/PHI classification
Index
FK
Retention
Audit requirement
```

Example:

| Field | Meaning | Required | Notes |
|---|---|---:|---|
| `patient_id` | Internal patient identity | Yes | Never reused |
| `encounter_id` | Clinical interaction | Usually | Required for encounter-bound clinical data |
| `status` | Lifecycle state | Yes | Must have state transition rules |
| `created_at` | Record creation timestamp | Yes | Technical timestamp |
| `occurred_at` | Real-world event timestamp | Domain-specific | Must not be replaced by `created_at` |

---

# 84. Most Important Distinction: Created Time vs Clinical Time

Every important clinical system should distinguish:

```text
created_at
updated_at
occurred_at
performed_at
observed_at
recorded_at
verified_at
signed_at
```

Example:

```text
Blood sample collected:
08:10

Analyzer produced result:
08:37

Technician verified:
08:45

Doctor viewed:
09:05

Database row created:
09:06
```

These are not the same event.

A scalable clinical data model preserves the distinction.

---

# 85. Final Recommended Domain Map

```text
SIMRS
│
├── IAM
│   ├── Users
│   ├── Roles
│   └── Permissions
│
├── ORGANIZATION
│   ├── Organization
│   ├── Facility
│   ├── Location
│   └── Provider
│
├── MASTER
│   ├── Terminology
│   ├── Service Catalog
│   ├── Tariff
│   ├── Medication
│   └── Unit
│
├── PATIENT
│   ├── MPI
│   ├── Identifier
│   ├── Contact
│   ├── Address
│   └── Merge
│
├── CARE
│   ├── Registration
│   ├── Appointment
│   ├── Queue
│   ├── Episode
│   ├── Encounter
│   ├── Admission
│   ├── Bed
│   ├── Transfer
│   └── Referral
│
├── CLINICAL
│   ├── Diagnosis
│   ├── Problem
│   ├── Procedure
│   ├── Observation
│   ├── Vital Sign
│   ├── Allergy
│   ├── Clinical Order
│   └── Clinical Document
│
├── NURSING
│   ├── Assessment
│   ├── Care Plan
│   └── Nursing Documentation
│
├── LAB
│   ├── Order
│   ├── Specimen
│   ├── Result
│   └── Verification
│
├── RADIOLOGY
│   ├── Order
│   ├── Study
│   ├── Series
│   ├── Report
│   └── PACS Reference
│
├── PHARMACY
│   ├── Medication
│   ├── Prescription
│   ├── Dispensing
│   └── Dispensing Items
│
├── INVENTORY
│   ├── Item
│   ├── Warehouse
│   ├── Lot
│   ├── Balance
│   └── Movement Ledger
│
├── MCU
│   ├── Package
│   ├── Package Version
│   ├── Package Item
│   ├── Episode
│   └── Item Execution
│
├── BILLING
│   ├── Charge
│   ├── Invoice
│   └── Invoice Item
│
├── FINANCE
│   ├── Payment
│   ├── Allocation
│   └── Adjustment
│
├── INSURANCE
│   ├── Payer
│   ├── Coverage
│   ├── Eligibility
│   └── Claim
│
├── DOCUMENT
│   └── Object/File Metadata
│
├── INTEGRATION
│   ├── Systems
│   ├── External IDs
│   ├── Outbox
│   └── Messages
│
├── AUDIT
│   ├── Audit Log
│   └── Security Event
│
└── REPORTING
    ├── Read Models
    ├── Materialized Views
    └── Analytics Facts/Dimensions
```

---

# 86. Final Architectural Principles

If this blueprint becomes the baseline architecture, the following rules should be frozen.

### Rule 1 — One Patient Master

No module may create its own patient identity.

### Rule 2 — Episode and Encounter are different

Episode groups a course of care; encounter represents a concrete interaction.

### Rule 3 — Care Delivery is extensible

Rawat Jalan, Rawat Inap, IGD, MCU, Daycare, Homecare, Dialysis, etc. reuse the same core.

### Rule 4 — MCU is an orchestrator

MCU owns package/workflow configuration, while Lab, Radiology, ECG, Procedure, and clinical assessment remain reusable services.

### Rule 5 — Clinical data is not CRUD data

Finalized clinical records are versioned/audited and corrected through controlled workflows.

### Rule 6 — Orders are not results

An order requests work; a result records what actually happened.

### Rule 7 — Prescription is not dispensing

A prescription is intent; dispensing is fulfillment.

### Rule 8 — Stock is a ledger

Balance is a projection of stock movements, not the only historical truth.

### Rule 9 — Clinical events generate charges

Billing should not be tightly coupled to UI operations.

### Rule 10 — External systems do not own internal identity

SATUSEHAT, BPJS, PACS, LIS, and other IDs are mapped through integration identifiers.

### Rule 11 — Audit is immutable

Do not make audit logs editable/deletable through normal application workflows.

### Rule 12 — Integration is asynchronous where possible

Use transaction + outbox + worker + retry + idempotency.

### Rule 13 — Reporting is a separate concern

Do not turn the OLTP database into a BI engine.

### Rule 14 — Do not over-engineer infrastructure

Start with a disciplined modular monolith. Split components only when domain boundaries and scaling evidence justify it.

### Rule 15 — Database is not the entire architecture

The database must be aligned with business workflow, authorization, application services, integration contracts, audit, operations, and disaster recovery.

---

# 87. Architecture Review Checklist

Before declaring the SIMRS database production-ready, verify:

- [ ] Single MPI exists.
- [ ] Duplicate patient workflow exists.
- [ ] Patient merge workflow exists.
- [ ] Registration is separate from encounter.
- [ ] Episode is separate from encounter.
- [ ] RJ/RI/IGD/MCU are service lines, not isolated patient databases.
- [ ] Admission and bed history are modeled.
- [ ] Triage is modeled.
- [ ] Clinical notes are versioned.
- [ ] Clinical signing/amendment is audited.
- [ ] Diagnosis has terminology mapping.
- [ ] Orders and results are separate.
- [ ] Lab specimen lifecycle exists.
- [ ] Radiology has accession/study/report separation.
- [ ] DICOM is not stored as ordinary DB blobs.
- [ ] Prescription and dispensing are separate.
- [ ] Inventory has movement ledger.
- [ ] Lot and expiry are modeled where needed.
- [ ] Charge is separate from invoice.
- [ ] Invoice is separate from payment.
- [ ] Claim is separate from payment.
- [ ] Tariffs are effective-dated.
- [ ] MCU package versions are immutable for historical episodes.
- [ ] External identifiers are separated from internal IDs.
- [ ] Outbox exists for critical integration.
- [ ] Idempotency exists for retryable external operations.
- [ ] Audit exists.
- [ ] Sensitive read access can be audited where required.
- [ ] Reporting does not depend on uncontrolled ad-hoc joins.
- [ ] Backup and restore are tested.
- [ ] Downtime procedure exists.
- [ ] Data retention is defined.
- [ ] Indexes are workload-driven.
- [ ] Partitioning is evidence-driven.
- [ ] State machines are documented.
- [ ] Business invariants are tested.

---

# 88. Closing Architecture Statement

The scalable SIMRS database should be understood as a **clinical transaction platform**, not a collection of CRUD modules.

The central model is:

```text
                 PATIENT
                    |
                    v
             CARE EPISODE
                    |
                    v
                ENCOUNTER
                    |
          +---------+---------+
          |                   |
          v                   v
       CLINICAL             ORDERS
       RECORDS                |
                              +-------+-------+-------+
                              |       |       |       |
                              v       v       v       v
                             LAB     RAD    PHARM    PROCEDURE
                              |       |       |       |
                              +-------+-------+-------+
                                      |
                                      v
                                    CHARGE
                                      |
                                      v
                                   INVOICE
                                      |
                              +-------+-------+
                              |               |
                              v               v
                           PAYMENT          CLAIM
```

And the cross-cutting foundation is:

```text
IAM
MASTER DATA
TERMINOLOGY
AUDIT
DOCUMENT
INTEGRATION
OUTBOX
OBSERVABILITY
BACKUP/DR
REPORTING
```

The most important consequence is that **adding a new service such as MCU, dialysis, homecare, executive clinic, occupational health, or daycare does not require creating another copy of Patient, Registration, Encounter, Billing, Lab, Pharmacy, and Radiology.**

That is the architectural property that makes the database scalable at the domain level.

---

## References / Standards Context

- SATUSEHAT Platform documents that SATUSEHAT uses HL7 FHIR for data models and APIs. citeturn0search0
- SATUSEHAT publishes interoperability resources including Patient, Encounter, EpisodeOfCare, Observation, DiagnosticReport, ImagingStudy, MedicationRequest, MedicationDispense, ServiceRequest, Specimen, ChargeItem, Invoice, Claim, and others. citeturn0search1
- SATUSEHAT's interoperability guidance separates service modules such as outpatient, emergency, inpatient, and pharmacy from thematic use cases, which supports keeping the internal model extensible rather than hardcoding every care pathway into one table. citeturn0search3
- SATUSEHAT's current documentation describes periodic updates to interoperability and terminology guidance, reinforcing the value of an internal terminology/mapping layer rather than coupling the core schema directly to external profiles. citeturn0search8

> **Important:** This document is an architecture blueprint, not a substitute for current Indonesian regulatory/legal requirements, hospital clinical governance, or the latest SATUSEHAT/BPJS technical specifications. Those requirements should be validated during detailed design and before production release.
