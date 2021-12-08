package database

var SELECT_USER_BUDGET_INFO = `
	select 
		user_id,
		full_name,
		username,
		total_money
	from
		users
	where 
		username = ? and password = crypt(?, password)
`
var SELECT_USER_CASH_LIST = `
	select
		cash_id,
		amount,
		summary,
		received_at
	from 
		cash
	where user_id = (select user_id from users where username = ? and password = crypt(?, password))
	order by received_at desc;
`

var SELECT_USER_EXP_LIST = `
	select
		expenditure_id,
		amount,
		summary,
		spent_at
	from 
		expenditures
	where user_id = (select user_id from users where username = ? and password = crypt(?, password))
	order by received_at desc;
`
//

/*
	select
		c.cash_id,
		c.amount,
		c.summary,
		c.received_at
	from 
		cash as c
	join users as u using (user_id)
	where
		u.username = 'jasurbek' and u.password = crypt('1001goog', password)
	order by c.received_at desc	
*/