CREATE TABLE purchases
(
    "id"            uuid      DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_uid"      uuid           NOT NULL,
    "amount"        DECIMAL(10, 2) NOT NULL,
    "purchase_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "created_at"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)