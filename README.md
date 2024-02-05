# Matching System

## Overview
The Matching System API allows users to interact with the matching system to create, query, and manage user profiles. Users can add themselves to the system, search for potential matches, remove their profiles, and express interest in other users.

## Base URL
```
http://localhost:8080
```

## Endpoints
### 1. Add Single Person and Match
#### Endpoint: ```/users```
#### Method: POST
#### Description:
- Add a single person to the system and find potential matches.
#### Request: 
- ```curl -X POST -H "Content-Type: application/json" -d '{"name":"userA", "height":170, "gender":"male", "dates":3}' http://localhost:8080/users```
#### Response:
- JSON representation of the matched user list.


### 2. Query Single People Matches
#### Endpoint: ```/users/:name/matches```
#### Method: GET
#### Description:
- Query potential matches for a single user.
#### Request:
- ```curl http://localhost:8080/users/userA/matches?N=5```
#### Response:
- JSON representation of the matched user list (up to the specified limit N).


### 3. Remove Single Person
#### Endpoint: ```/users/:name```
#### Method: DELETE
#### Description: 
- Remove a single person from the system.
#### Request:
- ```curl -X DELETE http://localhost:8080/users/userA```
#### Response:
- Success message.


### 4. Like Another User
#### Endpoint: ```/users/:name/like```
#### Method: POST
#### Description:
- Express interest in another user. If both users like each other, it triggers a match.
#### Request:
- ```curl -X POST -H "Content-Type: application/json" -d '{"likedName":"userB"}' http://localhost:8080/users/userA/like```
#### Response:
- Success message.


## Complete Flow Example
In this example, we will walk through the complete flow of using the Matching System API, from creating users to interactions between them.

### 1. Create UserA
- ```curl -X POST -H "Content-Type: application/json" -d '{"name":"userA", "height":170, "gender":"male", "dates":3}' http://localhost:8080/users```
### 2. Create UserB
- ```curl -X POST -H "Content-Type: application/json" -d '{"name":"userB", "height":160, "gender":"female", "dates":2}' http://localhost:8080/users```
### 3. UserA Query Matches
- ```curl http://localhost:8080/users/userA/matches?N=5```
### 4. UserB Query Matches
- ```curl http://localhost:8080/users/userB/matches?N=5```
### 5. UserA Likes UserB
- ```curl -X POST -H "Content-Type: application/json" -d '{"likedName":"userB"}' http://localhost:8080/users/userA/like```
### 6. UserB Likes UserA
- ```curl -X POST -H "Content-Type: application/json" -d '{"likedName":"userA"}' http://localhost:8080/users/userB/like```
### 7. UserA Query Matches (Discover UserB's Date Decreased)
- ```curl http://localhost:8080/users/userA/matches?N=5```
### 8. UserB Query Matches (Discover UserA's Date Decreased)
- ```curl http://localhost:8080/users/userB/matches?N=5```
### 9. Delete UserB
- ```curl -X DELETE http://localhost:8080/users/userB```
### 10. UserA Query Matches (Discover UserB Removed)
- ```curl http://localhost:8080/users/userA/matches?N=5```


## Building and Running
### Build
To build the Go application, use Docker. Open your terminal and run the following command:
- ```make docker-build```
    
    This command will utilize Docker to build your Go application and package it into a Docker image.
### Run
Once the build is complete, start the Docker container using the following command:
- ```make docker-run```

    This will launch a Docker container named MatchSystem and map the local port 8080 to port 8080 within the container.

    Your Go application should now be accessible at http://localhost:8080.

### Stop
To stop and remove the container, execute the following command:
- ```make docker-stop```

## Import and Run on Another Machine
If you want to run the application on a different machine, follow these steps:

### 1. Export Docker Image:
On the machine where you built the Docker image, export the image to a tar file:
- ```docker save -o dockerimage.tar mu8086/match-system```

### 2. Transfer to New Machine:
Transfer the your-image-name.tar file to the new machine, using SCP, FTP, or any other method.

### 3. Import Docker Image:
On the new machine, import the Docker image using the following command:
- ```docker load -i dockerimage.tar```

### 4. Run Docker Container:
Start the Docker container on the new machine:
- ```docker run -p 8080:8080 -d mu8086/match-system```

    This will run the container in the background, mapping local port 8080 to the container's 8080.

### 5. Access the Application:
Verify that your application is running by accessing http://localhost:8080.