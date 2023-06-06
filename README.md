# Project Setup README
This README provides instructions for setting up a project with a backend written in Go (Golang) and a frontend written in Vue.js, using local installation for development and Docker for containerization.

## Prerequisites
Before getting started, make sure you have the following software installed on your system:

- Go (Golang)
- Node.js, Yarn (Package Manager)
- Docker, docker-compose


## Local Setup

### Backend Setup
1. Clone the project repository from the version control systemChange into the backend directory:
```
$ cd backend
```

2. Install the Go dependencies using Go Modules:
```
$ go mod download
```

3. Create .env and add the configuration values as per your requirements.
4. Start the backend server:
```
$ go run main.go
```

The backend server should now be running on http://localhost:8000.

### Frontend Setup
1. Change into the frontend directory:\
```
$ cd ecommerce
```
2. Install the frontend dependencies:
```$ yarn```
3. Copy sample.env to .env and add the configuration values with http://localhost:8000.
4. Start the frontend development server:
```$ yarn serve```

The frontend development server should now be running on http://localhost:8080.

## Docker Deployment

### Instructions

1. Clone the project repository from the version control system.
2. Update .env files as per suggested in the Local Setup for frontend and backend projects.

3. Build the frontend and backend:
```
$ docker-compose up --build -d
```

4. Populate data into the database using the below command.
```
$ docker-compose exec web bash
# cd backend/uploadCSV
# go run main.go #
```


## Accessing the Application
You can now access the application by opening your web browser and navigating to http://localhost, Live Deployement URL
http://54.174.172.252


## Additional Information
- The backend API endpoints and logic can be implemented in the `backend` directory.
- The frontend user interface and components can be developed in the `ecommerce` directory.

Feel free to customize the project structure and files based on your specific requirements.

## Conclusion
Congratulations! You have successfully set up the project with a backend written in Go and a frontend written in Vue.js, using Docker for containerization. You can now start building your application by implementing the backend API endpoints and developing the frontend user interface.


