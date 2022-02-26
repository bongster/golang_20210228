CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (255) NOT NULL,
   balance bigint NOT NULL,
   curreny VARCHAR (3) NOT NULL,
   created_at timestamptz NOT NULL DEFAULT (now()),
   updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS transactions(
   id serial PRIMARY KEY,
   user_id integer not null references users,
   transaction_type varchar(255) not null,
   status varchar(50) not null,
   amount bigint NOT NULL,
   created_at timestamptz NOT NULL DEFAULT (now()),
   updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX IF NOT EXISTS transactions_amount_idx on transactions (amount);