FROM ubuntu:18.04

# Install
RUN apt-get update -y && apt-get install python3-minimal python3-pip libmysqlclient-dev -y

# Createuser
RUN useradd -ms /bin/bash app

# Copy app
RUN mkdir /ButlingButler
COPY app /ButlingButler/app
COPY config /ButlingButler/config
COPY run_server.py /ButlingButler/run_server.py
COPY requirements.txt /ButlingButler/requirements.txt

# Install Python modules
RUN pip3 install -r /ButlingButler/requirements.txt

# Change perms
RUN chown app:app -R /ButlingButler

EXPOSE 5000
WORKDIR /ButlingButler
USER app

CMD ["python3","-u", "run_server.py"]