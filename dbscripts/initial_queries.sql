PRAGMA foreign_keys = ON;

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    document VARCHAR(11) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    balance TEXT NOT NULL,
    usertype INTEGER NOT NULL
);

CREATE TABLE transactions (
    id INTEGER PRIMARY KEY,
    amount TEXT NOT NULL,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    timestamp DATE NOT NULL,
    -- in the video JPA is used and "id" is not used, but JPA maps the "User" class to its id.
    FOREIGN KEY(sender_id) REFERENCES users(id),
    FOREIGN KEY(receiver_id) REFERENCES users(id)
);