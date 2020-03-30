from flask import current_app as app
import requests

def get_static_map(location):
    """ 
    Input: Takes in a location ins tring format
    Output: Returns image from Google Map static lookup
    """
    # google map format string
    google_map_url = f"{app.config['GOOGLE_MAPS_BASE_URL']}?center={location}&zoom={app.config['GOOGLE_MAPS_ZOOM']}&size={app.config['GOOGLE_MAPS_SIZE']}&key={app.config['GOOGLE_MAPS_SIZE_API_KEY']}"

    # Download image as a stram
    resp = requests.get(google_map_url, stream=True)

    # Set decode_content value to True, otherwise the downloaded image file's size will be zero.
    resp.raw.decode_content = True
    return resp.raw