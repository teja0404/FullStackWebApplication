# FullStackWebApplication

## Dashboard

![image-20230202210205564](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210205564.png)

## Starting the application

In the project root, run: 
```
docker-compose build 
docker-compose up
```

## Overview

**Users can add the course by clicking on the `add to cart` button**

![image-20230202211709528](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202211709528.png)

**Course can be removed from the cart by clicking on `Remove from cart` button**

![image-20230202211731794](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202211731794.png)

**On adding the desired courses he can click the `Proceed` button and he will be navigated to payment option section**

![image-20230202210432218](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210432218.png)

**User should enter the name (Case sensitive & same will be reflected in the transaction history) and click on the `Gpay` button to open Payment gateway** 

![image-20230202210459587](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210459587.png)

**Check the final amount and click on `Pay`**

![image-20230202210532581](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210532581.png)

**User can check his/her payment history by clicking on the `Get Purchase history of customer` button**

![image-20230202210625698](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210625698.png)

**New course can be added by the Admin by entering all the required details**

![image-20230202210822587](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210822587.png)

**Newly added courses will be automatically reflected in the available courses section**

![image-20230202210949516](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202210949516.png)

**User cannot add the already existing course into the cart again**

![image-20230202211034241](C:\Users\IllaDurgaVaraSivaTej\AppData\Roaming\Typora\typora-user-images\image-20230202211034241.png)



## Features

- View all the available courses
- Add courses into cart
- Remove courses from cart
- Checkout the products in the cart
- Users purchase history can be viewed
- *Only admin can add or update the courses

## Frameworks used:

- Gin - GO framework
- React for frontend (Hooks)
- Postgres for DB
- Docker

### Why Gin instead of Mux ?

- https://golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http (Gin is easy to understand)

## Libraries:

- GORM - ORM in GO
- godotenv - to store config variables
- go get -u gorm.io/driver/postgres - Postgres support

## Resources:

- https://go.dev/doc/
- GO ORM - https://gorm.io/
- Gin - https://gin-gonic.com/docs/
- React Hooks - https://www.freecodecamp.org/news/react-hooks-fundamentals/#:~:text=React%20Hooks%20are%20simple%20JavaScript,updater%20function%20to%20update%20it.
- useStateHook & useEffectHook - https://www.youtube.com/watch?v=O6P86uwfdR0
- https://reactjs.org/docs/conditional-rendering.html

## Important Links:

- https://pkg.go.dev/github.com/gin-gonic/gin#H
- https://go.dev/doc/code
- https://go.dev/doc/gopath_code
- Connecting to postgres DB : https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- https://gorm.io/docs/query.html

## New terms (ignore):

- Closures
  - https://betterprogramming.pub/closures-made-simple-with-golang-69db3017cd7b
- FMT - Formatted I/O
- Struct - Class in Java

- Context
  - context is a standard package of Golang that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.
