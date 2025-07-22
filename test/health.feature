Feature: Health check endpoint

  Scenario: Calling the health check
    When I invoke the health check Lambda handler
    Then the response status code should be 200
    And the response body should contain "message"
