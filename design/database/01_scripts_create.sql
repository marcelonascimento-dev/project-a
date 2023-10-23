-- Create the project_a database
CREATE DATABASE project_a;

-- Connect to the project_a database
\c project_a

-- Create the product table
CREATE TABLE product (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT,
    category_id VARCHAR(255),
    creation_date TIMESTAMP NOT NULL,
    last_updated TIMESTAMP
);

-- Create the category table
CREATE TABLE category (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    creation_date TIMESTAMP NOT NULL,
    last_updated TIMESTAMP
);
