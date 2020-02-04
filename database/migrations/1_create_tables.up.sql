CREATE TABLE IF NOT EXISTS degrees (
    id                    SERIAL PRIMARY KEY,
    cels                  DOUBLE PRECISION NOT NULL,
    fahr                  DOUBLE PRECISION NOT NULL,
    UNIQUE (cels)
);
