CREATE TABLE service_calls
(
    user_id varchar(255) not null,
    model_id uuid not null,
    request_id uuid not null,
    cost int,
    call_time timestamp default now(),
    status varchar(255),
    metadata jsonb,
    PRIMARY KEY (request_id)
)