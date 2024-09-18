CREATE DATABASE IF NOT EXISTS `ecommerce`;
USE `ecommerce`;
-- Add other initialization SQL commands here.

DROP TABLE IF EXISTS `crm_user`;

CREATE TABLE
    `crm_user` (
        `user_id` int (11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
        `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
        `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
        `usr_username` varchar(32) NOT NULL DEFAULT '' COMMENT 'Username',
        `usr_phone` char(10) NOT NULL DEFAULT '' COMMENT 'Phone Number',
        `usr_create_at` int (11) NOT NULL DEFAULT 0 COMMENT 'Creation Time IP',
        `usr_created_ip_at` int (11) NOT NULL DEFAULT 0 COMMENT 'Creation Time',
        `usr_updated_at` int (11) NOT NULL DEFAULT 0 COMMENT 'Update Time',
        `usr_last_login_at` int (11) NOT NULL DEFAULT 0 COMMENT 'Last Login Time',
        `usr_last_login_ip_at` int (11) NOT NULL DEFAULT 0 COMMENT 'Last Login IP',
        `usr_status` tinyint (1) NOT NULL DEFAULT 0 COMMENT 'Status 1:enable 0:disable -1:deleted',
        PRIMARY KEY (`user_id`),
        KEY `idx_email` (`usr_email`),
        KEY `idx_username` (`usr_username`),
        KEY `idx_phone` (`usr_phone`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'Account';