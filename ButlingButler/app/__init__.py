from flask import Flask

def create_app(config_object):
    # Init flask
    app = Flask(__name__)
    app.config.from_object(config_object)

    with app.app_context(): 
        # Import blueprints
        from app.views.api import api

        # Register blueprints
        app.register_blueprint(api)

        return app