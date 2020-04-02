from flask import current_app as app
import time, json, hashlib, hmac
from datetime import datetime

def verify_slack_request(slack_signature=None, slack_request_timestamp=None, request_body=None) -> bool:
    """
    Input: Slack signature, slack timstamp of message, and body
    Ouput: Return true if verification is sucessful
    """
    # Ensure this isn't a replay attack
    # old message, ignore
    if round(time.time() - float(slack_request_timestamp)) > 60 * 60:
        print (f"[-] {datetime.now()} - Relay attack - {slack_signature}")
        return False

    # Form the basestring as stated in the Slack API docs. We need to make a bytestring.
    concatenated = (b"v0:%b:%b" % (slack_request_timestamp.encode("utf-8"), request_body))

    # Make the Signing Secret a bytestring too.
    signing_secret = bytes(app.config['SLACK_SGNING_SECRET'], 'utf-8')
    

    # Compare the the Slack provided signature to ours.
    # If they are equal, the request should be verified successfully.
    # Log the unsuccessful requests for further analysis
    # (along with another relevant info about the request). '''
    computed_signature = 'v0=' + hmac.new(signing_secret, msg=concatenated, digestmod=hashlib.sha256).hexdigest()
    if hmac.compare_digest(computed_signature, slack_signature):
        print (f"[+] {datetime.now()} - Verification sucessful")
        return True
    else:
        print (f"[-] {datetime.now()} - Verification failed")
        return False