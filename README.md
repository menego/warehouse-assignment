# Assignment

## Compromises

### Database

For the sake of speed I decided to implement a simple memory database instead of a real one. The choices for a production
environment would have been different, as by choosing a NoSql or relational solution depending on the specific use case.

### Architecture

Given that this is an assignment in which I need to show all the functionalities in one place and that it should be 
easily reproducible, I chose to implement a single monolith application. Ideally, given the time, I would have also provided
the correct way to deploy such a solution on the right cloud service (e.g. Docker container and then Fargate, Elastic Beanstalk or K8s).
In a different situation I would have opted for a serverless solution (e.g. Lambdas + Api Gateway + DB) with the right 
infrastructure as code (e.g. terraform).
Moreover, no authentication and authorization is provided here despite the OpenAPI spec saying otherwise. Also here,
I would have opted for an Open Id Connect protocol and proper middlewares (e.g. Auth 0).

### Tests

Testing, especially in case of TDD, takes a great deal of time (and pays out in future maintenance). For this assignment
I covered only the readers/json package to give you an idea of how I approach this matter when I develop.

## Execution

### Server

<ol>
    <li>Build the program:</li> go build cmd/warehouse/main.go
    <li>Give execution permission (linux and mac only):</li> chmod +x main 
    <li>Run the program:</li> ./main 
</ol>
The program will load the data in present in the <i>assets</i> folder and start an http server on localhost:3000, you can use
the postman collection in the <i>postman_collection</i> folder to use the available endpoints.

### Test

To run the tests execute the following command `go test ./...`.