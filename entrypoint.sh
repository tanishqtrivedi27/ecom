#!/bin/sh

# Run migrations
echo "Running migrations..."
/app/db

# Run the main application
echo "Starting application..."
/app/ecom
