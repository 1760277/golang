create table company(
	company_id VARCHAR PRIMARY KEY,
	branch_name VARCHAR,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR
);

create table branch(
	branch_number VARCHAR PRIMARY KEY,
	company_id VARCHAR,
	branch_name VARCHAR UNIQUE,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR,
	version VARCHAR,
	CONSTRAINT fk_company 
	FOREIGN KEY(company_id) 
	REFERENCES company(company_id)
);

alter table company 
ADD CONSTRAINT fk_company FOREIGN KEY(branch_name) REFERENCES branch(branch_name);

create table administrator(
	admin_id VARCHAR PRIMARY KEY,
	company_id VARCHAR,
	admin_name VARCHAR UNIQUE NOT NULL,
	admin_password VARCHAR NOT NULL,
	admin_mail_address VARCHAR,
	admin_registration_time TIME,
	admin_password_init_flag BOOLEAN,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR,
	version VARCHAR
	-- CONSTRAINT fk_company 
	-- FOREIGN KEY(company_id) 
	-- REFERENCES company(company_id)
);

create table token(
	access_token VARCHAR PRIMARY KEY,
	access_token_expiring_date DATE NOT NULL,
	access_token_status BOOLEAN,
	refresh_token VARCHAR,
	refresh_token_expiring_date DATE NOT NULL,
	refresh_token_status BOOLEAN,
	scope INT,
	consumer_id VARCHAR,
	company_id VARCHAR,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR,
	version VARCHAR,
	CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES company(company_id),
	CONSTRAINT fk_customer FOREIGN KEY(consumer_id) REFERENCES customer(consumer_id)
);

create table file(
	file_id VARCHAR PRIMARY KEY,
	file_path VARCHAR,
	file_name VARCHAR UNIQUE NOT NULL,
	file_total_pages INT,
	consumer_id VARCHAR,
	company_id VARCHAR,
	agent_id VARCHAR,
	file_upload_time TIME,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR,
	version VARCHAR,
	CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES company(company_id),
	CONSTRAINT fk_customer FOREIGN KEY(consumer_id) REFERENCES customer(consumer_id),
	CONSTRAINT fk_sales FOREIGN KEY(agent_id) REFERENCES sales(agent_id)
);

create table customer(
	consumer_id VARCHAR PRIMARY KEY,
	company_id VARCHAR,
	consumer_name VARCHAR UNIQUE NOT NULL,
	consumer_name_kana VARCHAR UNIQUE,
	consumer_birth DATE,
	consumer_phone_number1 VARCHAR NOT NULL,
	consumer_phone_number2 VARCHAR,
	consumer_mail_address VARCHAR,
	consumer_postal_code VARCHAR,
	consumer_address VARCHAR,
	branch_number VARCHAR,
	agent_id VARCHAR,
	consumer_registration_date DATE,
	create_date DATE,
	update_date DATE,
	update_pmg_id VARCHAR,
	request_id VARCHAR,
	version VARCHAR,
	CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES company(company_id),
	CONSTRAINT fk_branch FOREIGN KEY(branch_number) REFERENCES branch(branch_number),
	CONSTRAINT fk_sales FOREIGN KEY(agent_id) REFERENCES sales(agent_id)
);

CREATE TABLE public.sales
(
    agent_id character varying COLLATE pg_catalog."default" NOT NULL,
    company_id character varying COLLATE pg_catalog."default",
    agent_name character varying COLLATE pg_catalog."default" NOT NULL,
    agent_password character varying COLLATE pg_catalog."default" NOT NULL,
    agent_mail_address character varying COLLATE pg_catalog."default",
    agent_registration_time time without time zone,
    agent_password_init_flag boolean,
    create_date date,
    update_date date,
    update_pmg_id character varying COLLATE pg_catalog."default",
    request_id character varying COLLATE pg_catalog."default",
    version character varying COLLATE pg_catalog."default",
    CONSTRAINT sales_pkey PRIMARY KEY (agent_id),
    CONSTRAINT sales_agent_name_key UNIQUE (agent_name),
    CONSTRAINT fk_company FOREIGN KEY (company_id)
        REFERENCES public.company (company_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)