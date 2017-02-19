# toolo-skill

Alexa skill to query my [toolo-api](https://github.com/Mumakil/toolo-api) running in https://toolo-api.herokuapp.com to ask if you can have beers in Töölö or if there is a game today in one of the nearby sports arenas.

Built with `go` on top of [github.com/mikeflynn/go-alexa/skillserver](https://github.com/mikeflynn/go-alexa/tree/master/skillserver).

## Running

For your personal use, you can just run this in heroku. Only configuration needed is the Amazon Alexa Application ID, which you can get when you create the app in the api console. Set it as the `APP_ID` env in the heroku app and then point the Alexa skill to `https://<your heroku url>/echo/beers`. You can find the intent schemas and sample utterances in this repository for the skill configuration.

## License

MIT
