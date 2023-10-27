CREATE TABLE products (
	id          serial  primary key,
	name        text    not null,
	retailer    text    not null,
	price       decimal not null default 0,
	description text             default ''
);

INSERT INTO products (name, retailer, price, description) VALUES
('Test Product 1', 'Unknown', 9.99, ''),
('Test Product 2', 'HS Flensburg', 9.90, '');
