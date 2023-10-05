# Cloud Engineering: Implementing and testing microservices in Go
In this example we implement microservices of a webshop. This includes
the following services:

* [User Service](src/user-service/): Authentication features like registration and login.
* [Product Service](src/product-service/): Holds detailed information about products like prices, sellers, etc.

## Developing

### Unit Testing
Mocks are generated using the [gomock](https://github.com/uber-go/mock)
framework and located at `_mocks` in every microservice source folder.
