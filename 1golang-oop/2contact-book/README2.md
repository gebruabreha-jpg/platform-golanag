Explicit Dependencies:-
All dependencies are visible in constructors.
Applied in this project:
  Every NewXxx() function takes exactly what it needs as parameters.

  Examples:
    NewPostgresContactRepository(db *sql.DB)
    NewContactService(contacts repository.ContactRepository)
    NewContactHandler(contacts service.ContactServiceInterface)

  What this means:
  - You can see all dependencies at a glance
  - No hidden global state or singletons
  - No package secretly creates its own dependencies
  - Testing is easy — just pass a mock

  Anti-pattern (not used here):
    func Create() {
        db := sql.Open(...)  // hidden dependency
        repo := NewRepo(db)  // hidden dependency
    }

Dependencies are injected via constructors
ALL objects are created in main.go (the composition root):
cfg := config.Load()                          // Config
log := logger.New()                           // Logger
db, err := database.Connect(cfg.DatabaseURL)  // DB
contactRepo := repository.NewPostgresContactRepository(db)  // Repo
contactSvc := service.NewContactService(contactRepo)       // Service
contactHandler := handler.NewContactHandler(contactSvc)    // Handler



The * (pointer) is the standard Go convention for constructors. It means:
Allocate a new struct in memory
Return a pointer/mem add to it (not a copy)
The caller gets a reference to the same object



New signals "this function creates a new instance"
// Bad — unclear if this creates or modifies
PostgresContactRepository(db *sql.DB) *PostgresContactRepository

// Good — clear intent
NewPostgresContactRepository(db *sql.DB) *PostgresContactRepository


Go Constructor Naming Convention:-
New + TypeNameNew 
NewPostgresContactRepository
NewContactService
NewContactHandler




1. Object (Struct)
An object in Go is a struct — a collection of fields that hold data.

// This is an object (struct definition)
type ContactService struct {
    contacts repository.ContactRepository
}
This defines the shape of an object. It doesn't exist in memory yet — it's just a blueprint.

2. Pointer (*)
A pointer is a variable that stores the memory address of another variable.

// Without pointer — creates a COPY
svc := ContactService{contacts: repo}
// svc is a value type — passing it around copies all data

// With pointer — creates a REFERENCE
svc := &ContactService{contacts: repo}
// svc is *ContactService — passing it around shares the same object
The * means "pointer to this type." It's like giving someone the address of a house instead of building a new house.

3. Constructor (NewXxx)
A constructor is a function that creates a new object and returns a pointer to it.

// This is a constructor
func NewContactService(contacts repository.ContactRepository) *ContactService {
    return &ContactService{contacts: contacts}
}
It does 3 things:

Allocates memory for the struct
Sets the fields with the values you pass
Returns a pointer (*ContactService) to that memory
How They Differ
Concept	What it is	Example
Object (struct)	Blueprint/definition of data	type ContactService struct { ... }
Pointer (*)	Reference to memory address	*ContactService
Constructor (NewXxx)	Function that creates and initializes an object	NewContactService(repo)
In Your Project
// 1. Struct definition (blueprint) — no memory allocated yet
type ContactService struct {
    contacts repository.ContactRepository
}

// 2. Constructor (creates the object in memory, returns pointer)
func NewContactService(contacts repository.ContactRepository) *ContactService {
    return &ContactService{contacts: contacts}  // & means "pointer to"
}

// 3. Usage (object now exists in memory)
contactSvc := service.NewContactService(contactRepo)
// contactSvc is *ContactService — a pointer to the object
Why Pointers?
In Go, structs are value types by default. Without pointers:

svc1 := ContactService{contacts: repo}
svc2 := svc1  // COPIES all data — svc2 is a completely separate object
svc2.contacts = nil  // svc1 is NOT affected
With pointers:

svc1 := NewContactService(repo)
svc2 := svc1  // Both point to the SAME object
svc2.contacts = nil  // svc1.contacts is ALSO nil

conclusion:-
Composition Root pattern ✅
Dependency Inversion ✅
Interface-based repositories ✅
Constructor injection ✅
Framework-agnostic shared packages ✅
Parameterized SQL queries ✅
Sentinel errors ✅
DTO separation ✅