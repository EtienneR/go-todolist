## About

With this minimalist todolist, you can post / delete a message.  
Built with the help of three library below :
- Gin-Gonic (micro framework) 
- Gorp (ORM) 
- SQL Driver


## Installation

Install each library :  
<code>go get github.com/gin-gonic/gin</code>  
<code>go get gopkg.in/gorp.v1</code>  
<code>go get github.com/go-sql-driver/mysql</code>  

Then, create a new database "todolist" in your MySQL / MariaDB server :

```sql
CREATE DATABASE IF NOT EXISTS todolist;
```


## Sources

* Gin Gonic : https://github.com/gin-gonic/gin
* Gorp : https://github.com/go-gorp/gorp
* Driver SQL : https://github.com/go-sql-driver/mysql