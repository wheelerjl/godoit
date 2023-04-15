# godoit

## Getting Started

- Start the service with `go run main.go`
    - Exampple starting the service with shell environemnt variables
        - Start the service with `GODOIT_PORT=8080 GODOIT_DEBUG=true GODOIT_ENVIRONMENT=dev go run main.go`
    - Example `.vscode/launch.json` file to run in godoit in debug mode with the built in debugger

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
            },
        }
    ]
}
```