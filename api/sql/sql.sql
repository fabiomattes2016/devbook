CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios(
    id int AUTOINCREMENT PRIMARY KEY,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    senha varchar(150) NOT NULL,
    criadoEm timestamp DEFAULT CURRENT_TIMESTAMP()
) ENGINE = INNODB;