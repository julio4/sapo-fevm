FROM python:3
WORKDIR /usr/src/app
COPY . .
RUN apt-get update
RUN apt-get install tree
CMD ["python", "./describe-container.py"]