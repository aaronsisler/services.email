Feature: Health check endpoint

  Scenario: Calling the health check
    When I invoke the health check Lambda handler
    Then the health response status code should be 200
    And the health response body should contain "message"
