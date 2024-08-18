CREATE DATABASE IF NOT EXISTS bank;

USE bank;

CREATE TABLE IF NOT EXISTS `accounts` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `owner` VARCHAR(255) NOT NULL,
  `balance` INT NOT NULL,
  `currency` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `entries` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `account_id` INT,
  `amount` INT NOT NULL COMMENT 'can be positive or negative',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `transfers` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` INT,
  `to_account_id` INT,
  `amount` INT NOT NULL COMMENT 'must be positive',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX IF NOT EXISTS `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX IF NOT EXISTS `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX IF NOT EXISTS `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX IF NOT EXISTS `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);
