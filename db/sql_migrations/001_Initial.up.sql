

CREATE TABLE IF NOT EXISTS user_profiles (
	id bigserial NOT NULL,
	"name" text NULL,
	CONSTRAINT user_profiles_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	username text NULL,
	first_name text NULL,
	last_name text NULL,
	profile_id int8 NULL,
	email text NULL,
	active bool NULL,
	created timestamptz NULL,
	updated timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT fk_users_profile FOREIGN KEY (profile_id) REFERENCES user_profiles(id)
);

CREATE TABLE IF NOT EXISTS applications (
	id text NOT NULL,
	"name" text NULL,
	category int8 NULL,
	description text NULL,
	created timestamptz NULL,
	updated timestamptz NULL,
	owner_id int8 NULL,
	lead_id int8 NULL,
	CONSTRAINT applications_pkey PRIMARY KEY (id),
	CONSTRAINT fk_applications_lead FOREIGN KEY (lead_id) REFERENCES users(id),
	CONSTRAINT fk_applications_owner FOREIGN KEY (owner_id) REFERENCES users(id)
);