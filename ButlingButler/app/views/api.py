from flask import redirect, render_template, flash, Blueprint, request, session, url_for, abort,send_file
from flask import current_app as app
from app.helpers.google_maps import get_static_map
import time
import requests
import urllib.parse
import json


# Add blueprints
api = Blueprint('api', __name__, url_prefix="/askjeeves", template_folder='templates')

# def verify_signing_key(slack_signing_secret, slack_token, slack_timestamp, slack_signature) -> bool:
#     """
#     """
#     if abs(time.time() - slack_timestamp) > 60 * 5:
#         return False

#     sig_basestring = "v0:" + slack_timestamp + ":token=" + slack_token

#     my_signature = 'v0=' + hmac.compute_hash_sha256(
#                     slack_signing_secret,
#                     sig_basestring
#                 ).hexdigest()

#     if hmac.compare(my_signature, slack_signature):
#         return True
#     return False
    


@api.route('/GoogleMaps', methods=['GET'])
def get_google_map():
    """
    Input: Takes in a location via the query string in the URL
    Output: Returns a Google Map static image (PNG) to client
    """
    # Get query string
    query_string = request.args
    image = get_static_map(query_string['location'])
    return send_file(image, mimetype='image/png')



@api.route('/UserResponse', methods=['POST'])
def user_response():
    # Make sure HTTP header includes Slack headers
    #if request.headers.get("X-Slack-Signature") and request.headers.get("X-Slack-Request-Timestamp") and request.headers["Content-Type"] == "application/x-www-form-urlencoded":
    # Get URL encoded form data
    payload = json.loads(request.form['payload'])

    # Unpack values from fields
    temp_dict = dict()
    for field in payload['message']['blocks'][3]['fields']:
        temp_dict[field['text'].split("*\n")[0][1:]] = field['text'].split("*\n")[1]
    temp_dict['Username'] = payload['user']['username']
    temp_dict['user_selection'] = payload['actions'][0]['value']
    print (temp_dict)


    return abort(404)

