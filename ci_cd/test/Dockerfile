# syntax=docker/dockerfile:1

FROM python:3.10.8-slim-buster

RUN useradd -rm -d /home/github -s /bin/bash -u 1000 github

USER github

WORKDIR /home/github

ENV PATH=/home/github/.local/bin:$PATH

COPY requirements.txt .

RUN pip install --upgrade pip

RUN pip install -r requirements.txt