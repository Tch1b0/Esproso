FROM python:latest

COPY /snake.py ./snake.py

EXPOSE 3000:3000

CMD ["python", "snake.py"]