# About

dto is a library for creating data transfer objects (DTOs) in Go.
dto uses reflection to create a new struct with the same fields as the original struct.

it’s similar to entities but the differences are:

- DTO use as a struct to store objects from request body, request param, query param, or response object
- as a response object, DTO has a responsibility to respond to the entity object to its object and respond as the JSON in the handlers
- don’t use DTO as a variable on the repository, and vice versa. don’t use an entity as a request, or response object