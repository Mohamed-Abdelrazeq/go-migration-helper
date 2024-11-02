# Go Migration Helper

Go Migration Helper is a tool to manage database migrations for your Go projects. It provides a simple interface to initialize migrations, add new migration files, migrate all files, roll back recent changes, and reset the database.

## Features

- **Init Migration**: Initialize a new migration.
- **Add New Migration File**: Add a new migration file.
- **Migrate All Files**: Apply all pending migrations.
- **Roll Back Recent File**: Roll back the most recent migration.
- **Reset the DB**: Reset the database to its initial state.

## Usage

### Init Migration

To initialize a new migration, run:
```sh
go-migration-helper init
```

### Add New Migration File

To add a new migration file, run:
```sh
go-migration-helper add <filename>
```

### Migrate All Files

To apply all pending migrations, run:
```sh
go-migration-helper migrate
```

### Roll Back Recent File

To roll back the most recent migration, run:
```sh
go-migration-helper rollback
```

### Reset the DB

To reset the database to its initial state, run:
```sh
go-migration-helper reset
```

## Installation

To install Go Migration Helper, use:
```sh
go get github.com/yourusername/go-migration-helper
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss what you would like to change.

## License

This project is licensed under the MIT License.
