
# Player stats table

team_id (key)	string
player_id (key)	string
points_per_game	number
fg_percentage	number
ft_percentage	number
steals_per_game	number
fouls_per_game	number

# Players table

team_id (key) string
player_name string
player_id string

# Player contract table

player_id (key)	string
created_at	string (UTC)
updated_at	string (UTC)
salary	string
contract_length	number
contract_type	string
team_id (key)	string

# Team Budget table

team_id (key)	string
budget	number
year	number
created_at	string (UTC)
updated_at	string (UTC)

# Team table

team_id (key)	string
team_name	string
league_id (foreign key related to other table)	string