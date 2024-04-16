CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "totalSalesDept" FLOAT DEFAULT (0) NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW(),
);

CREATE TABLE IF NOT EXISTS managers(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "departmentId" uuid NOT NULL,
    "firstName" VARCHAR NOT NULL,
    "lastName" VARCHAR NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_departments FOREIGN KEY(departmentId) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS items(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "departmentId" uuid NOT NULL,
    "name" VARCHAR NOT NULL,
    "price" FLOAT NOT NULL,
    "qty" INT NOT NULL,
	"qtySold" INT NOT NULL DEFAULT (0),
	"category" VARCHAR NOT NULL DEFAULT ("NONE"),
    "promo" BOOLEAN  NOT NULL DEFAULT (FALSE),
    "promoPrice" FLOAT NOT NULL DEFAULT (0),
    "totalSalesItem" FLOAT DEFAULT(0) NOT NULL,
	"aisle" VARCHAR NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_departments FOREIGN KEY(departmentId) REFERENCES departments(id)
);