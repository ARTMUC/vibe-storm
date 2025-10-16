# Local Development Setup

This guide explains how to set up the local development environment using Docker Compose for the MySQL database.

## Prerequisites

- Docker and Docker Compose installed on your machine

## Starting the Local Database

To start the MySQL database for local development:

```bash
docker-compose up -d
```

This will start a MySQL container with the following configuration:
- Database name: `vibe_storm`
- Root password: `password`
- User: `vibe_user`
- Password: `vibe_password`
- Port: `3306` (exposed to host)

## Connecting to the Database

You can connect to the database using any MySQL client:

```bash
mysql -h localhost -P 3306 -u vibe_user -p vibe_storm
```

Or use the root user:
```bash
mysql -h localhost -P 3306 -u root -p
```

## Data Persistence

The database data is persisted in a Docker volume named `mysql_data`. This means that even if you stop and remove the container, your data will be preserved when you start it again.

To see the volumes:
```bash
docker volume ls
```

To remove the volume and all data:
```bash
docker-compose down -v
```

## Environment Variables

Update your `.env` file with the following database configuration:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=vibe_user
DB_PASSWORD=vibe_password
DB_NAME=vibe_storm
```

## Stopping the Database

To stop the database:

```bash
docker-compose down
```

To stop and remove the data volume (this will delete all data):

```bash
docker-compose down -v
