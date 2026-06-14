CREATE TABLE IF NOT EXISTS catalog_items(
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(200) NOT NULL,
    short_description TEXT,
    full_description TEXT,
    image_url TEXT,
    price FLOAT,
    brand_id UUID REFERENCES brands(id),
    category_id UUID REFERENCES categories(id)
);

INSERT INTO catalog_items(
    id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id
    )
VALUES(
    'e0000001-0000-0000-0000-000000000001',
    'test1 brand',
    'qwe',
    'qwertyuiopasdfghjkl',
    NULL,
    299,
    'b0000001-0000-0000-0000-000000000001',
    'c0000001-0000-0000-0000-000000000001'
    )
ON CONFLICT DO NOTHING;

INSERT INTO catalog_items(id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id)
VALUES(
    'e0000002-0000-0000-0000-000000000001',
    'tesе2 item',
    'abdu',
    'qwertyuijvndslicdsklcndslkcnsdjvcsdklcvnsd',
    NULL,
    299,
    'b0000002-0000-0000-0000-000000000001', 'c0000001-0000-0000-0000-000000000002')
ON CONFLICT DO NOTHING;

INSERT INTO catalog_items(id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id)
VALUES(
    'e0000003-0000-0000-0000-000000000001',
    'test3 item',
    'qwe',
    'wertyuilkjhgfdsdfghjk',
    NULL,
    4344343,
    'b0000003-0000-0000-0000-000000000001', 'c0000001-0000-0000-0000-000000000003')
ON CONFLICT DO NOTHING;

INSERT INTO catalog_items(id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id)
VALUES(
    'e0000004-0000-0000-0000-000000000001',
    'test4 item',
    'dslkncldsncjksdnckjlsdnc',
    'wertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjk',
    NULL,
    4344343,
    'b0000004-0000-0000-0000-000000000001', 'c0000001-0000-0000-0000-000000000004')
ON CONFLICT DO NOTHING;

INSERT INTO catalog_items(id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id)
VALUES(
    'e0000005-0000-0000-0000-000000000001',
    'test5 item',
    'dslknclnckjlsdnc',
    'wertyuilgfdsdfghjkwertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjkwertyuilkjhgfdsdfghjk',
    NULL,
    5,
    'b0000005-0000-0000-0000-000000000001', 'c0000001-0000-0000-0000-000000000005')
ON CONFLICT DO NOTHING;