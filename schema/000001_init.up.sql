CREATE TABLE service_calls
(
    /*?
    Нужен ли Primary Key?
      ?*/
    user_id varchar(255) not null
    model_id uuid not null
    request_id uuid not null
    cost int
    status varchar(255)
    metadata jsonb
)