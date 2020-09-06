set search_path to daily;

CREATE TABLE daily.users
(
    id        bigint DEFAULT base_id_generator(10)  not null primary key,
    name      varchar(255) not null,
    avatar    varchar(255) default 'http://47.98.227.139:81/user_base_avatar.jpg' not null ,
    password  varchar(255) not null ,
    sex       int not null ,
    create_at varchar(255) not null ,
    login_at  varchar(255)
);

CREATE UNIQUE INDEX daily_users_name_uindex ON daily.users (name);

CREATE TABLE daily.issues
(
    id        bigint default base_id_generator(10) not null primary key,
    name      varchar(255) not null ,
    user_id   bigint not null ,
    level     int not null ,
    time      varchar(255) not null ,
    finish_if    boolean,
    check_if  boolean,
    check_id  bigint,
    create_at varchar(255) not null ,
    update_at varchar(255)
);


create function daily.base_id_generator(a int)
    returns bigint as
$$
select (
                           a::text ||
                           extract(years from now()) ||
                           extract(months from now()) ||
                           extract(days from now()) ||
--                            lpad((random() * 100000)::text, 5, '00')
                           (SELECT num FROM generate_series(100, 999) AS t(num) ORDER BY random() LIMIT 1)
           )::bigint;

$$ language sql strict;


INSERT INTO daily.users(name, password, sex,create_at,login_at) values ('zhy','123456','2',current_timestamp,current_timestamp);
INSERT INTO daily.users(name, password, sex,create_at,login_at) values ('pbc','123456','1',current_timestamp,current_timestamp);