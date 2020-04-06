from werkzeug.security import generate_password_hash, check_password_hash
from . import db

class UserResponse(db.Model):
    """Model for user responses."""

    __tablename__ = 'user-responses'

    # {
    #   'EventID': '123abc', 
    #   'Username': 'bbornholm2194', 
    #   'Timestamp': 'Sun Mar 29 19:40:54 +0000 2020', 
    #   'Location': 'Minnesota, US, NA', 
    #   'IPaddress': '128.101.101.101', 
    #   'VPNhash': '8ce264e0b99905f9db21c6c1a70eb3d88aebd0c5f9f604de5174b9a27a0c4ee5', 
    #   'Device': 'macOS', 
    #   'Hostname': 'starwars', 
    #   'user_selection': 'legitimate_login'
    # }
    ID = db.Column(db.Integer, autoincrement=True)
    EventID = db.Column(db.String(100), primary_key=True)
    Username = db.Column(db.String(100), nullable=False, unique=False)
    Timestamp = db.Column(db.String(100), nullable=False, unique=False)
    Location = db.Column(db.String(100), nullable=False, unique=False)
    IPaddress = db.Column(db.String(100), nullable=False, unique=False)
    VPNHash = db.Column(db.String(100), nullable=False, unique=False)
    Device = db.Column(db.String(100), nullable=False, unique=False)
    Hostname = db.Column(db.String(100), nullable=False, unique=False)
    Selection = db.Column(db.String(100), nullable=False, unique=False)



class Users(db.Model):
    """Model for user accounts."""

    __tablename__ = 'users'

    # ID of entry in table
    id = db.Column(db.Integer, primary_key=True)
    # E-mail for user
    username = db.Column(db.String(120), unique=True, nullable=False)
    # Password for user
    password = db.Column(db.String(120), nullable=False)

    def set_password(self, password):
        """Create hashed password."""
        self.password = generate_password_hash(password, method='sha256')

    def check_password(self, password):
        """Check hashed password."""
        return check_password_hash(self.password, password)

    def __repr__(self):
        return '<User {}>'.format(self.username)

    def __str__(self):
        return "User(id='%s')" % self.id