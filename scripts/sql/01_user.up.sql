CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50) UNIQUE
    password VARCHAR(50)
    age int(50)
);
