CREATE TABLE IF NOT EXISTS payments(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   courses VARCHAR (50) NOT NULL,
   bill integer NOT NULL,
   date VARCHAR (50) NOT NULL,
   status VARCHAR (50) NOT NULL,
   client_secret VARCHAR (50) NOT NULL
);