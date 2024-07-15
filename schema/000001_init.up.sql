CREATE TABLE service_calls
(
    user_id varchar(255) not null,
    model_id uuid not null,
    request_id uuid not null,
    cost int,
    status varchar(255),
    call_time timestamp,
    metadata jsonb,
    PRIMARY KEY (request_id)
)