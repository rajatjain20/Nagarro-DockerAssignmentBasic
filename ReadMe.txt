###Project Overview:

	I have built a multiservice application "UserDataApp", where there will be three services - 

	1. frontend:
		This is a Golang based web application, which will have two pages to add user and to get user details. This will communicate to backend application to add user data and get data back from it.

	2. backend:
		This is a Golang based web application, which will save the data, provided through webapi, into Database. User will be able to view that data stored in DB through another webapi.
		
	3. Database:
		This is MSSQL database. In this, we have 'USERDATA' database and a table named "USERINFO". Data will be stored in this table through backend.

###Setting up project:
	
	1. Checkout the application code locally.
	
	2. Execute CLI and go to the "UserDataApp" directory:
		> cd ".\Assignment-Basic\UserDataApp"
	
	3. We need two network for this application to work as frontend communicates with backend (network 1) 
		and backend communicates with database (network 2)
	
	4. Create Networks using below cmd:
		> docker network create -d bridge frontend_network
		> docker network create -d bridge backend_network
		
	5. Image Creation :
		- database image:
			> docker build -t msdb_i:v0 .\database
			
		- backend image:
			> docker build -t backend_i:v0 .\backend
			
		- frontend image:
			> docker build -t frontend_i:v0 .\frontend
	
	6. Container Creation:
		- Database Container:
			> docker run -e MSSQL_SA_PASSWORD=admin@123 -p 1433:1433 --name msdb_c -d msdb_i:v0
			
			- connect this container with backend_network: (we want a dns name as "mssqldevdb", so connecting it separately)
			> docker network connect --alias mssqldevdb backend_network msdb_c
			
			(we will be creating containers using .env-dev files. to use other env files, alias name would be similar to what
			 mentioned in respective env files - like for test it is "mssqltstdb" and for prod "mssqlprddb")
			
		- Backend Container:
			> docker run --env-file ./backend/config/.env-dev -p 3030:3030 -d --network backend_network --name backend_c backend_i:v0
			
		- Frontend Container:
			> docker run --env-file ./frontend/config/.env-dev -p 4040:4040 -d --network frontend_network --name frontend_c frontend_i:v0
	
Note: Please check the screenshots of running application in the dcoument.
	
	
###Urls to verify the application is running and connected to the database container:
Url1 – localhost:4040
Url2 – localhost:4040/getUserInfo
Url3 – localhost:4040/addUser


###Other Useful Commands:
	To Stop containers:
               > docker stop msdb_c
               > docker stop backend_c
			   > docker stop frontend_c
			   
	To remove containers:
               > docker rm msdb_c
               > docker rm backend_c
			   > docker rm frontend_c
			   
	To remove images:
	           > docker image rm msdb_i:v0
               > docker image rm backend_i:v0
			   > docker image rm frontend_i:v0
			   
	To remove network:
			   > docker network rm frontend_network
			   > docker network rm backend_network
			   
###Docker Compose:
	(Make sure to execute above other useful commands before executing compose commands to avoid any name or port related issues)
	
	1. Go to the "UserDataApp" directory:
		> cd ".\Assignment-Basic\UserDataApp"
	
	2. Execute below command to run docker compose:
		> docker compose up -d
	
	Try executing below urls on host machine:
	Url1 – localhost:4040
	Url2 – localhost:4040/getUserInfo
	Url3 – localhost:4040/addUser
	
	3. To stop and remove all containers and network, created through this compose command
		> docker compose down

###Pushing Docker images to a Docker registry:
	
	Commands:	
		> docker login
		> docker push rajatjain20/frontend_i:v0
		> docker push rajatjain20/backend_i:v0
		> docker push rajatjain20/mssqldb_i:v0

###Implement environment variable management in your Docker Compose file to handle different environments(development, testing, production). 
	
	To implement this, we have a base compose file (docker-compose-base.yml) and environment specific compose files as below:
	
	i.	Development Environment: docker-compose-dev.yml
	ii.	Test Environment: docker-compose-test.yml
	iii.Production Environment: docker-compose-prod.yml

	All these environment specific docker compose files are using their specific env files to access environment variables.
	Base docker compose creates frontend and backend images as those are common in all the environments and creates their own
	docker images for database for each environments. This will also separate the database for each environments.
	Also we will have different networks for each environments.
	
	Commands:
		> docker compose -f docker-compose-base.yml -f docker-compose-dev.yml up -d
		> docker compose -f docker-compose-base.yml -f docker-compose-test.yml up -d
		> docker compose -f docker-compose-base.yml -f docker-compose-prod.yml up -d
		
	Other commands:
		> docker compose -f docker-compose-base.yml -f docker-compose-dev.yml down
		> docker compose -f docker-compose-base.yml -f docker-compose-test.yml down
		> docker compose -f docker-compose-base.yml -f docker-compose-prod.yml down
	
	
