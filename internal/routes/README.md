# About

This package appears to contain several groups of routes for a web server application built with the Fiber framework in Go. These routes are organized into groups based on the type of functionality they provide.

The routes are grouped as follows:

- `/register:` This group contains a route for registering a new user.
- `/login:` This group contains a route for logging in a user.
- `/user:` This group contains routes for logging out, updating a user's profile, and showing a user's profile. These routes require authorization.
- `/products:` This group contains routes for getting the details of a product and showing all products.
- `/categories:` This group contains a route for getting all categories.
- `/store:` This group contains a route for registering a new store. These routes require authorization.
- `/login-store:` This group contains a route for logging in a store.
- `/balance:` This group contains routes for topping up a user's balance and getting a user's balance. These routes require authorization.
- `/product:` This group contains routes for creating a new product and showing a store's products. These routes require authorization.
- `/admin:` This group contains routes for showing all stores and getting a store by ID, as well as routes for creating and showing categories. These routes require authorization.

Each group is prefixed with /api and the routes within each group are specified with HTTP methods such as POST, PUT, and GET. Many of the routes also have middleware functions applied to them, which are responsible for handling authentication and authorization.




