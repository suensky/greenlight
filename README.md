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