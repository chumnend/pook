# Design

## Proposed Project Structure

```bash
pook
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go  # HTTP request handlers
│   ├── models
│   │   └── models.go    # Data structures and models
│   └── routes
│       └── routes.go    # Application routes setup
├── pkg
│   └── utils
│       └── utils.go     # Utility functions
├── web
|    |_ react            # Frontend application
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```
