Feature: Email endpoint

  Scenario: Bad request sends the correct response
    Given I have a request with the missing from field
    When I invoke the email Lambda handler
    Then the response status code should be 400
    And the response body should contain "errors"
    And the "error" should have the correct fields
