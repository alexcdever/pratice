-- drop database if exists blog;
-- create database blog ;


drop table if exists post;
CREATE TABLE post
(
    id        bigserial,
    title     varchar(255) NOT NULL default '',
    draft     boolean      NOT NULL default true,
    content   text         not null,
    filename  varchar(255) NOT NULL default '',
    md5       CHAR(32)     NOT NULL default '',
    created_at time         not null default current_timestamp,
    updated_at time         not null default current_timestamp,
    deleted_at time         ,
    primary key (id)
);

drop table if exists tag;
CREATE TABLE tag
(
    id        bigserial,
    post_id    bigint       not null,
    tag       varchar(255) NOT NULL,
    created_at time         not null default CURRENT_TIMESTAMP,
    updated_at time         not null default CURRENT_TIMESTAMP,
    deleted_at time         ,
    primary key (id)
);

drop table if exists category;
CREATE TABLE category
(
    id        bigserial,
    post_id    bigint       not null,
    category  varchar(255) NOT NULL,
    created_at time         not null default CURRENT_TIMESTAMP,
    updated_at time         not null default CURRENT_TIMESTAMP,
    deleted_at time         ,
    primary key (id)
);