# Banking Application Backend

**Work under progress**

**Tech Stack : Go, PostgreSQL, AWS**

BASE URL : http://ad950a068f10c4ce19f7a995869ec22c-616cedf622eb1400.elb.eu-north-1.amazonaws.com/

DB Schema documentation : https://dbdocs.io/sohamkanji02/simple_bank

Routes :

- /users (POST request) -> Creates user account.
- /users/login (POST request) -> Logins to user account. Sends back PASETO token, with a validity for 15 mins.
  **Rest of the paths should be sent the PASETO token for authentication.
- /accounts (POST request) -> Creates an account.
- /accounts/:id (GET request) -> Fetches the account with the given id.
- /accounts (GET request) -> Lists all the accounts created by the current user.
- /accounts/:id (DELETE request) -> Deletes the account with the given id.
- /accounts/:id (PUT request) -> Updates the account with the given id.
- /tranfsers (POST request) -> Transfers money from one account to another.
