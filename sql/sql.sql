CREATE DATABASE IF NOT EXISTS ljones;
USE ljones;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criado_em timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE categorias(
    id int auto_increment primary key,
    nome varchar(50) not null,
    tipo varchar(20) not null,

    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE

) ENGINE=INNODB;