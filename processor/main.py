from config import cfg
from trades_consumer import Consumer
from queue import Queue
from threading import Thread

if __name__ == "__main__":
    buffer = Queue()
    consumer = Consumer(buffer, "trades", **cfg.rabbit_login.dict())
    consumer.connect_queue(
        queue_name="consumer",
        exchange="trades",
        routing_keys=["BTC-PERP", "ETH-PERP"],
        is_durable=True,
        auto_delete=True
    )
    th = Thread(consumer.consume())
    th.start()

    while buffer.not_empty:
        trade = buffer.get()
        print(trade)
        buffer.task_done()
