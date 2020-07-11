FROM python:latest
RUN mkdir app/
ADD . /app/
WORKDIR /app
RUN pip install grpcio-tools
CMD ["python", "grpc.client.py"]