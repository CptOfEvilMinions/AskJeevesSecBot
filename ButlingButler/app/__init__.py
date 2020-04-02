from app.helpers.create_db_users import create_users
from app.helpers.jwt_auth import authenticate, identity
from flask_sqlalchemy import SQLAlchemy
from flask_jwt_extended import JWTManager
from flask import Flask
import slack

# init SQLAlchemy
db = SQLAlchemy()

# Init JWT
jwt = JWTManager()

def create_app(config_object):
    # Init flask
    app = Flask(__name__)
    app.config.from_object(config_object)

    with app.app_context(): 
        # Init DB
        db.init_app(app)

        # Add JWT
        jwt.init_app(app)

        # Import blueprints
        from app.views.api import api
        from app.views.status import status
        from app.views.auth import auth

        # Register blueprints
        app.register_blueprint(api)
        app.register_blueprint(status)
        app.register_blueprint(auth)

        # Initialize Global db
        db.create_all()

        # Create users
        create_users(db)

        # Initi slack
        app.slack_client = slack.WebClient(app.config['SLACK_TOKEN'])

        return app