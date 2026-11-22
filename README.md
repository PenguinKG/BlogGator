# Blog Agre**Gator**

Gator is a blog aggregator developped in Go.

It is recommended that you run it on macOS / Linux or WSL 2.

To be able to use Gator you need both Go 1.27+ and PostgreSQL installed.

If you don't have them already, you can install them by running:

## **Go**
`curl -sS https://webi.sh/golang | sh; \
source ~/.config/envman/PATH.env`

## **Postgre**
### *macOS with [brew](https://brew.sh/):*
`brew install postgresql@15`

### *Linux / WSL (Debian):*
`sudo apt update
sudo apt install postgresql postgresql-contrib`

## **Install Gator**
To install Gator just clone the repository
`git clone https://github.com/PenguinKG/BlogGator.git`

Access the directory and do all the setup inside it:
`cd BlogGator`

## **Setup**
Once installed, Gator's setup is pretty simple.
First update your Postgres password (only if on Linux / WSL) \[Be sure to remember your password]:
`sudo passwd [your_password_here]`

Then start the Postgres server:
### *macOS*
`brew services start postgresql@15`

### *Linux / WSL*
`sudo service postgresql start`

After that, enter the psql shell:
### *macOS*
`psql postgres`

### *Linux / WSL*
`sudo -u postgres psql`

Once on the psql shell, create your database. I called mine gator:
`CREATE DATABASE gator;`

Connect to the new database:
`\c gator`

Set the user password (Linux / WSL only):
`ALTER USER postgres PASSWORD '[your_password_here]';`

Install Goose:
`go install github.com/pressly/goose/v3/cmd/goose@latest`

Now get your connection string. It should look something like this (the default port is 5432):
`protocol://username:password@host:port/database`
- Remember that macOS users don't need the password field, only username.

Now run the 5 database schemas (by running the following command 5 times):
`goose -dir sql/schema postgres <connection_string> up`

Install the CLI with:
`go install .`

Finally, in your Home directory, create a .gatorconfig.json file containing the following:
`{"db_url":<connection_string>, "current_user_name": ""}`

## **Usage**
Now that you've finished installing and setting up Gator, you can finally use it.

All commands in Gator follow the same formula
`BlogGator <command> <parameters>`

Here are the available commands and their parameters:
# **User Commands**
- `register <username>` - Creates a new user and logs in
- `login <username>` - Switches the active user
- `users` - Lists all registered users
# **Feed Commands**
- `addfeed <name> <url>` - Adds a new feed and follows it for the current user
- `feeds` - Lists all feeds and who created them
- `follow <url>` - Follows a feed for the current user
- `following` - Lists all feeds followed by the current user
- `unfollow <url>` - Unfollows a feed
# **Aggregation & Browsing**
- `agg <time_between_requests>` - Starts the aggregator loop
- `browse [limit]` - Views recent posts for followed feeds.
