# Complete-Go-for-Professional-Developers

To run the application, use:
```sh
go run main.go
```

docker compose up --build

export PATH=$HOME/go/bin:$PATH

ls -l ~/go/bin | grep goose

psql -U postgres -h localhost -p 5432

> **Note:**
> If you see `zsh: command not found: psql`, you need to install PostgreSQL client tools.
> - On macOS: `brew install postgresql`
> - On Ubuntu: `sudo apt-get install postgresql-client`
> - On Windows: Download from https://www.postgresql.org/download/

Test:
ls cd store
go test .