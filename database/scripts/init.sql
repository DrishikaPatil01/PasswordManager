create database if not exists passwordManagerService;

use passwordManagerService;

create table if not exists USER(
    USER_ID varchar(255) NOT NULL, 
    USERNAME varchar(255) NOT NULL, 
    PASSWORD varchar(255) NOT NULL, 
    PRIMARY KEY (USER_ID) 
);

create table if not exists CREDENTIALS(
    CREDENTIAL_ID varchar(255) NOT NULL, 
    USER_ID varchar(255) NOT NULL,
    USERNAME varchar(255) NOT NULL, 
    PASSWORD varchar(255) NOT NULL, 
    OPTIONAL varchar(255),
    TITLE varchar(255),
    PRIMARY KEY (CREDENTIAL_ID),
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID)
);