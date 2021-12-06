--MOCK DATA

insert into users (full_name, username, password) values ('Jasurbek Shamsitdionov', 'jasurbek', crypt('1001goog', gen_salt('bf')));
insert into users (full_name, username, password) values ('Doniyor Kuchimov', 'doniyor', crypt('8878', gen_salt('bf')));
insert into users (full_name, username, password) values ('Zohidjon Hamidov', 'zokhidjan', crypt('zohid_88', gen_salt('bf')));

insert into cash (user_id, amount, summary) values (1, 77600000.00, 'salary');
insert into cash (user_id, amount, summary) values (9, 1000000.00, 'friend');

insert into expenditures (user_id, amount, summary) values (1, 7500, 'fanta C');
insert into expenditures (user_id, amount, summary) values (9, 60000.00, 'CocaCola');