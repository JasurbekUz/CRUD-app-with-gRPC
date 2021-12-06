-- getting basic info
select 
	user_id,
	username,
	total_money
from
	users
where 
	password = crypt('1001goog0', password);

-- getting cash informations
select
	u.user_id,
	u.username,
	u.total_money,
	c.amount,
	c.summary,
	c.received_at
from 
	cash as c
join users as u using (user_id)
where
	u.password = crypt('1001goog', password)
order by c.received_at desc
;

-- getting expenditure informations
select
	u.user_id,
	u.username,
	u.total_money,
	e.amount,
	e.summary,
	e.spent_at
from 
	expenditures as e
join users as u using (user_id)
where
	u.password = crypt('1001goog', password)
order by e.spent_at desc
;



