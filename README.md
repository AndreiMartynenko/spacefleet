# spacefleet

In terminal
Error 1049 (42000): Unknown database 'spacefleet'
Error 1049 (42000): Unknown database 'spacefleet'

When you start the server and see this error, make sure that the MySQL database is installed and also make sure that you have created a table with the exact name. In our case, 'spacefleet'.

mysql> create DATABASE spacefleet;
mysql> show DATABASES;

For testing purposes 

username root
password - password

Database  

information_schema
mysql              
erformance_schema 
spacefleet        


Source:
https://dev.mysql.com/doc/mysql-getting-started/en/#mysql-getting-started-installing

Also try to run the application on different web engines such as Firefox, Safari, Google etc

To test the Server's response USE cURL 

- Get All Spacesips: GET /spaceship/
Exapmple:
curl http://localhost:8080/spaceship/

- Get Spacesip by id: GET /spaceship/id
Exapmple:
curl http://localhost:8080/spaceship/1

- Add Spaceship: POST /spaceship/ 
Content-Type: multipart/form-data
Example:
curl -F name=Devastator -F class=Star%20Destroyer -F crew=35000 -F image=image.png -F value=1999.99 -F status=operational http://localhost:8080/spaceship/
Example with armoments:
curl -F name=Devastator -F class=Star%20Destroyer -F crew=35000 -F image=image.png -F value=1999.99 -F status=operational -F armaments='[{"title":"B1", "qty":"8"}, {"title":"D5", "qty": "30"}]' http://localhost:8080/spaceship/

Filter spaceships: GET /spaceship/[id]?[params] or /spaceship/?[params] 
Example:
curl "http://localhost:8080/spaceship/5?class=true&status=true"

Delete spaceship: DELETE /spaceship/[id]
Example delete spaceship with id 1
curl -X DELETE http://localhost:8080/spaceship/1

Update spaceship: PUT /spaceship/[id]?[params]
Note that armaments data get completly replaced
Exampe update spaceship with id 1 update name:
curl -X PUT http://localhost:8080/spaceship/1?name=Destroyer1
Exampe update spaceship with id 1 update name class and crew:
curl -X PUT "http://localhost:8080/spaceship/1?name=Destroyer1&class=Planet_Destroyer&crew=4000" 
Exampe update spaceship with id 1 update armaments with [{"title" : "gun", "qty" : "5"}, {"title" : "cannon", "qty" : "3"}]:
curl -X PUT http://localhost:8080/spaceship/1?armaments=%5B%7B%22title%22%3A%22gun%22%2C%22qty%22%3A%225%22%7D%2C%7B%22title%22%3A%22cannon%22%2C%22qty%22%3A%223%22%7D%5D
Exampe update spaceship with id 1 update name, crew and armaments with [{"title" : "gun", "qty" : "5"}, {"title" : "cannon", "qty" : "3"}]:
curl -X PUT 'http://localhost:8080/spaceship/1?name=Destroyer1&crew=4000&armaments=%5B%7B%22title%22%3A%22gun%22%2C%22qty%22%3A%225%22%7D%2C%7B%22title%22%3A%22cannon%22%2C%22qty%22%3A%223%22%7D%5D'

