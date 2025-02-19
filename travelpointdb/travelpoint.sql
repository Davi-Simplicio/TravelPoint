CREATE DATABASE travel_point IF NOT EXISTS;
\c travel_point;

CREATE TABLE address (
    ID SERIAL PRIMARY KEY,
    Longitude DECIMAL(10, 6),
    Latitude DECIMAL(10, 6),
    City VARCHAR(255),
    State VARCHAR(50),
    Country VARCHAR(255),
    Cep VARCHAR(255),
    Neighborhood VARCHAR(255),
    Street VARCHAR(255),
    Number VARCHAR(255)
);

CREATE TABLE calendar (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    Status BOOLEAN
);

CREATE TABLE date (
    ID SERIAL PRIMARY KEY,
    Date DATE,
    CalendarId SERIAL,
    HasEvent BOOLEAN,
    FOREIGN KEY (CalendarId) REFERENCES calendar(ID) ON DELETE CASCADE
);

CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    LastName VARCHAR(255),
    BirthDate DATE,
    Email VARCHAR(255) UNIQUE,
    Password VARCHAR(255),
    PhoneNumber VARCHAR(255),
    IsOwner BOOLEAN,
    CalendarId SERIAL,
    AddressId SERIAL,
    FOREIGN KEY (CalendarId) REFERENCES calendar(ID) ON DELETE SET NULL,
    FOREIGN KEY (AddressId) REFERENCES address(ID) ON DELETE SET NULL
);
