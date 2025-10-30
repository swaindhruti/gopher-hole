#!/bin/bash

# Database Setup Script for Blog Site Server
# This script creates the database schema from your Go models

set -e  # Exit on error

echo "🗄️  Blog Site Server - Database Setup"
echo "======================================"
echo ""

# Load environment variables from .env file
if [ -f .env ]; then
    echo "📝 Loading environment variables from .env..."
    export $(cat .env | grep -v '^#' | xargs)
else
    echo "❌ Error: .env file not found!"
    echo "Please create a .env file with DATABASE_URL"
    exit 1
fi

# Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
    echo "❌ Error: DATABASE_URL is not set in .env file"
    echo "Example: DATABASE_URL=postgres://user:password@localhost:5432/blogdb"
    exit 1
fi

echo "✅ Database URL configured"
echo ""

# Run the migration
echo "🚀 Running database migrations..."
echo ""

if command -v psql &> /dev/null; then
    # Use psql to run the migration
    psql "$DATABASE_URL" -f migrations/001_initial_schema.sql
    
    if [ $? -eq 0 ]; then
        echo ""
        echo "✅ Database schema created successfully!"
        echo ""
        echo "📊 Tables created:"
        echo "  • users"
        echo "  • blogs"
        echo "  • comments"
        echo ""
        echo "🎉 Your database is ready to use!"
    else
        echo ""
        echo "❌ Error running migrations"
        exit 1
    fi
else
    echo "❌ Error: psql command not found"
    echo "Please install PostgreSQL client tools"
    exit 1
fi
