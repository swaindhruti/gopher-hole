# Database Setup Guide

This guide explains how to sync your Go models with the PostgreSQL database.

## Quick Start (Recommended)

### Step 1: Configure `.env` File

Create or update your `.env` file in the project root:

```bash
# Example .env file
DATABASE_URL=postgres://username:password@localhost:5432/blogdb
PORT=8080
```

Replace:
- `username` - Your PostgreSQL username
- `password` - Your PostgreSQL password  
- `localhost` - Your database host (usually localhost)
- `5432` - PostgreSQL port (default is 5432)
- `blogdb` - Your database name

### Step 2: Create the Database (if not exists)

```bash
# Connect to PostgreSQL
sudo -u postgres psql

# Create database
CREATE DATABASE blogdb;

# Create user (if needed)
CREATE USER myuser WITH PASSWORD 'mypassword';
GRANT ALL PRIVILEGES ON DATABASE blogdb TO myuser;

# Exit
\q
```

### Step 3: Run the Setup Script

```bash
cd /path/to/blog-site-server

# Run the setup script
./scripts/setup-db.sh
```

This will automatically create all tables based on your models!

---

## Manual Setup (Alternative)

If you prefer to run SQL manually:

```bash
# From project root
cd /path/to/blog-site-server

# Run migration file
psql $DATABASE_URL -f migrations/001_initial_schema.sql

# Or connect and paste SQL
psql $DATABASE_URL
```

Then paste the contents of `migrations/001_initial_schema.sql`

---

## Verify Tables Were Created

```bash
# Connect to database
psql $DATABASE_URL

# List all tables
\dt

# Describe a specific table
\d users
\d blogs
\d comments

# Exit
\q
```

You should see:
```
 Schema |   Name   | Type  |  Owner
--------+----------+-------+---------
 public | blogs    | table | myuser
 public | comments | table | myuser
 public | users    | table | myuser
```

---

## Understanding the Model-to-Database Mapping

### Go Model → PostgreSQL Table

Your Go models are mapped to database tables as follows:

#### User Model (`internal/models/user.go`)
```go
type User struct {
    ID           int64     `json:"id"`           // → id SERIAL PRIMARY KEY
    Username     string    `json:"username"`     // → username VARCHAR(100)
    FullName     string    `json:"name"`         // → full_name VARCHAR(255)
    Email        string    `json:"email"`        // → email VARCHAR(255)
    Role         string    `json:"role"`         // → role VARCHAR(50)
    PasswordHash string    `json:"-"`            // → password_hash VARCHAR(255)
    Bio          string    `json:"bio"`          // → bio TEXT
    AvatarURL    string    `json:"avatar_url"`   // → avatar_url VARCHAR(500)
    CreatedAt    time.Time `json:"created_at"`   // → created_at TIMESTAMP
    UpdatedAt    time.Time `json:"updated_at"`   // → updated_at TIMESTAMP
    IsActive     bool      `json:"is_active"`    // → is_active BOOLEAN
}
```

#### Blog Model (`internal/models/blog.go`)
```go
type Blog struct {
    ID         int64      `json:"id"`          // → id SERIAL PRIMARY KEY
    Title      string     `json:"title"`       // → title VARCHAR(255)
    Content    string     `json:"content"`     // → content TEXT
    CoverImage string     `json:"cover_image"` // → cover_image VARCHAR(500)
    AuthorID   int64      `json:"author_id"`   // → author_id INTEGER
    CreatedAt  time.Time  `json:"created_at"`  // → created_at TIMESTAMP
    UpdatedAt  *time.Time `json:"updated_at"`  // → updated_at TIMESTAMP
}
```

#### Comment Model (`internal/models/comment.go`)
```go
type Comment struct {
    ID        int64     `json:"id"`         // → id SERIAL PRIMARY KEY
    PostID    int64     `json:"post_id"`    // → post_id INTEGER
    UserID    int64     `json:"user_id"`    // → user_id INTEGER
    Content   string    `json:"content"`    // → content TEXT
    CreatedAt time.Time `json:"created_at"` // → created_at TIMESTAMP
    UpdatedAt time.Time `json:"updated_at"` // → updated_at TIMESTAMP
}
```

---

## Common Issues & Solutions

### Issue 1: "DATABASE_URL not set"
**Solution**: Make sure your `.env` file exists and contains `DATABASE_URL`

### Issue 2: "database does not exist"
**Solution**: Create the database first:
```bash
sudo -u postgres createdb blogdb
```

### Issue 3: "permission denied"
**Solution**: Grant privileges:
```sql
GRANT ALL PRIVILEGES ON DATABASE blogdb TO your_user;
```

### Issue 4: "relation already exists"
**Solution**: Tables already exist. Either:
- Drop existing tables: `DROP TABLE comments, blogs, users CASCADE;`
- Or skip migration if schema matches

---

## Advanced: Auto-Migration with Code (Optional)

If you want to auto-create tables from Go code, you can add migration logic to your app. Here's an example:

Create `internal/db/migrate.go`:

```go
package db

import "log"

func RunMigrations() error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (...)`,
        `CREATE TABLE IF NOT EXISTS blogs (...)`,
        `CREATE TABLE IF NOT EXISTS comments (...)`,
    }
    
    for _, query := range queries {
        if _, err := DB.Exec(query); err != nil {
            log.Printf("Migration error: %v", err)
            return err
        }
    }
    
    log.Println("✅ Migrations completed")
    return nil
}
```

Then call it from `main.go`:
```go
db.InitializeDB()
db.RunMigrations()  // Add this
```

---

## Next Steps

After setting up the database:

1. ✅ Run the server: `go run main.go`
2. ✅ Test endpoints with curl or Postman
3. ✅ Start building your blog application!

```bash
# Start the server
go run main.go

# In another terminal, test it
curl http://localhost:8080/blogs
```
