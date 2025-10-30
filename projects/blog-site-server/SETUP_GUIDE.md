# Complete Setup & Run Guide

## 🎯 Complete Workflow: From Models to Running Server

### Step-by-Step Process

#### 1️⃣ **Update Your Models** (Already Done ✅)
Your Go models are in `internal/models/`:
- `blog.go` - Blog post structure
- `user.go` - User structure  
- `comment.go` - Comment structure

#### 2️⃣ **Configure Environment**

```bash
# Copy the example env file
cp .env.example .env

# Edit with your credentials
nano .env
```

Update to something like:
```properties
DATABASE_URL=postgres://myuser:mypassword@localhost:5432/blogdb
PORT=8080
```

#### 3️⃣ **Create PostgreSQL Database**

```bash
# Option A: Using createdb command
createdb blogdb

# Option B: Using psql
sudo -u postgres psql -c "CREATE DATABASE blogdb;"

# Option C: Interactive
sudo -u postgres psql
postgres=# CREATE DATABASE blogdb;
postgres=# \q
```

#### 4️⃣ **Create Tables from Models**

```bash
# Easiest way - run the setup script
./scripts/setup-db.sh
```

This automatically creates all tables matching your Go models!

#### 5️⃣ **Verify Database Schema**

```bash
# Connect to database
psql $DATABASE_URL

# List tables
\dt

# Check users table structure
\d users

# Check blogs table structure  
\d blogs

# Check comments table structure
\d comments

# Exit
\q
```

#### 6️⃣ **Install Go Dependencies**

```bash
go mod download
# or
go mod tidy
```

#### 7️⃣ **Run the Server**

```bash
go run main.go
```

You should see:
```
Database connection established
Starting blog site server on port 8080...
Available endpoints:
  POST   /blogs
  GET    /blogs
  ...
```

#### 8️⃣ **Test Your API**

```bash
# Open another terminal

# Test GET all blogs (should return empty array initially)
curl http://localhost:8080/blogs

# Create a user first
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "full_name": "John Doe",
    "email": "john@example.com",
    "password_hash": "$2a$10$...",
    "role": "author",
    "bio": "Tech blogger"
  }'

# Create a blog post
curl -X POST http://localhost:8080/blogs \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Blog Post",
    "content": "This is an amazing blog post about Go!",
    "cover_image": "https://example.com/image.jpg",
    "author_id": 1
  }'

# Get all blogs again
curl http://localhost:8080/blogs

# Get specific blog
curl http://localhost:8080/blogs/1
```

---

## 🔄 When You Update Models

If you modify your Go models, you need to update the database:

### Process:

1. **Update Go model** (e.g., add a field to `Blog`)
   ```go
   type Blog struct {
       // existing fields...
       Tags string `json:"tags"` // NEW FIELD
   }
   ```

2. **Create a new migration file**
   ```sql
   -- migrations/002_add_tags_to_blogs.sql
   ALTER TABLE blogs ADD COLUMN tags VARCHAR(255);
   ```

3. **Run the migration**
   ```bash
   psql $DATABASE_URL -f migrations/002_add_tags_to_blogs.sql
   ```

4. **Update repository if needed**
   ```go
   // internal/repository/blog.go
   // Update Create and GetByID methods to include Tags
   ```

5. **Restart server**
   ```bash
   go run main.go
   ```

---

## 🛠️ Quick Commands Reference

```bash
# Setup everything
cp .env.example .env          # Copy env template
nano .env                     # Configure database
./scripts/setup-db.sh         # Create tables
go run main.go                # Start server

# Development workflow
go mod tidy                   # Update dependencies
go fmt ./...                  # Format code
go build                      # Build binary
./blog-site-server            # Run binary

# Database commands
psql $DATABASE_URL            # Connect to DB
psql $DATABASE_URL -c "\dt"   # List tables
psql $DATABASE_URL -f file.sql # Run SQL file

# Testing
curl http://localhost:8080/blogs              # GET
curl -X POST http://localhost:8080/blogs ...  # POST
curl -X PUT http://localhost:8080/blogs/1 ... # PUT
curl -X DELETE http://localhost:8080/blogs/1  # DELETE
```

---

## 📂 Project Structure

```
blog-site-server/
├── .env                      # Your configuration (git-ignored)
├── .env.example              # Template for .env
├── main.go                   # Application entry point
├── go.mod                    # Go dependencies
├── migrations/               # SQL migration files
│   └── 001_initial_schema.sql
├── scripts/                  # Helper scripts
│   └── setup-db.sh          # Database setup script
├── internal/
│   ├── models/              # Go structs (your models)
│   │   ├── blog.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── repository/          # Database operations
│   │   ├── blog.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── handlers/            # HTTP handlers
│   │   ├── blog-handler.go
│   │   ├── user-handler.go
│   │   └── comment-handler.go
│   ├── routes/              # Route registration
│   │   └── routes.go
│   └── db/                  # Database connection
│       └── db.go
└── docs/
    ├── ARCHITECTURE.md      # Architecture guide
    └── DATABASE_SETUP.md    # This guide
```

---

## 🎓 Understanding the Flow

```
1. Go Models (structs)
   ↓
2. SQL Migration Files (create tables matching models)
   ↓
3. Repository Layer (CRUD operations using SQL)
   ↓
4. Handler Layer (HTTP endpoints)
   ↓
5. Routes (map URLs to handlers)
   ↓
6. Main (wire everything together)
```

**Your models define the structure**, and you create matching database tables using SQL migrations!

---

## ✅ Checklist

- [ ] PostgreSQL installed and running
- [ ] Database created
- [ ] `.env` file configured with DATABASE_URL
- [ ] Tables created (ran `./scripts/setup-db.sh`)
- [ ] Go dependencies installed (`go mod download`)
- [ ] Server starts successfully (`go run main.go`)
- [ ] Endpoints respond correctly (tested with curl)

---

## 🆘 Troubleshooting

**Server won't start?**
- Check `.env` file exists and DATABASE_URL is correct
- Verify PostgreSQL is running: `sudo systemctl status postgresql`
- Test database connection: `psql $DATABASE_URL`

**Tables not found?**
- Run setup script: `./scripts/setup-db.sh`
- Or manually: `psql $DATABASE_URL -f migrations/001_initial_schema.sql`

**Permission denied?**
- Make script executable: `chmod +x scripts/setup-db.sh`
- Check database user has privileges

Happy coding! 🚀
