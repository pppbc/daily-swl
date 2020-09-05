set search_path to daily;

CREATE TABLE daily.users
(
    id        bigint primary key,
    name      varchar(255),
    avatar    varchar(255),
    password  varchar(255),
    sex       int,
    create_at varchar(255),
    login_at  varchar(255)
);

CREATE TABLE daily.issues
(
    id        bigint primary key,
    name      varchar(255),
    user_id   bigint,
    level     int,
    time      varchar(255),
    finish_if    boolean,
    check_if  boolean,
    check_id  bigint,
    create_at varchar(255),
    update_at varchar(255)
);