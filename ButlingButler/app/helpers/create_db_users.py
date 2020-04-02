from werkzeug.security import generate_password_hash
from flask import current_app as app
from datetime import datetime

def create_users(db):
    print ('[*] {0} - Adding users to datbase'.format( datetime.now() ))
    # Create users
    from app.model import Users

    # Add AskJeeves user

    existing_user = Users.query.filter_by(username=app.config['ASKJEEVES_USERNAME']).first()
    if existing_user is None:
        user = Users(
            username=app.config['ASKJEEVES_USERNAME'],
            password=generate_password_hash(app.config['ASKJEEVES_PASSWORD'], method='sha256'),
        )
        db.session.add(user)

        # Commit user to database
        db.session.commit()

    # Query all users
    users = Users.query.all()
    print (users)