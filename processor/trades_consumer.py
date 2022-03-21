
from queue import Queue
from pyamqp.rabbit.receiver import Receiver


class Consumer(Receiver):
    def __init__(self,
                 trades_buffer: Queue,
                 exchange: str,
                 host: str,
                 port: int = 5672,
                 username: str = 'guest',
                 password: str = 'guest'):
        super().__init__(host, port, username, password, threaded=True)
        self.exchange = exchange
        self.trades_buffer = trades_buffer

    def on_message(self, message, details):
        self.trades_buffer.put(message)


