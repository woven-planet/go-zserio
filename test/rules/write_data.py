if __name__ == "__main__":
    import sys
    import zserio
    import data

    sys.stdout.buffer.write(zserio.serialize_to_bytes(data.new()))
