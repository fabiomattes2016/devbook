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
CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE = INNODB;