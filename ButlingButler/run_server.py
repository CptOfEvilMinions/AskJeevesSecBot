from app import create_app
from config import config
import os

"""
Detect if running in Docker
"""
def is_docker():
    path = '/proc/self/cgroup'
    return (
        os.path.exists('/.dockerenv') or
        os.path.isfile(path) and any('docker' in line for line in open(path))
    )

"""
Detect if running in Heroku
"""
def is_heroku():
    if os.environ.get('ON_HEROKU') == "1":
        return True
    return False

CONFIG = None
if is_heroku() == True:
    CONFIG = config.HerokuConfig
elif is_docker() == True:
    CONFIG = config.DockerConfig
else:
    CONFIG = config.DevelopmentConfig

# Init app
app = create_app(CONFIG)

if __name__ == "__main__":
    app.run( host=CONFIG.HOST, port=CONFIG.PORT, debug=CONFIG.DEBUG )