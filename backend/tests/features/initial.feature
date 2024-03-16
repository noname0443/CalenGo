Feature: Start Server
    Scenario: Starting server and testing it
        When I start the server on "0.0.0.0:3000"
        And The server is working
    
    Scenario: Note handler tests
        When I GET note /api/v1/note/1
        Then I got line "decode response: unexpected status code: 501"

        When I POST note /api/v1/note
        """
        {
            "name": "note1",
            "description": "note1",
            "startTime": "",
            "endTime": ""
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I PUT note /api/v1/note
        """
        {
            "name": "note1",
            "description": "note2",
            "startTime": "",
            "endTime": ""
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I DELETE note /api/v1/note
        """
        {
            "name": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"

    Scenario: User handler tests
        When I GET user /api/v1/user/1
        Then I got line "decode response: unexpected status code: 501"

        When I POST user /api/v1/user
        """
        {
            "name": "note1",
            "password": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I PUT user /api/v1/user
        """
        {
            "name": "note1",
            "password": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I DELETE user /api/v1/user
        """
        {
            "name": "note1",
            "password": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"