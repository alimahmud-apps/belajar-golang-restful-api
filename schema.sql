CREATE DATABASE IF NOT EXISTS belajar_golang_restful_api;
USE belajar_golang_restful_api;

CREATE TABLE catagory(
`id` INTEGER NOT NULL AUTO_INCREMENT,
`name` VARCHAR(200) NOT NULL,
PRIMARY KEY(`id`)
) ENGINE=innodb
