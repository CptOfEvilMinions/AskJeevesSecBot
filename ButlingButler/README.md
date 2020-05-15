# Butling Butler

Do to the way the Slack API is designed it will `PUSH` the user's response to a URL endpoint. This means that you need your machine 

## Deploy to Heroku
### Init Heroku
1. `brew tap heroku/brew && brew install heroku`
1. `heroku autocomplete`
1. `printf "$(heroku autocomplete:script zsh)" >> ~/.zshrc; source ~/.zshrc`
1. `heroku plugins:install @heroku-cli/plugin-manifest`

### Create Heroku app
1. `cd AskJeevesSecBot/ButlingButler`
1. `heroku create butlingbutler`

### Create Postgres database
1. `heroku addons:create heroku-postgresql:database`

### Add env vars
1. `heroku config:set ASKJEEVES_USERNAME=<ASKJEEVES_USERNAME>`
1. `heroku config:set ASKJEEVES_PASSWORD=<ASKJEEVES_PASSWORD>`
1. `heroku config:set GOOGLE_MAPS_SIZE_API_KEY=<GOOGLE_MAPS_SIZE_API_KEY>`
1. `heroku config:set SLACK_TOKEN=<SLACK_TOKEN>`
1. `heroku config:set SLACK_SGNING_SECRET=<SLACK_SGNING_SECRET>`

### Push Docker image to Heroku
1. `heroku container:login`
1. `heroku container:push web --app butlingbutler`
1. `heroku container:release web --app=butlingbutler`
1. `heroku open --app butlingbutler` 

## Helpful Heroku commands
1. `heroku logs --tail --app=butlingbutler`
1. `heroku apps:destroy butlingbutler --confirm butlingbutler`

## References
### Heroku
* [Heroku Download and install](https://devcenter.heroku.com/articles/heroku-cli)
* [Docker Up and Running 15: Deploy Docker on Heroku](https://www.youtube.com/watch?v=tTwGdUTR5h8)
* [How do I delete/destroy an app on Heroku?](https://help.heroku.com/LGKL6LTN/how-do-i-delete-destroy-an-app-on-heroku)
* [Provisioning Heroku Postgres](https://devcenter.heroku.com/articles/heroku-postgresql#provisioning-heroku-postgres)
* [Using the Heroku CLI to set env vars](https://devcenter.heroku.com/articles/config-vars#using-the-heroku-cli)
* [How do I use the $PORT environment variable in container based apps?](https://help.heroku.com/PPBPA231/how-do-i-use-the-port-environment-variable-in-container-based-apps)
* [Container Registry & Runtime (Docker Deploys)](https://devcenter.heroku.com/articles/container-registry-and-runtime)
* [Heroku Logging](https://devcenter.heroku.com/articles/logging#view-logs)

### Flask
* [How to return a 404 error in Flask application](https://code-maven.com/flask-return-404)
* [Implementing a RESTful Web API with Python & Flask](http://blog.luisrei.com/articles/flaskrest.html)
* [How do I generate a random string (of length X, a-z only) in Python? [duplicate]](https://stackoverflow.com/questions/1957273/how-do-i-generate-a-random-string-of-length-x-a-z-only-in-python)
* [How to return images in flask response? [duplicate]](https://stackoverflow.com/questions/8637153/how-to-return-images-in-flask-response)
* [How To Download Image File From Url Use Python Requests Or Wget Module](https://www.dev2qa.com/how-to-download-image-file-from-url-use-python-requests-or-wget-module/)
* [Query strings in Flask | Learning Flask Ep. 11](https://pythonise.com/series/learning-flask/flask-query-strings)
* [Flask Configuration](https://exploreflask.com/en/latest/configuration.html)
* [Different ways to Remove a key from Dictionary in Python | del vs dict.pop()](https://thispointer.com/different-ways-to-remove-a-key-from-dictionary-in-python/)
* [Convert sqlalchemy row object to python dict](https://stackoverflow.com/questions/1958219/convert-sqlalchemy-row-object-to-python-dict)
* [unable to create autoincrementing primary key with flask-sqlalchemy](https://stackoverflow.com/questions/20848300/unable-to-create-autoincrementing-primary-key-with-flask-sqlalchemy)
* [How to query all rows of a table](https://stackoverflow.com/questions/51612876/how-to-query-all-rows-of-a-table/51612954)
* [Implementing a RESTful Web API with Python & Flask](http://blog.luisrei.com/articles/flaskrest.html)
* [Modular Applications with Blueprints](https://flask.palletsprojects.com/en/1.1.x/blueprints/)
* [Python: converting a list of dictionaries to json](https://stackoverflow.com/questions/21525328/python-converting-a-list-of-dictionaries-to-json)
* [Flask-JWT](https://pythonhosted.org/Flask-JWT/)
* [flask_jwt_extended](https://flask-jwt-extended.readthedocs.io/en/stable/basic_usage/)
* [Github - vimalloc/flask-jwt-extended](https://github.com/vimalloc/flask-jwt-extended/blob/master/examples/database_blacklist/app.py)
* [Full-stack tutorial — 3: Flask + jwt](https://medium.com/@riken.mehta/full-stack-tutorial-3-flask-jwt-e759d2ee5727)
* [Simple JWT Authentication with Flask-JWT](https://blog.tecladocode.com/simple-jwt-authentication-with-flask-jwt/)

### Database things
* [mysql_config not found when installing mysqldb python interface](https://stackoverflow.com/questions/7475223/mysql-config-not-found-when-installing-mysqldb-python-interface)
* [ImportError: No module named psycopg2](https://www.odoo.com/forum/help-1/question/importerror-no-module-named-psycopg2-39160)

### Slack
* [Slackmojis](https://slackmojis.com/categories/7-party-parrot-emojis)
* [Upload custom emoji to express your team’s culture](https://slack.com/slack-tips/upload-custom-slack-emoji-to-express-your-unique-office-culture)
* [Simple Python code to send message to Slack channel (without packages)](https://keestalkstech.com/2019/10/simple-python-code-to-send-message-to-slack-channel-without-packages/)
* [python-cn/flask-slackbot](https://github.com/python-cn/flask-slackbot/blob/master/flask_slackbot/base.py)
* [Slack - chat.update](https://api.slack.com/methods/chat.update#arg_attachments)
* [Github - slackapi/python-slackclient](https://github.com/slackapi/python-slackclient)
* [Github -slackapi/node-slack-interactive-messages](https://github.com/slackapi/node-slack-interactive-messages/tree/master/examples/express-all-interactions)
* [Github - slackapi/python-message-menu-example](https://github.com/slackapi/python-message-menu-example/blob/master/example.py)
* [Verify Slack requests in AWS Lambda and Python](https://janikarhunen.fi/verify-slack-requests-in-aws-lambda-and-python)
* [Epoch & Unix Timestamp Conversion Tools](https://www.epochconverter.com/)
* [Slack Interactive Messages: POST request payload has an unexpected format](https://stackoverflow.com/questions/52959991/slack-interactive-messages-post-request-payload-has-an-unexpected-format)
* [Flask - How do I read the raw body in a POST request when the content type is “application/x-www-form-urlencoded”](https://stackoverflow.com/questions/17640687/flask-how-do-i-read-the-raw-body-in-a-post-request-when-the-content-type-is-a)
* [Get the data received in a Flask request](https://stackoverflow.com/questions/10434599/get-the-data-received-in-a-flask-request)