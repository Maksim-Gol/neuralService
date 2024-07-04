CREATE TABLE users
(
    id serial not null unique
    username varchar(255) not null
)



CREATE TABLE models
(
    id serial not null unique
    name varchar(255) not null
    call_cost int not null
)



CREATE TABLE calls
(
    id serial not null unique
    user_id int references users(id) on delete cascade not null
    model_id int references models(id) on delete cascade not null
    call_time timestamp not null
)