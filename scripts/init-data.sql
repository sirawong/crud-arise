-- Initialize sample data for Product Management API
-- This script runs automatically when PostgreSQL container starts for the first time

-- Create tables if they don't exist (in case GORM hasn't run yet)
CREATE TABLE IF NOT EXISTS categories (
                                          id UUID PRIMARY KEY,
                                          name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
    );

CREATE TABLE IF NOT EXISTS products (
                                        id UUID PRIMARY KEY,
                                        name VARCHAR(255) NOT NULL,
    description TEXT,
    sku VARCHAR(100) UNIQUE NOT NULL,
    price DECIMAL(10,2) NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0,
    image_url VARCHAR(255),
    category_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
    );

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON categories(deleted_at);
CREATE INDEX IF NOT EXISTS idx_products_deleted_at ON products(deleted_at);

-- Insert Categories (with conflict handling)
INSERT INTO categories (id, name, created_at, updated_at) VALUES
                                                              ('550e8400-e29b-41d4-a716-446655440001', 'Electronics', NOW(), NOW()),
                                                              ('550e8400-e29b-41d4-a716-446655440002', 'Clothing', NOW(), NOW()),
                                                              ('550e8400-e29b-41d4-a716-446655440003', 'Books', NOW(), NOW()),
                                                              ('550e8400-e29b-41d4-a716-446655440004', 'Home & Garden', NOW(), NOW()),
                                                              ('550e8400-e29b-41d4-a716-446655440005', 'Sports & Outdoors', NOW(), NOW())
    ON CONFLICT (id) DO NOTHING;

-- Insert Products (with conflict handling)
INSERT INTO products (id, name, description, sku, price, stock, image_url, category_id, created_at, updated_at) VALUES
                                                                                                                    (
                                                                                                                        '550e8400-e29b-41d4-a716-446655440101',
                                                                                                                        'iPhone 15 Pro',
                                                                                                                        'Latest Apple iPhone with advanced camera system and A17 Pro chip',
                                                                                                                        'IPHONE15PRO-001',
                                                                                                                        1199.99,
                                                                                                                        50,
                                                                                                                        'https://example.com/images/iphone15pro.jpg',
                                                                                                                        '550e8400-e29b-41d4-a716-446655440001',
                                                                                                                        NOW(),
                                                                                                                        NOW()
                                                                                                                    ),
                                                                                                                    (
                                                                                                                        '550e8400-e29b-41d4-a716-446655440102',
                                                                                                                        'Samsung Galaxy S24',
                                                                                                                        'Premium Android smartphone with AI features and excellent camera',
                                                                                                                        'GALAXY-S24-001',
                                                                                                                        999.99,
                                                                                                                        75,
                                                                                                                        'https://example.com/images/galaxys24.jpg',
                                                                                                                        '550e8400-e29b-41d4-a716-446655440001',
                                                                                                                        NOW(),
                                                                                                                        NOW()
                                                                                                                    ),
                                                                                                                    (
                                                                                                                        '550e8400-e29b-41d4-a716-446655440103',
                                                                                                                        'Classic Cotton T-Shirt',
                                                                                                                        'Comfortable 100% cotton t-shirt available in multiple colors',
                                                                                                                        'TSHIRT-COTTON-001',
                                                                                                                        29.99,
                                                                                                                        200,
                                                                                                                        'https://example.com/images/cotton-tshirt.jpg',
                                                                                                                        '550e8400-e29b-41d4-a716-446655440002',
                                                                                                                        NOW(),
                                                                                                                        NOW()
                                                                                                                    ),
                                                                                                                    (
                                                                                                                        '550e8400-e29b-41d4-a716-446655440104',
                                                                                                                        'The Go Programming Language',
                                                                                                                        'Comprehensive guide to Go programming by Alan Donovan and Brian Kernighan',
                                                                                                                        'BOOK-GO-LANG-001',
                                                                                                                        45.99,
                                                                                                                        30,
                                                                                                                        'https://example.com/images/go-programming-book.jpg',
                                                                                                                        '550e8400-e29b-41d4-a716-446655440003',
                                                                                                                        NOW(),
                                                                                                                        NOW()
                                                                                                                    ),
                                                                                                                    (
                                                                                                                        '550e8400-e29b-41d4-a716-446655440105',
                                                                                                                        'Wireless Bluetooth Headphones',
                                                                                                                        'High-quality over-ear headphones with noise cancellation and 30-hour battery',
                                                                                                                        'HEADPHONES-BT-001',
                                                                                                                        149.99,
                                                                                                                        100,
                                                                                                                        'https://example.com/images/bluetooth-headphones.jpg',
                                                                                                                        '550e8400-e29b-41d4-a716-446655440001',
                                                                                                                        NOW(),
                                                                                                                        NOW()
                                                                                                                    )
    ON CONFLICT (sku) DO NOTHING;