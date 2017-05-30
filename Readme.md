# GoDataberus

GoDataberus is a "Database Driver API" in order to reduce the complexity of the CRUD operations in some common databases.
The idea is use Databerus in order to connect to a external DB (using an available DB driver) using the same API independently of the database  

# Architecture

GoDataberus has been developed in Golang and contains:
 - Connection Redis store: Redis to store the database's connection data. It return an uuid to use in the following api calls. The connection data could be a local database, external cloud provider, or whatever.
 - GoDataberus go package: API to use the basic CRUD with the common databases. It based on different drivers for the databases and implements the basic functions for these databases.

![Image of architecture](architecture.png)

# How to use GoDataberus
## API Definition

| Endpoint     | Method     | Description |
| :------------- | :------------- | :------------- |

## Available DB Backend connections
## Register a Database Backend connection
## Use the Backend with basic CRUD operations


# Build Dockerfile

```
TODO
```

# Run container

```
TODO
```

