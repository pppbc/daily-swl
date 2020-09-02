set search_path to daily;

CREATE TABLE daily.users
(
    id        bigint primary key,
    name      varchar(255),
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
    status    varchar(255),
    check_if boolean,
    check_id bigint,
    create_at varchar(255),
    update_at varchar(255)
);