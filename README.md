# Go Gin Project Structure

A well-organized project structure is essential when working with the Go Gin web framework. Below is a recommended structure for a Go Gin project, which can help you maintain a clean, scalable, and maintainable codebase.

## Recommended Project Structure

```plaintext
myapp/
├── cmd/
│   └── myapp/
│       ├── main.go
├── config/
│   ├── config.go
├── controllers/
│   ├── user_controller.go
│   └── auth_controller.go
├── models/
│   ├── user.go
│   └── auth.go
├── middlewares/
│   ├── auth_middleware.go
├── routes/
│   ├── router.go
├── services/
│   ├── user_service.go
│   └── auth_service.go
├── repositories/
│   ├── user_repository.go
│   └── auth_repository.go
├── utils/
│   ├── hash_util.go
│   └── jwt_util.go
├── docs/
│   └── swagger.yaml
├── go.mod
└── go.sum
```

## Explanation

- **cmd/**: Contains the application's entry point. It’s common to have a subfolder for each application (e.g., `myapp`).
    - `main.go`: The main function starts the Gin server and handles initialization (e.g., loading configuration, setting up routes, etc.).

- **config/**: Contains configuration-related files and functions.
    - `config.go`: Manages application configuration, like loading environment variables, setting up database connections, etc.

- **controllers/**: Contains the logic that handles incoming HTTP requests and produces responses.
    - `user_controller.go`: Manages user-related routes and logic.
    - `auth_controller.go`: Manages authentication-related routes and logic.

- **models/**: Defines the data models and database schema.
    - `user.go`: Contains the user model and database interactions.
    - `auth.go`: Contains models related to authentication, like tokens.

- **middlewares/**: Contains middleware functions that can be applied globally or to specific routes.
    - `auth_middleware.go`: Handles authentication checks before allowing access to certain routes.

- **routes/**: Defines the routing of the application.
    - `router.go`: Registers all the routes and associated controllers.

- **services/**: Contains the business logic that is used by controllers. It’s where most of the application’s logic should reside.
    - `user_service.go`: Handles business logic related to users.
    - `auth_service.go`: Handles authentication logic, such as generating tokens.

- **repositories/**: Manages data access, separating the data layer from the business logic. This can make testing easier and improve maintainability.
    - `user_repository.go`: Handles database operations related to users.
    - `auth_repository.go`: Handles authentication data operations.

- **utils/**: Contains utility functions that can be used throughout the project.
    - `hash_util.go`: Handles hashing operations (e.g., password hashing).
    - `jwt_util.go`: Manages JWT token creation and validation.

- **docs/**: Contains documentation files, such as API documentation.
    - `swagger.yaml`: API documentation in Swagger format.

- **go.mod**: The `go.mod` file manages dependencies and module information.

- **go.sum**: Contains the checksums of the module's dependencies.

## Additional Notes

- **Testing**: You may also want to create a `tests/` directory to keep your test files organized.
- **Database Migrations**: Consider adding a `migrations/` directory if you’re using tools like `golang-migrate` to manage database migrations.
- **Environment Variables**: Consider using a `.env` file for local environment variables, loaded using a library like `godotenv`.
