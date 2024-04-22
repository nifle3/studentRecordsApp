CREATE TABLE IF NOT EXIST Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,

    CONSTRAINT users_email_unique UNIQUE (email)
)

CREATE TABLE IF NOT EXIST Student (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    passport_seria INT NOT NULL,
    passport_number INT NOT NULL,
    birthday DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_country_code SMALLINT NOT NULL,
    phone_city_code VARCHAR(10) NOT NULL,
    phone_number VARCHAR(10) NOT NULL,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    house INT NOT NULL,
    apartment_number INT NOT NULL,

    CONSTRAINT student_email_unique UNIQUE (email),
)

CREATE TABLE IF NOT EXIST Application (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID REFERENCES Student(id) ON DELETE SET NULL,
    contact_info VARCHAR(50) NOT NULL,
    application_text TEXT NOT NULL,
    application_status VARCHAR(50) NOT NULL,
    )

CREATE TABLE IF NOT EXIST StudentReport (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    enroll_year INT NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    enroll_order_number INT NOT NULL,
)

CREATE TABLE IF NOT EXIST StudentDocument (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID REFERENCES Student(id) ON DELETE CASCADE,
    document_name VARCHAR(100) NOT NULL,
    document_type VARCHAR(100) NOT NULL,
    document_link_s3 VARCHAR(100) NOT NULL,
)