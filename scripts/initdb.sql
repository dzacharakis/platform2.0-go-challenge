DROP DATABASE IF EXISTS gwichallenge;
DROP ROLE IF EXISTS dimitris;

CREATE ROLE dimitris LOGIN PASSWORD 'gwidbpass';
ALTER ROLE dimitris WITH CREATEDB CREATEROLE;
CREATE DATABASE gwichallenge OWNER dimitris;

\connect gwichallenge dimitris;

CREATE TYPE variable_type AS ENUM (
  'categorical', -- places objects described in a set of data, into one or more groups, or categories
  'quantitative' -- takes numerical values for which arithmetic operations make sense
);

-- Let's examine the DEFINITIONS provided:

-- 1. AUDIENCE: is a series of characteristics
-- "gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month"
-- This means any characteristics of objects described in a set of data, for example, "length of calls" in a set of "calls made to a customer service center".
-- In statistics is called VARIABLE, so variable == audience.

-- 2. CHART: has a title, axes and data
-- In statistics bar graphs and pie charts display and help us understand the distribution of a variable quickly.

-- 3. INSIGHT: a small piece of text that provides some insight into a topic, e.g,"40% of millenials spend more than 3 hours on social media daily"
-- This may be rephrased to: 
-- one-axis: percent of people aged 25 to 40 ("age group"), here: 40% 
-- another-axis: "time spent on socials", here: more than 3hours/day. 
-- So, an insight is just the knowledge gained after learning from data.


-- ENTITY SET

-- chart, variable/audience, insight - a specialization of a higher-level entity set "asset" (superclass-subclass relationship)
CREATE TABLE asset (
	asset_id SERIAL,
  asset_type CHAR(8) CHECK (asset_type IN ('chart', 'variable', 'insight')),
  PRIMARY KEY (asset_id, asset_type),
  UNIQUE (asset_id)
);

-- audience = variable 
CREATE TABLE variable (
  asset_id INT,
  asset_type CHAR(8) DEFAULT 'variable' CHECK (asset_type = 'variable'),
  name VARCHAR,
  var_type variable_type, -- quantitative, categorical
  possible_values VARCHAR NOT NULL DEFAULT '',
	unit VARCHAR NOT NULL DEFAULT '',
  PRIMARY KEY (asset_id, asset_type),
  FOREIGN KEY (asset_id, asset_type) REFERENCES asset (asset_id, asset_type),
  UNIQUE (name)
);

CREATE TABLE chart (
  asset_id INT NOT NULL,
  asset_type CHAR(8) DEFAULT 'chart' CHECK (asset_type = 'chart'),
  title VARCHAR,
  x_name VARCHAR, -- variable e.g., age group (25-34 years)
  y_name VARCHAR, -- variable e.g., highest education level
  PRIMARY KEY (asset_id, asset_type),
  UNIQUE (asset_id),
  FOREIGN KEY (asset_id, asset_type) REFERENCES asset (asset_id, asset_type),
  FOREIGN KEY (x_name) REFERENCES variable (name),
  FOREIGN KEY (y_name) REFERENCES variable (name)
) ;

CREATE TABLE chart_data (
  id SERIAL,
  chart_id INT NOT NULL, -- One-to-Many association: a chart has one or more data. Data belongs to one and only one chart.
  x_value VARCHAR NOT NULL,
  y_value VARCHAR NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (chart_id) REFERENCES chart (asset_id)
);

-- insight is just knowledge gained after learning from chart data
CREATE TABLE insight (
  asset_id INT,
  asset_type CHAR(8) DEFAULT 'insight' CHECK (asset_type = 'insight'),
  chart_id INT NOT NULL,  -- One-to-Many association: a chart has one or more insights. An insight refers to one and only one chart.
  statement VARCHAR NOT NULL,
  PRIMARY KEY (asset_id, asset_type),
  FOREIGN KEY (asset_id, asset_type) REFERENCES asset (asset_id, asset_type),
  FOREIGN KEY (chart_id) REFERENCES chart (asset_id)
) ;

CREATE TABLE users (
  id SERIAL,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (email)
);

-- RELATIONSHIP SET

-- Many-to-Many from users to chart
-- A user might have zero or more assets as favourites. An asset might have zero or more users who saved it as a favourite.
CREATE TABLE favourite (
  id SERIAL,
	user_id INT NOT NULL,
	asset_id INT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (user_id, asset_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (asset_id) REFERENCES asset (asset_id)
);

-- INSERT DATA
INSERT INTO public.users (id, username, password, email) VALUES 
(DEFAULT, 'dimitris', 'test', 'dimitris@example.com'),
(DEFAULT, 'zacharakis', 'test2 ', 'zacharakis@example.com');

INSERT INTO public.asset ("asset_id", "asset_type") VALUES 
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'variable'),
(DEFAULT, 'chart'),
(DEFAULT, 'chart'),
(DEFAULT, 'insight'),
(DEFAULT, 'insight'),
(DEFAULT, 'insight'),
(DEFAULT, 'insight');


INSERT INTO public.variable (asset_id, name, var_type, possible_values, unit) VALUES
(1,'Gender', 'categorical','Male, Female', DEFAULT),
(2,'Birth Country', 'categorical', 'Greece, England, France', DEFAULT),
(3,'Age group', 'categorical', 'Children: 00-14, Youth: 15-19, Adults: 20-64, Seniors: (65 and over)', DEFAULT),
(4,'Education', 'categorical', 'Less than High School, High School graduate, Bachelor degree, Advanced degree', DEFAULT),
(5,'Time spent daily on socials', 'quantitative', DEFAULT,'hours per day'),
(6,'Purchases last month', 'quantitative', DEFAULT, '# purchases per month'),
(7,'Monthly rent', 'quantitative', DEFAULT, 'money per month'),
(8,'Length of a call', 'quantitative', DEFAULT, 'seconds');

INSERT INTO public.chart (asset_id, title, x_name, y_name) VALUES 
(9, 'Educational attainment of people aged 25 to 34 years.', 'Age group', 'Education'),
(10, 'Daily social media usage of people aged 25 to 40 years.', 'Age group', 'Time spent daily on socials');

INSERT INTO public.chart_data (chart_id, x_value, y_value) VALUES 
(9, '13%', 'Not High School graduate'),
(9, '30%', 'High School degree'),
(9, '23%', 'Bachelor''s degree'),
(9, '7%', 'Advanced degree'),
(10,'10%', 'less than 1 hour'),
(10,'30%', 'from 1 hour to 2 hours'),
(10,'20%', 'from 2 hours to 3 hours'),
(10,'40%', 'more than 3 hours');


INSERT INTO public.insight (asset_id, chart_id, statement) VALUES 
(11, 9, '30% of people aged 25 to 34 have High School degree but no higher degree.'),
(12 ,9, '23% of people aged 25 to 34 have Bachelor''s degree but no higher degree.'),
(13 ,9, '7% of people aged 25 to 34 have Advanced degree.'),
(14, 10, '40% of people aged 25 to 40 spend more than 3 hours on on social media daily.');

INSERT INTO public.favourite (user_id, asset_id) VALUES 
(1, 8),
(1, 9),
(1, 10),
(1, 14);