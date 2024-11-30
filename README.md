# Jobs REST API
In this project, we develop A **REST API** with Go (or Golang) that performs **CRUD** (i.e., CREATE, READ, UPDATE AND DELETE) operations on jobs data saved on a PostgreSQL database.  

**Tech Stack:** Go, PostgreSQL, Postman  

The project consists of these **qualities**:  
- We work with structs and slices to create our model and store and pass data around.  
- We create a router to register our API endpoints and handler functions.
- We make use of technologies such as the PostgreSQL database for persisting records, Postman for making API requests, third party packages such as gorilla/mux for routing, and other packages for loading environment variables, making http requests, encoding and decoding JSON, performing string conversions and checking types.
- We refactor our code and create a modular file structure.

# Set up Evironment
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


# Run Application  
Follow the steps below to **run** the application.  

(1) Type in command line:  
`go build && ./jobs-list`    

(2) Open your local browser:  
Visit http://localhost:8000/jobs  

Alternatively, you can open **Postman** (desktop agent) and send a GET request:  
`GET` http://localhost:8000/jobs  


# Examples  
We **recommend** using Postman to make requests to the REST API. Open Postman (desktop agent) in your browser and send a **GET**, **POST**, **PUT**, or **DELETE** request. We provide some **examples** below.  

To **CREATE** a new job:   
`POST` http://localhost:8000/jobs  

Body:  
{  
&nbsp; &nbsp; &nbsp; &nbsp;“title”:”Software Engineer”,  
&nbsp; &nbsp; &nbsp; &nbsp;“company”:”LLM AI”,  
&nbsp; &nbsp; &nbsp; &nbsp;”location”:”San Francisco, CA”,  
&nbsp; &nbsp; &nbsp; &nbsp;“type”:”Full-time”  
}    

Assume the “id” returned for the POST request is 8.  

To **READ** all the jobs:  
`GET` http://localhost:8000/jobs  

To **read** a specific job, specify its ID:  
`GET` http://localhost:8000/jobs/1  

To **UPDATE** the job created above:  
`PUT` http://localhost:8000/jobs  

Body:  
{  
&nbsp; &nbsp; &nbsp; &nbsp;“id”:8,  
&nbsp; &nbsp; &nbsp; &nbsp;“title”:”Machine Learning Engineer”,  
&nbsp; &nbsp; &nbsp; &nbsp;“company”:”LLM AI”,  
&nbsp; &nbsp; &nbsp; &nbsp;”location”:”New York, NY”,  
&nbsp; &nbsp; &nbsp; &nbsp;“type”:”Full-time”  
}  

To **DELETE** the job above:  
`DELETE` http://localhost:8000/jobs/8  

For more examples of using our API with Postman, view our [demo](https://github.com/nabilshadman/go-rest-api-jobs-list-postgres/tree/main/demo).  

