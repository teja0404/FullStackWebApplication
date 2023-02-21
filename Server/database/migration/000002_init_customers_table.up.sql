CREATE TABLE IF NOT EXISTS customers(
   id serial PRIMARY KEY,
   name VARCHAR (75) UNIQUE NOT NULL,
   age integer,
   email VARCHAR (50) NOT NULL,
   gender VARCHAR (10) NOT NULL,
   is_deleted boolean
);