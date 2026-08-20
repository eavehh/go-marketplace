-- Seed: несколько заказов для аккаунта test_account
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
        'a0000000-0000-0000-0000-000000000003',
        'test_account',
        4580.00,
        'Draft',
        'Иван',
        'Иванов',
        'ivan@test.com',
        'Ленина 1',
        'Москва',
        'Московская',
        '101000',
        'CreditCard',
        'Pending',
        'Иван Иванов',
        '4111111111111111',
        '12/26',
        '456',
        'System',
        NOW()
    ),
    (
        'a0000000-0000-0000-0000-000000000004',
        'test_account',
        12450.00,
        'Submitted',
        'Иван',
        'Иванов',
        'ivan@test.com',
        'Пушкина 5',
        'Санкт-Петербург',
        'Ленинградская',
        '190000',
        'BankTransfer',
        'Pending',
        NULL,
        NULL,
        NULL,
        NULL,
        'System',
        NOW()
    ),
    (
        'a0000000-0000-0000-0000-000000000005',
        'test_account',
        2990.00,
        'Cancelled',
        'Иван',
        'Иванов',
        'ivan@test.com',
        'Ленина 1',
        'Москва',
        'Московская',
        '101000',
        'CreditCard',
        'Refunded',
        'Иван Иванов',
        '4111111111111111',
        '12/25',
        '123',
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
        'a0000000-0000-0000-0000-000000000003',
        'Наушники Sony WH-1000XM4',
        1,
        4580.00
    ),
    (
        'a0000000-0000-0000-0000-000000000004',
        'Робот-пылесос Xiaomi Mi Robot',
        1,
        8990.00
    ),
    (
        'a0000000-0000-0000-0000-000000000004',
        'Электрический чайник Bosch TWK8611P',
        1,
        3460.00
    ),
    (
        'a0000000-0000-0000-0000-000000000005',
        'Настольная лампа Yeelight LED',
        1,
        2990.00
    );