drop database if exists blog;
CREATE database if not exists blog default charset utf8mb4 collate utf8MB4_general_ci;
use blog;
CREATE TABLE post
(
    id         bigint auto_increment,
    title      varchar(255) NOT NULL default '',
    draft      boolean      NOT NULL default true,
    content    longtext     not null,
    filename   varchar(255) NOT NULL default '',
    md5        CHAR(32)     NOT NULL default '',
    createdAt datetime     not null default CURRENT_TIMESTAMP,
    updatedAt datetime     not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deletedAt datetime     ,
    primary key (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
CREATE TABLE tag
(
    id         bigint auto_increment,
    postId     bigint       not null,
    tag        varchar(255) NOT NULL,
    createdAt datetime     not null default CURRENT_TIMESTAMP,
    updatedAt datetime     not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deletedAt datetime     ,
    primary key (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
CREATE TABLE category
(
    id         bigint auto_increment,
    postId     bigint       not null,
    category   varchar(255) NOT NULL,
    createdAt datetime     not null default CURRENT_TIMESTAMP,
    updatedAt datetime     not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deletedAt datetime     ,
    primary key (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;