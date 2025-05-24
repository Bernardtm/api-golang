-- Create Status Table
CREATE TABLE default_schema.status (
    status_uuid UUID NOT NULL DEFAULT default_schema.uuid_generate_v7() PRIMARY KEY,
    name VARCHAR(50) NOT NULL unique,
    creation_date DATE DEFAULT CURRENT_DATE,
    modification_date DATE NULL
);


-- Menu Table
CREATE TABLE default_schema.menus (
    menu_uuid UUID NOT NULL DEFAULT default_schema.uuid_generate_v7() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
	url VARCHAR(255) NOT NULL,
    order_index INTEGER NOT NULL,
    creation_date DATE DEFAULT CURRENT_DATE,
    modification_date DATE NULL,
    status_uuid UUID NOT NULL REFERENCES default_schema.status(status_uuid)
);

