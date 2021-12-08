package database

var CREATE_NEW_USER = `
	insert into 
		users (full_name, username, password) 
	values 
		(?, ?, crypt(?, gen_salt('bf'))) 
	returning 
		user_id, full_name, username, total_money
`

var CREATE_NEW_CASH = `
	insert into 
		cash (user_id, amount, summary) 
	values 
		((select user_id from users where username = ? and password = crypt(?, password)), ?, ?)
	returning
		cash_id, amount, summary, received_at

`

var CREATE_NEW_EXP = `
	insert into 
		expenditures (user_id, amount, summary) 
	values 
		((select user_id from users where username = ? and password = crypt(?, password)), ?, ?)
	returning
		expenditure_id, amount, summary, spent_at

`

/*insert into 
		cash (user_id, amount, summary) 
	values 
		(select user_id from users where username = 'jasurbek' and password = crypt('1001goog', password)), 5000.00, 'bonus')
	returning
		cash_id, amount, amount, summary, received_at;*/