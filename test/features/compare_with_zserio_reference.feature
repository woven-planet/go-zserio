Feature:  Compare go-zserio with the zserio Reference Implementation
  Scenario: Re-encode a zserio Binary with Go
    Given I have a zserio binary created using the reference implementation
    When I decode this binary with go-zerio and encode it again
    Then The binary content should be the same