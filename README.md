# Setup Guide for CVWO ForumApp Backend

## Introduction

Welcome to the setup guide for the backend of ForumApp built using Golang and PostgreSQL. This document will guide you through the process of setting up the application on your local server.

## Prerequisites
Before you begin, ensure you have the following installed:
- GoLang: [GoLang Installation Guide](https://golang.org/doc/install)
- PostgreSQL: [PostgreSQL Installation Guide](https://www.postgresql.org/download/)

## Installation and Configuration

### Cloning the Repository
Clone the ForumApp backend repository from GitHub using the following commands:
```bash
git clone https://github.com/yizhong187/CVWO.git
cd CVWO
```

### Setting Up the Database
1. Install PostgreSQL if you haven't already.
2. Create a new database named `forum`:
   ```sql
   CREATE DATABASE forum;
   ```
   Run this command in your PostgreSQL command line interface or through a tool like pgAdmin.
3. From the root of your project directory, import the schema file `schema.sql` located in the `database` folder::
   ```bash
   psql -U [username] -d forum -f database/schema.sql
   ```
   Replace `[username]` with your PostgreSQL username. By default, it is `postgres`.
4. After setting up the schema, import the data from the `data_backup.sql` file, also located in the `database` folder:
   ```bash
   psql -U [username] -d forum -f database/data.sql
   ```
   This will populate the database with test data.

### Configuring the Application
1. Set up the necessary environment variables in a `.env` file or your environment:
   - `PORT`: The port number on which the backend server will run. By default, it's set to 8080. You can change this if you have another service running on this port or if you prefer to use a different port.
   - `DB_HOST`: The hostname or IP address of the database server. For local development, it's usually set to localhost. 
   - `DB_PORT`: The port number your PostgreSQL database is listening on. The default PostgreSQL port is `5432`.
   - `DB_USER`: The username for accessing the PostgreSQL database. By default, it's set to `postgres`, which is the default superuser account in PostgreSQL. Use a different account if you have created specific users for your database.
   - `DB_NAME`: The name of the database that the application will connect to. Unless you changed the database name, it will be set as forum. Ensure that you have already created this database on your PostgreSQL server.
   - `DB_SSLMODE`: This sets the SSL mode for connecting to the database. By default, it's set to `disable` for local development.
   - `DB_PASSWORD`: The password for the database user. 

   - The tables are set as follow based on the schema provided: 
     - `DB_USERS_TABLE` = `public.users`
     - `DB_THREADS_TABLE` = `public.threads`
     - `DB_SUBFORUMS_TABLE` = `public.subforums`
     - `DB_REPLIES_TABLE` = `public.replies`

   - `SECRET_KEY`: An important key used for hashing user passwords. It should be a complex and unique string for security purposes, never shared or exposed publicly.


### Running the Application
Start the application using:
```bash
go run .
```

## Usage
Base URL: By default, it is set to `localhost:8080/v2`
This application supports various CRUD operations through its RESTful API. 
Use tools like `curl`, Postman, or Thunder Client to interact with and test the API.

For the test data, several subforums have been set up. Additionally, 2 normal users and 1 super user have also been included. Note that the passwords in the database are hashed.
1. Normal user 1
   - name: timmy
   - type: normal
   - password: password 
2. Normal user 2
   - name: yizhong187
   - type: normal
   - password: password 
3. Super user 1
   - name: god
   - type: super
   - password: god'spassword 

<br>

### User Related Endpoints

1. Sign up
   - Description: For registration of new users.
   - Method: 'POST'
   - Endpoint: '/signup'
   - Request Body: 
     - 'name': name
     - 'password': password
   - Reponse: Confirmation of creation.

2. Login
   - Description: For login of existing users.
   - Method: 'POST'
   - Endpoint: '/login'
   - Request Body: 
     - 'name': name
     - 'password': password
   - Response: User information, with JWT.

3. Review user
   - Decription: Review the account details of the existing logged in user.
   - Method: 'GET'
   - Endpoint: '/user'
   - Response: User information.

4. Update user
   - Description: Update the account details of the existing logged in user.
   - Method: 'PUT'
   - Endpoint: '/user'
   - Request Body: 
     - 'name': name
     - 'password': password
   - Response: Confirmation of user update.

5. Logout
   - Decription: To logout of current session.
   - Method: 'GET'
   - Endpoint: '/logout'
   - Response: Expired JWT Cookie to revoke login session.

6. Review user posts.
   - Decription: Review all posts (threads and replies) of existing user.
   - Method: 'GET'
   - Endpoint: '/user/{userName}/posts'
   - Response: One list of threads and one list of post.


<br>

### Subforum Related Endpoints

1. List subforums
   - Description: Retrieve a list of all subforums.
   - Method: 'GET'
   - Endpoint: '/subforums'
   - Response: A list of all subforums.

2. Review subforum
   - Description: Review one subforum.
   - Method: 'GET'
   - Endpoint: '/subforums/{subforumID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of the subforum.
   - Response: Detailed information of the subforum.

<br>

### Thread Related Endpoints

1. List threads of a subforum
   - Description: Retrieve a list of all threads in a specific subforum.
   - Method: 'GET'
   - Endpoint: '/subforums/{subforumID}/threads'
   - URL Parameters: 
     - 'subforumID': unique identifier of the subforum.
   - Response: A list of all threads in the subforum.

2. Review thread 
   - Description: Review one thread.
   - Method: 'GET'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of the subforum.
     - 'threadID': unique identifier of thread
   - Response: Detailed information of the thread. 

3. Post new thread
   - Description: Post a new thread on a subforum. User has to be logged in.
   - Method: 'POST'
   - Endpoint: '/subforums/{subforumID}/threads'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
   - Request Body: 
     - 'title': title of new thread
     - 'content': content of new thread
   - Response: Confirmation of thread posted. 

4. Update existing thread
   - Description: Update the details of an existing thread. Original thread creator has to be logged in.
   - Method: 'PUT'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
     - 'threadID': unique identifier of thread
   - Request Body: 
     - 'title': new title of thread
     - 'content': new content of thread
   - Response: Confirmation of thread update.

5. Delete existing thread
   - Description: Delete an existing thread and all replies under it. Original thread creator has to be logged in.
   - Method: 'DELETE'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
     - 'threadID': unique identifier of thread
   - Response: Confirmation of thread deletion.

<br>

### Reply Related Endpoints

1. List replies of a thread
   - Description: Retrieve a list of all replies in a specific thread.
   - Method: 'GET'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}/replies'
   - URL Parameters: 
     - 'subforumID': unique identifier of the subforum.
     - 'threadID': unique identifier of thread
   - Response: A list of all replies in the thread.

2. Post new reply
   - Description: Post a new reply in a thread. User has to be logged in.
   - Method: 'POST'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}/replies'
   - URL Parameters: 
     - 'subforumID': unique identifier of the subforum.
     - 'threadID': unique identifier of thread
   - Request Body: 
     - 'content': content of new reply
   - Response: Confirmation of reply posted.

3. Update existing reply
   - Description: Update the details of an existing reply. Original replier has to be logged in.
   - Method: 'PUT'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}/replies/{replyID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
     - 'threadID': unique identifier of thread
     - 'replyID': unique identifier of reply
   - Request Body: 
     - 'content': new content of reply
   - Response: Confirmation of reply update.

4. Delete existing reply
   - Description: Delete an existing reply. Original replier has to be logged in.
   - Method: 'DELETE'
   - Endpoint: '/subforums/{subforumID}/threads/{threadID}/replies/{replyID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
     - 'threadID': unique identifier of thread
     - 'replyID': unique identifier of reply
   - Response: Confirmation of reply deletion.

<br>

### SUPERUSER Related Endpoints: 

1. Create subforum 
   - Description: Create new subforum. User must be SUPERUSER.
   - Method: 'POST'
   - Endpoint: '/superuser/subforums'
   - Request Body: 
     - 'name': name of new subforum
     - 'description': description of new subforum
     - 'photoURL': link to cover photo of new subforum
   - Response: Confirmation of subforum creation.

2. Update existing subforum 
   - Description: Update the details of an existing subforum. User must be SUPERUSER.
   - Method: 'PUT'
   - Endpoint: '/superuser/subforums/{subforumID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
   - Request Body: 
     - 'name': new name of subforum
     - 'description': new description of subforum
     - 'photoURL': new link to cover photo of subforum
   - Response: Confirmation of subforum creation.

3. List users 
   - Description: Retrieve a list of all users.
   - Method: 'GET'
   - Endpoint: '/users'
   - Response: A list of all users.


## Troubleshooting
Encounter an issue? Check out these common problems and solutions:
- **Database Connection Error**: Ensure your PostgreSQL is running and the connection string is correct.
- **Build Failures**: Verify that you have the correct version of GoLang installed.

For more help, please open an issue on the [GitHub issues page](https://github.com/yizhong187/CVWO/issues).
