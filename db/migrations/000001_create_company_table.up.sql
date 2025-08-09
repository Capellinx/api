CREATE TABLE company (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    industry VARCHAR(100) NOT NULL,
    website VARCHAR(255) NOT NULL,
    company_size VARCHAR(50) NOT NULL,
    cnpj VARCHAR(14) NULL,
    country VARCHAR(100) NULL,
    currency VARCHAR(10) NULL,
    address VARCHAR(255) NULL,
    postal_code VARCHAR(20) NULL,
    city VARCHAR(100) NULL,
    state VARCHAR(100) NULL,
    active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);