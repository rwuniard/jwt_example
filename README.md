# jwt_example


## Running the code
You can run with this on the command line after the dependencies are setup.<br/>
It will re-compile/run it if there is any code changes.<br/>
CompileDaemon -command="./jwt_example"<br/>
<br/>
Send these requests:<br/>
http://localhost:3000/signup<br/>
POST request through Postman with body message<br/>
{
    "email": "test@test.user.com,
    "password": "abc1234"
}
<br/>
<br/>
http://localhost:3000/login<br/>
POST request through Postman with body message<br/>
{
    "email": "test@test.user.com,
    "password": "abc1234"
}
It will return a cookie with Authorization value of JWT
<br/>
<br/>
http://localhost:3000/validate<br/>
GET request Postman with body message<br/>
This will allow to do any operation afterwards. 
<br/>
<br/>
http://localhost:3000/logout<br/>
GET request Postman with body message<br/>
This will remove the cookie Authorization value. 

## Setting up the dependencies
Watches your go files in a directory and invokes go build if file changed.<br/>  
Github location:<br/>
https://github.com/githubnemo/CompileDaemon<br/>
  
Install:<br/>
go get github.com/githubnemo/CompileDaemon<br/>
go install github.com/githubnemo/CompileDaemon<br/>
Dependency will be added to your go.mod<br/>

The executable will be put in your $HOME/go/bin<br/>
<br/>
How to use it<br/>
CompileDaemon -command=“./<Your_executable_filename>”<br/>
Every time you change your code, the CompileDaemon will compile and run your executable.<br/> 
<br/>

### Web Framework - Gin
GitHub location:<br/>
https://github.com/gin-gonic/gin<br/>

Install:<br/>
go get -u github.com/gin-gonic/gin<br/>
Dependency will be added to your go.mod.<br/>
<br/>

### Go dot env
Github location:<br/>
https://github.com/joho/godotenv<br/>

Install:<br/> 
go get github.com/joho/godotenv<br/>
<br/>
Dependency will be added to your go.mod.<br/>
<br/>

### Gorm - to connect to postgres
https://gorm.io/docs/

Install these:<br/>
go get -u gorm.io/gorm<br/>
go get -u gorm.io/driver/postgres<br/>
<br/>

### Bcrypt
https://pkg.go.dev/golang.org/x/crypto/bcrypt<br/>
<br/>

### JWT
https://github.com/golang-jwt/jwt

Install:<br/>
go get -u github.com/golang-jwt/jwt/v4<br/>
