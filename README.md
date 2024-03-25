# Golang intern assignment 3
## Create database (docker)
```
make create-db
```

## Start the app and migration db 
```
make run 
```
## Test the app 
Run play ground: <br>
```
localhost:3000/
```

Create new user:<br>
```
mutation {
  createUser(input:{
    username: "user2",
    password: "123456",
    email:"usr2@techvify.com.vn"
  })
}
```

Login as user: <br> 
```
mutation {
  login(input:{username: "user2", password: "123456"})
}
```

Login as admin (ROLE ADMIN auto generate when migarate database): <br>
```
mutation {
  login(input:{username: "admin", password: "password"})
}
```

Fetch 100 first sample earthquake (JUST ADMIN): <br>
```
localhost:3000/fetchgeo
```

Multifilter earthquake: <br>
```
query ListEarthquakes($filter: EarthquakeFilterByDay, $pagination: PaginationInput){
	listEarthquakesByDay(filter: $filter, pagination: $pagination) {
    id,
    mag
  }
}
```
Variables: <br>
```
{
  "filter": {
    "magnitude": 0.5,
    "place": "CA",
    "dayAgo": 1
  },
  "pagination": {
    "limit": 6,
    "offset": 10
  }
}

```
Header: <br>
```
{
    "Authorization": <User's token>
}
```

Log out: <br>
```
query {
    logout
}

```
## Database Schema
Earthquake database <br>
![This is database design of this repo](/assets/images/earthquake.png)

User database <br>
![This is database design of this repo](/assets/images/authen-system.png)

