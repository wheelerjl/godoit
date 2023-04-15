# godoit

## Getting Started

- Starting the service
    - From the command line with environment variables
        - `GODOIT_PORT=8080 GODOIT_DEBUG=true GODOIT_ENVIRONMENT=dev go run main.go`, adding extra `KEY=VALUE`'s as needed
    - Example `.vscode/launch.json` file to run in debug mode with vscode's built in debugger
- Starting postgres
    - From the command line
        - `docker run --name postgres-local -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=godoit -p 5432:5432 -d postgres`
- Connecting to postgres from previous command
    - From the command line using psql
        - `psql -h localhost -p 5432 -U postgres -d godoit`

```JSON
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch", 
            "mode": "debug",
            "program": "${workspaceFolder}",
            "showLog": true,
            "env": {
                "GODOIT_PORT": "8080",
                "GODOIT_DEBUG": "true",
                "GODOIT_ENVIRONMENT": "dev",
                "GODOIT_DB_HOST": "localhost",
                "GODOIT_DB_USER": "postgres",
                "GODOIT_DB_PASS": "postgres",
            },
        }
    ]
}
```