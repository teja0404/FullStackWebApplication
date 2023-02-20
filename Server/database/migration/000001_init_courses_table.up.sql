CREATE TABLE IF NOT EXISTS courses(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   instructor_name VARCHAR (50) NOT NULL,
   price integer,
   description VARCHAR (50) NOT NULL,
   duration integer
);