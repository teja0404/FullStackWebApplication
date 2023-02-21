CREATE TABLE IF NOT EXISTS courses(
   id serial PRIMARY KEY,
   name VARCHAR (75) NOT NULL,
   instructor_name VARCHAR (75) NOT NULL,
   price integer,
   description VARCHAR (200) NOT NULL,
   duration integer
);