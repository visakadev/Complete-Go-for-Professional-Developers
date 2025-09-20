# Complete-Go-for-Professional-Developers

To run the application, use:
```sh
go run main.go
```

docker compose up --build

export PATH=$HOME/go/bin:$PATH

## Install goose migration tool:
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

> **Troubleshooting goose installation:**
> If you get `zsh: command not found: goose` after installation:
> 1. Check if goose is installed: `ls -l ~/go/bin | grep goose`
> 2. Verify GOPATH: `go env GOPATH`
> 3. Add to PATH permanently by adding this to your `~/.zshrc`:
>    ```sh
>    export PATH=$HOME/go/bin:$PATH
>    ```
> 4. Reload shell: `source ~/.zshrc`
> 5. Or run directly: `~/go/bin/goose`

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

add a new migration

```
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
```


handler, database, SQL, APP, route.GO