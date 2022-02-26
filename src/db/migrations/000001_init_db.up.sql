CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL
);

CREATE TABLE IF NOT EXISTS balances(
   id serial PRIMARY KEY,
   user_id integer not null references users,
   balance numeric not null
);

CREATE TABLE IF NOT EXISTS transactions(
   id serial PRIMARY KEY,
   user_id integer not null references users,
   transaction_type varchar(255) not null,
   amount numeric not null
);

CREATE INDEX IF NOT EXISTS transactions_amount_idx on transactions (amount);