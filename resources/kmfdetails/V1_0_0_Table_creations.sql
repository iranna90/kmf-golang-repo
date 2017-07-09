CREATE TABLE persons (
  id          BIGSERIAL   NOT NULL,
  firstName   VARCHAR(25) NOT NULL,
  lastName    VARCHAR(25) NOT NULL,
  lastUpdated TIMESTAMP   NOT NULL,
  CONSTRAINT person_first_last_name UNIQUE (firstName, lastName)
);

CREATE TABLE address (
  id          BIGSERIAL                      NOT NULL,
  person_ref  BIGINT REFERENCES persons (id) NOT NULL,
  phoneNumber BIGINT                         NOT NULL,
  fullAddress VARCHAR(500)                   NOT NULL
);

CREATE TABLE transactions (
  id              BIGSERIAL                      NOT NULL,
  person_ref      BIGINT REFERENCES persons (id) NOT NULL,
  numberOfListers SMALLINT                       NOT NULL,
  day             TIMESTAMP                      NOT NULL,
  personName      VARCHAR(25)                    NOT NULL
);

CREATE TABLE paid (
  id         BIGSERIAL                      NOT NULL,
  person_ref BIGINT REFERENCES persons (id) NOT NULL,
  amount     INT                            NOT NULL,
  paidTo     VARCHAR(25)                    NOT NULL,
  day        TIMESTAMP                      NOT NULL
);

CREATE TABLE balance (
  id          BIGSERIAL                      NOT NULL,
  person_ref  BIGINT REFERENCES persons (id) NOT NULL,
  amount      BIGINT                         NOT NULL,
  lastUpdated TIMESTAMP                      NOT NULL,
  CONSTRAINT person_id_unique UNIQUE (person_ref)
);

