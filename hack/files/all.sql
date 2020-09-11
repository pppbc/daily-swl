set search_path to daily;

CREATE TABLE daily.users
(
    id        bigint DEFAULT base_id_generator(10)  not null primary key,
    name      varchar(255) not null,
    avatar    varchar(255) default 'http://116.62.41.19:81/base_avatar_male.jpeg' not null ,
    password  varchar(255) not null ,
    sex       int not null ,
    create_at int64 not null ,
    login_at  int64
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
    create_at int64 not null ,
    update_at int64
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


INSERT INTO daily.users(name, password, avatar, sex,create_at,login_at) values ('zhy','123456','http://116.62.41.19:81/base_avatar_female.jpeg','2',current_timestamp,current_timestamp);
INSERT INTO daily.users(name, password, avatar, sex,create_at,login_at) values ('pbc','123456','http://116.62.41.19:81/base_avatar_male.jpeg','1',current_timestamp,current_timestamp);