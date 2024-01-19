# CVWO
This is a repo for CVWO's forum webapp backend.


## API Overview
Base URL: [does not exist]


## User Related Endpoints

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
   - Decription: Review the details of the existing logged in user.
   - Method: 'GET'
   - Endpoint: '/user'
   - Response: User information.

4. Review user posts.
   - Decription: Review all posts (threads and replies) of the existing logged in user.
   - Method: 'GET'
   - Endpoint: '/user/posts'
   - Response: One list of threads and one list of post.

5. Update user
   - Description: Update the details of the existing logged in user.
   - Method: 'PUT'
   - Endpoint: '/user'
   - Request Body: 
     - 'name': name
     - 'password': password
   - Response: Confirmation of user update.

6. Logout
   - Decription: To logout of current session.
   - Method: 'GET'
   - Endpoint: '/logout'
   - Response: Expired JWT Cookie to revoke login session.

7. List users
   - Description: Retrieve a list of all users.
   - Method: 'GET'
   - Endpoint: '/users'
   - Response: A list of all users.


## Subforum Related Endpoints

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


## Thread Related Endpoints

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


## Reply Related Endpoints

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


## SUPERUSER Related Endpoints: 

1. Create subforum [WORK IN PROGRESS]
   - Description: Create new subforum. User must be SUPERUSER.
   - Method: 'POST'
   - Endpoint: '/subforum'
   - Request Body: 
     - 'name': name of new subforum
     - 'description': description of new subforum
     - 'photoURL': link to cover photo of new subforum
   - Response: Confirmation of subforum creation.

2. Update existing subforum [WORK IN PROGRESS]
   - Description: Update the details of an existing subforum. User must be SUPERUSER.
   - Method: 'PUT'
   - Endpoint: '/subforum/{subforumID}'
   - URL Parameters: 
     - 'subforumID': unique identifier of subforum
   - Request Body: 
     - 'name': new name of subforum
     - 'description': new description of subforum
     - 'photoURL': new link to cover photo of subforum
   - Response: Confirmation of subforum creation.














   - Description: 
   - Method: ''
   - Endpoint: ''
   - Request Body: 
     - 
   - Response: .

