# File Server Go
This is an open API for a file server written in Go. To work, it is enough to have some dependencies and specify some environment variables.

## Required dependencies:
- Go >= 1.26
- Docker >= 29.5
- Docker-compose >= 5.1

## Required variables:
For the project to work, you need to add the following variables to the system environment variable:
- `FILE_PATH=/path/to/your/directiory`: this is the path to the directory where the file server will upload the data.
- `DB_URl=postgres://Your_login:Your_Password@Your_Host:Your_Port/file_db`: this is the standard url for connecting to the Postgres database. **You need to replace login, password, host and port with your own!**
