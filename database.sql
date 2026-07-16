DROP DATABASE IF EXISTS Game_Store;

CREATE DATABASE Game_Store;

USE Game_Store;

CREATE TABLE users (
   user_id     BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   email       VARCHAR(255)    NOT NULL,
   password    VARCHAR(255)    NOT NULL,         
   name        VARCHAR(100),
   role        VARCHAR(20)     NOT NULL DEFAULT 'customer',
   created_at  DATETIME        DEFAULT CURRENT_TIMESTAMP,

   PRIMARY KEY (user_id),
   
   UNIQUE KEY uk_users_email (email)
) ENGINE=InnoDB;

CREATE TABLE user_profiles (
   profile_id  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   user_id     BIGINT UNSIGNED NOT NULL,
   user_name   VARCHAR(100),
   phone       VARCHAR(20),

   PRIMARY KEY (profile_id),

   UNIQUE KEY uk_profile_user (user_id),

   CONSTRAINT fk_profile_user
       FOREIGN KEY (user_id) REFERENCES users (user_id)
       ON DELETE CASCADE
) ENGINE=InnoDB;


CREATE TABLE categories (
   category_id  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   name         VARCHAR(100)    NOT NULL,
   description  VARCHAR(255),

   PRIMARY KEY (category_id)
) ENGINE=InnoDB;


CREATE TABLE products (
   product_id   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   category_id  BIGINT UNSIGNED NOT NULL,
   title        VARCHAR(150)    NOT NULL,
   price        DECIMAL(12,2)   NOT NULL DEFAULT 0,
   stock        INT             NOT NULL DEFAULT 0,
   created_at   DATETIME        DEFAULT CURRENT_TIMESTAMP,

   PRIMARY KEY (product_id),

   CONSTRAINT fk_product_category
       FOREIGN KEY (category_id) REFERENCES categories (category_id)
) ENGINE=InnoDB;


CREATE TABLE orders (
   order_id     BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   user_id      BIGINT UNSIGNED NOT NULL,
   order_date   DATETIME        DEFAULT CURRENT_TIMESTAMP,
   status       VARCHAR(20)     NOT NULL DEFAULT 'pending',
   total_price  DECIMAL(12,2)   NOT NULL DEFAULT 0,

   PRIMARY KEY (order_id),

   CONSTRAINT fk_order_user
       FOREIGN KEY (user_id) REFERENCES users (user_id)
) ENGINE=InnoDB;


CREATE TABLE order_items (
   order_item_id  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   order_id       BIGINT UNSIGNED NOT NULL,
   product_id     BIGINT UNSIGNED NOT NULL,
   qty            INT             NOT NULL DEFAULT 1,
   price          DECIMAL(12,2)   NOT NULL DEFAULT 0,

   PRIMARY KEY (order_item_id),

   CONSTRAINT fk_item_order
       FOREIGN KEY (order_id) REFERENCES orders (order_id)
       ON DELETE CASCADE,

   CONSTRAINT fk_item_product
       FOREIGN KEY (product_id) REFERENCES products (product_id)
) ENGINE=InnoDB;


CREATE TABLE reviews (
   review_id    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   user_id      BIGINT UNSIGNED NOT NULL,
   product_id   BIGINT UNSIGNED NOT NULL,
   rating       INT NOT NULL,
   comment      TEXT,
   created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,

   PRIMARY KEY (review_id),

   UNIQUE KEY uk_review_user_product (user_id, product_id),

   CONSTRAINT fk_review_user
      FOREIGN KEY (user_id)
      REFERENCES users (user_id)
      ON DELETE CASCADE,

   CONSTRAINT fk_review_product
      FOREIGN KEY (product_id)
      REFERENCES products (product_id)
      ON DELETE CASCADE
) ENGINE=InnoDB;
