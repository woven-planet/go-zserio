from testdata.index_operator.schema.api import IndexOperator, Block, BlockHeader


def new():
    headers: list[BlockHeader] = [
        BlockHeader(num_items_=3),
        BlockHeader(num_items_=0),
        BlockHeader(num_items_=1),
    ]

    array = IndexOperator(
        num_blocks_=3,
        headers_=headers,
        blocks_=[
            Block(
                header_=headers[0],
                items_=[1, 2, 3],
                condition_item_=100,
            ),            
            Block(
				header_=headers[1],
                # len(items_) is 0, since BlockHeader.num_items_ is 0.
				# For the same reason, condition_item_ won't be serialized.
                items_=[],
                condition_item_=1234, # This value should be ignored.
            ),
            Block(
                header_=headers[2],
                items_=[10],
                condition_item_=200,
            ),
        ]
    )
    return array
