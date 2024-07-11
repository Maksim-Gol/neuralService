CREATE TABLE service_calls
(
    user_id varchar(255) not null default '42',
    model_id uuid not null default '550e8400-e29b-41d4-a716-446655440001'::uuid,
    request_id uuid not null default '550e8400-e29b-41d4-a716-446655440000'::uuid,
    cost int default 0,
    status varchar(255) default 'def',
    call_time timestamp default now(),
    metadata jsonb default '{}'::jsonb,
    PRIMARY KEY (request_id)
)