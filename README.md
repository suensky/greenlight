# greenlight

## PostgreSQL
### Installation
```bash
brew install postgresql@16
```

### Run
Run once
```
LC_ALL="C" /opt/homebrew/opt/postgresql@16/bin/postgres -D /opt/homebrew/var/postgresql@16
```

Run as a backend service and restart at login
```
brew services start postgresql@16
```

### Setup (One-time)
1. Connect to PSQL
```
psql postgres
```
2. Create database
```
CREATE DATABASE greenlight;
```
3. Connect to the database
```
\c greenlight
```
4. Create a new user
```
create role greenlight with login password '<any password>';
```
5. Create an extension
```
create extension if not exists citext;
```

### Start and connect to DB
```
psql --host=localhost --dbname=greenlight --username=greenlight
```

### Exit DB
```
\q
# or
exit
```

### Migration
- Run migrations up
```bash
migrate -path=./migrations -database=<$DB_DSN> up
```

- Run migration down, i.e., rollback optional n versions. n=1, rollback the last version. n not exists: rollback all.
```bash
migrate -path=./migrations -database=<$DB_DSN> down <n>
```

- Check migration version
```bash
migrate -path=./migrations -database=<$DB_DSN> version
```

- Migrate up or down to a specific version, e.g, 1
```bash
migrate -path=./migrations -database=<$DB_DSN> goto 1
```

- Note 1, for PostgreSQL 15+, need to run `GRANT ALL ON SCHEMA public TO <dd_user>;` to avoid issues like 
```bash
error: pq: permission denied for schema public
```

- Note 2, if migration has dirty database errors, such as
```
error: Dirty database version 2. Fix and force version.
```
Fix by running `migrate -path=./migrations -database=<$DB_DSN> force <VERSION>` 