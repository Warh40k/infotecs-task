CREATE TABLE wallets
(
    id VARCHAR PRIMARY KEY,
    balance FLOAT CONSTRAINT nn_balance CHECK (balance >= 0)
);
CREATE TABLE transactions
(
    id VARCHAR PRIMARY KEY,
    "from" VARCHAR REFERENCES wallets ON DELETE RESTRICT,
    "to" VARCHAR REFERENCES wallets ON DELETE RESTRICT,
    amount DECIMAL CONSTRAINT nn_amount CHECK (amount >= 0),
    time timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);