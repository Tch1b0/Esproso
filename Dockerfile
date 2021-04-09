FROM python:3.7.2

WORKDIR /app

COPY /snake.py ./snake.py
COPY /requirements.txt ./requirements.txt

RUN pip install -r requirements.txt

CMD ["python", "snake.py"]