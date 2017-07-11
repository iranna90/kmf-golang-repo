-- noinspection SqlDialectInspectionForFile

CREATE TABLE persons (
  id           BIGSERIAL   NOT NULL,
  first_name   VARCHAR(25) NOT NULL,
  last_name    VARCHAR(25) NOT NULL,
  last_updated TIMESTAMP   NOT NULL,
  CONSTRAINT uniq_id UNIQUE (id),
  CONSTRAINT person_first_last_name UNIQUE (first_name, last_name)
);

CREATE TABLE address (
  id           BIGSERIAL                      NOT NULL,
  person_ref   BIGINT REFERENCES persons (id) NOT NULL,
  phone_number BIGINT                         NOT NULL,
  full_address VARCHAR(500)                   NOT NULL
);

CREATE TABLE daily_transactions (
  id               BIGSERIAL                      NOT NULL,
  person_ref       BIGINT REFERENCES persons (id) NOT NULL,
  number_of_liters SMALLINT                       NOT NULL,
  day              TIMESTAMP                      NOT NULL,
  person_name      VARCHAR(25)                    NOT NULL
);

CREATE TABLE amount_paid (
  id         BIGSERIAL                      NOT NULL,
  person_ref BIGINT REFERENCES persons (id) NOT NULL,
  amount     INT                            NOT NULL,
  paid_to    VARCHAR(25)                    NOT NULL,
  day        TIMESTAMP                      NOT NULL
);

CREATE TABLE total_balance (
  id           BIGSERIAL                      NOT NULL,
  person_ref   BIGINT REFERENCES persons (id) NOT NULL,
  amount       BIGINT                         NOT NULL,
  last_updated TIMESTAMP                      NOT NULL,
  CONSTRAINT person_id_unique UNIQUE (person_ref)
);
