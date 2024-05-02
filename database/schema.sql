SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------

--
-- Table structure for table `users`
-- 

CREATE SCHEMA IF NOT EXISTS pricecompare DEFAULT CHARACTER SET utf8mb4;

USE pricecompare;

CREATE TABLE listing (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  query VARCHAR(500) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
)engine=innodb;

CREATE TABLE products (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  product_name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  listing_id INT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (listing_id) REFERENCES listing(id)
)engine=innodb;

