Feature: Email endpoint

  Scenario: Bad request sends the correct response
    Given I have a request with the missing from field
    When I invoke the email Lambda handler
    Then the email response status code should be 400
    And the email response body should contain "errors"
    And the "error" should have the correct fields

  Scenario: Good request sends the correct response
    Given I have a request with all the required fields
    When I invoke the email Lambda handler
    Then the email response status code should be 200
    And the email response body should contain "message"
    And the email sender should be called
