# Jobs REST API
In this project, we develop A **REST API** with Go (or Golang) that performs CRUD (i.e., CREATE, READ, UPDATE AND DELETE) operations on jobs data saved on a PostgreSQL database.  

**Tech Stack:** Go, PostgreSQL, Postman.

The project consists of these **qualities**:  
- We work with structs and slices to create our model and store and pass data around.  
- We create a router to register our API endpoints and handler functions.
- We make use of technologies such as the PostgreSQL database for persisting records, Postman for making API requests, third party packages such as gorilla/mux for routing, and other packages for loading environment variables, making http requests, encoding and decoding JSON, performing string conversions and checking types.
- We refactor our code and create a modular file structure.

## Environment Setup
To run the application on your local machine, set up your **software** environment first following the steps below.  

(1) Install [Go](https://go.dev/doc/install)  

(2) Install Go third party packages  
-	Get [gorilla/mux](https://github.com/gorilla/mux)
-	Get [gotenv](https://github.com/subosito/gotenv)
-	Get [pq](https://github.com/lib/pq)

(3) Get [Postman](https://www.postman.com/) desktop agent  

(4) Get [ElephantSQL](https://www.elephantsql.com/)  
-	Set up a (free) database instance
-	Create a new table in your instance called `jobs` with PostgreSQL syntax
-	Create a `.env` file in your project folder
-	Create a variable in your .env file called `ELEPHANTSQL_URL` and assign to it your database URL in quotes


## Run Application

To start and test the application:

### 1. Build and Run

Launch the application by running the following command in your terminal:

```bash
go build && ./jobs-list
```

### 2. Access the API

You can access the API in two ways:

#### Browser
Open your web browser and navigate to:
```
http://localhost:8000/jobs
```

#### Postman
Alternatively, use Postman to send requests:
```http
GET http://localhost:8000/jobs
```

Note: The application runs on port 8000 by default. Make sure this port is available on your system.


## Examples

We recommend using Postman to make requests to the REST API. You can use either the desktop agent or browser version to send GET, POST, PUT, or DELETE requests.

### Create a Job

To create a new job posting:

```http
POST http://localhost:8000/jobs
```

Request body:
```json
{
    "title": "Software Engineer",
    "company": "LLM AI",
    "location": "San Francisco, CA",
    "type": "Full-time"
}
```

The server will return a response containing the newly created job's ID (assumed to be 8 in subsequent examples).

### Read Jobs

To retrieve all jobs:

```http
GET http://localhost:8000/jobs
```

To retrieve a specific job by ID:

```http
GET http://localhost:8000/jobs/1
```

### Update a Job

To update an existing job:

```http
PUT http://localhost:8000/jobs
```

Request body:
```json
{
    "id": 8,
    "title": "Machine Learning Engineer",
    "company": "LLM AI",
    "location": "New York, NY",
    "type": "Full-time"
}
```

### Delete a Job

To delete a job by ID:

```http
DELETE http://localhost:8000/jobs/8
```

For more detailed examples of using our API with Postman, check out our [demo](https://github.com/nabilshadman/go-rest-api-jobs-list-postgres/tree/main/demo).


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE.txt) file for details.

## Citation  
If you use this work in your research, please cite:  

```bibtex  
@misc{jobs-rest-api,
  author = {Shadman, Nabil},
  title = {Jobs REST API with Go and PostgreSQL},
  year = {2021},
  publisher = {GitHub},
  url = {https://github.com/nabilshadman/go-rest-api-jobs-list-postgres}
}
```
