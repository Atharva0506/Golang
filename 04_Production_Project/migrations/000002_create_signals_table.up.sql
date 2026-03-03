-- Create the 'symbol' ENUM type
CREATE TYPE symbol AS ENUM ('SOL', 'ETH');

-- Create the 'action' ENUM type
CREATE TYPE action AS ENUM ('buy', 'sell');

-- Create the 'signals' table
CREATE TABLE IF NOT EXISTS signals (
    id UUID PRIMARY KEY,
    symbol symbol NOT NULL,
    action action NOT NULL,
    price BIGINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
