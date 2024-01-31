CREATE EXTENSION "uuid-ossp";
CREATE TABLE wallets
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    balance decimal CONSTRAINT nn_balance CHECK (balance >= 0)
);
CREATE TABLE transactions
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "from" uuid REFERENCES wallets ON DELETE RESTRICT,
    "to" uuid REFERENCES wallets ON DELETE RESTRICT,
    amount decimal CONSTRAINT nn_amount CHECK (amount >= 0),
    time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);