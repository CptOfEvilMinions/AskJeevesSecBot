from flask_jwt_extended import create_access_token
from app.helpers.jwt_auth import authenticate
from flask import request, Blueprint, jsonify
from app.model import Users
from datetime import datetime

# Add blueprints
auth = Blueprint('auth', __name__, url_prefix="/auth", template_folder='templates')

# Create status route
@auth.route('/login', methods=['POST'])
def login():
    if not request.is_json:
        return jsonify({"msg": "Missing JSON in request"}), 400

    # Get username and password
    username = request.json.get('username', None)
    password = request.json.get('password', None)

    if not username:
        return jsonify({"msg": "Missing username parameter"}), 400
    if not password:
        return jsonify({"msg": "Missing password parameter"}), 400

    # Do DB lookup
    if authenticate(username,password) == False :
        return jsonify({"msg": "Bad username or password"}), 401

    # Identity can be any data that is json serializable
    access_token = create_access_token(identity=username)
    return jsonify(access_token=access_token), 200
