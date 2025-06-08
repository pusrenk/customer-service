# PostgreSQL Setup

This repository contains a simple PostgreSQL setup using Docker Compose.

## Prerequisites

- Docker
- Docker Compose
- DBeaver (or any other PostgreSQL client)

## Getting Started

1. Start the PostgreSQL container:
```bash
docker-compose up -d
```

2. Connect to PostgreSQL using DBeaver:
   - Host: localhost
   - Port: 5432
   - Database: customer_service
   - Username: postgres
   - Password: postgres

## Connection Details

- **Host**: localhost
- **Port**: 5432
- **Database**: customer_service
- **Username**: postgres
- **Password**: postgres

## Stopping the Database

To stop the PostgreSQL container:
```bash
docker-compose down
```

To stop and remove all data (including the volume):
```bash
docker-compose down -v
```