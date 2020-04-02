from flask import request, Blueprint, jsonify
from flask_jwt_extended import jwt_required, get_jwt_identity
from datetime import datetime

# Add blueprints
status = Blueprint('status', __name__, url_prefix="/status", template_folder='templates')

# Create status route
@status.route('/', methods=['GET'])
@jwt_required
def check_status():
    return '{"status": "running"}'

# Create status route
@status.route('/identity', methods=['GET'])
@jwt_required
def get_identity():
    current_user = get_jwt_identity()
    return jsonify(logged_in_as=current_user), 200