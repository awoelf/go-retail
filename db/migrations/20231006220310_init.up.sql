CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "name" TEXT NOT NULL,
    "totalSalesDept" FLOAT NOT NULL DEFAULT 0,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS managers (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "departmentId" uuid REFERENCES departments(id),
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS items (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "departmentId" uuid REFERENCES departments(id),
    "name" TEXT NOT NULL,
    "price" FLOAT NOT NULL,
    "qty" INT NOT NULL,
	"qtySold" INT NOT NULL DEFAULT 0,
	"category" TEXT NOT NULL,
    "promo" BOOLEAN NOT NULL DEFAULT FALSE,
    "promoPrice" FLOAT NOT NULL DEFAULT 0,
    "totalSalesItem" FLOAT NOT NULL DEFAULT 0,
	"aisle" TEXT NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	"type" TEXT NOT NULL,
	"paymentMethod" TEXT NOT NULL,
	"items" uuid ARRAY,
	"qtyItems" INT NOT NULL DEFAULT 0,
	"totalCost" FLOAT NOT NULL DEFAULT 0,
	"savings" FLOAT NOT NULL DEFAULT 0,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW()
)