# godoit

## Getting Started

- Starting the service
    - From the command line with environment variables
        - `GODOIT_PORT=8080 GODOIT_DEBUG=true GODOIT_ENVIRONMENT=dev go run main.go`
    - Example `.vscode/launch.json` file to run in debug mode with vscode's built in debugger

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