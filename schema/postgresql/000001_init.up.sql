CREATE EXTENSION "uuid-ossp";
CREATE TABLE wallets
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    balance decimal CONSTRAINT nn_balance CHECK (balance >= 0)
);
CREATE TABLE transactions
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "from" uuid references wallets on delete RESTRICT,
    "to" uuid references wallets on delete RESTRICT,
    amount decimal CONSTRAINT nn_amount CHECK (amount >= 0)
);