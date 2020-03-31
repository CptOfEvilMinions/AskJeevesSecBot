from flask_sqlalchemy import SQLAlchemy
from flask import Flask
import slack

# init SQLAlchemy
db = SQLAlchemy()

def create_app(config_object):
    # Init flask
    app = Flask(__name__)
    app.config.from_object(config_object)

    with app.app_context(): 
        # Init DB
        db.init_app(app)

        # Import blueprints
        from app.views.api import api

        # Register blueprints
        app.register_blueprint(api)

        # Initialize Global db
        db.create_all()

        # Create users
        # create_users(db)

        # Initi slack
        app.slack_client = slack.WebClient(app.config['SLACK_TOKEN'])

        return app