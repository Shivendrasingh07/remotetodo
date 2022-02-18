CREATE TABLE if not exists  users(
                                     userid SERIAL PRIMARY KEY,
                                     name text NOT NULL,
                                     email text NOT NULL,
                                     password TEXT NOT NULL ,
                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()

);
