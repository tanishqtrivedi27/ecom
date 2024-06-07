#!/bin/sh

echo "Running database migrations..."
/app/db

echo "Starting the application..."
/app/ecom
