# FullStackWebApplication

## Application is up and running and hosted using GCP Cloud run

It can be accessed with the [link](https://fullstackwebapplication-client-nvywvuaozq-uc.a.run.app/)

## Cloud Run Dashboard

Both Client and Server images are running

![image](https://user-images.githubusercontent.com/38076041/219547237-da9659b2-b56d-4cdb-84a0-558a8959f0a2.png)

## Stripe webhook endpoint
![image](https://user-images.githubusercontent.com/38076041/219547376-3a343cfe-0f92-49ec-b52f-f7c9c7199565.png)

## Dashboard

![image](https://user-images.githubusercontent.com/38076041/216379940-e771b858-2f91-48e8-900d-d81fcc2a0f7b.png)


## Starting the application

In the project root, run: 
```
docker-compose build 
docker-compose up

Use 4242 4242 4242 4242 as card number
12/24 MM/YY and any 3 digit number as CVV
```

## Payment screen 
![image](https://user-images.githubusercontent.com/38076041/218306398-3c19a406-9c11-4c2c-a3a0-144a2541a3da.png)

## Payment success screen
![image](https://user-images.githubusercontent.com/38076041/218306423-4b635bfc-01a7-422d-83ed-daf86a8973c0.png)

## Stripe dashboard
![image](https://user-images.githubusercontent.com/38076041/218306460-c78eae94-8243-475f-a7e0-2cf325e970e7.png)


## Overview

**Users can add the course by clicking on the `add to cart` button**

![image](https://user-images.githubusercontent.com/38076041/216380111-e8d4f9d2-88b9-49c0-ac5f-422bf2f345b0.png)

**Course can be removed from the cart by clicking on `Remove from cart` button**

![image](https://user-images.githubusercontent.com/38076041/216380414-f881798d-132e-455b-801f-d1658b8d03d4.png)

**On adding the desired courses he can click the `Proceed` button and he will be navigated to payment option section**

![image](https://user-images.githubusercontent.com/38076041/216380548-03723f2d-c06b-4ff5-a674-7882a1c1d2bc.png)

**User should enter the name (Case sensitive & same will be reflected in the transaction history) and click on the `Gpay` button to open Payment gateway** 

![image](https://user-images.githubusercontent.com/38076041/216380683-23a3a84f-a6bd-474a-bb5d-4e8118cff612.png)

**Check the final amount and click on `Pay`**

![image](https://user-images.githubusercontent.com/38076041/216380776-752e8381-bce1-478d-8a99-a4a8b6fbf0c6.png)

**User can check his/her payment history by clicking on the `Get Purchase history of customer` button**

![image](https://user-images.githubusercontent.com/38076041/216380864-82df58f1-4128-4c0d-9fe2-9fea9c96ee58.png)

**New course can be added by the Admin by entering all the required details**

![image](https://user-images.githubusercontent.com/38076041/216381169-627a3915-7807-4c12-af65-ac56d342c1b6.png)

**Newly added courses will be automatically reflected in the available courses section**

![image](https://user-images.githubusercontent.com/38076041/216381250-cc0219f1-1322-4902-af29-5dcf29f08a62.png)

**User cannot add the already existing course into the cart again**

![image](https://user-images.githubusercontent.com/38076041/216381348-c735d048-ee4a-4d6d-b1a8-149399320cfe.png)



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
