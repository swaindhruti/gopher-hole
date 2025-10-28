# Blog Site Server Architecture

## Project Structure

```
blog-site-server/
├── main.go                    # Application entry point
├── internal/
│   ├── db/                    # Database connection management
│   │   └── db.go
│   ├── models/                # Data structures (DTOs)
│   │   ├── blog.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── repository/            # Database operations (CRUD)
│   │   ├── blog.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── handlers/              # HTTP handlers (controllers)
│   │   ├── blog-handler.go
│   │   ├── user-handler.go
│   │   └── comment-handler.go
│   └── routes/                # Route registration
│       └── routes.go
└── .env                       # Environment variables
```

## Architecture Pattern: Repository Pattern

### Layer Responsibilities

#### 1. **Models Layer** (`internal/models/`)
- **Purpose**: Define data structures
- **Contains**: Struct definitions with JSON tags
- **No business logic or database operations**

```go
type Blog struct {
    ID         int64     `json:"id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    // ...
}
```

#### 2. **Repository Layer** (`internal/repository/`)
- **Purpose**: Database operations (CRUD)
- **Responsibilities**:
  - Execute SQL queries
  - Handle database transactions
  - Return models or errors
- **Key Features**:
  - Uses PostgreSQL parameterized queries (`$1`, `$2`)
  - Auto-generates timestamps
  - Returns IDs using `RETURNING` clause

```go
type BlogRepository struct {
    db *sql.DB
}

func (r *BlogRepository) Create(blog *models.Blog) error
func (r *BlogRepository) GetByID(id int64) (*models.Blog, error)
func (r *BlogRepository) GetAll() ([]*models.Blog, error)
func (r *BlogRepository) Update(blog *models.Blog) error
func (r *BlogRepository) Delete(id int64) error
```

#### 3. **Handler Layer** (`internal/handlers/`)
- **Purpose**: HTTP request/response handling
- **Responsibilities**:
  - Parse HTTP requests
  - Validate input
  - Call repository methods
  - Return JSON responses
  - Handle HTTP status codes

```go
type BlogHandler struct {
    repo *repository.BlogRepository
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request)
func (h *BlogHandler) GetBlog(w http.ResponseWriter, r *http.Request)
// ...
```

#### 4. **Routes Layer** (`internal/routes/`)
- **Purpose**: Route registration
- **Maps URLs to handlers**
- **Uses Go 1.22+ pattern matching**

```go
mux.HandleFunc("POST /blogs", blogHandler.CreateBlog)
mux.HandleFunc("GET /blogs/{id}", blogHandler.GetBlog)
```

#### 5. **Database Layer** (`internal/db/`)
- **Purpose**: Database connection management
- **Singleton pattern** for database pool
- **Provides** `InitializeDB()`, `CloseDB()`, `GetDB()`

## Dependency Injection Flow

```
main.go
  └─> db.InitializeDB()
  └─> database := db.GetDB()
  └─> blogRepo := repository.NewBlogRepository(database)
  └─> blogHandler := handlers.NewBlogHandler(blogRepo)
  └─> mux := routes.Setup(blogHandler, ...)
```

## API Endpoints

### Blog Endpoints
- `POST /blogs` - Create a new blog
- `GET /blogs` - Get all blogs
- `GET /blogs/{id}` - Get blog by ID
- `PUT /blogs/{id}` - Update blog
- `DELETE /blogs/{id}` - Delete blog

### User Endpoints
- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `PUT /users/{id}` - Update user
- `DELETE /users/{id}` - Delete user

### Comment Endpoints
- `POST /comments` - Create a new comment
- `GET /blogs/{blogID}/comments` - Get all comments for a blog
- `GET /comments/{id}` - Get comment by ID
- `PUT /comments/{id}` - Update comment
- `DELETE /comments/{id}` - Delete comment

## Key Optimizations

### 1. **PostgreSQL Compatibility**
- Changed `?` placeholders to `$1`, `$2`, etc.
- Uses `RETURNING` clause to get generated IDs

### 2. **Proper HTTP Handlers**
- JSON encoding/decoding
- HTTP status codes (200, 201, 204, 400, 404, 500)
- Path parameter extraction using `r.PathValue()`

### 3. **Consistent Naming**
- Repository methods: `Create`, `GetByID`, `GetAll`, `Update`, `Delete`
- Handler methods match HTTP operations
- Private fields in structs (lowercase `db`)

### 4. **Security**
- Password hashes never returned in API responses
- Input validation in handlers

### 5. **Auto Timestamps**
- `created_at` and `updated_at` managed by repository layer
- Uses `time.Now()` automatically

## Running the Server

```bash
# Set environment variables
export DATABASE_URL="postgres://user:password@localhost:5432/blogdb"
export PORT=8080

# Run the server
go run main.go
```

## Database Schema Required

```sql
CREATE TABLE blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    cover_image VARCHAR(500),
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    full_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    bio TEXT,
    avatar_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES blogs(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

## Benefits of This Architecture

✅ **Separation of Concerns** - Each layer has a single responsibility  
✅ **Testability** - Easy to mock repositories for testing  
✅ **Maintainability** - Changes isolated to specific layers  
✅ **Scalability** - Can add caching, service layer, etc.  
✅ **Clean Code** - Clear dependency flow  
✅ **Type Safety** - Strong typing throughout  
