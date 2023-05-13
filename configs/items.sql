-- Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

-- Item table to store general product information
CREATE TABLE Item (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    asin VARCHAR(10) NOT NULL UNIQUE,
    sku VARCHAR(255) NOT NULL UNIQUE,
    brand VARCHAR(255),
    title VARCHAR(500),
    product_group VARCHAR(255),
    product_type VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ItemAttributes table to store product attributes
CREATE TABLE ItemAttributes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    binding VARCHAR(255),
    item_height REAL,
    item_length REAL,
    item_width REAL,
    item_weight REAL,
    item_dimensions_unit VARCHAR(10),
    package_height REAL,
    package_length REAL,
    package_width REAL,
    package_weight REAL,
    package_dimensions_unit VARCHAR(10),
    release_date DATE,
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ItemImage table to store product images
CREATE TABLE ItemImage (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    image_url VARCHAR(500),
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ItemSummaryByMarketplace table to store marketplace-specific summary information
CREATE TABLE ItemSummaryByMarketplace (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    marketplace_id VARCHAR(255),
    sales_rank INT,
    main_image_url VARCHAR(500),
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Issue table to store issue details
CREATE TABLE Issue (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    code VARCHAR(255),
    message TEXT,
    severity VARCHAR(255),
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ItemOfferByMarketplace table to store marketplace-specific offer information
CREATE TABLE ItemOfferByMarketplace (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    marketplace_id VARCHAR(255),
    list_price REAL,
    currency_code VARCHAR(3),
    package_quantity INT,
    availability_status VARCHAR(255),
    fulfillment_channel VARCHAR(255),
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ItemProcurement table to store procurement information of a product
CREATE TABLE ItemProcurement (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    item_id BIGINT,
    external_product_id VARCHAR(255),
    external_product_id_type VARCHAR(50),
    FOREIGN KEY (item_id) REFERENCES Item(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- FulfillmentAvailability table to store fulfillment availability information
CREATE TABLE FulfillmentAvailability (
    id INT PRIMARY KEY AUTO_INCREMENT,
    offer_id INT,
    fulfillment_center_id VARCHAR(255),
    quantity_available INT,
    FOREIGN KEY (offer_id) REFERENCES ItemOfferByMarketplace(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);