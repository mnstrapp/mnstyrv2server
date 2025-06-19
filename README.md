# mnstr v2 server

## Development

### Requirements

- [Go](https://go.dev) 1.24.3+
- [Goose](https://github.com/pressly/goose?tab=readme-ov-file#goose) for database migrations
- [Direnv](https://direnv.net/docs/installation.html) (Optional) for os environment variables
- [PostgreSQL]()

### Configuration

Their are two ways to configure the server:

1. Using operating system environment variables
1. Using flags at runtime

The former requires you to setup the variables before runtime, the latter as flags at runtime.

#### OS Environment Variables

Here's a list of OS Env Variables that mnstr v2 server uses:

- **MNSTR_DATABASE_URL**: used to connect to a PostgreSQL server
- **MNSTR_HOST**: used as the host the server will run on
- **MNSTR_PORT**: used as the port the server will run on

We also have some Goose OS Env Variables to setup:

- **GOOSE_DBSTRING**: this should be the same as your **MNSTR_DATABASE_URL**
- **GOOSE_DRIVER**: this has to equal **postgres**
- **GOOSE_MIGRATION_DIR**: this has to equal **migrations**

How they are setup varies on the operating system. We only support *nix (Linux, Mac OS) at the moment, but are open to other OS testing.

##### Using Bash or Zsh

Add the following to your `.bashrc` (Bash) or `.zshrc` (Zsh):

```
# mnstr v2 server
export MNSTR_DATABASE_URL="postgresql://user:password@host:port/dbname"
export MNSTR_PORT=8080
export MNSTR_HOST=0.0.0.0

# Goose
export GOOSE_DBSTRING=$MNSTR_DATABASE_URL
export GOOSE_DRIVER=postgres
```

##### (Alternatively) Using Direnv

If you installed [Direnv](https://direnv.net/docs/installation.html), you can use a `.envrc` for setting up the application options. Just rename the `.envrc.example` file to `.envrc` and edit it with the correct values.

In your terminal, `cd` into the project directory and run: `direnv allow .`.

### Database Migrations

#### Install Requirements

We'll need to install the Goose CLI in order to manage database migrations. In your terminal run:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

#### Migration Management

We use Goose to manage the database migrations. All below commands assume that they are run in the terminal from the projects root directory.

##### Update Database

To update your database to the latest migrations run:

```
goose up
```

##### Create a Migration

To create a new database migration, run:

```
goose create <migration_name> sql
```

##### Rollback one Change

To rollback the latest database migration, run:

```
goose down
```

##### Rollback all Changes

To rollback all the database migrations, run:

```
goose down-to 0
```

##### Migration Status

To see what the latest database migration status in, run:

```
goose status
```

##### More Advanced Migrations

You can find all the migration possibillities (up-to, down-to, up-by-one, etc.) [on the Goose README](https://github.com/pressly/goose?tab=readme-ov-file#usage).

##### About Migrations

To see example database migrations, [visit the Migration section](https://github.com/pressly/goose?tab=readme-ov-file#migrations) on the Goose README.

### Building and Running Instructions

In order to build and run the server, the following steps need to be followed:

1. [Migrate the database](#database-migrations)
1. Build the server via: `go build -o build/mnstrv2server .`
1. Run the binary via: `./build/mnstrv2server`
