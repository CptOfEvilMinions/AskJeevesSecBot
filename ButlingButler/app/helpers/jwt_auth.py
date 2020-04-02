from flask import current_app as app

def authenticate(username, password):
    from app.model import Users

    user = Users.query.filter_by(username=username).first()

    if not user:
        return False
    if not user.check_password(password):
        return False
    return user

def identity(payload):
    from app.model import Users
    user_id = payload['identity']
    return Users.find_by_id(user_id)