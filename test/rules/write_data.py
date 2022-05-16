import importlib
import sys
import zserio

if __name__ == "__main__":
    data = importlib.import_module(sys.argv[1])
    sys.stdout.buffer.write(zserio.serialize_to_bytes(data.new()))
