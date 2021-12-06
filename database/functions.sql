-- trigger function for users

create or replace function users_trigger_func() returns trigger language plpgsql as $$

  declare

     u_name varchar;

  begin

    select u.username from users as u into u_name where u.username = new.username;
    
      if found then

          return null;

      end if;          

    return new;
    
  end;

$$
;

create trigger user_trigger
before insert on users
for each row execute procedure users_trigger_func();


-- trigger for controlling user cashes
create or replace function cash_controller_trigger_func() returns trigger language plpgsql as $$
    
    begin

        update users
        set total_money = total_money + (select amount from cash where cash_id = new.cash_id)
        where user_id = new.user_id;

        return new;
        
    end;
$$
;

create trigger cash_controller_trigger 
after insert on cash
for each row execute procedure cash_controller_trigger_func();

-- trigger for controlling user expenditures
delete from users where user_id = 3;

create or replace function expenditures_controller_trigger_func() returns trigger language plpgsql as $$
    
    declare

        u_total_money decimal(10, 2);

    begin

        select total_money from users into u_total_money where user_id = new.user_id;

        if u_total_money > (select amount from expenditures where expenditure_id = new.expenditure_id) then 


            update users
            set total_money = total_money - (select amount from expenditures where expenditure_id = new.expenditure_id)
            where user_id = new.user_id;

            return new;

        end if;

        return null;

    end;
$$
;

create trigger expenditures_controller_trigger 
after insert on expenditures
for each row execute procedure expenditures_controller_trigger_func();

-- trigger for controlling expenditures +
create or replace function expenditures_trigger_func() returns trigger language plpgsql as $$

  declare

     u_total_money decimal(10, 2);

  begin

    select total_money from users as u into u_total_money where u.user_id = new.user_id;
    
      if u_total_money > new.amount then

          return new;

      end if;          

    return null;
    
  end;

$$
;

create trigger expenditures_trigger
before insert on expenditures
for each row execute procedure expenditures_trigger_func();

-- trigger for controlling cashes -
create or replace function cash_trigger_func() returns trigger language plpgsql as $$

  declare

     u_total_money decimal(10, 2);
     limit_of_total_money decimal(10, 2) = 99999999.99;

  begin

    select total_money from users as u into u_total_money where u.user_id = new.user_id;
    
      if limit_of_total_money < u_total_money + new.amount then

          return new;

      end if;          

    return null;
    
  end;

$$
;

create trigger cash_trigger
before insert on cash
for each row execute procedure cash_trigger_func();