FROM ubuntu:18.04

# Install
RUN apt-get update -y && apt-get install python3-minimal python3-pip -y

# Copy app
RUN mkdir /ButlingButler
COPY app /ButlingButler/app
COPY config /ButlingButler/config
COPY run_server.py /ButlingButler/run_server.py
COPY requirements.txt /ButlingButler/requirements.txt

# Install Python modules
RUN pip3 install -r /ButlingButler/requirements.txt

EXPOSE 5000
WORKDIR /ButlingButler

CMD ["python3","-u", "run_server.py"]