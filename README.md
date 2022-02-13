# Gelder Bot
[Telegram bot](https://t.me/gelderbot) that helps to organize expenses and income
## Stack
- Go with [telebot](https://github.com/tucnak/telebot)
- MongoDB
- Docker
## Running
### Environment variables
```bash
APP_URL=<app ip address or url>
BOT_TOKEN=<telegram bot token>
DATABASE_COLLECTION=<collection where documents are stored>
DATABASE_NAME=<database name>
MONGO_URI=<URI of your MongoDB database>
INITIAL_DIRECTORY=<root of the project, like /gelder-bot>
```
### MongoDB
I recommend use [MongoDB Atlas](https://www.mongodb.com/atlas/database)
### Use the already built container
```bash
docker pull ghcr.io/gelerum/gelder-bot:main
docker run --name some-gelder-bot -e APP_URL=<url> -e BOT_TOKEN=<token>...  -d gelder-bot
```
### Build container by yourself
```bash
git clone https://github.com/gelerum/gelder-bot.git
cd gelder-bot
docker build . -t gelder-bot
docker run --name some-gelder-bot -e APP_URL=<url> -e BOT_TOKEN=<token>...  -d gelder-bot
```
### Locally Build and run app natively
```bash
export APP_URL=<url>
export BOT_TOKEN=<token>
...
git clone https://github.com/gelerum/gelder-bot.git
cd gelder-bot
go build cmd/bot/main.go
./bot
```
## License
Usage is provided under the [MIT License](https://opensource.org/licenses/mit-license.php). See LICENSE for the full details.
