#!/usr/bin/env bash

# Exit on any error
set -e

# Function to log messages to stdout
log() {
    echo "[INFO] $1" >&1
}

# Function to log errors to stderr
error() {
    echo "[ERROR] $1" >&2
    exit 1
}

# Step 0: Check if running as root
if [ "$EUID" -ne 0 ]; then
    error "This script must be run as root"
fi

# Step 1: Perform git pull
log "Performing git pull..."
if ! git pull origin main; then
    error "Git pull failed"
fi

# Step 2: Perform go build
log "Building Go application..."
if ! go build -buildvcs=false -o app; then
    error "Go build failed"
fi

# Step 3: Setting appropriate permissions and ownerships
log "Setting proper folder ownership for www-data..."
if ! chown -R www-data:www-data /sites/rsvp; then
    error "Setting ownership to www-data for /sites/rsvp failed"
fi

# Step 4: Setting ownership of git to user: nomana
log "Setting ownership of git to nomana..."
if ! chown -R nomana:nomana /sites/rsvp/.git /sites/rsvp/.gitignore; then
    error "Changing ownership to nomana for .git failed"
fi

# Step 5: Restart app service
log "Restarting app service..."
if ! systemctl restart abantibiplob.fun.service; then
    error "Failed to restart app service"
fi
systemctl status --no-block abantibiplob.fun.service

# Step 6: Restart nginx service
log "Restarting nginx service..."
if ! systemctl restart nginx; then
    error "Failed to restart nginx service"
fi
systemctl status --no-block nginx

log "Deployment completed successfully"
