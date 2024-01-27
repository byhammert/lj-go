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

CREATE TABLE contas(
    id int auto_increment primary key,
    nome varchar(50) not null,
    saldo decimal(15,2),
    imagem varchar(50),
    status_conta varchar(20) not null,
    tipo varchar(20) not null,

    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE

) ENGINE=INNODB;

CREATE TABLE lancamentos(
    id int auto_increment primary key,
    descricao varchar(60) not null,
    detalhe text,
    valor decimal(15,2) not null,
    data_compra timestamp not null,
    data_vencimento timestamp not null default current_timestamp(),
    data_pagamento timestamp null default null,
    tipo varchar(20) not null,
    forma_pagamento varchar(20) not null,

    id_categoria int not null,
    id_usuario int not null,
    id_conta int not null,
    FOREIGN KEY (id_categoria)
    REFERENCES categorias(id)
    ON DELETE CASCADE,

    
    FOREIGN KEY (id_conta)
    REFERENCES contas(id)
    ON DELETE CASCADE,

    FOREIGN KEY (id_usuario)
    REFERENCES usuarios(id)
    ON DELETE CASCADE
)ENGINE=INNODB;