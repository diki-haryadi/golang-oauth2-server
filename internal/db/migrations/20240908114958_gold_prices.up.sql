CREATE TABLE gold_prices
(
    "id"          uuid      DEFAULT uuid_generate_v4() PRIMARY KEY,
    "type"        VARCHAR(10),
    "price"       DECIMAL(10, 2) NOT NULL,
    "recorded_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (recorded_at),
    "created_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)