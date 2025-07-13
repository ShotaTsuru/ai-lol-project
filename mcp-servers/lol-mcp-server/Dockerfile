FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt ./
RUN apt-get update && apt-get install -y git
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install --no-cache-dir git+https://github.com/jlowin/fastmcp.git

COPY . .

EXPOSE 8080

CMD ["python", "main.py"] 