-- noinspection SqlDialectInspectionForFile

CREATE TABLE persons (
  id           BIGSERIAL   NOT NULL,
  person_id    VARCHAR(50) NOT NULL,
  first_name   VARCHAR(25) NOT NULL,
  last_name    VARCHAR(25) NOT NULL,
  last_updated TIMESTAMP   NOT NULL,
  CONSTRAINT uniq_id UNIQUE (person_id)
);

CREATE TABLE address (
  id           BIGSERIAL                      NOT NULL,
  person_ref   BIGINT REFERENCES persons (id) NOT NULL,
  phone_number BIGINT                         NOT NULL,
  full_address VARCHAR(500)                   NOT NULL,
  CONSTRAINT unique_per_person UNIQUE (person_ref)
);

CREATE TABLE daily_transactions (
  id                 BIGSERIAL                      NOT NULL,
  person_ref         BIGINT REFERENCES persons (id) NOT NULL,
  number_of_liters   INT                            NOT NULL,
  total_price_of_day INT                            NOT NULL,
  day                TIMESTAMP                      NOT NULL,
  person_name        VARCHAR(25)                    NOT NULL
);

CREATE TABLE payment_details (
  id               BIGSERIAL                      NOT NULL,
  person_ref       BIGINT REFERENCES persons (id) NOT NULL,
  amount_payed     INT                            NOT NULL,
  paid_to          VARCHAR(25)                    NOT NULL,
  day              TIMESTAMP                      NOT NULL
);

CREATE TABLE total_balance (
  id           BIGSERIAL                      NOT NULL,
  person_ref   BIGINT REFERENCES persons (id) NOT NULL,
  amount       BIGINT                         NOT NULL,
  last_updated TIMESTAMP                      NOT NULL,
  CONSTRAINT person_id_unique UNIQUE (person_ref)
);
