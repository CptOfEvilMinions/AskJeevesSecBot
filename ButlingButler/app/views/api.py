from flask import Blueprint, request, abort, send_file, Response, make_response
from app.helpers.google_maps import get_static_map
from app.helpers.slack import verify_slack_request
from flask_jwt_extended import jwt_required
from app.model import db,UserResponse
from flask import current_app as app
import json


# Add blueprints
api = Blueprint('api', __name__, url_prefix="/askjeeves", template_folder='templates')


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


@api.route('/', methods=['POST'])
@api.route('/UserResponse', methods=['POST'])
def user_response():
    if request.headers.get("X-Slack-Signature") and request.headers.get("X-Slack-Request-Timestamp") and request.headers["Content-Type"] == "application/x-www-form-urlencoded":
        request_body = request.get_data()
        slack_signature = request.headers.get('X-Slack-Signature', None)
        slack_request_timestamp = request.headers.get('X-Slack-Request-Timestamp', None)
        if verify_slack_request(slack_signature, slack_request_timestamp, request_body):
            # Get URL encoded form data
            payload = json.loads(request.form['payload'])

            # Unpack values from fields
            temp_dict = dict()
            for field in payload['message']['blocks'][3]['fields']:
                temp_dict[field['text'].split("*\n")[0][1:]] = field['text'].split("*\n")[1]
            temp_dict['Username'] = payload['user']['username']
            temp_dict['user_selection'] = payload['actions'][0]['value']

            # Create DB entry
            userResponse = UserResponse(
                EventID=temp_dict['EventID'],
                Username=temp_dict['Username'],
                Timestamp=temp_dict['Timestamp'],
                Location=temp_dict['Location'],
                IPaddress=temp_dict['IPaddress'],
                VPNHash=temp_dict['VPNhash'],
                Device=temp_dict['Device'],
                Hostname=temp_dict['Hostname'],
                Selection=temp_dict['user_selection']
            )

            # Commit DB entry
            db.session.add(userResponse)
            db.session.commit()

            # remove blocks
            del payload['message']['blocks']

            selection = payload['actions'][0]['value']

            msg_text = str()
            if selection == "legitimate_login":
                msg_text = ":partyparrot:"
            else:
                msg_text = ":rotating-light-red:  :rotating-light-red:  :rotating-light-red:  Alerting security team :rotating-light-red:  :rotating-light-red:  :rotating-light-red: "

    
            response = app.slack_client.chat_update(
                channel=payload["channel"]["id"],
                ts=payload['container']["message_ts"],
                text=msg_text,
                blocks=[],
                attachments=[]
            )
            return make_response("", 200)

    return abort(404)


@api.route('/GetUserResponse', methods=['GET'])
@jwt_required
def get_user_responses():
    """
    Input: Request to get all the user responses in MySQL database
    Output: Return JSON list of all user responses
    """
    # Request all user responses from DB
    userResponses = db.session.query(UserResponse).all()

    # Delete all entries
    for userResponse in userResponses:
        db.session.delete(userResponse)
    db.session.commit()

    # Create list of dicts of each DB entry
    userResponseLst = list()
    for userResponse in userResponses:
        temp = userResponse.__dict__
        del temp['_sa_instance_state']
        userResponseLst.append(temp)        

    # return user responses as JSON
    return json.dumps(userResponseLst)
