-- Orders table
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    account_name VARCHAR(100) NOT NULL,
    total_amount DECIMAL(18,2) NOT NULL DEFAULT 0,

-- Order Status
current_order_status VARCHAR(50) NOT NULL DEFAULT 'Draft',

-- Contact Info (embedded)
contact_first_name VARCHAR(100) NOT NULL,
contact_last_name VARCHAR(100) NOT NULL,
contact_email VARCHAR(255) NOT NULL,

-- Delivery Address (embedded)
address_street VARCHAR(200) NOT NULL,
address_city VARCHAR(100) NOT NULL,
address_region VARCHAR(100) NOT NULL,
address_postal_code VARCHAR(20) NOT NULL,

-- Payment
current_payment_method VARCHAR(50) NOT NULL DEFAULT 'CreditCard',
current_payment_status VARCHAR(50) NOT NULL DEFAULT 'Pending',

-- Card Details (optional, embedded)
card_name VARCHAR(100),
card_number VARCHAR(20),
card_expiration VARCHAR(10),
card_cvv VARCHAR(10),

-- Audit
created_by VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_modified_by VARCHAR(100),
    last_modified_at TIMESTAMP WITH TIME ZONE
);

-- Order Items table
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    catalog_item_name VARCHAR(200) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(18, 2) NOT NULL
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_orders_account_name ON orders (account_name);

CREATE INDEX IF NOT EXISTS idx_orders_status ON orders (current_order_status);

CREATE INDEX IF NOT EXISTS idx_orders_payment_status ON orders (current_payment_status);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items (order_id);

-- Seed data
INSERT INTO
    orders (
        id,
        account_name,
        total_amount,
        current_order_status,
        contact_first_name,
        contact_last_name,
        contact_email,
        address_street,
        address_city,
        address_region,
        address_postal_code,
        current_payment_method,
        current_payment_status,
        card_name,
        card_number,
        card_expiration,
        card_cvv,
        created_by,
        created_at
    )
VALUES (
        'a0000000-0000-0000-0000-000000000001',
        'test_account',
        15670.00,
        'Paid',
        'Иван',
        'Иванов',
        'ivan@test.com',
        'Ленина 1',
        'Москва',
        'Московская',
        '101000',
        'CreditCard',
        'Completed',
        'Иван Иванов',
        '4111111111111111',
        '12/25',
        '123',
        'System',
        NOW()
    ),
    (
        'a0000000-0000-0000-0000-000000000002',
        'admin@example.com',
        3290.00,
        'Submitted',
        'Анна',
        'Петрова',
        'anna@example.com',
        'Тверская 12',
        'Москва',
        'Московская',
        '101001',
        'BankTransfer',
        'Pending',
        NULL,
        NULL,
        NULL,
        NULL,
        'System',
        NOW()
    );

INSERT INTO
    order_items (
        order_id,
        catalog_item_name,
        quantity,
        unit_price
    )
VALUES (
        'a0000000-0000-0000-0000-000000000001',
        'Мультиварка Redmond RMC-M90',
        1,
        5890.00
    ),
    (
        'a0000000-0000-0000-0000-000000000001',
        'Бюджетный смартфон с хорошей камерой',
        1,
        7990.00
    ),
    (
        'a0000000-0000-0000-0000-000000000001',
        'Фен Polaris PHD 2077',
        1,
        1790.00
    ),
    (
        'a0000000-0000-0000-0000-000000000002',
        'Кофеварка Polaris PCM 1516E',
        1,
        3290.00
    );