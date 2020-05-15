FROM ubuntu:18.04

# Install
RUN apt-get update -y && apt-get install python3-minimal python3-pip libmysqlclient-dev libpq-dev -y

# Createuser
RUN useradd -ms /bin/bash app

# Copy app
RUN mkdir /ButlingButler
COPY requirements.txt /ButlingButler/requirements.txt

# Install Python modules
RUN pip3 install -r /ButlingButler/requirements.txt

COPY app /ButlingButler/app
COPY config/config.py /ButlingButler/config/config.py
COPY run_server.py /ButlingButler/run_server.py

# Change perms
RUN chown app:app -R /ButlingButler

WORKDIR /ButlingButler
USER app

CMD ["python3","-u", "run_server.py"]