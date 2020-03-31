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
    EvnetID = db.Column(db.String(100), primary_key=True)
    Username = db.Column(db.String(100), nullable=False, unique=False)
    Timestamp = db.Column(db.String(100), nullable=False, unique=False)
    Location = db.Column(db.String(100), nullable=False, unique=False)
    IPaddress = db.Column(db.String(100), nullable=False, unique=False)
    VPNHash = db.Column(db.String(100), nullable=False, unique=False)
    Device = db.Column(db.String(100), nullable=False, unique=False)
    Hostname = db.Column(db.String(100), nullable=False, unique=False)
    Selection = db.Column(db.String(100), nullable=False, unique=False)

    # ID = db.Column(db.Integer, autoincrement=True)
    # EvnetID = db.Column(db.String, primary_key=True)
    # Username = db.Column(db.String, nullable=False, unique=False)
    # Timestamp = db.Column(db.String, nullable=False, unique=False)
    # Location = db.Column(db.String, nullable=False, unique=False)
    # IPaddress = db.Column(db.String, nullable=False, unique=False)
    # VPNHash = db.Column(db.String, nullable=False, unique=False)
    # Device = db.Column(db.String, nullable=False, unique=False)
    # Hostname = db.Column(db.String, nullable=False, unique=False)
    # Selection = db.Column(db.String, nullable=False, unique=False)
    