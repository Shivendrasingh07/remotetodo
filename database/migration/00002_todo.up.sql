CREATE TABLE if not exists todos(
                                    id SERIAL PRIMARY KEY,
                                    userid int references users(userid) NOT NULL,
                                    task TEXT NOT NULL,
                                    detail TEXT,
                                    date TIMESTAMP NOT NULL,
                                    createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);