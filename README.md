# REST API Go

This Go REST API offers a streamlined approach to developing RESTful services with minimal dependencies, making it an excellent starting point for any new project. 
By leveraging Go's robust standard library and efficient concurrency model, this project provides a solid foundation for building scalable web applications.

## Project

The structure of the API is modular, with each component encapsulated into smaller services. Each service handles a specific function or domain within the application, allowing for greater maintainability and scalability. 
Services are independently packaged as main packages to allow flexible integration and deployment strategies without enforcing a rigid project structure.

Example: 
```
.
users.go      // Service for users
users_test.go // Tests for the users service
```

## Run
To initiate the project on your local machine, ensure you have Go installed. 
The project can be started using a simple command which executes a Makefile that defines the necessary steps to get the service up and running. 
This command not only builds the application but also starts it:

```bash
make run 
// or Docker
```

Before running the project, ensure that all necessary environment variables are correctly configured. 
These variables are defined in the config.go file and should be set or injected at runtime. 
For environment variable management, it is recommended to use direnv, which allows for dynamic loading of environment settings based on the current directory.

## Test
To validate the integrity and functionality of the application, run the included tests using the command:

```bash
make test
```

## Deploy
Deployment involves building a Docker image and deploying it to a cloud provider of your choice. 
This process is simplified through the use of Docker, allowing for consistent deployment environments and easy scalability.

```bash
docker build -t golang .
docker run -p 8080:8080 golang
```

Deploy your Docker container to your preferred cloud platform to make your API accessible over the internet.
