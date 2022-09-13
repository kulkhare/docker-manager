Docker Manger is an easy-to-install web application It provides a simple interface that enables you to manage your containers, applications, and images directly from your machine without having to use cli.

clone the repository

cd into the cloned repository and run ```go mod tidy``` after that run ```go run main.go```

You can change the the port to run on in ```.env``` default port is ```8087```

The project follow MVC pattern, all the routes are defined inside ```/routes/routes.go``` and all the controllers are defined inside ```/controllers``` folder