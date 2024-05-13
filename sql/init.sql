CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    user_role VARCHAR(20) NOT NULL,

    CONSTRAINT users_email_unique UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS Students (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NULL,
    passport_seria INT NOT NULL,
    passport_number INT NOT NULL,
    birth_date DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    house INT NOT NULL,
    apartment INT NOT NULL,
    enroll_year DATE NOT NULL,
    link_photo VARCHAR(100) NULL,
    specialization VARCHAR(100) NOT NULL,
    course INT NOT NULL DEFAULT 1,
    _group INT NOT NULL,

    CONSTRAINT student_email_unique UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS PhoneNumbers (
    id UUID PRIMARY KEY,
    student_id UUID NOT NULL REFERENCES Students(id) ON DELETE CASCADE ON UPDATE CASCADE,
    country_code VARCHAR(3) NOT NULL,
    city_code VARCHAR(10) NOT NULL,
    code VARCHAR(10) NOT NULL,
    description VARCHAR(50) NULL
);

CREATE TABLE IF NOT EXISTS Applications (
    id UUID PRIMARY KEY,
    student_id UUID REFERENCES Students(id) ON DELETE SET NULL,
    _name VARCHAR(50) NOT NULL,
    contact_info VARCHAR(50) NOT NULL,
    _text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at DATE NOT NULL,
    link VARCHAR(50) NULL
);

CREATE TABLE IF NOT EXISTS StudentsDocuments (
    id UUID PRIMARY KEY,
    student_id UUID REFERENCES Students(id) ON DELETE CASCADE,
    _name VARCHAR(100) NOT NULL,
    _type VARCHAR(100) NOT NULL,
    link VARCHAR(100) NOT NULL,
    created_at DATE NOT NULL
);
